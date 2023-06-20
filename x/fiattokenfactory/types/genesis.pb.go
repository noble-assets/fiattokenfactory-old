// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/fiattokenfactory/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the fiattokenfactory module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BlacklistedList      []Blacklisted      `protobuf:"bytes,2,rep,name=blacklisted_list,json=blacklistedList,proto3" json:"blacklisted_list"`
	Paused               *Paused            `protobuf:"bytes,3,opt,name=paused,proto3" json:"paused,omitempty"`
	MasterMinter         *MasterMinter      `protobuf:"bytes,4,opt,name=master_minter,json=masterMinter,proto3" json:"master_minter,omitempty"`
	MintersList          []Minters          `protobuf:"bytes,5,rep,name=minters_list,json=mintersList,proto3" json:"minters_list"`
	Pauser               *Pauser            `protobuf:"bytes,6,opt,name=pauser,proto3" json:"pauser,omitempty"`
	Blacklister          *Blacklister       `protobuf:"bytes,7,opt,name=blacklister,proto3" json:"blacklister,omitempty"`
	Owner                *Owner             `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	MinterControllerList []MinterController `protobuf:"bytes,9,rep,name=minter_controller_list,json=minterControllerList,proto3" json:"minter_controller_list"`
	MintingDenom         *MintingDenom      `protobuf:"bytes,10,opt,name=minting_denom,json=mintingDenom,proto3" json:"minting_denom,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bceddd93438df2a, []int{0}
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

func (m *GenesisState) GetBlacklistedList() []Blacklisted {
	if m != nil {
		return m.BlacklistedList
	}
	return nil
}

func (m *GenesisState) GetPaused() *Paused {
	if m != nil {
		return m.Paused
	}
	return nil
}

func (m *GenesisState) GetMasterMinter() *MasterMinter {
	if m != nil {
		return m.MasterMinter
	}
	return nil
}

func (m *GenesisState) GetMintersList() []Minters {
	if m != nil {
		return m.MintersList
	}
	return nil
}

func (m *GenesisState) GetPauser() *Pauser {
	if m != nil {
		return m.Pauser
	}
	return nil
}

func (m *GenesisState) GetBlacklister() *Blacklister {
	if m != nil {
		return m.Blacklister
	}
	return nil
}

func (m *GenesisState) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GenesisState) GetMinterControllerList() []MinterController {
	if m != nil {
		return m.MinterControllerList
	}
	return nil
}

func (m *GenesisState) GetMintingDenom() *MintingDenom {
	if m != nil {
		return m.MintingDenom
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "circle.fiattokenfactory.v1.GenesisState")
}

func init() {
	proto.RegisterFile("circle/fiattokenfactory/v1/genesis.proto", fileDescriptor_9bceddd93438df2a)
}

var fileDescriptor_9bceddd93438df2a = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0xb6, 0x16, 0x70, 0x8b, 0x40, 0xd6, 0x84, 0xac, 0x1d, 0xc2, 0x18, 0x12, 0xcb,
	0x61, 0x24, 0x5a, 0x39, 0x20, 0x71, 0x42, 0x65, 0x12, 0x42, 0x5a, 0x05, 0x0a, 0x17, 0xc4, 0x81,
	0x28, 0xc9, 0xdc, 0xcc, 0x5a, 0x6c, 0x57, 0xb6, 0x37, 0xd8, 0xb7, 0xe0, 0x63, 0xed, 0xb8, 0x23,
	0x27, 0x84, 0xda, 0xaf, 0xc1, 0x01, 0xf9, 0x0f, 0x59, 0xb6, 0x0a, 0x37, 0xa7, 0xc6, 0xd6, 0xef,
	0x79, 0xfa, 0xbe, 0x8f, 0xdf, 0x17, 0x44, 0x25, 0x11, 0x65, 0x8d, 0x93, 0x19, 0xc9, 0x95, 0xe2,
	0xa7, 0x98, 0xcd, 0xf2, 0x52, 0x71, 0x71, 0x91, 0x9c, 0x1f, 0x24, 0x15, 0x66, 0x58, 0x12, 0x19,
	0xcf, 0x05, 0x57, 0x1c, 0x6e, 0x5b, 0x32, 0xbe, 0x4d, 0xc6, 0xe7, 0x07, 0xdb, 0xfb, 0x1e, 0x97,
	0xa2, 0xce, 0xcb, 0xd3, 0x9a, 0x48, 0x85, 0x8f, 0xad, 0x53, 0x47, 0x5a, 0x38, 0x3a, 0xf6, 0xd0,
	0x34, 0xd7, 0x60, 0x46, 0x09, 0xbb, 0xe6, 0xc7, 0x3e, 0xde, 0x80, 0x59, 0xc9, 0x99, 0x12, 0xbc,
	0xae, 0x1b, 0x4d, 0xb4, 0x56, 0x23, 0xbb, 0x54, 0x43, 0x98, 0x22, 0xac, 0xca, 0x8e, 0x31, 0xe3,
	0xd4, 0xf1, 0xcf, 0x3d, 0x3c, 0xff, 0xc6, 0x9a, 0x0a, 0xf6, 0x3c, 0xdc, 0x3c, 0x17, 0x39, 0x95,
	0x9d, 0xc0, 0x33, 0xd9, 0xa4, 0xbc, 0x16, 0xfc, 0xf7, 0xd7, 0x5b, 0x15, 0xaf, 0xb8, 0xf9, 0x4c,
	0xf4, 0x97, 0xbd, 0xdd, 0xfd, 0xd3, 0x07, 0xa3, 0x77, 0x76, 0x00, 0x3e, 0xa9, 0x5c, 0x61, 0xf8,
	0x06, 0x0c, 0x6c, 0x21, 0x28, 0xd8, 0x09, 0xa2, 0xe1, 0x78, 0x37, 0xfe, 0xff, 0x40, 0xc4, 0x1f,
	0x0d, 0x39, 0xd9, 0xbc, 0xfc, 0xf5, 0xa4, 0x97, 0x3a, 0x1d, 0xfc, 0x0c, 0x1e, 0xb5, 0x86, 0x21,
	0xd3, 0x3f, 0xe8, 0xce, 0xce, 0x46, 0x34, 0x1c, 0xef, 0xf9, 0xbc, 0x26, 0xd7, 0x1a, 0x67, 0xf8,
	0xb0, 0x65, 0x73, 0x44, 0xa4, 0x82, 0xaf, 0x75, 0x6d, 0xba, 0x77, 0xb4, 0xd1, 0xa5, 0x36, 0x4d,
	0xa6, 0x4e, 0x01, 0xa7, 0xe0, 0xc1, 0x8d, 0x31, 0x42, 0x9b, 0xc6, 0x22, 0xf2, 0x59, 0x4c, 0x8d,
	0x60, 0x6a, 0xf8, 0x74, 0x44, 0x5b, 0x27, 0x78, 0x04, 0x46, 0x6e, 0x62, 0x6c, 0x83, 0x7d, 0xd3,
	0xe0, 0x33, 0xaf, 0x9b, 0xe5, 0x5d, 0x73, 0x43, 0x27, 0xbf, 0xd1, 0x98, 0x40, 0x83, 0x8e, 0x8d,
	0x09, 0xd7, 0x98, 0x80, 0xef, 0xc1, 0xb0, 0xb5, 0x4d, 0xe8, 0xae, 0x31, 0xe8, 0x98, 0xb4, 0x48,
	0xdb, 0x5a, 0xf8, 0x0a, 0xf4, 0xcd, 0xb0, 0xa2, 0x7b, 0xc6, 0xe4, 0xa9, 0xcf, 0xe4, 0x83, 0x06,
	0x53, 0xcb, 0xc3, 0x13, 0xf0, 0x78, 0x65, 0xe7, 0x6c, 0x2e, 0xf7, 0x4d, 0x2e, 0xfb, 0xeb, 0x73,
	0x79, 0xdb, 0x08, 0x5d, 0x40, 0x5b, 0xf4, 0xd6, 0xbd, 0x49, 0x4a, 0x3f, 0x63, 0x7b, 0xff, 0x10,
	0xe8, 0xf0, 0x8c, 0x56, 0x70, 0xa8, 0xf9, 0x74, 0x44, 0x5b, 0xa7, 0xc9, 0xd7, 0xcb, 0x45, 0x18,
	0x5c, 0x2d, 0xc2, 0xe0, 0xf7, 0x22, 0x0c, 0x7e, 0x2c, 0xc3, 0xde, 0xd5, 0x32, 0xec, 0xfd, 0x5c,
	0x86, 0xbd, 0x2f, 0x87, 0x15, 0x51, 0x27, 0x67, 0x45, 0x5c, 0x72, 0x9a, 0x58, 0xef, 0x19, 0x61,
	0x09, 0xe3, 0x45, 0x8d, 0x5f, 0xac, 0xec, 0xda, 0xf7, 0xd5, 0xf5, 0x53, 0x17, 0x73, 0x2c, 0x8b,
	0x81, 0xd9, 0xb2, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x2e, 0x29, 0x9d, 0x80, 0x05,
	0x00, 0x00,
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
	if m.MintingDenom != nil {
		{
			size, err := m.MintingDenom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.MinterControllerList) > 0 {
		for iNdEx := len(m.MinterControllerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinterControllerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.Owner != nil {
		{
			size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Blacklister != nil {
		{
			size, err := m.Blacklister.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Pauser != nil {
		{
			size, err := m.Pauser.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.MintersList) > 0 {
		for iNdEx := len(m.MintersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.MasterMinter != nil {
		{
			size, err := m.MasterMinter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Paused != nil {
		{
			size, err := m.Paused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlacklistedList) > 0 {
		for iNdEx := len(m.BlacklistedList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlacklistedList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BlacklistedList) > 0 {
		for _, e := range m.BlacklistedList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Paused != nil {
		l = m.Paused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MasterMinter != nil {
		l = m.MasterMinter.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MintersList) > 0 {
		for _, e := range m.MintersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Pauser != nil {
		l = m.Pauser.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Blacklister != nil {
		l = m.Blacklister.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Owner != nil {
		l = m.Owner.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MinterControllerList) > 0 {
		for _, e := range m.MinterControllerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MintingDenom != nil {
		l = m.MintingDenom.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedList", wireType)
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
			m.BlacklistedList = append(m.BlacklistedList, Blacklisted{})
			if err := m.BlacklistedList[len(m.BlacklistedList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
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
			if m.Paused == nil {
				m.Paused = &Paused{}
			}
			if err := m.Paused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterMinter", wireType)
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
			if m.MasterMinter == nil {
				m.MasterMinter = &MasterMinter{}
			}
			if err := m.MasterMinter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintersList", wireType)
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
			m.MintersList = append(m.MintersList, Minters{})
			if err := m.MintersList[len(m.MintersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pauser", wireType)
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
			if m.Pauser == nil {
				m.Pauser = &Pauser{}
			}
			if err := m.Pauser.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklister", wireType)
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
			if m.Blacklister == nil {
				m.Blacklister = &Blacklister{}
			}
			if err := m.Blacklister.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			if m.Owner == nil {
				m.Owner = &Owner{}
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterControllerList", wireType)
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
			m.MinterControllerList = append(m.MinterControllerList, MinterController{})
			if err := m.MinterControllerList[len(m.MinterControllerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingDenom", wireType)
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
			if m.MintingDenom == nil {
				m.MintingDenom = &MintingDenom{}
			}
			if err := m.MintingDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
