// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/storage/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	DepositAccount string `protobuf:"bytes,1,opt,name=deposit_account,json=depositAccount,proto3" json:"deposit_account,omitempty"`
	ProofWindow    int64  `protobuf:"varint,2,opt,name=proof_window,json=proofWindow,proto3" json:"proof_window,omitempty"`
	// Chunk size of a file is divided into
	// The value cannot be smaller than 1 to avoid zero division
	ChunkSize    int64  `protobuf:"varint,3,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	MissesToBurn int64  `protobuf:"varint,4,opt,name=misses_to_burn,json=missesToBurn,proto3" json:"misses_to_burn,omitempty"`
	PriceFeed    string `protobuf:"bytes,5,opt,name=price_feed,json=priceFeed,proto3" json:"price_feed,omitempty"`
	// Life span of a contract in blocks
	MaxContractAgeInBlocks int64 `protobuf:"varint,6,opt,name=max_contract_age_in_blocks,json=maxContractAgeInBlocks,proto3" json:"max_contract_age_in_blocks,omitempty"`
	PricePerTbPerMonth     int64 `protobuf:"varint,7,opt,name=price_per_tb_per_month,json=pricePerTbPerMonth,proto3" json:"price_per_tb_per_month,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6380cb4192ac15, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDepositAccount() string {
	if m != nil {
		return m.DepositAccount
	}
	return ""
}

func (m *Params) GetProofWindow() int64 {
	if m != nil {
		return m.ProofWindow
	}
	return 0
}

func (m *Params) GetChunkSize() int64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *Params) GetMissesToBurn() int64 {
	if m != nil {
		return m.MissesToBurn
	}
	return 0
}

func (m *Params) GetPriceFeed() string {
	if m != nil {
		return m.PriceFeed
	}
	return ""
}

func (m *Params) GetMaxContractAgeInBlocks() int64 {
	if m != nil {
		return m.MaxContractAgeInBlocks
	}
	return 0
}

func (m *Params) GetPricePerTbPerMonth() int64 {
	if m != nil {
		return m.PricePerTbPerMonth
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "canine_chain.storage.Params")
}

func init() { proto.RegisterFile("canine_chain/storage/params.proto", fileDescriptor_9a6380cb4192ac15) }

var fileDescriptor_9a6380cb4192ac15 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0xaa, 0xda, 0x40,
	0x14, 0x86, 0x13, 0xb5, 0x16, 0xa7, 0x62, 0x61, 0x10, 0x09, 0x42, 0x53, 0x2d, 0x85, 0xba, 0xa9,
	0x01, 0xbb, 0x73, 0xa7, 0x85, 0x42, 0xc1, 0x82, 0x58, 0xa1, 0xd0, 0xcd, 0x30, 0x99, 0x1c, 0x93,
	0xa9, 0x66, 0x4e, 0x98, 0x99, 0xa0, 0xf5, 0x29, 0xba, 0xec, 0xb2, 0x4f, 0xd0, 0xe7, 0xb8, 0x4b,
	0x97, 0x77, 0x79, 0xd1, 0x17, 0xb9, 0x38, 0x91, 0xcb, 0x5d, 0x25, 0x7c, 0xff, 0x77, 0x7e, 0x06,
	0x7e, 0x32, 0x14, 0x5c, 0x49, 0x05, 0x4c, 0x64, 0x5c, 0xaa, 0xc8, 0x58, 0xd4, 0x3c, 0x85, 0xa8,
	0xe0, 0x9a, 0xe7, 0x66, 0x5c, 0x68, 0xb4, 0x48, 0xbb, 0xcf, 0x95, 0xf1, 0x4d, 0xe9, 0x77, 0x53,
	0x4c, 0xd1, 0x09, 0xd1, 0xf5, 0xaf, 0x72, 0xdf, 0xfd, 0xaf, 0x91, 0xe6, 0xd2, 0x1d, 0xd3, 0x0f,
	0xe4, 0x75, 0x02, 0x05, 0x1a, 0x69, 0x19, 0x17, 0x02, 0x4b, 0x65, 0x03, 0x7f, 0xe0, 0x8f, 0x5a,
	0xab, 0xce, 0x0d, 0xcf, 0x2a, 0x4a, 0x87, 0xa4, 0x5d, 0x68, 0xc4, 0x0d, 0xdb, 0x4b, 0x95, 0xe0,
	0x3e, 0xa8, 0x0d, 0xfc, 0x51, 0x7d, 0xf5, 0xca, 0xb1, 0x1f, 0x0e, 0xd1, 0x37, 0x84, 0x88, 0xac,
	0x54, 0x5b, 0x66, 0xe4, 0x11, 0x82, 0xba, 0x13, 0x5a, 0x8e, 0x7c, 0x97, 0x47, 0xa0, 0xef, 0x49,
	0x27, 0x97, 0xc6, 0x80, 0x61, 0x16, 0x59, 0x5c, 0x6a, 0x15, 0x34, 0x9c, 0xd2, 0xae, 0xe8, 0x1a,
	0xe7, 0xa5, 0x56, 0xd7, 0x92, 0x42, 0x4b, 0x01, 0x6c, 0x03, 0x90, 0x04, 0x2f, 0xdc, 0x5b, 0x5a,
	0x8e, 0x7c, 0x01, 0x48, 0xe8, 0x94, 0xf4, 0x73, 0x7e, 0x60, 0x02, 0x95, 0xd5, 0x5c, 0x58, 0xc6,
	0x53, 0x60, 0x52, 0xb1, 0x78, 0x87, 0x62, 0x6b, 0x82, 0xa6, 0x2b, 0xec, 0xe5, 0xfc, 0xf0, 0xf9,
	0x26, 0xcc, 0x52, 0xf8, 0xaa, 0xe6, 0x2e, 0xa5, 0x13, 0xd2, 0xab, 0xaa, 0x0b, 0xd0, 0xcc, 0xc6,
	0xee, 0x93, 0xa3, 0xb2, 0x59, 0xf0, 0xd2, 0xdd, 0x51, 0x97, 0x2e, 0x41, 0xaf, 0xe3, 0x25, 0xe8,
	0x6f, 0xd7, 0x64, 0xda, 0xf8, 0xfb, 0xef, 0xad, 0x37, 0x5f, 0xdc, 0x9d, 0x43, 0xff, 0x74, 0x0e,
	0xfd, 0x87, 0x73, 0xe8, 0xff, 0xb9, 0x84, 0xde, 0xe9, 0x12, 0x7a, 0xf7, 0x97, 0xd0, 0xfb, 0x39,
	0x49, 0xa5, 0xcd, 0xca, 0x78, 0x2c, 0x30, 0x8f, 0x7e, 0x71, 0xb1, 0xe5, 0xbb, 0x05, 0x8f, 0x4d,
	0x54, 0x8d, 0xf1, 0xb1, 0xda, 0xeb, 0xf0, 0xb4, 0x98, 0xfd, 0x5d, 0x80, 0x89, 0x9b, 0x6e, 0x85,
	0x4f, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5c, 0x15, 0x9f, 0xd6, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PricePerTbPerMonth != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PricePerTbPerMonth))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxContractAgeInBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxContractAgeInBlocks))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PriceFeed) > 0 {
		i -= len(m.PriceFeed)
		copy(dAtA[i:], m.PriceFeed)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceFeed)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MissesToBurn != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MissesToBurn))
		i--
		dAtA[i] = 0x20
	}
	if m.ChunkSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChunkSize))
		i--
		dAtA[i] = 0x18
	}
	if m.ProofWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ProofWindow))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DepositAccount) > 0 {
		i -= len(m.DepositAccount)
		copy(dAtA[i:], m.DepositAccount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositAccount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ProofWindow != 0 {
		n += 1 + sovParams(uint64(m.ProofWindow))
	}
	if m.ChunkSize != 0 {
		n += 1 + sovParams(uint64(m.ChunkSize))
	}
	if m.MissesToBurn != 0 {
		n += 1 + sovParams(uint64(m.MissesToBurn))
	}
	l = len(m.PriceFeed)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxContractAgeInBlocks != 0 {
		n += 1 + sovParams(uint64(m.MaxContractAgeInBlocks))
	}
	if m.PricePerTbPerMonth != 0 {
		n += 1 + sovParams(uint64(m.PricePerTbPerMonth))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofWindow", wireType)
			}
			m.ProofWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkSize", wireType)
			}
			m.ChunkSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChunkSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissesToBurn", wireType)
			}
			m.MissesToBurn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissesToBurn |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceFeed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxContractAgeInBlocks", wireType)
			}
			m.MaxContractAgeInBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxContractAgeInBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PricePerTbPerMonth", wireType)
			}
			m.PricePerTbPerMonth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PricePerTbPerMonth |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
