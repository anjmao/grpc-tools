// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug.proto

package debug

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{0}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(dst, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

type GetVersionReply struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionReply) Reset()         { *m = GetVersionReply{} }
func (m *GetVersionReply) String() string { return proto.CompactTextString(m) }
func (*GetVersionReply) ProtoMessage()    {}
func (*GetVersionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{1}
}
func (m *GetVersionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionReply.Unmarshal(m, b)
}
func (m *GetVersionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionReply.Marshal(b, m, deterministic)
}
func (dst *GetVersionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionReply.Merge(dst, src)
}
func (m *GetVersionReply) XXX_Size() int {
	return xxx_messageInfo_GetVersionReply.Size(m)
}
func (m *GetVersionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionReply proto.InternalMessageInfo

func (m *GetVersionReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{2}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{3}
}
func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (dst *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(dst, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetStreamRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStreamRequest) Reset()         { *m = GetStreamRequest{} }
func (m *GetStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GetStreamRequest) ProtoMessage()    {}
func (*GetStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{4}
}
func (m *GetStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStreamRequest.Unmarshal(m, b)
}
func (m *GetStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStreamRequest.Marshal(b, m, deterministic)
}
func (dst *GetStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStreamRequest.Merge(dst, src)
}
func (m *GetStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GetStreamRequest.Size(m)
}
func (m *GetStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStreamRequest proto.InternalMessageInfo

type GetStreamReply struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStreamReply) Reset()         { *m = GetStreamReply{} }
func (m *GetStreamReply) String() string { return proto.CompactTextString(m) }
func (*GetStreamReply) ProtoMessage()    {}
func (*GetStreamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_debug_a35c2808ca6542ca, []int{5}
}
func (m *GetStreamReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStreamReply.Unmarshal(m, b)
}
func (m *GetStreamReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStreamReply.Marshal(b, m, deterministic)
}
func (dst *GetStreamReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStreamReply.Merge(dst, src)
}
func (m *GetStreamReply) XXX_Size() int {
	return xxx_messageInfo_GetStreamReply.Size(m)
}
func (m *GetStreamReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStreamReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetStreamReply proto.InternalMessageInfo

func (m *GetStreamReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*GetVersionRequest)(nil), "debug.GetVersionRequest")
	proto.RegisterType((*GetVersionReply)(nil), "debug.GetVersionReply")
	proto.RegisterType((*PingRequest)(nil), "debug.PingRequest")
	proto.RegisterType((*PingReply)(nil), "debug.PingReply")
	proto.RegisterType((*GetStreamRequest)(nil), "debug.GetStreamRequest")
	proto.RegisterType((*GetStreamReply)(nil), "debug.GetStreamReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (Debug_GetStreamClient, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionReply, error) {
	out := new(GetVersionReply)
	err := c.cc.Invoke(ctx, "/debug.Debug/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/debug.Debug/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (Debug_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/debug.Debug/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_GetStreamClient interface {
	Recv() (*GetStreamReply, error)
	grpc.ClientStream
}

type debugGetStreamClient struct {
	grpc.ClientStream
}

func (x *debugGetStreamClient) Recv() (*GetStreamReply, error) {
	m := new(GetStreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetStream(*GetStreamRequest, Debug_GetStreamServer) error
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/debug.Debug/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).GetStream(m, &debugGetStreamServer{stream})
}

type Debug_GetStreamServer interface {
	Send(*GetStreamReply) error
	grpc.ServerStream
}

type debugGetStreamServer struct {
	grpc.ServerStream
}

func (x *debugGetStreamServer) Send(m *GetStreamReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Debug_GetVersion_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Debug_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Debug_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "debug.proto",
}

func init() { proto.RegisterFile("debug.proto", fileDescriptor_debug_a35c2808ca6542ca) }

var fileDescriptor_debug_a35c2808ca6542ca = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x49, 0x4d, 0x2a,
	0x4d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x84, 0xb9, 0x04, 0xdd,
	0x53, 0x4b, 0xc2, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0xb4, 0xb9, 0xf8, 0x91, 0x05, 0x0b, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0xcb, 0x20, 0x7c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x49, 0x9d, 0x8b, 0x3b, 0x20, 0x33, 0x2f,
	0x1d, 0xaa, 0x17, 0xa4, 0x30, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x15, 0xa6, 0x10, 0xca, 0x55,
	0x52, 0xe5, 0xe2, 0x84, 0x28, 0x84, 0x9a, 0x87, 0x43, 0x99, 0x10, 0x97, 0x80, 0x7b, 0x6a, 0x49,
	0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0xcc, 0x41, 0x2a, 0x5c, 0x7c, 0x48, 0x62, 0x20, 0xfd, 0x42,
	0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x50, 0xcd, 0x60, 0xb6, 0xd1, 0x2e, 0x46, 0x2e, 0x56, 0x17,
	0x90, 0xaf, 0x84, 0x1c, 0xb8, 0xb8, 0x10, 0x1e, 0x10, 0x92, 0xd0, 0x83, 0x78, 0x1c, 0xc3, 0xa3,
	0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x84, 0xf4, 0xb8, 0x58, 0x40, 0x8e, 0x15,
	0x12, 0x82, 0xaa, 0x40, 0xf2, 0xa2, 0x94, 0x00, 0x8a, 0x18, 0x44, 0xbd, 0x3d, 0x17, 0x27, 0xdc,
	0x85, 0x42, 0xe2, 0x08, 0x63, 0x51, 0xfc, 0x21, 0x25, 0x8a, 0x29, 0x01, 0xd6, 0x6e, 0xc0, 0x98,
	0xc4, 0x06, 0x8e, 0x16, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x62, 0xfe, 0xae, 0xa5,
	0x01, 0x00, 0x00,
}
