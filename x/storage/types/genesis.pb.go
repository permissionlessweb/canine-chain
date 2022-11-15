// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine-chain/storage/genesis.proto

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

// GenesisState defines the storage module's genesis state.
type GenesisState struct {
	Params          Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ContractsList   []Contracts   `protobuf:"bytes,2,rep,name=contractsList,proto3" json:"contractsList"`
	ActiveDealsList []ActiveDeals `protobuf:"bytes,4,rep,name=activeDealsList,proto3" json:"activeDealsList"`
	ProvidersList   []Providers   `protobuf:"bytes,5,rep,name=providersList,proto3" json:"providersList"`
	PayBlocksList   []PayBlocks   `protobuf:"bytes,6,rep,name=payBlocksList,proto3" json:"payBlocksList"`
	ClientUsageList []ClientUsage `protobuf:"bytes,7,rep,name=clientUsageList,proto3" json:"clientUsageList"`
	StraysList      []Strays      `protobuf:"bytes,8,rep,name=straysList,proto3" json:"straysList"`
	FidCidList      []FidCid      `protobuf:"bytes,3,rep,name=fidCidList,proto3" json:"fidCidList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_529bdbc9d39fd888, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetContractsList() []Contracts {
	if m != nil {
		return m.ContractsList
	}
	return nil
}

func (m *GenesisState) GetActiveDealsList() []ActiveDeals {
	if m != nil {
		return m.ActiveDealsList
	}
	return nil
}

func (m *GenesisState) GetProvidersList() []Providers {
	if m != nil {
		return m.ProvidersList
	}
	return nil
}

func (m *GenesisState) GetPayBlocksList() []PayBlocks {
	if m != nil {
		return m.PayBlocksList
	}
	return nil
}

func (m *GenesisState) GetClientUsageList() []ClientUsage {
	if m != nil {
		return m.ClientUsageList
	}
	return nil
}

func (m *GenesisState) GetStraysList() []Strays {
	if m != nil {
		return m.StraysList
	}
	return nil
}

func (m *GenesisState) GetFidCidList() []FidCid {
	if m != nil {
		return m.FidCidList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "jackaldao.canine.storage.GenesisState")
}

func init() {
	proto.RegisterFile("canine-chain/storage/genesis.proto", fileDescriptor_529bdbc9d39fd888)
}

var fileDescriptor_529bdbc9d39fd888 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0x86, 0x5b, 0x3f, 0x44, 0xd3, 0x4f, 0x63, 0xd2, 0xb8, 0x20, 0x2c, 0x2a, 0xa2, 0x44, 0x16,
	0xd2, 0x26, 0xb8, 0x37, 0x11, 0x0c, 0x6e, 0x4c, 0x24, 0x12, 0x36, 0x6e, 0x9a, 0xc3, 0xcc, 0x50,
	0x46, 0x4a, 0xa7, 0x69, 0x07, 0x62, 0xef, 0xc2, 0xcb, 0x62, 0xc9, 0xd2, 0x95, 0x31, 0x70, 0x21,
	0x7e, 0xe9, 0x4c, 0xa7, 0xfc, 0xa4, 0x3f, 0xbb, 0x2e, 0x9e, 0xf7, 0xe9, 0x39, 0xe7, 0x1d, 0xa3,
	0x8b, 0x20, 0xa0, 0x01, 0x19, 0xa0, 0x15, 0xd0, 0xc0, 0x89, 0x39, 0x8b, 0xc0, 0x23, 0x8e, 0x47,
	0x02, 0x12, 0xd3, 0xd8, 0x0e, 0x23, 0xc6, 0x99, 0xd9, 0xfa, 0x09, 0x68, 0x0d, 0x3e, 0x06, 0x66,
	0x4b, 0xda, 0xce, 0xb8, 0xf6, 0x4b, 0x8f, 0x79, 0x4c, 0x40, 0x4e, 0xfa, 0x25, 0xf9, 0xf6, 0xeb,
	0x42, 0x67, 0x08, 0x11, 0x6c, 0x32, 0x65, 0xfb, 0x6d, 0x21, 0x82, 0x58, 0xc0, 0x23, 0x40, 0x5c,
	0x51, 0xef, 0x0a, 0x29, 0x40, 0x9c, 0xee, 0x88, 0x8b, 0x09, 0xf8, 0xd5, 0xba, 0x30, 0x62, 0x3b,
	0x8a, 0x49, 0xa4, 0xa8, 0x5e, 0xc9, 0x5c, 0x89, 0xbb, 0xf0, 0x19, 0x5a, 0x57, 0xff, 0x15, 0xf9,
	0x94, 0x04, 0xdc, 0xdd, 0xc6, 0xe0, 0x91, 0xca, 0x3d, 0x63, 0x1e, 0x41, 0xa2, 0x5c, 0xc5, 0xe7,
	0x5d, 0x52, 0xec, 0x22, 0x8a, 0x25, 0xd3, 0xfd, 0xdf, 0x30, 0x9e, 0x7d, 0x91, 0x07, 0x9f, 0x71,
	0xe0, 0xc4, 0xfc, 0x68, 0x34, 0xe5, 0xb1, 0x5a, 0x7a, 0x47, 0xef, 0xdf, 0x0f, 0x3b, 0x76, 0x59,
	0x01, 0xf6, 0x54, 0x70, 0xa3, 0xc6, 0xfe, 0xef, 0x2b, 0xed, 0x7b, 0x96, 0x32, 0xbf, 0x19, 0xcf,
	0xf3, 0x4b, 0x7e, 0xa5, 0x31, 0x6f, 0x3d, 0xea, 0xdc, 0xf5, 0xef, 0x87, 0x6f, 0xca, 0x35, 0x63,
	0x85, 0x67, 0xa6, 0xeb, 0xbc, 0x39, 0x37, 0x5e, 0xc8, 0xa3, 0x7f, 0x4e, 0x6f, 0x2e, 0x94, 0x0d,
	0xa1, 0xec, 0x95, 0x2b, 0x3f, 0x9d, 0x03, 0x99, 0xf4, 0xd6, 0x91, 0xce, 0x99, 0x57, 0x24, 0xa4,
	0x8f, 0xeb, 0xe6, 0x9c, 0x2a, 0x5c, 0xcd, 0x79, 0x95, 0x17, 0x42, 0x48, 0x46, 0xa2, 0x4c, 0x21,
	0x6c, 0xd6, 0x0a, 0x15, 0x9e, 0x0b, 0x2f, 0xf3, 0xe9, 0xe2, 0xb2, 0xf7, 0x79, 0x5a, 0xbb, 0x50,
	0x3e, 0xa9, 0x5b, 0x7c, 0x7c, 0x0e, 0xa8, 0xc5, 0x6f, 0x1c, 0xe6, 0xc4, 0x30, 0xe4, 0x2b, 0x11,
	0xc6, 0xa7, 0xc2, 0x58, 0x51, 0xf2, 0x4c, 0xb0, 0x99, 0xec, 0x22, 0x99, 0x7a, 0x96, 0x14, 0x8f,
	0x29, 0x16, 0x9e, 0xbb, 0x3a, 0xcf, 0x44, 0xb0, 0xca, 0x73, 0x4e, 0x8e, 0x26, 0xfb, 0xa3, 0xa5,
	0x1f, 0x8e, 0x96, 0xfe, 0xef, 0x68, 0xe9, 0xbf, 0x4f, 0x96, 0x76, 0x38, 0x59, 0xda, 0x9f, 0x93,
	0xa5, 0xfd, 0x78, 0xef, 0x51, 0xbe, 0xda, 0x2e, 0x6c, 0xc4, 0x36, 0x8e, 0xf4, 0x0e, 0x30, 0x30,
	0x47, 0x8a, 0x9d, 0x5f, 0xf9, 0x8b, 0xe6, 0x49, 0x48, 0xe2, 0x45, 0x53, 0x3c, 0xe8, 0x0f, 0x0f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0x67, 0xa8, 0x9b, 0x55, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StraysList) > 0 {
		for iNdEx := len(m.StraysList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StraysList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ClientUsageList) > 0 {
		for iNdEx := len(m.ClientUsageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientUsageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PayBlocksList) > 0 {
		for iNdEx := len(m.PayBlocksList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PayBlocksList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ProvidersList) > 0 {
		for iNdEx := len(m.ProvidersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProvidersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ActiveDealsList) > 0 {
		for iNdEx := len(m.ActiveDealsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActiveDealsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.FidCidList) > 0 {
		for iNdEx := len(m.FidCidList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FidCidList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContractsList) > 0 {
		for iNdEx := len(m.ContractsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ContractsList) > 0 {
		for _, e := range m.ContractsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FidCidList) > 0 {
		for _, e := range m.FidCidList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActiveDealsList) > 0 {
		for _, e := range m.ActiveDealsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProvidersList) > 0 {
		for _, e := range m.ProvidersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PayBlocksList) > 0 {
		for _, e := range m.PayBlocksList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientUsageList) > 0 {
		for _, e := range m.ClientUsageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StraysList) > 0 {
		for _, e := range m.StraysList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractsList = append(m.ContractsList, Contracts{})
			if err := m.ContractsList[len(m.ContractsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FidCidList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FidCidList = append(m.FidCidList, FidCid{})
			if err := m.FidCidList[len(m.FidCidList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDealsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveDealsList = append(m.ActiveDealsList, ActiveDeals{})
			if err := m.ActiveDealsList[len(m.ActiveDealsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidersList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvidersList = append(m.ProvidersList, Providers{})
			if err := m.ProvidersList[len(m.ProvidersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayBlocksList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayBlocksList = append(m.PayBlocksList, PayBlocks{})
			if err := m.PayBlocksList[len(m.PayBlocksList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientUsageList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientUsageList = append(m.ClientUsageList, ClientUsage{})
			if err := m.ClientUsageList[len(m.ClientUsageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StraysList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StraysList = append(m.StraysList, Strays{})
			if err := m.StraysList[len(m.StraysList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
