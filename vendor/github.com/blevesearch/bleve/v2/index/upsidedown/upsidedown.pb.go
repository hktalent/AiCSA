// Code generated by protoc-gen-gogo.
// source: upsidedown.proto
// DO NOT EDIT!

/*
Package upsidedown is a generated protocol buffer package.

It is generated from these files:

	upsidedown.proto

It has these top-level messages:

	BackIndexTermsEntry
	BackIndexStoreEntry
	BackIndexRowValue
*/
package upsidedown

import proto "github.com/golang/protobuf/proto"
import math "math"

import io "io"
import fmt "fmt"
import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type BackIndexTermsEntry struct {
	Field            *uint32  `protobuf:"varint,1,req,name=field" json:"field,omitempty"`
	Terms            []string `protobuf:"bytes,2,rep,name=terms" json:"terms,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BackIndexTermsEntry) Reset()         { *m = BackIndexTermsEntry{} }
func (m *BackIndexTermsEntry) String() string { return proto.CompactTextString(m) }
func (*BackIndexTermsEntry) ProtoMessage()    {}

func (m *BackIndexTermsEntry) GetField() uint32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BackIndexTermsEntry) GetTerms() []string {
	if m != nil {
		return m.Terms
	}
	return nil
}

type BackIndexStoreEntry struct {
	Field            *uint32  `protobuf:"varint,1,req,name=field" json:"field,omitempty"`
	ArrayPositions   []uint64 `protobuf:"varint,2,rep,name=arrayPositions" json:"arrayPositions,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BackIndexStoreEntry) Reset()         { *m = BackIndexStoreEntry{} }
func (m *BackIndexStoreEntry) String() string { return proto.CompactTextString(m) }
func (*BackIndexStoreEntry) ProtoMessage()    {}

func (m *BackIndexStoreEntry) GetField() uint32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BackIndexStoreEntry) GetArrayPositions() []uint64 {
	if m != nil {
		return m.ArrayPositions
	}
	return nil
}

type BackIndexRowValue struct {
	TermsEntries     []*BackIndexTermsEntry `protobuf:"bytes,1,rep,name=termsEntries" json:"termsEntries,omitempty"`
	StoredEntries    []*BackIndexStoreEntry `protobuf:"bytes,2,rep,name=storedEntries" json:"storedEntries,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *BackIndexRowValue) Reset()         { *m = BackIndexRowValue{} }
func (m *BackIndexRowValue) String() string { return proto.CompactTextString(m) }
func (*BackIndexRowValue) ProtoMessage()    {}

func (m *BackIndexRowValue) GetTermsEntries() []*BackIndexTermsEntry {
	if m != nil {
		return m.TermsEntries
	}
	return nil
}

func (m *BackIndexRowValue) GetStoredEntries() []*BackIndexStoreEntry {
	if m != nil {
		return m.StoredEntries
	}
	return nil
}

func (m *BackIndexTermsEntry) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Terms = append(m.Terms, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipUpsidedown(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpsidedown
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	return nil
}
func (m *BackIndexStoreEntry) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayPositions", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ArrayPositions = append(m.ArrayPositions, v)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipUpsidedown(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpsidedown
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	return nil
}
func (m *BackIndexRowValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TermsEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthUpsidedown
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TermsEntries = append(m.TermsEntries, &BackIndexTermsEntry{})
			if err := m.TermsEntries[len(m.TermsEntries)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthUpsidedown
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredEntries = append(m.StoredEntries, &BackIndexStoreEntry{})
			if err := m.StoredEntries[len(m.StoredEntries)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipUpsidedown(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpsidedown
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func skipUpsidedown(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUpsidedown
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUpsidedown(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUpsidedown = fmt.Errorf("proto: negative length found during unmarshaling")
)

func (m *BackIndexTermsEntry) Size() (n int) {
	var l int
	_ = l
	if m.Field != nil {
		n += 1 + sovUpsidedown(uint64(*m.Field))
	}
	if len(m.Terms) > 0 {
		for _, s := range m.Terms {
			l = len(s)
			n += 1 + l + sovUpsidedown(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackIndexStoreEntry) Size() (n int) {
	var l int
	_ = l
	if m.Field != nil {
		n += 1 + sovUpsidedown(uint64(*m.Field))
	}
	if len(m.ArrayPositions) > 0 {
		for _, e := range m.ArrayPositions {
			n += 1 + sovUpsidedown(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackIndexRowValue) Size() (n int) {
	var l int
	_ = l
	if len(m.TermsEntries) > 0 {
		for _, e := range m.TermsEntries {
			l = e.Size()
			n += 1 + l + sovUpsidedown(uint64(l))
		}
	}
	if len(m.StoredEntries) > 0 {
		for _, e := range m.StoredEntries {
			l = e.Size()
			n += 1 + l + sovUpsidedown(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUpsidedown(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUpsidedown(x uint64) (n int) {
	return sovUpsidedown(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackIndexTermsEntry) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BackIndexTermsEntry) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Field == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintUpsidedown(data, i, uint64(*m.Field))
	}
	if len(m.Terms) > 0 {
		for _, s := range m.Terms {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackIndexStoreEntry) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BackIndexStoreEntry) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Field == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintUpsidedown(data, i, uint64(*m.Field))
	}
	if len(m.ArrayPositions) > 0 {
		for _, num := range m.ArrayPositions {
			data[i] = 0x10
			i++
			i = encodeVarintUpsidedown(data, i, uint64(num))
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackIndexRowValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BackIndexRowValue) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TermsEntries) > 0 {
		for _, msg := range m.TermsEntries {
			data[i] = 0xa
			i++
			i = encodeVarintUpsidedown(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.StoredEntries) > 0 {
		for _, msg := range m.StoredEntries {
			data[i] = 0x12
			i++
			i = encodeVarintUpsidedown(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Upsidedown(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Upsidedown(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUpsidedown(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
