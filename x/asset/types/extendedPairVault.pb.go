// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/extendedPairVault.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type ExtendedPairVault struct {
	Id                  uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppMappingId        uint64                                 `protobuf:"varint,2,opt,name=app_mapping_id,json=appMappingId,proto3" json:"app_mapping_id,omitempty" yaml:"app_mapping_id"`
	PairId              uint64                                 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	LiquidationRatio    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=liquidation_ratio,json=liquidationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_ratio" yaml:"liquidation_ratio"`
	StabilityFee        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=stability_fee,json=stabilityFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stability_fee" yaml:"stability_fee"`
	ClosingFee          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=closing_fee,json=closingFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"closing_fee" yaml:"closing_fee"`
	LiquidationPenalty  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_penalty" yaml:"liquidation_penalty"`
	DrawDownFee         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=draw_down_fee,json=drawDownFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"draw_down_fee" yaml:"draw_down_fee"`
	IsVaultActive       bool                                   `protobuf:"varint,9,opt,name=is_vault_active,json=isVaultActive,proto3" json:"is_vault_active,omitempty" yaml:"active_flag"`
	DebtCeiling         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=debt_ceiling,json=debtCeiling,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_ceiling" yaml:"debt_ceiling"`
	DebtFloor           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=debt_floor,json=debtFloor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_floor" yaml:"debt_floor"`
	IsPsmPair           bool                                   `protobuf:"varint,12,opt,name=is_psm_pair,json=isPsmPair,proto3" json:"is_psm_pair,omitempty" yaml:"is_psm_pair"`
	MinCr               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=min_cr,json=minCr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_cr" yaml:"min_cr"`
	PairName            string                                 `protobuf:"bytes,14,opt,name=pair_name,json=pairName,proto3" json:"pair_name,omitempty" yaml:"pair_name"`
	AssetOutOraclePrice bool                                   `protobuf:"varint,15,opt,name=asset_out_oracle_price,json=assetOutOraclePrice,proto3" json:"asset_out_oracle_price,omitempty" yaml:"asset_out_oracle_price"`
	AssetOutPrice       uint64                                 `protobuf:"varint,16,opt,name=asset_out_price,json=assetOutPrice,proto3" json:"asset_out_price,omitempty" yaml:"asset_out_price"`
	MinUsdValueLeft     uint64                                 `protobuf:"varint,17,opt,name=min_usd_value_left,json=minUsdValueLeft,proto3" json:"min_usd_value_left,omitempty" yaml:"min_usd_value_left"`
}

func (m *ExtendedPairVault) Reset()         { *m = ExtendedPairVault{} }
func (m *ExtendedPairVault) String() string { return proto.CompactTextString(m) }
func (*ExtendedPairVault) ProtoMessage()    {}
func (*ExtendedPairVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_23dd38fcddb231cd, []int{0}
}
func (m *ExtendedPairVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedPairVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedPairVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedPairVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedPairVault.Merge(m, src)
}
func (m *ExtendedPairVault) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedPairVault) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedPairVault.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedPairVault proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtendedPairVault)(nil), "comdex.asset.v1beta1.ExtendedPairVault")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/extendedPairVault.proto", fileDescriptor_23dd38fcddb231cd)
}

var fileDescriptor_23dd38fcddb231cd = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x4e, 0xca, 0xbd, 0xb9, 0xcd, 0xe4, 0xaf, 0x99, 0x86, 0xca, 0xb7, 0xd2, 0xb5, 0xcb, 0x2c,
	0x50, 0x25, 0xb8, 0xb1, 0x2a, 0x24, 0x16, 0x2c, 0xb8, 0x22, 0xfd, 0x91, 0x5a, 0x15, 0x1a, 0x46,
	0x22, 0x0b, 0x36, 0xa3, 0x89, 0x3d, 0x49, 0x87, 0xda, 0x1e, 0xe3, 0x99, 0x24, 0xcd, 0x82, 0x0d,
	0x4f, 0xc0, 0x63, 0xf0, 0x28, 0x5d, 0x76, 0x89, 0x58, 0x58, 0x90, 0xbe, 0x81, 0x9f, 0x00, 0xcd,
	0xd8, 0xa5, 0x4e, 0xcb, 0x26, 0x62, 0x13, 0x9f, 0x9f, 0x6f, 0xbe, 0xef, 0x9c, 0x33, 0xf6, 0x09,
	0xf8, 0xdc, 0x13, 0xa1, 0xcf, 0x6e, 0x5d, 0x2a, 0x25, 0x53, 0xee, 0xfc, 0x68, 0xcc, 0x14, 0x3d,
	0x72, 0xd9, 0xad, 0x62, 0x91, 0xcf, 0xfc, 0x21, 0xe5, 0xc9, 0x88, 0xce, 0x02, 0xd5, 0x8f, 0x13,
	0xa1, 0x04, 0xec, 0xe5, 0xe8, 0xbe, 0x41, 0xf7, 0x0b, 0xf4, 0x7e, 0x6f, 0x2a, 0xa6, 0xc2, 0x00,
	0x5c, 0x6d, 0xe5, 0x58, 0xf4, 0x6b, 0x03, 0x74, 0x4f, 0x9f, 0xf3, 0xc0, 0x36, 0xd8, 0xe2, 0xbe,
	0x55, 0x3d, 0xa8, 0x1e, 0xbe, 0xc2, 0x5b, 0xdc, 0x87, 0x1f, 0x40, 0x9b, 0xc6, 0x31, 0x09, 0x69,
	0x1c, 0xf3, 0x68, 0x4a, 0xb8, 0x6f, 0x6d, 0xe9, 0xdc, 0xe0, 0x6d, 0x96, 0x3a, 0x1f, 0x2f, 0x69,
	0x18, 0x7c, 0x85, 0xd6, 0xf3, 0x08, 0x37, 0x69, 0x1c, 0x7f, 0x9b, 0xfb, 0xe7, 0x3e, 0xfc, 0x0c,
	0xbc, 0x89, 0x29, 0x4f, 0xf4, 0xc9, 0x8f, 0xcc, 0x49, 0x98, 0xa5, 0x4e, 0x3b, 0x3f, 0x59, 0x24,
	0x10, 0xae, 0x69, 0xeb, 0xdc, 0x87, 0x0b, 0xd0, 0x0d, 0xf8, 0xcf, 0x33, 0xee, 0x53, 0xc5, 0x45,
	0x44, 0x12, 0xfd, 0xb0, 0x5e, 0x1d, 0x54, 0x0f, 0xeb, 0x83, 0x8b, 0xbb, 0xd4, 0xa9, 0xfc, 0x99,
	0x3a, 0x9f, 0x4e, 0xb9, 0xba, 0x9e, 0x8d, 0xfb, 0x9e, 0x08, 0x5d, 0x4f, 0xc8, 0x50, 0xc8, 0xe2,
	0xf1, 0x5e, 0xfa, 0x37, 0xae, 0x5a, 0xc6, 0x4c, 0xf6, 0x4f, 0x98, 0x97, 0xa5, 0x8e, 0x95, 0x8b,
	0xbc, 0x20, 0x44, 0x78, 0xa7, 0x14, 0xc3, 0xfa, 0x17, 0xde, 0x80, 0x96, 0x54, 0x74, 0xcc, 0x03,
	0xae, 0x96, 0x64, 0xc2, 0x98, 0xf5, 0xda, 0x88, 0x9e, 0x6d, 0x2c, 0xda, 0xcb, 0x45, 0xd7, 0xc8,
	0x10, 0x6e, 0xfe, 0xeb, 0x9f, 0x31, 0x06, 0x19, 0x68, 0x78, 0x81, 0x90, 0x7a, 0x5e, 0x5a, 0xaa,
	0x66, 0xa4, 0x4e, 0x36, 0x96, 0x82, 0xb9, 0x54, 0x89, 0x0a, 0x61, 0x50, 0x78, 0x5a, 0xe6, 0x17,
	0xb0, 0x5b, 0xee, 0x3d, 0x66, 0x11, 0x0d, 0xd4, 0xd2, 0x7a, 0x63, 0xe4, 0x2e, 0x37, 0x96, 0xdb,
	0x7f, 0x39, 0xce, 0x82, 0x12, 0x61, 0x58, 0x8a, 0x0e, 0xf3, 0x20, 0xfc, 0x09, 0xb4, 0xfc, 0x84,
	0x2e, 0x88, 0x2f, 0x16, 0x91, 0xe9, 0x73, 0xfb, 0xff, 0x8d, 0x74, 0x8d, 0x0c, 0xe1, 0x86, 0xf6,
	0x4f, 0xc4, 0x22, 0xd2, 0xad, 0x7e, 0x0d, 0x3a, 0x5c, 0x92, 0xb9, 0x7e, 0x83, 0x09, 0xf5, 0x14,
	0x9f, 0x33, 0xab, 0x7e, 0x50, 0x3d, 0xdc, 0x1e, 0xec, 0x3d, 0xcd, 0x29, 0x8f, 0x93, 0x49, 0x40,
	0xa7, 0x08, 0xb7, 0xb8, 0x34, 0xef, 0xfb, 0x37, 0x26, 0x08, 0xaf, 0x41, 0xd3, 0x67, 0x63, 0x45,
	0x3c, 0xc6, 0x03, 0x1e, 0x4d, 0x2d, 0x60, 0x4a, 0x3d, 0xdd, 0xa0, 0xd4, 0xf3, 0x48, 0x65, 0xa9,
	0xb3, 0x5b, 0x94, 0x5a, 0xe2, 0xd2, 0x95, 0xb2, 0xb1, 0x3a, 0xce, 0x3d, 0x38, 0x06, 0xc0, 0x64,
	0x27, 0x81, 0x10, 0x89, 0xd5, 0x30, 0x3a, 0xc7, 0x1b, 0xeb, 0x74, 0x4b, 0x3a, 0x86, 0x09, 0xe1,
	0xba, 0x76, 0xce, 0xb4, 0x0d, 0xbf, 0x04, 0x0d, 0x2e, 0x49, 0x2c, 0x43, 0xa2, 0x3f, 0x2b, 0xab,
	0xf9, 0x7c, 0x12, 0xa5, 0x24, 0xc2, 0x75, 0x2e, 0x87, 0x32, 0xd4, 0x0b, 0x00, 0x8e, 0x40, 0x2d,
	0xe4, 0x11, 0xf1, 0x12, 0xab, 0x65, 0xea, 0xfa, 0xb0, 0xf1, 0x55, 0xb5, 0x72, 0x81, 0x9c, 0x05,
	0xe1, 0xd7, 0x21, 0x8f, 0x8e, 0x13, 0x78, 0x04, 0xea, 0xe6, 0x4b, 0x8f, 0x68, 0xc8, 0xac, 0xb6,
	0xa1, 0xee, 0x65, 0xa9, 0xb3, 0x53, 0x5a, 0x02, 0x3a, 0x85, 0xf0, 0xb6, 0xb6, 0xbf, 0xa3, 0x21,
	0x83, 0x23, 0xb0, 0x67, 0x76, 0x18, 0x11, 0x33, 0x45, 0x44, 0x42, 0xbd, 0x80, 0x91, 0x38, 0xe1,
	0x1e, 0xb3, 0x3a, 0xa6, 0x9b, 0x4f, 0xb2, 0xd4, 0x79, 0x57, 0xdc, 0xeb, 0x7f, 0xe2, 0x10, 0xde,
	0x35, 0x89, 0xab, 0x99, 0xba, 0x32, 0xe1, 0xa1, 0x8e, 0xc2, 0x01, 0xe8, 0x3c, 0xe1, 0x73, 0xc2,
	0x1d, 0xb3, 0x95, 0xf6, 0xb3, 0xd4, 0xd9, 0x7b, 0x4e, 0x58, 0x30, 0xb5, 0x1e, 0x99, 0x72, 0x8e,
	0x0b, 0x00, 0x75, 0x83, 0x33, 0xe9, 0x93, 0x39, 0x0d, 0x66, 0x8c, 0x04, 0x6c, 0xa2, 0xac, 0xae,
	0xa1, 0x79, 0x97, 0xa5, 0xce, 0xdb, 0xa7, 0x21, 0xac, 0x63, 0x10, 0xee, 0x84, 0x3c, 0xfa, 0x41,
	0xfa, 0x23, 0x1d, 0xba, 0x64, 0x13, 0x35, 0xf8, 0xfe, 0xee, 0x6f, 0xbb, 0xf2, 0xfb, 0xca, 0xae,
	0xdc, 0xad, 0xec, 0xea, 0xfd, 0xca, 0xae, 0xfe, 0xb5, 0xb2, 0xab, 0xbf, 0x3d, 0xd8, 0x95, 0xfb,
	0x07, 0xbb, 0xf2, 0xc7, 0x83, 0x5d, 0xf9, 0xd1, 0x5d, 0x1b, 0xbe, 0xde, 0xee, 0xef, 0xc5, 0x64,
	0xc2, 0x3d, 0x4e, 0x83, 0xc2, 0x77, 0x1f, 0xff, 0x1d, 0xcc, 0x4d, 0x8c, 0x6b, 0x66, 0xbd, 0x7f,
	0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd6, 0x7e, 0x92, 0x3a, 0x06, 0x00, 0x00,
}

func (m *ExtendedPairVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedPairVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedPairVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinUsdValueLeft != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.MinUsdValueLeft))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.AssetOutPrice != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AssetOutPrice))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.AssetOutOraclePrice {
		i--
		if m.AssetOutOraclePrice {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.PairName) > 0 {
		i -= len(m.PairName)
		copy(dAtA[i:], m.PairName)
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(len(m.PairName)))
		i--
		dAtA[i] = 0x72
	}
	{
		size := m.MinCr.Size()
		i -= size
		if _, err := m.MinCr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.IsPsmPair {
		i--
		if m.IsPsmPair {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.DebtFloor.Size()
		i -= size
		if _, err := m.DebtFloor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.DebtCeiling.Size()
		i -= size
		if _, err := m.DebtCeiling.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.IsVaultActive {
		i--
		if m.IsVaultActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.DrawDownFee.Size()
		i -= size
		if _, err := m.DrawDownFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ClosingFee.Size()
		i -= size
		if _, err := m.ClosingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.StabilityFee.Size()
		i -= size
		if _, err := m.StabilityFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LiquidationRatio.Size()
		i -= size
		if _, err := m.LiquidationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PairId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.AppMappingId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AppMappingId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtendedPairVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtendedPairVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtendedPairVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.Id))
	}
	if m.AppMappingId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.AppMappingId))
	}
	if m.PairId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.PairId))
	}
	l = m.LiquidationRatio.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.StabilityFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.ClosingFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DrawDownFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsVaultActive {
		n += 2
	}
	l = m.DebtCeiling.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DebtFloor.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsPsmPair {
		n += 2
	}
	l = m.MinCr.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = len(m.PairName)
	if l > 0 {
		n += 1 + l + sovExtendedPairVault(uint64(l))
	}
	if m.AssetOutOraclePrice {
		n += 2
	}
	if m.AssetOutPrice != 0 {
		n += 2 + sovExtendedPairVault(uint64(m.AssetOutPrice))
	}
	if m.MinUsdValueLeft != 0 {
		n += 2 + sovExtendedPairVault(uint64(m.MinUsdValueLeft))
	}
	return n
}

func sovExtendedPairVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtendedPairVault(x uint64) (n int) {
	return sovExtendedPairVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtendedPairVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedPairVault
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
			return fmt.Errorf("proto: ExtendedPairVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedPairVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppMappingId", wireType)
			}
			m.AppMappingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppMappingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StabilityFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StabilityFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawDownFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DrawDownFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVaultActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsVaultActive = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtCeiling", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtCeiling.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtFloor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtFloor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPsmPair", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPsmPair = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutOraclePrices", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AssetOutOraclePrice = bool(v != 0)
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutPrices", wireType)
			}
			m.AssetOutPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOutPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUsdValueLeft", wireType)
			}
			m.MinUsdValueLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUsdValueLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedPairVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedPairVault
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
func skipExtendedPairVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
				return 0, ErrInvalidLengthExtendedPairVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtendedPairVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtendedPairVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtendedPairVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtendedPairVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtendedPairVault = fmt.Errorf("proto: unexpected end of group")
)
