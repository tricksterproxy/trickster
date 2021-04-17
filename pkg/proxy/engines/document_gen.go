/*
 * Copyright 2018 The Trickster Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package engines

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tricksterproxy/trickster/pkg/proxy/ranges/byterange"

	"github.com/tinylib/msgp/msgp"
)

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
		case "content_length":
			z.ContentLength, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "content_type":
			z.ContentType, err = dc.ReadString()
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
		case "ranges":
			err = z.Ranges.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "range_parts":
			var zb0004 uint32
			zb0004, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.StoredRangeParts == nil {
				z.StoredRangeParts = make(map[string]*byterange.MultipartByteRange, zb0004)
			} else if len(z.StoredRangeParts) > 0 {
				for key := range z.StoredRangeParts {
					delete(z.StoredRangeParts, key)
				}
			}
			for zb0004 > 0 {
				zb0004--
				var za0004 string
				var za0005 *byterange.MultipartByteRange
				za0004, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					za0005 = nil
				} else {
					if za0005 == nil {
						za0005 = new(byterange.MultipartByteRange)
					}
					err = za0005.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.StoredRangeParts[za0004] = za0005
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
	// map header, size 9
	// write "status_code"
	err = en.Append(0x89, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
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
	// write "content_length"
	err = en.Append(0xae, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ContentLength)
	if err != nil {
		return
	}
	// write "content_type"
	err = en.Append(0xac, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.ContentType)
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
	// write "ranges"
	err = en.Append(0xa6, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73)
	if err != nil {
		return
	}
	err = z.Ranges.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "range_parts"
	err = en.Append(0xab, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.StoredRangeParts)))
	if err != nil {
		return
	}
	for za0004, za0005 := range z.StoredRangeParts {
		err = en.WriteString(za0004)
		if err != nil {
			return
		}
		if za0005 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = za0005.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HTTPDocument) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "status_code"
	o = append(o, 0x89, 0xab, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65)
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
	// string "content_length"
	o = append(o, 0xae, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendInt64(o, z.ContentLength)
	// string "content_type"
	o = append(o, 0xac, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.ContentType)
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
	// string "ranges"
	o = append(o, 0xa6, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73)
	o, err = z.Ranges.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "range_parts"
	o = append(o, 0xab, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.StoredRangeParts)))
	for za0004, za0005 := range z.StoredRangeParts {
		o = msgp.AppendString(o, za0004)
		if za0005 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = za0005.MarshalMsg(o)
			if err != nil {
				return
			}
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
		case "content_length":
			z.ContentLength, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "content_type":
			z.ContentType, bts, err = msgp.ReadStringBytes(bts)
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
		case "ranges":
			bts, err = z.Ranges.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "range_parts":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.StoredRangeParts == nil {
				z.StoredRangeParts = make(map[string]*byterange.MultipartByteRange, zb0004)
			} else if len(z.StoredRangeParts) > 0 {
				for key := range z.StoredRangeParts {
					delete(z.StoredRangeParts, key)
				}
			}
			for zb0004 > 0 {
				var za0004 string
				var za0005 *byterange.MultipartByteRange
				zb0004--
				za0004, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					za0005 = nil
				} else {
					if za0005 == nil {
						za0005 = new(byterange.MultipartByteRange)
					}
					bts, err = za0005.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.StoredRangeParts[za0004] = za0005
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
	s += 5 + msgp.BytesPrefixSize + len(z.Body) + 15 + msgp.Int64Size + 13 + msgp.StringPrefixSize + len(z.ContentType) + 15
	if z.CachingPolicy == nil {
		s += msgp.NilSize
	} else {
		s += z.CachingPolicy.Msgsize()
	}
	s += 7 + z.Ranges.Msgsize() + 12 + msgp.MapHeaderSize
	if z.StoredRangeParts != nil {
		for za0004, za0005 := range z.StoredRangeParts {
			_ = za0005
			s += msgp.StringPrefixSize + len(za0004)
			if za0005 == nil {
				s += msgp.NilSize
			} else {
				s += za0005.Msgsize()
			}
		}
	}
	return
}
