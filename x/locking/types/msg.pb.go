// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/locking/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type MsgLockTokens struct {
	Owner    string                                  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Duration time.Duration                           `protobuf:"bytes,2,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	Coin     github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=coin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *MsgLockTokens) Reset()         { *m = MsgLockTokens{} }
func (m *MsgLockTokens) String() string { return proto.CompactTextString(m) }
func (*MsgLockTokens) ProtoMessage()    {}
func (*MsgLockTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_bed6c441190fad9b, []int{0}
}
func (m *MsgLockTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLockTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLockTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLockTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLockTokens.Merge(m, src)
}
func (m *MsgLockTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgLockTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLockTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLockTokens proto.InternalMessageInfo

func (m *MsgLockTokens) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgLockTokens) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *MsgLockTokens) GetCoin() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Coin
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

type MsgLockTokensResponse struct {
}

func (m *MsgLockTokensResponse) Reset()         { *m = MsgLockTokensResponse{} }
func (m *MsgLockTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLockTokensResponse) ProtoMessage()    {}
func (*MsgLockTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bed6c441190fad9b, []int{1}
}
func (m *MsgLockTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLockTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLockTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLockTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLockTokensResponse.Merge(m, src)
}
func (m *MsgLockTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLockTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLockTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLockTokensResponse proto.InternalMessageInfo

type MsgBeginUnlockingTokens struct {
	Owner  string                                  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	LockId uint64                                  `protobuf:"varint,2,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty" yaml:"lock_id"`
	Coin   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=coin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *MsgBeginUnlockingTokens) Reset()         { *m = MsgBeginUnlockingTokens{} }
func (m *MsgBeginUnlockingTokens) String() string { return proto.CompactTextString(m) }
func (*MsgBeginUnlockingTokens) ProtoMessage()    {}
func (*MsgBeginUnlockingTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_bed6c441190fad9b, []int{2}
}
func (m *MsgBeginUnlockingTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBeginUnlockingTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBeginUnlockingTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBeginUnlockingTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBeginUnlockingTokens.Merge(m, src)
}
func (m *MsgBeginUnlockingTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgBeginUnlockingTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBeginUnlockingTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBeginUnlockingTokens proto.InternalMessageInfo

func (m *MsgBeginUnlockingTokens) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgBeginUnlockingTokens) GetLockId() uint64 {
	if m != nil {
		return m.LockId
	}
	return 0
}

func (m *MsgBeginUnlockingTokens) GetCoin() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Coin
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

type MsgBeginUnlockingTokensResponse struct {
}

func (m *MsgBeginUnlockingTokensResponse) Reset()         { *m = MsgBeginUnlockingTokensResponse{} }
func (m *MsgBeginUnlockingTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBeginUnlockingTokensResponse) ProtoMessage()    {}
func (*MsgBeginUnlockingTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bed6c441190fad9b, []int{3}
}
func (m *MsgBeginUnlockingTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBeginUnlockingTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBeginUnlockingTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBeginUnlockingTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBeginUnlockingTokensResponse.Merge(m, src)
}
func (m *MsgBeginUnlockingTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBeginUnlockingTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBeginUnlockingTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBeginUnlockingTokensResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLockTokens)(nil), "comdex.locking.v1beta1.MsgLockTokens")
	proto.RegisterType((*MsgLockTokensResponse)(nil), "comdex.locking.v1beta1.MsgLockTokensResponse")
	proto.RegisterType((*MsgBeginUnlockingTokens)(nil), "comdex.locking.v1beta1.MsgBeginUnlockingTokens")
	proto.RegisterType((*MsgBeginUnlockingTokensResponse)(nil), "comdex.locking.v1beta1.MsgBeginUnlockingTokensResponse")
}

func init() { proto.RegisterFile("comdex/locking/v1beta1/msg.proto", fileDescriptor_bed6c441190fad9b) }

var fileDescriptor_bed6c441190fad9b = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0xb5, 0xa5, 0xc0, 0xf1, 0x6d, 0x01, 0x0d, 0x19, 0xec, 0x60, 0x09, 0xa8, 0x04, 0xb9,
	0x53, 0xca, 0x80, 0xc4, 0x68, 0x58, 0x10, 0x64, 0xb1, 0x60, 0x61, 0x00, 0xf9, 0xe3, 0x72, 0x3d,
	0x25, 0xbe, 0x67, 0xe5, 0x1c, 0x20, 0x03, 0x03, 0xff, 0x01, 0x23, 0x7f, 0x52, 0x27, 0xd4, 0x91,
	0xc9, 0xa0, 0x64, 0x41, 0x8c, 0x19, 0x99, 0x90, 0xef, 0xa3, 0x34, 0x52, 0x41, 0x64, 0x60, 0xf2,
	0x9d, 0xdf, 0xef, 0xbd, 0xdf, 0x87, 0x9f, 0x71, 0x37, 0x83, 0x22, 0x67, 0xef, 0xe8, 0x18, 0xb2,
	0x91, 0x90, 0x9c, 0xbe, 0xe9, 0xa7, 0xac, 0x4a, 0xfa, 0xb4, 0x50, 0x9c, 0x94, 0x13, 0xa8, 0xc0,
	0xbb, 0x6e, 0x10, 0xc4, 0x22, 0x88, 0x45, 0x74, 0xae, 0x72, 0xe0, 0xa0, 0x21, 0xb4, 0x39, 0x19,
	0x74, 0xc7, 0xe7, 0x00, 0x7c, 0xcc, 0xa8, 0xbe, 0xa5, 0xd3, 0x21, 0xcd, 0xa7, 0x93, 0xa4, 0x12,
	0x20, 0x5d, 0x3d, 0x03, 0x55, 0x80, 0xa2, 0x69, 0xa2, 0xd8, 0x11, 0x59, 0x06, 0xc2, 0xd6, 0xc3,
	0x0f, 0x1b, 0xf8, 0xc2, 0x40, 0xf1, 0x67, 0x90, 0x8d, 0x9e, 0xc3, 0x88, 0x49, 0xe5, 0xdd, 0xc6,
	0xa7, 0xe0, 0xad, 0x64, 0x93, 0x36, 0xea, 0xa2, 0xdd, 0xb3, 0xd1, 0xe5, 0x65, 0x1d, 0x9c, 0x9f,
	0x25, 0xc5, 0xf8, 0x61, 0xa8, 0x5f, 0x87, 0xb1, 0x29, 0x7b, 0xfb, 0xf8, 0x8c, 0xe3, 0x6a, 0x6f,
	0x74, 0xd1, 0xee, 0xb9, 0xbd, 0x1b, 0xc4, 0x88, 0x21, 0x4e, 0x0c, 0x79, 0x6c, 0x01, 0x51, 0xff,
	0xa0, 0x0e, 0x5a, 0x3f, 0xea, 0xc0, 0x73, 0x2d, 0xf7, 0xa0, 0x10, 0x15, 0x2b, 0xca, 0x6a, 0xb6,
	0xac, 0x83, 0x4b, 0x66, 0xbe, 0xab, 0x85, 0x9f, 0xbe, 0x06, 0x28, 0x3e, 0x9a, 0xee, 0xbd, 0xc2,
	0x5b, 0x8d, 0xe2, 0xf6, 0xa6, 0x65, 0x31, 0x96, 0x48, 0x63, 0xc9, 0xa5, 0x43, 0x1e, 0x81, 0x90,
	0x11, 0x6d, 0x58, 0x7e, 0xd6, 0xc1, 0x1d, 0x2e, 0xaa, 0xfd, 0x69, 0x4a, 0x32, 0x28, 0xa8, 0xf5,
	0x6f, 0x1e, 0x3d, 0x95, 0x8f, 0x68, 0x35, 0x2b, 0x99, 0xd2, 0x0d, 0xb1, 0x9e, 0x1b, 0xee, 0xe0,
	0x6b, 0x2b, 0x11, 0xc4, 0x4c, 0x95, 0x20, 0x15, 0x0b, 0x3f, 0x23, 0xbc, 0x33, 0x50, 0x3c, 0x62,
	0x5c, 0xc8, 0x17, 0xd2, 0x7e, 0x90, 0x35, 0x63, 0xba, 0x8b, 0x4f, 0x37, 0x8d, 0xaf, 0x45, 0xae,
	0x53, 0xda, 0x8a, 0xbc, 0x65, 0x1d, 0x5c, 0x34, 0x48, 0x5b, 0x08, 0xe3, 0xed, 0xe6, 0xf4, 0x24,
	0xff, 0xef, 0x4e, 0x6f, 0xe2, 0xe0, 0x0f, 0x7e, 0x9c, 0xe7, 0xbd, 0xef, 0x08, 0x6f, 0x0e, 0x14,
	0xf7, 0x52, 0x8c, 0x8f, 0x2d, 0xc5, 0x2d, 0x72, 0xf2, 0x56, 0x92, 0x95, 0xe0, 0x3a, 0xbd, 0x7f,
	0x82, 0x39, 0x2e, 0xef, 0x3d, 0xbe, 0x72, 0x4c, 0x8b, 0xa5, 0xa2, 0x7f, 0x99, 0x71, 0x92, 0xf2,
	0xce, 0x83, 0x35, 0x1b, 0x1c, 0x7d, 0xf4, 0xf4, 0x60, 0xee, 0xa3, 0xc3, 0xb9, 0x8f, 0xbe, 0xcd,
	0x7d, 0xf4, 0x71, 0xe1, 0xb7, 0x0e, 0x17, 0x7e, 0xeb, 0xcb, 0xc2, 0x6f, 0xbd, 0xec, 0xaf, 0xc4,
	0xda, 0x0c, 0xef, 0xc1, 0x70, 0x28, 0x32, 0x91, 0x8c, 0xed, 0x9d, 0xfe, 0xfe, 0x85, 0x75, 0xca,
	0xe9, 0xb6, 0x5e, 0xfa, 0xfb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0x89, 0xa0, 0xd8, 0xe1,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	LockTokens(ctx context.Context, in *MsgLockTokens, opts ...grpc.CallOption) (*MsgLockTokensResponse, error)
	BeginUnlockTokens(ctx context.Context, in *MsgBeginUnlockingTokens, opts ...grpc.CallOption) (*MsgBeginUnlockingTokensResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) LockTokens(ctx context.Context, in *MsgLockTokens, opts ...grpc.CallOption) (*MsgLockTokensResponse, error) {
	out := new(MsgLockTokensResponse)
	err := c.cc.Invoke(ctx, "/comdex.locking.v1beta1.Msg/LockTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BeginUnlockTokens(ctx context.Context, in *MsgBeginUnlockingTokens, opts ...grpc.CallOption) (*MsgBeginUnlockingTokensResponse, error) {
	out := new(MsgBeginUnlockingTokensResponse)
	err := c.cc.Invoke(ctx, "/comdex.locking.v1beta1.Msg/BeginUnlockTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	LockTokens(context.Context, *MsgLockTokens) (*MsgLockTokensResponse, error)
	BeginUnlockTokens(context.Context, *MsgBeginUnlockingTokens) (*MsgBeginUnlockingTokensResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) LockTokens(ctx context.Context, req *MsgLockTokens) (*MsgLockTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockTokens not implemented")
}
func (*UnimplementedMsgServer) BeginUnlockTokens(ctx context.Context, req *MsgBeginUnlockingTokens) (*MsgBeginUnlockingTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginUnlockTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_LockTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLockTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LockTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.locking.v1beta1.Msg/LockTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LockTokens(ctx, req.(*MsgLockTokens))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BeginUnlockTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBeginUnlockingTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BeginUnlockTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.locking.v1beta1.Msg/BeginUnlockTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BeginUnlockTokens(ctx, req.(*MsgBeginUnlockingTokens))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.locking.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LockTokens",
			Handler:    _Msg_LockTokens_Handler,
		},
		{
			MethodName: "BeginUnlockTokens",
			Handler:    _Msg_BeginUnlockTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/locking/v1beta1/msg.proto",
}

func (m *MsgLockTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLockTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLockTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMsg(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLockTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLockTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLockTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBeginUnlockingTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBeginUnlockingTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBeginUnlockingTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.LockId != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.LockId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBeginUnlockingTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBeginUnlockingTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBeginUnlockingTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLockTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovMsg(uint64(l))
	l = m.Coin.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgLockTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBeginUnlockingTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.LockId != 0 {
		n += 1 + sovMsg(uint64(m.LockId))
	}
	l = m.Coin.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgBeginUnlockingTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLockTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgLockTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLockTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgLockTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgLockTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLockTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgBeginUnlockingTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgBeginUnlockingTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBeginUnlockingTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockId", wireType)
			}
			m.LockId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgBeginUnlockingTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgBeginUnlockingTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBeginUnlockingTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
