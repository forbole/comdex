// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/bandoracle/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	Creator        string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OracleScriptId string                                   `protobuf:"bytes,2,opt,name=oracle_script_id,json=oracleScriptId,proto3" json:"oracle_script_id,omitempty"`
	SourceChannel  string                                   `protobuf:"bytes,3,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	AskCount       string                                   `protobuf:"bytes,5,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	MinCount       string                                   `protobuf:"bytes,6,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	FeeLimit       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
	PrepareGas     string                                   `protobuf:"bytes,9,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty"`
	ExecuteGas     string                                   `protobuf:"bytes,10,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_51400c598ccb3e01, []int{0}
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

func (m *Params) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Params) GetOracleScriptId() string {
	if m != nil {
		return m.OracleScriptId
	}
	return ""
}

func (m *Params) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *Params) GetAskCount() string {
	if m != nil {
		return m.AskCount
	}
	return ""
}

func (m *Params) GetMinCount() string {
	if m != nil {
		return m.MinCount
	}
	return ""
}

func (m *Params) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func (m *Params) GetPrepareGas() string {
	if m != nil {
		return m.PrepareGas
	}
	return ""
}

func (m *Params) GetExecuteGas() string {
	if m != nil {
		return m.ExecuteGas
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "comdex.bandoracle.v1beta1.Params")
}

func init() {
	proto.RegisterFile("comdex/bandoracle/v1beta1/params.proto", fileDescriptor_51400c598ccb3e01)
}

var fileDescriptor_51400c598ccb3e01 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x6e, 0xda, 0x30,
	0x1c, 0xc7, 0x13, 0xd0, 0x80, 0x18, 0x0d, 0x4d, 0xd1, 0x0e, 0x81, 0x49, 0x01, 0x4d, 0xda, 0x94,
	0x0b, 0xf1, 0xd8, 0xf6, 0x04, 0x70, 0x40, 0x95, 0xaa, 0xaa, 0xa2, 0xb7, 0x5e, 0x22, 0xc7, 0x71,
	0x82, 0x45, 0x62, 0x47, 0xb6, 0x53, 0xd1, 0xb7, 0xe8, 0x73, 0xf4, 0x49, 0x38, 0x72, 0x6c, 0x2f,
	0x6d, 0x05, 0x2f, 0x52, 0xc5, 0x0e, 0x42, 0x9c, 0x92, 0xdf, 0xf7, 0xfb, 0xb1, 0x7f, 0xff, 0x0c,
	0x7e, 0x63, 0x5e, 0x24, 0x64, 0x0b, 0x63, 0xc4, 0x12, 0x2e, 0x10, 0xce, 0x09, 0x7c, 0x98, 0xc5,
	0x44, 0xa1, 0x19, 0x2c, 0x91, 0x40, 0x85, 0x0c, 0x4b, 0xc1, 0x15, 0x77, 0x87, 0x86, 0x0b, 0xcf,
	0x5c, 0xd8, 0x70, 0xa3, 0xef, 0x19, 0xcf, 0xb8, 0xa6, 0x60, 0xfd, 0x67, 0x0e, 0x8c, 0x7c, 0xcc,
	0x65, 0xc1, 0x25, 0x8c, 0x91, 0x3c, 0x5f, 0x89, 0x39, 0x65, 0xc6, 0xff, 0xf9, 0xda, 0x02, 0x9d,
	0x5b, 0x9d, 0xc1, 0xf5, 0x40, 0x17, 0x0b, 0x82, 0x14, 0x17, 0x9e, 0x3d, 0xb1, 0x03, 0x67, 0x75,
	0x0a, 0xdd, 0x00, 0x7c, 0x33, 0xc9, 0x22, 0x89, 0x05, 0x2d, 0x55, 0x44, 0x13, 0xaf, 0xa5, 0x91,
	0x81, 0xd1, 0xef, 0xb4, 0x7c, 0x95, 0xb8, 0xbf, 0xc0, 0x40, 0xf2, 0x4a, 0x60, 0x12, 0xe1, 0x35,
	0x62, 0x8c, 0xe4, 0x5e, 0x5b, 0x73, 0x5f, 0x8d, 0xba, 0x30, 0xa2, 0xfb, 0x03, 0x38, 0x48, 0x6e,
	0x22, 0xcc, 0x2b, 0xa6, 0xbc, 0x2f, 0x9a, 0xe8, 0x21, 0xb9, 0x59, 0xd4, 0x71, 0x6d, 0x16, 0x94,
	0x35, 0x66, 0xc7, 0x98, 0x05, 0x65, 0xc6, 0x5c, 0x03, 0x27, 0x25, 0x24, 0xca, 0x69, 0x41, 0x95,
	0xd7, 0x9d, 0xb4, 0x83, 0xfe, 0xdf, 0x61, 0x68, 0x7a, 0x0c, 0xeb, 0x1e, 0x4f, 0xe3, 0x08, 0x17,
	0x9c, 0xb2, 0xf9, 0x9f, 0xdd, 0xdb, 0xd8, 0x7a, 0x7e, 0x1f, 0x07, 0x19, 0x55, 0xeb, 0x2a, 0x0e,
	0x31, 0x2f, 0x60, 0x33, 0x10, 0xf3, 0x99, 0xca, 0x64, 0x03, 0xd5, 0x63, 0x49, 0xa4, 0x3e, 0x20,
	0x57, 0xbd, 0x94, 0x90, 0xeb, 0xfa, 0x72, 0x77, 0x0c, 0xfa, 0xa5, 0x20, 0x25, 0x12, 0x24, 0xca,
	0x90, 0xf4, 0x1c, 0x5d, 0x08, 0x68, 0xa4, 0x25, 0x92, 0x35, 0x40, 0xb6, 0x04, 0x57, 0xca, 0x00,
	0xc0, 0x00, 0x8d, 0xb4, 0x44, 0x72, 0x7e, 0xb3, 0x3b, 0xf8, 0xf6, 0xfe, 0xe0, 0xdb, 0x1f, 0x07,
	0xdf, 0x7e, 0x3a, 0xfa, 0xd6, 0xfe, 0xe8, 0x5b, 0x2f, 0x47, 0xdf, 0xba, 0xff, 0x7f, 0x51, 0x4f,
	0xbd, 0xd1, 0x29, 0x4f, 0x53, 0x8a, 0x29, 0xca, 0x9b, 0x18, 0x5e, 0xbc, 0x05, 0x5d, 0x61, 0xdc,
	0xd1, 0x2b, 0xfb, 0xf7, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x56, 0x23, 0x86, 0x2d, 0x02, 0x00,
	0x00,
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
	if len(m.ExecuteGas) > 0 {
		i -= len(m.ExecuteGas)
		copy(dAtA[i:], m.ExecuteGas)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ExecuteGas)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.PrepareGas) > 0 {
		i -= len(m.PrepareGas)
		copy(dAtA[i:], m.PrepareGas)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PrepareGas)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.MinCount) > 0 {
		i -= len(m.MinCount)
		copy(dAtA[i:], m.MinCount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MinCount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AskCount) > 0 {
		i -= len(m.AskCount)
		copy(dAtA[i:], m.AskCount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AskCount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleScriptId) > 0 {
		i -= len(m.OracleScriptId)
		copy(dAtA[i:], m.OracleScriptId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OracleScriptId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Creator)))
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
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.OracleScriptId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AskCount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MinCount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.PrepareGas)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ExecuteGas)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptId", wireType)
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
			m.OracleScriptId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
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
			m.AskCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
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
			m.MinCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGas", wireType)
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
			m.PrepareGas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGas", wireType)
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
			m.ExecuteGas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
