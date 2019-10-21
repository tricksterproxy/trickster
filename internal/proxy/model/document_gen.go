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

package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CachingPolicy) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "is_fresh":
			z.IsFresh, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "nocache":
			z.NoCache, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "notransform":
			z.NoTransform, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "freshness_lifetime":
			z.FreshnessLifetime, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "can_revalidate":
			z.CanRevalidate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "must_revalidate":
			z.MustRevalidate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "last_modified":
			z.LastModified, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "expires":
			z.Expires, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "date":
			z.Date, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "local_date":
			z.LocalDate, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "etag":
			z.ETag, err = dc.ReadString()
			if err != nil {
				return
			}
		case "if_none_match_value":
			z.IfNoneMatchValue, err = dc.ReadString()
			if err != nil {
				return
			}
		case "if_match_value":
			z.IfMatchValue, err = dc.ReadString()
			if err != nil {
				return
			}
		case "if_modified_since_time":
			z.IfModifiedSinceTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "if_unmodified_since_time":
			z.IfUnmodifiedSinceTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CachingPolicy) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "is_fresh"
	err = en.Append(0x8f, 0xa8, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsFresh)
	if err != nil {
		return
	}
	// write "nocache"
	err = en.Append(0xa7, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.NoCache)
	if err != nil {
		return
	}
	// write "notransform"
	err = en.Append(0xab, 0x6e, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteBool(z.NoTransform)
	if err != nil {
		return
	}
	// write "freshness_lifetime"
	err = en.Append(0xb2, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.FreshnessLifetime)
	if err != nil {
		return
	}
	// write "can_revalidate"
	err = en.Append(0xae, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.CanRevalidate)
	if err != nil {
		return
	}
	// write "must_revalidate"
	err = en.Append(0xaf, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.MustRevalidate)
	if err != nil {
		return
	}
	// write "last_modified"
	err = en.Append(0xad, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastModified)
	if err != nil {
		return
	}
	// write "expires"
	err = en.Append(0xa7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Expires)
	if err != nil {
		return
	}
	// write "date"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Date)
	if err != nil {
		return
	}
	// write "local_date"
	err = en.Append(0xaa, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LocalDate)
	if err != nil {
		return
	}
	// write "etag"
	err = en.Append(0xa4, 0x65, 0x74, 0x61, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.ETag)
	if err != nil {
		return
	}
	// write "if_none_match_value"
	err = en.Append(0xb3, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.IfNoneMatchValue)
	if err != nil {
		return
	}
	// write "if_match_value"
	err = en.Append(0xae, 0x69, 0x66, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.IfMatchValue)
	if err != nil {
		return
	}
	// write "if_modified_since_time"
	err = en.Append(0xb6, 0x69, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.IfModifiedSinceTime)
	if err != nil {
		return
	}
	// write "if_unmodified_since_time"
	err = en.Append(0xb8, 0x69, 0x66, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.IfUnmodifiedSinceTime)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CachingPolicy) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "is_fresh"
	o = append(o, 0x8f, 0xa8, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x73, 0x68)
	o = msgp.AppendBool(o, z.IsFresh)
	// string "nocache"
	o = append(o, 0xa7, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65)
	o = msgp.AppendBool(o, z.NoCache)
	// string "notransform"
	o = append(o, 0xab, 0x6e, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendBool(o, z.NoTransform)
	// string "freshness_lifetime"
	o = append(o, 0xb2, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.FreshnessLifetime)
	// string "can_revalidate"
	o = append(o, 0xae, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.CanRevalidate)
	// string "must_revalidate"
	o = append(o, 0xaf, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.MustRevalidate)
	// string "last_modified"
	o = append(o, 0xad, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendTime(o, z.LastModified)
	// string "expires"
	o = append(o, 0xa7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73)
	o = msgp.AppendTime(o, z.Expires)
	// string "date"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendTime(o, z.Date)
	// string "local_date"
	o = append(o, 0xaa, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendTime(o, z.LocalDate)
	// string "etag"
	o = append(o, 0xa4, 0x65, 0x74, 0x61, 0x67)
	o = msgp.AppendString(o, z.ETag)
	// string "if_none_match_value"
	o = append(o, 0xb3, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.IfNoneMatchValue)
	// string "if_match_value"
	o = append(o, 0xae, 0x69, 0x66, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.IfMatchValue)
	// string "if_modified_since_time"
	o = append(o, 0xb6, 0x69, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.IfModifiedSinceTime)
	// string "if_unmodified_since_time"
	o = append(o, 0xb8, 0x69, 0x66, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.IfUnmodifiedSinceTime)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CachingPolicy) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "is_fresh":
			z.IsFresh, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "nocache":
			z.NoCache, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "notransform":
			z.NoTransform, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "freshness_lifetime":
			z.FreshnessLifetime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "can_revalidate":
			z.CanRevalidate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "must_revalidate":
			z.MustRevalidate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "last_modified":
			z.LastModified, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "expires":
			z.Expires, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "date":
			z.Date, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "local_date":
			z.LocalDate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "etag":
			z.ETag, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "if_none_match_value":
			z.IfNoneMatchValue, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "if_match_value":
			z.IfMatchValue, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "if_modified_since_time":
			z.IfModifiedSinceTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "if_unmodified_since_time":
			z.IfUnmodifiedSinceTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CachingPolicy) Msgsize() (s int) {
	s = 1 + 9 + msgp.BoolSize + 8 + msgp.BoolSize + 12 + msgp.BoolSize + 19 + msgp.IntSize + 15 + msgp.BoolSize + 16 + msgp.BoolSize + 14 + msgp.TimeSize + 8 + msgp.TimeSize + 5 + msgp.TimeSize + 11 + msgp.TimeSize + 5 + msgp.StringPrefixSize + len(z.ETag) + 20 + msgp.StringPrefixSize + len(z.IfNoneMatchValue) + 15 + msgp.StringPrefixSize + len(z.IfMatchValue) + 23 + msgp.TimeSize + 25 + msgp.TimeSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HTTPDocument) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "status_code":
			z.StatusCode, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "status":
			z.Status, err = dc.ReadString()
			if err != nil {
				return
			}
		case "headers":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Headers == nil {
				z.Headers = make(map[string][]string, zb0002)
			} else if len(z.Headers) > 0 {
				for key := range z.Headers {
					delete(z.Headers, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 []string
				za0001, err = dc.ReadString()
				if err != nil {
					return
				}
				var zb0003 uint32
				zb0003, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(za0002) >= int(zb0003) {
					za0002 = (za0002)[:zb0003]
				} else {
					za0002 = make([]string, zb0003)
				}
				for za0003 := range za0002 {
					za0002[za0003], err = dc.ReadString()
					if err != nil {
						return
					}
				}
				z.Headers[za0001] = za0002
			}
		case "body":
			z.Body, err = dc.ReadBytes(z.Body)
			if err != nil {
				return
			}
		case "caching_policy":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.CachingPolicy = nil
			} else {
				if z.CachingPolicy == nil {
					z.CachingPolicy = new(CachingPolicy)
				}
				err = z.CachingPolicy.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *HTTPDocument) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "status_code"
	err = en.Append(0x85, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.StatusCode)
	if err != nil {
		return
	}
	// write "status"
	err = en.Append(0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Status)
	if err != nil {
		return
	}
	// write "headers"
	err = en.Append(0xa7, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Headers)))
	if err != nil {
		return
	}
	for za0001, za0002 := range z.Headers {
		err = en.WriteString(za0001)
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(za0002)))
		if err != nil {
			return
		}
		for za0003 := range za0002 {
			err = en.WriteString(za0002[za0003])
			if err != nil {
				return
			}
		}
	}
	// write "body"
	err = en.Append(0xa4, 0x62, 0x6f, 0x64, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Body)
	if err != nil {
		return
	}
	// write "caching_policy"
	err = en.Append(0xae, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79)
	if err != nil {
		return
	}
	if z.CachingPolicy == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.CachingPolicy.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HTTPDocument) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "status_code"
	o = append(o, 0x85, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, z.StatusCode)
	// string "status"
	o = append(o, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "headers"
	o = append(o, 0xa7, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Headers)))
	for za0001, za0002 := range z.Headers {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendArrayHeader(o, uint32(len(za0002)))
		for za0003 := range za0002 {
			o = msgp.AppendString(o, za0002[za0003])
		}
	}
	// string "body"
	o = append(o, 0xa4, 0x62, 0x6f, 0x64, 0x79)
	o = msgp.AppendBytes(o, z.Body)
	// string "caching_policy"
	o = append(o, 0xae, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79)
	if z.CachingPolicy == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.CachingPolicy.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HTTPDocument) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "status_code":
			z.StatusCode, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "headers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Headers == nil {
				z.Headers = make(map[string][]string, zb0002)
			} else if len(z.Headers) > 0 {
				for key := range z.Headers {
					delete(z.Headers, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 []string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(za0002) >= int(zb0003) {
					za0002 = (za0002)[:zb0003]
				} else {
					za0002 = make([]string, zb0003)
				}
				for za0003 := range za0002 {
					za0002[za0003], bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				}
				z.Headers[za0001] = za0002
			}
		case "body":
			z.Body, bts, err = msgp.ReadBytesBytes(bts, z.Body)
			if err != nil {
				return
			}
		case "caching_policy":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.CachingPolicy = nil
			} else {
				if z.CachingPolicy == nil {
					z.CachingPolicy = new(CachingPolicy)
				}
				bts, err = z.CachingPolicy.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HTTPDocument) Msgsize() (s int) {
	s = 1 + 12 + msgp.IntSize + 7 + msgp.StringPrefixSize + len(z.Status) + 8 + msgp.MapHeaderSize
	if z.Headers != nil {
		for za0001, za0002 := range z.Headers {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.ArrayHeaderSize
			for za0003 := range za0002 {
				s += msgp.StringPrefixSize + len(za0002[za0003])
			}
		}
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Body) + 15
	if z.CachingPolicy == nil {
		s += msgp.NilSize
	} else {
		s += z.CachingPolicy.Msgsize()
	}
	return
}
