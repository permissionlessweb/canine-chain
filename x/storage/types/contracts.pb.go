// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/contracts.proto

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

type Contracts struct {
	Cid        string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Priceamt   string `protobuf:"bytes,2,opt,name=priceamt,proto3" json:"priceamt,omitempty"`
	Pricedenom string `protobuf:"bytes,3,opt,name=pricedenom,proto3" json:"pricedenom,omitempty"`
	Merkle     string `protobuf:"bytes,5,opt,name=merkle,proto3" json:"merkle,omitempty"`
	Signee     string `protobuf:"bytes,6,opt,name=signee,proto3" json:"signee,omitempty"`
	Duration   string `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Filesize   string `protobuf:"bytes,8,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Fid        string `protobuf:"bytes,9,opt,name=fid,proto3" json:"fid,omitempty"`
	Creator    string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Contracts) Reset()         { *m = Contracts{} }
func (m *Contracts) String() string { return proto.CompactTextString(m) }
func (*Contracts) ProtoMessage()    {}
func (*Contracts) Descriptor() ([]byte, []int) {
	return fileDescriptor_739ff281a83347de, []int{0}
}
func (m *Contracts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contracts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contracts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contracts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contracts.Merge(m, src)
}
func (m *Contracts) XXX_Size() int {
	return m.Size()
}
func (m *Contracts) XXX_DiscardUnknown() {
	xxx_messageInfo_Contracts.DiscardUnknown(m)
}

var xxx_messageInfo_Contracts proto.InternalMessageInfo

func (m *Contracts) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Contracts) GetPriceamt() string {
	if m != nil {
		return m.Priceamt
	}
	return ""
}

func (m *Contracts) GetPricedenom() string {
	if m != nil {
		return m.Pricedenom
	}
	return ""
}

func (m *Contracts) GetMerkle() string {
	if m != nil {
		return m.Merkle
	}
	return ""
}

func (m *Contracts) GetSignee() string {
	if m != nil {
		return m.Signee
	}
	return ""
}

func (m *Contracts) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Contracts) GetFilesize() string {
	if m != nil {
		return m.Filesize
	}
	return ""
}

func (m *Contracts) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *Contracts) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Contracts)(nil), "jackaldao.canine.storage.Contracts")
}

func init() { proto.RegisterFile("storage/contracts.proto", fileDescriptor_739ff281a83347de) }

var fileDescriptor_739ff281a83347de = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x0a, 0x69, 0xe3, 0x09, 0x79, 0x80, 0x13, 0x83, 0x85, 0x98, 0x18, 0x20, 0x19,
	0x78, 0x03, 0x90, 0x78, 0x00, 0x46, 0x36, 0xd7, 0xbe, 0x04, 0xd3, 0xc4, 0x8e, 0x1c, 0x57, 0x02,
	0x9e, 0x82, 0xc7, 0x62, 0xec, 0xc8, 0x88, 0x92, 0x99, 0x77, 0x40, 0xb1, 0x6b, 0xd4, 0xed, 0xbe,
	0xef, 0x3f, 0xe9, 0x4e, 0x3f, 0x3d, 0x1f, 0xbc, 0x75, 0xa2, 0xc1, 0x4a, 0x5a, 0xe3, 0x9d, 0x90,
	0x7e, 0x28, 0x7b, 0x67, 0xbd, 0x65, 0xf0, 0x2a, 0xe4, 0x46, 0xb4, 0x4a, 0xd8, 0x52, 0x0a, 0xa3,
	0x0d, 0x96, 0xfb, 0xcd, 0xab, 0x5f, 0x42, 0x8b, 0x87, 0xb4, 0xcd, 0x4e, 0xe9, 0x42, 0x6a, 0x05,
	0xe4, 0x92, 0x5c, 0x17, 0x4f, 0xf3, 0xc8, 0x2e, 0xe8, 0xaa, 0x77, 0x5a, 0xa2, 0xe8, 0x3c, 0x1c,
	0x05, 0xfd, 0xcf, 0x8c, 0x53, 0x1a, 0x66, 0x85, 0xc6, 0x76, 0xb0, 0x08, 0xe9, 0x81, 0x61, 0x67,
	0x34, 0xef, 0xd0, 0x6d, 0x5a, 0x84, 0x93, 0x90, 0xed, 0x69, 0xf6, 0x83, 0x6e, 0x0c, 0x22, 0xe4,
	0xd1, 0x47, 0x9a, 0x6f, 0xa9, 0xad, 0x13, 0x5e, 0x5b, 0x03, 0xcb, 0x78, 0x2b, 0xf1, 0x9c, 0xd5,
	0xba, 0xc5, 0x41, 0x7f, 0x20, 0xac, 0x62, 0x96, 0x78, 0xfe, 0xba, 0xd6, 0x0a, 0x8a, 0xf8, 0x75,
	0xad, 0x15, 0x03, 0xba, 0x94, 0x0e, 0x85, 0xb7, 0x0e, 0x8e, 0x83, 0x4d, 0x78, 0xff, 0xf8, 0x35,
	0x72, 0xb2, 0x1b, 0x39, 0xf9, 0x19, 0x39, 0xf9, 0x9c, 0x78, 0xb6, 0x9b, 0x78, 0xf6, 0x3d, 0xf1,
	0xec, 0xf9, 0xa6, 0xd1, 0xfe, 0x65, 0xbb, 0x2e, 0xa5, 0xed, 0xaa, 0x58, 0xd7, 0xad, 0x12, 0xb6,
	0x8a, 0x7d, 0x55, 0x6f, 0x55, 0xea, 0xd6, 0xbf, 0xf7, 0x38, 0xac, 0xf3, 0x50, 0xec, 0xdd, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x0a, 0x99, 0x40, 0x73, 0x01, 0x00, 0x00,
}

func (m *Contracts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contracts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contracts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fid) > 0 {
		i -= len(m.Fid)
		copy(dAtA[i:], m.Fid)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Fid)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Filesize) > 0 {
		i -= len(m.Filesize)
		copy(dAtA[i:], m.Filesize)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Filesize)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Duration) > 0 {
		i -= len(m.Duration)
		copy(dAtA[i:], m.Duration)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Duration)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Signee) > 0 {
		i -= len(m.Signee)
		copy(dAtA[i:], m.Signee)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Signee)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Merkle) > 0 {
		i -= len(m.Merkle)
		copy(dAtA[i:], m.Merkle)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Merkle)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Pricedenom) > 0 {
		i -= len(m.Pricedenom)
		copy(dAtA[i:], m.Pricedenom)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Pricedenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Priceamt) > 0 {
		i -= len(m.Priceamt)
		copy(dAtA[i:], m.Priceamt)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Priceamt)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintContracts(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContracts(dAtA []byte, offset int, v uint64) int {
	offset -= sovContracts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contracts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Priceamt)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Pricedenom)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Merkle)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Signee)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Filesize)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	l = len(m.Fid)
	if l > 0 {
		n += 1 + l + sovContracts(uint64(l))
	}
	return n
}

func sovContracts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContracts(x uint64) (n int) {
	return sovContracts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contracts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContracts
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
			return fmt.Errorf("proto: Contracts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contracts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priceamt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Priceamt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pricedenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pricedenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Merkle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Merkle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filesize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filesize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContracts
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
				return ErrInvalidLengthContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContracts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContracts
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
func skipContracts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContracts
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
					return 0, ErrIntOverflowContracts
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
					return 0, ErrIntOverflowContracts
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
				return 0, ErrInvalidLengthContracts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContracts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContracts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContracts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContracts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContracts = fmt.Errorf("proto: unexpected end of group")
)