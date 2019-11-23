// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloServerRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloServerRequest) Reset()         { *m = HelloServerRequest{} }
func (m *HelloServerRequest) String() string { return proto.CompactTextString(m) }
func (*HelloServerRequest) ProtoMessage()    {}
func (*HelloServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *HelloServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloServerRequest.Unmarshal(m, b)
}
func (m *HelloServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloServerRequest.Marshal(b, m, deterministic)
}
func (m *HelloServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloServerRequest.Merge(m, src)
}
func (m *HelloServerRequest) XXX_Size() int {
	return xxx_messageInfo_HelloServerRequest.Size(m)
}
func (m *HelloServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloServerRequest proto.InternalMessageInfo

func (m *HelloServerRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloServerResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloServerResponse) Reset()         { *m = HelloServerResponse{} }
func (m *HelloServerResponse) String() string { return proto.CompactTextString(m) }
func (*HelloServerResponse) ProtoMessage()    {}
func (*HelloServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *HelloServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloServerResponse.Unmarshal(m, b)
}
func (m *HelloServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloServerResponse.Marshal(b, m, deterministic)
}
func (m *HelloServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloServerResponse.Merge(m, src)
}
func (m *HelloServerResponse) XXX_Size() int {
	return xxx_messageInfo_HelloServerResponse.Size(m)
}
func (m *HelloServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloServerResponse proto.InternalMessageInfo

func (m *HelloServerResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*HelloServerRequest)(nil), "helloworld.HelloServerRequest")
	proto.RegisterType((*HelloServerResponse)(nil), "helloworld.HelloServerResponse")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0xab, 0xa4, 0xc7, 0x25, 0x04, 0x56, 0x17, 0x9c, 0x5a, 0x54, 0x96,
	0x5a, 0x04, 0x33, 0x11, 0xb7, 0x7a, 0x7d, 0x2e, 0x61, 0x14, 0xf5, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0xb8, 0x35, 0x18, 0xcd, 0x66, 0xe4, 0x62, 0x77, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x12,
	0xb2, 0xe3, 0xe2, 0x08, 0x4e, 0xac, 0x04, 0xeb, 0x17, 0x92, 0xd0, 0x43, 0xf2, 0x23, 0xb2, 0x77,
	0xa4, 0xc4, 0xb0, 0xc8, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x05, 0x70, 0x71, 0x23, 0x59, 0x2e,
	0x24, 0x87, 0xa1, 0x10, 0xc5, 0x17, 0x52, 0xf2, 0x38, 0xe5, 0x21, 0xae, 0x56, 0x62, 0x70, 0x32,
	0xe0, 0x92, 0xce, 0xcc, 0xd7, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0x2d, 0x46, 0xd2, 0xe4, 0xc4, 0x0f, 0xd6, 0x15, 0x0e, 0x62, 0x07, 0x80, 0xa2, 0x21, 0x80,
	0x31, 0x89, 0x0d, 0x1c, 0x1f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x1f, 0xe4, 0xa9,
	0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	HelloServer(ctx context.Context, in *HelloServerRequest, opts ...grpc.CallOption) (*HelloServerResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) HelloServer(ctx context.Context, in *HelloServerRequest, opts ...grpc.CallOption) (*HelloServerResponse, error) {
	out := new(HelloServerResponse)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/HelloServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	HelloServer(context.Context, *HelloServerRequest) (*HelloServerResponse, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) HelloServer(ctx context.Context, req *HelloServerRequest) (*HelloServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloServer not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_HelloServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).HelloServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/HelloServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).HelloServer(ctx, req.(*HelloServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "HelloServer",
			Handler:    _Greeter_HelloServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
