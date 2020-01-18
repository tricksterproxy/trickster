package byterange

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MultipartByteRange) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "range":
			err = z.Range.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "content":
			z.Content, err = dc.ReadBytes(z.Content)
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
func (z *MultipartByteRange) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "range"
	err = en.Append(0x82, 0xa5, 0x72, 0x61, 0x6e, 0x67, 0x65)
	if err != nil {
		return
	}
	err = z.Range.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "content"
	err = en.Append(0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Content)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MultipartByteRange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "range"
	o = append(o, 0x82, 0xa5, 0x72, 0x61, 0x6e, 0x67, 0x65)
	o, err = z.Range.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "content"
	o = append(o, 0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendBytes(o, z.Content)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MultipartByteRange) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "range":
			bts, err = z.Range.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "content":
			z.Content, bts, err = msgp.ReadBytesBytes(bts, z.Content)
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
func (z *MultipartByteRange) Msgsize() (s int) {
	s = 1 + 6 + z.Range.Msgsize() + 8 + msgp.BytesPrefixSize + len(z.Content)
	return
}
