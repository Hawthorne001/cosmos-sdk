// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the protocolpool module's genesis state.
type GenesisState struct {
	// ContinuousFund defines the continuous funds at genesis.
	ContinuousFund []*ContinuousFund `protobuf:"bytes,1,rep,name=continuous_fund,json=continuousFund,proto3" json:"continuous_fund,omitempty"`
	// Budget defines the budget proposals at genesis.
	Budget []*Budget `protobuf:"bytes,2,rep,name=budget,proto3" json:"budget,omitempty"`
	// last_balance contains the amount of tokens yet to be distributed, will be zero if
	// there are no funds to distribute.
	LastBalance DistributionAmount `protobuf:"bytes,3,opt,name=last_balance,json=lastBalance,proto3" json:"last_balance"`
	// distributions contains the list of distributions to be made to continuous
	// funds and budgets. It contains time in order to distribute to non-expired
	// funds only.
	Distributions []*Distribution `protobuf:"bytes,4,rep,name=distributions,proto3" json:"distributions,omitempty"`
	// params defines the parameters of this module, currently only contains the
	// denoms that will be used for continuous fund distributions.
	Params *Params `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_72560a99455b4146, []int{0}
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

func (m *GenesisState) GetContinuousFund() []*ContinuousFund {
	if m != nil {
		return m.ContinuousFund
	}
	return nil
}

func (m *GenesisState) GetBudget() []*Budget {
	if m != nil {
		return m.Budget
	}
	return nil
}

func (m *GenesisState) GetLastBalance() DistributionAmount {
	if m != nil {
		return m.LastBalance
	}
	return DistributionAmount{}
}

func (m *GenesisState) GetDistributions() []*Distribution {
	if m != nil {
		return m.Distributions
	}
	return nil
}

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type Distribution struct {
	// time at which this distribution was made, in order to distribute to non-expired funds only
	// and funds that existed at that time. Because we don't distribute right away, we keep track
	// of the time of distribution.
	Time *time.Time `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time,omitempty"`
	// amount is the list of coins to be distributed.
	Amount DistributionAmount `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_72560a99455b4146, []int{1}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func (m *Distribution) GetTime() *time.Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Distribution) GetAmount() DistributionAmount {
	if m != nil {
		return m.Amount
	}
	return DistributionAmount{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.protocolpool.v1.GenesisState")
	proto.RegisterType((*Distribution)(nil), "cosmos.protocolpool.v1.Distribution")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/genesis.proto", fileDescriptor_72560a99455b4146)
}

var fileDescriptor_72560a99455b4146 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0xeb, 0xd3, 0x30,
	0x18, 0x87, 0x9b, 0xad, 0xee, 0x90, 0x4d, 0x85, 0x22, 0x52, 0x7b, 0xe8, 0xe6, 0x18, 0x32, 0x3c,
	0xa4, 0x6c, 0x8a, 0x17, 0x4f, 0x56, 0x51, 0xf1, 0xa2, 0x74, 0x9e, 0xbc, 0x8c, 0xb4, 0xcd, 0x4a,
	0xb1, 0xcd, 0x5b, 0x96, 0x64, 0xe8, 0x97, 0x90, 0x7d, 0xac, 0x1d, 0x77, 0xf4, 0xa4, 0xb2, 0x7d,
	0x01, 0x3f, 0x82, 0x34, 0xe9, 0xa4, 0x03, 0x0b, 0xf2, 0xbf, 0xa5, 0x6f, 0x9f, 0xdf, 0xc3, 0xfb,
	0xe6, 0x0d, 0x9e, 0x25, 0x20, 0x4a, 0x10, 0x41, 0xb5, 0x05, 0x09, 0x09, 0x14, 0x15, 0x40, 0x11,
	0xec, 0x16, 0x41, 0xc6, 0x38, 0x13, 0xb9, 0x20, 0xba, 0xee, 0xdc, 0x37, 0x14, 0x69, 0x53, 0x64,
	0xb7, 0xf0, 0xa6, 0x1d, 0x69, 0xf9, 0xb5, 0x62, 0x0d, 0xed, 0xdd, 0xcb, 0x20, 0x03, 0x7d, 0x0c,
	0xea, 0x53, 0x53, 0x7d, 0x60, 0x92, 0x6b, 0xf3, 0xa3, 0xad, 0xf7, 0xc6, 0x19, 0x40, 0x56, 0x30,
	0x23, 0x8d, 0xd5, 0x26, 0x90, 0x79, 0xc9, 0x84, 0xa4, 0x65, 0x65, 0x80, 0xe9, 0xef, 0x1e, 0x1e,
	0xbd, 0x31, 0xfd, 0xad, 0x24, 0x95, 0xcc, 0x79, 0x8f, 0xef, 0x26, 0xc0, 0x65, 0xce, 0x15, 0x28,
	0xb1, 0xde, 0x28, 0x9e, 0xba, 0x68, 0xd2, 0x9f, 0x0f, 0x97, 0x8f, 0xc8, 0xbf, 0x1b, 0x27, 0x2f,
	0xff, 0xe2, 0xaf, 0x15, 0x4f, 0xa3, 0x3b, 0xc9, 0xd5, 0xb7, 0xf3, 0x0c, 0x0f, 0x62, 0x95, 0x66,
	0x4c, 0xba, 0x3d, 0xed, 0xf1, 0xbb, 0x3c, 0xa1, 0xa6, 0xa2, 0x86, 0x76, 0x56, 0x78, 0x54, 0x50,
	0x21, 0xd7, 0x31, 0x2d, 0x28, 0x4f, 0x98, 0xdb, 0x9f, 0xa0, 0xf9, 0x70, 0xf9, 0xb8, 0x2b, 0xfd,
	0x2a, 0x17, 0x72, 0x9b, 0xc7, 0x4a, 0xe6, 0xc0, 0x5f, 0x94, 0xa0, 0xb8, 0x0c, 0xed, 0xc3, 0x8f,
	0xb1, 0x15, 0x0d, 0x6b, 0x4b, 0x68, 0x24, 0xce, 0x3b, 0x7c, 0x3b, 0x6d, 0x81, 0xc2, 0xb5, 0x75,
	0x4f, 0xb3, 0xff, 0xb1, 0x46, 0xd7, 0xd1, 0x7a, 0xb0, 0x8a, 0x6e, 0x69, 0x29, 0xdc, 0x5b, 0xba,
	0xb5, 0xce, 0xc1, 0x3e, 0x68, 0x2a, 0x6a, 0xe8, 0xe9, 0x37, 0x84, 0x47, 0x6d, 0xaf, 0xf3, 0x14,
	0xdb, 0xf5, 0x5a, 0x5c, 0xa4, 0x35, 0x1e, 0x31, 0x3b, 0x23, 0x97, 0x9d, 0x91, 0x8f, 0x97, 0x9d,
	0x85, 0xf6, 0xfe, 0xe7, 0x18, 0x45, 0x9a, 0x76, 0xde, 0xe2, 0x01, 0xd5, 0x73, 0xba, 0xbd, 0x1b,
	0xde, 0x4c, 0x93, 0x0f, 0x9f, 0x1f, 0x4e, 0x3e, 0x3a, 0x9e, 0x7c, 0xf4, 0xeb, 0xe4, 0xa3, 0xfd,
	0xd9, 0xb7, 0x8e, 0x67, 0xdf, 0xfa, 0x7e, 0xf6, 0xad, 0x4f, 0x0f, 0x8d, 0x52, 0xa4, 0x9f, 0x49,
	0x0e, 0xc1, 0x97, 0xeb, 0xb7, 0xa9, 0x1f, 0x66, 0x3c, 0xd0, 0xb5, 0x27, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x1c, 0x62, 0x0f, 0xfd, 0x02, 0x00, 0x00,
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
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Distributions) > 0 {
		for iNdEx := len(m.Distributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Distributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.LastBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Budget) > 0 {
		for iNdEx := len(m.Budget) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Budget[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ContinuousFund) > 0 {
		for iNdEx := len(m.ContinuousFund) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContinuousFund[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Time != nil {
		n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Time):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintGenesis(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.ContinuousFund) > 0 {
		for _, e := range m.ContinuousFund {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Budget) > 0 {
		for _, e := range m.Budget {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.LastBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Distributions) > 0 {
		for _, e := range m.Distributions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Time)
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContinuousFund", wireType)
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
			m.ContinuousFund = append(m.ContinuousFund, &ContinuousFund{})
			if err := m.ContinuousFund[len(m.ContinuousFund)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Budget", wireType)
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
			m.Budget = append(m.Budget, &Budget{})
			if err := m.Budget[len(m.Budget)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBalance", wireType)
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
			if err := m.LastBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
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
			m.Distributions = append(m.Distributions, &Distribution{})
			if err := m.Distributions[len(m.Distributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Distribution) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
			if m.Time == nil {
				m.Time = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
