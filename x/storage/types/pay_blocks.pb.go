// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine-chain/storage/pay_blocks.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PayBlocks struct {
	Blockid   string `protobuf:"bytes,1,opt,name=blockid,proto3" json:"blockid,omitempty"`
	Bytes     string `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Blocktype string `protobuf:"bytes,3,opt,name=blocktype,proto3" json:"blocktype,omitempty"`
	Blocknum  string `protobuf:"bytes,4,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
}

func (m *PayBlocks) Reset()         { *m = PayBlocks{} }
func (m *PayBlocks) String() string { return proto.CompactTextString(m) }
func (*PayBlocks) ProtoMessage()    {}
func (*PayBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_afcab5148682d329, []int{0}
}
func (m *PayBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PayBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayBlocks.Merge(m, src)
}
func (m *PayBlocks) XXX_Size() int {
	return m.Size()
}
func (m *PayBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_PayBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_PayBlocks proto.InternalMessageInfo

func (m *PayBlocks) GetBlockid() string {
	if m != nil {
		return m.Blockid
	}
	return ""
}

func (m *PayBlocks) GetBytes() string {
	if m != nil {
		return m.Bytes
	}
	return ""
}

func (m *PayBlocks) GetBlocktype() string {
	if m != nil {
		return m.Blocktype
	}
	return ""
}

func (m *PayBlocks) GetBlocknum() string {
	if m != nil {
		return m.Blocknum
	}
	return ""
}

func init() {
	proto.RegisterType((*PayBlocks)(nil), "jackaldao.canine.storage.PayBlocks")
}

func init() {
	proto.RegisterFile("canine-chain/storage/pay_blocks.proto", fileDescriptor_afcab5148682d329)
}

var fileDescriptor_afcab5148682d329 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x4e, 0xcc, 0xcb,
	0xcc, 0x4b, 0xd5, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f,
	0xd5, 0x2f, 0x48, 0xac, 0x8c, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xc8, 0x4a, 0x4c, 0xce, 0x4e, 0xcc, 0x49, 0x49, 0xcc, 0xd7, 0x83, 0x68, 0xd0,
	0x83, 0x2a, 0x55, 0x2a, 0xe5, 0xe2, 0x0c, 0x48, 0xac, 0x74, 0x02, 0x2b, 0x16, 0x92, 0xe0, 0x62,
	0x07, 0x6b, 0xcb, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x44, 0xb8,
	0x58, 0x93, 0x2a, 0x4b, 0x52, 0x8b, 0x25, 0x98, 0xc0, 0xe2, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x27,
	0x58, 0x41, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x33, 0x58, 0x06, 0x21, 0x20, 0x24, 0xc5, 0xc5, 0x01,
	0xe6, 0xe4, 0x95, 0xe6, 0x4a, 0xb0, 0x80, 0x25, 0xe1, 0x7c, 0x27, 0x9f, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x87, 0xb8, 0x5a, 0x37, 0x25, 0x31, 0x5f, 0x1f, 0xc5, 0x9f, 0x15, 0x70, 0x9f, 0x82,
	0x2c, 0x2a, 0x4e, 0x62, 0x03, 0xfb, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x2a, 0x7f,
	0x52, 0x0e, 0x01, 0x00, 0x00,
}

func (m *PayBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PayBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Blocknum) > 0 {
		i -= len(m.Blocknum)
		copy(dAtA[i:], m.Blocknum)
		i = encodeVarintPayBlocks(dAtA, i, uint64(len(m.Blocknum)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Blocktype) > 0 {
		i -= len(m.Blocktype)
		copy(dAtA[i:], m.Blocktype)
		i = encodeVarintPayBlocks(dAtA, i, uint64(len(m.Blocktype)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintPayBlocks(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blockid) > 0 {
		i -= len(m.Blockid)
		copy(dAtA[i:], m.Blockid)
		i = encodeVarintPayBlocks(dAtA, i, uint64(len(m.Blockid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPayBlocks(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayBlocks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PayBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Blockid)
	if l > 0 {
		n += 1 + l + sovPayBlocks(uint64(l))
	}
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + sovPayBlocks(uint64(l))
	}
	l = len(m.Blocktype)
	if l > 0 {
		n += 1 + l + sovPayBlocks(uint64(l))
	}
	l = len(m.Blocknum)
	if l > 0 {
		n += 1 + l + sovPayBlocks(uint64(l))
	}
	return n
}

func sovPayBlocks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayBlocks(x uint64) (n int) {
	return sovPayBlocks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PayBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayBlocks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PayBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayBlocks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayBlocks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blockid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayBlocks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayBlocks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocktype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayBlocks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayBlocks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocktype = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocknum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayBlocks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayBlocks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocknum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayBlocks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayBlocks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPayBlocks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayBlocks
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayBlocks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPayBlocks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayBlocks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayBlocks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayBlocks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayBlocks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayBlocks = fmt.Errorf("proto: unexpected end of group")
)
