/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cache

import (
	"sync"
	"time"

	"github.com/Comcast/trickster/internal/util/log"
)

//go:generate msgp

// IndexKey is the key under which the index will write itself to its associated cache
const IndexKey = "cache.index"

var indexLock = sync.Mutex{}

// Index maintains metadata about a Cache when Retention enforcement is managed internally,
// like memory or bbolt. It is not used for independently managed caches like Redis.
type Index struct {
	// CacheSize represents the size of the cache in bytes
	CacheSize int `msg="cache_size"`
	// ObjectCount represents the count of objects in the Cache
	ObjectCount int `msg="object_count`
	// Objects is a map of Objects in the Cache
	Objects map[string]*Object `msg="objects"`

	name           string                             `msg="-"`
	bulkRemoveFunc func([]string, bool)               `msg="-"`
	reapInterval   time.Duration                      `msg="-"`
	flushInterval  time.Duration                      `msg="-"`
	flushFunc      func(cacheKey string, data []byte) `msg="-"`
	maxCacheSize   int                                `msg="-"`
}

// ToBytes returns a serialized byte slice representing the Index
func (idx *Index) ToBytes() []byte {
	bytes, _ := idx.MarshalMsg(nil)
	return bytes
}

// IndexFromBytes returns a deserialized Cache Object from a seralized byte slice
func IndexFromBytes(data []byte) (*Index, error) {
	i := &Index{}
	_, err := i.UnmarshalMsg(data)
	return i, err
}

// Object contains metadataa about an item in the Cache
type Object struct {
	// Key represents the name of the Object and is the accessor in a hashed collection of Cache Objects
	Key string `msg:"key"`
	// Expiration represents the time that the Object expires from Cache
	Expiration time.Time `msg:"expiration"`
	// LastWrite is the time the object was last Written
	LastWrite time.Time `msg:"lastwrite"`
	// LastAccess is the time the object was last Accessed
	LastAccess time.Time `msg:"lastaccess"`
	// Size the size of the Object in bytes
	Size int `msg:"size"`
	// Value is the value of the Object stored in the Cache
	// It is used by Caches but not by the Index
	Value []byte `msg:"value,omitempty"`
}

// ToBytes returns a serialized byte slice representing the Object
func (o *Object) ToBytes() []byte {
	bytes, _ := o.MarshalMsg(nil)
	return bytes
}

// ObjectFromBytes returns a deserialized Cache Object from a seralized byte slice
func ObjectFromBytes(data []byte) (*Object, error) {
	o := &Object{}
	_, err := o.UnmarshalMsg(data)
	return o, err
}

// NewIndex returns a new Index based on the provided inputs
func NewIndex(name string, indexData []byte, bulkRemoveFunc func([]string, bool), reapInterval time.Duration, flushInterval time.Duration, flushFunc func(cacheKey string, data []byte)) *Index {
	var i *Index
	if len(indexData) > 0 {
		i = &Index{}
		i.UnmarshalMsg(indexData)
		i.name = name
		i.reapInterval = reapInterval
		i.bulkRemoveFunc = bulkRemoveFunc
		i.flushInterval = flushInterval
		i.flushFunc = flushFunc
	} else {
		i = &Index{name: name, reapInterval: reapInterval, bulkRemoveFunc: bulkRemoveFunc, flushInterval: flushInterval, flushFunc: flushFunc}
		i.Objects = make(map[string]*Object)
	}

	if flushInterval > 0 && flushFunc != nil {
		go i.flusher()
	} else {
		log.Warn("cache index flusher did not start", log.Pairs{"cacheName": name, "flushInterval": flushInterval})
	}

	if reapInterval > 0 {
		go i.reaper()
	} else {
		log.Warn("cache reaper did not start", log.Pairs{"cacheName": name, "reapInterval": reapInterval})
	}

	return i
}

// UpdateObjectAccessTime updates the LastAccess for the object with the provided key
func (idx *Index) UpdateObjectAccessTime(key string) {
	indexLock.Lock()
	defer indexLock.Unlock()
	if _, ok := idx.Objects[key]; ok {
		idx.Objects[key].LastAccess = time.Now()
	}
}

// UpdateObject writes or updates the Index Metadata for the provided Object
func (idx *Index) UpdateObject(obj Object) {

	key := obj.Key
	if key == "" {
		return
	}

	indexLock.Lock()
	defer indexLock.Unlock()

	obj.Size = len(obj.Value)
	obj.Value = nil
	obj.LastAccess = time.Now()
	obj.LastWrite = obj.LastAccess

	if o, ok := idx.Objects[key]; ok {
		idx.CacheSize += (o.Size - idx.Objects[key].Size)
	} else {
		idx.CacheSize += obj.Size
		idx.ObjectCount++
	}

	idx.Objects[key] = &obj
}

// RemoveObject removes an Object's Metadata from the Index
func (idx *Index) RemoveObject(key string, noLock bool) {

	if !noLock {
		indexLock.Lock()
		defer indexLock.Unlock()
	}

	if o, ok := idx.Objects[key]; ok {
		idx.CacheSize -= o.Size
		idx.ObjectCount--
		delete(idx.Objects, key)
	}

}

// flusher continually iterates through the cache to find expired elements and removes them
func (idx *Index) flusher() {
	for {
		time.Sleep(idx.flushInterval)
		indexLock.Lock()
		bytes, err := idx.MarshalMsg(nil)
		indexLock.Unlock()
		if err != nil {
			log.Warn("unable to serialize index for flushing", log.Pairs{"detail": err.Error()})
			continue
		}
		idx.flushFunc(IndexKey, bytes)
	}
}

// reaper continually iterates through the cache to find expired elements and removes them
func (idx *Index) reaper() {
	for {
		idx.reap()
		time.Sleep(idx.reapInterval)
	}
}

// reap makes a single iteration through the cache index to to find and remove expired elements
// and evict least-recently-accessed elements to maintain the Maximum allowed Cache Size
func (idx *Index) reap() {

	indexLock.Lock()
	defer indexLock.Unlock()

	removals := make([]string, 0, 0)
	remainders := make([]*Object, 0, idx.ObjectCount)

	now := time.Now()

	for _, o := range idx.Objects {
		if o.Key == IndexKey {
			continue
		}
		if o.Expiration.Before(now) {
			removals = append(removals, o.Key)
		} else {
			remainders = append(remainders, o)
		}
	}

	if len(removals) > 0 {
		idx.bulkRemoveFunc(removals, true)
	}

	if idx.CacheSize > idx.maxCacheSize {
		// Sort Remainders by LastAccess
		// Remove until idx.CacheSize is back under the threshold
	}

}
