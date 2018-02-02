// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestRequest
	TestReply
*/
package test

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

// TestRequest 请求结构
type TestRequest struct {
	Typeid int64 `protobuf:"varint,1,opt,name=typeid" json:"typeid,omitempty"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestRequest) GetTypeid() int64 {
	if m != nil {
		return m.Typeid
	}
	return 0
}

// TestReply 响应结构
type TestReply struct {
	// repeated重复(数组)
	UserList []*TestReply_GetData `protobuf:"bytes,1,rep,name=user_list,json=userList" json:"user_list,omitempty"`
}

func (m *TestReply) Reset()                    { *m = TestReply{} }
func (m *TestReply) String() string            { return proto.CompactTextString(m) }
func (*TestReply) ProtoMessage()               {}
func (*TestReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestReply) GetUserList() []*TestReply_GetData {
	if m != nil {
		return m.UserList
	}
	return nil
}

// 返回数据类型
type TestReply_GetData struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *TestReply_GetData) Reset()                    { *m = TestReply_GetData{} }
func (m *TestReply_GetData) String() string            { return proto.CompactTextString(m) }
func (*TestReply_GetData) ProtoMessage()               {}
func (*TestReply_GetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *TestReply_GetData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestReply_GetData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "test.TestRequest")
	proto.RegisterType((*TestReply)(nil), "test.TestReply")
	proto.RegisterType((*TestReply_GetData)(nil), "test.TestReply.GetData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Test service

type TestClient interface {
	// 定义 SayTest 方法
	SayTest(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) SayTest(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error) {
	out := new(TestReply)
	err := grpc.Invoke(ctx, "/test.Test/SayTest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestServer interface {
	// 定义 SayTest 方法
	SayTest(context.Context, *TestRequest) (*TestReply, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_SayTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).SayTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/SayTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).SayTest(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayTest",
			Handler:    _Test_SayTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x54, 0xb9, 0xb8, 0x43, 0x52,
	0x8b, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x4a, 0x2a, 0x0b,
	0x52, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0xa5, 0x02, 0x2e, 0x4e,
	0x88, 0xb2, 0x82, 0x9c, 0x4a, 0x21, 0x13, 0x2e, 0xce, 0xd2, 0xe2, 0xd4, 0xa2, 0xf8, 0x9c, 0xcc,
	0xe2, 0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x71, 0x3d, 0xb0, 0xc9, 0x70, 0x35, 0x7a,
	0xee, 0xa9, 0x25, 0x2e, 0x89, 0x25, 0x89, 0x41, 0x1c, 0x20, 0x95, 0x3e, 0x99, 0xc5, 0x25, 0x52,
	0xba, 0x5c, 0xec, 0x50, 0x41, 0x21, 0x3e, 0x2e, 0x26, 0xb8, 0x0d, 0x4c, 0x99, 0x29, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x91,
	0x39, 0x17, 0x0b, 0xc8, 0x34, 0x21, 0x7d, 0x2e, 0xf6, 0xe0, 0xc4, 0x4a, 0x30, 0x53, 0x10, 0xd9,
	0x12, 0xb0, 0x7b, 0xa5, 0xf8, 0xd1, 0xec, 0x55, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xcf, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x86, 0x6f, 0x1b, 0xec, 0x00, 0x00, 0x00,
}