// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hellovoicy.proto

package hellovoicy

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// 呼び出しメッセージ
type HelloVoicyRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloVoicyRequest) Reset()         { *m = HelloVoicyRequest{} }
func (m *HelloVoicyRequest) String() string { return proto.CompactTextString(m) }
func (*HelloVoicyRequest) ProtoMessage()    {}
func (*HelloVoicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99e77eda75db571, []int{0}
}

func (m *HelloVoicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloVoicyRequest.Unmarshal(m, b)
}
func (m *HelloVoicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloVoicyRequest.Marshal(b, m, deterministic)
}
func (m *HelloVoicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloVoicyRequest.Merge(m, src)
}
func (m *HelloVoicyRequest) XXX_Size() int {
	return xxx_messageInfo_HelloVoicyRequest.Size(m)
}
func (m *HelloVoicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloVoicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloVoicyRequest proto.InternalMessageInfo

func (m *HelloVoicyRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// リターンメッセージ
type HelloVoicyReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloVoicyReply) Reset()         { *m = HelloVoicyReply{} }
func (m *HelloVoicyReply) String() string { return proto.CompactTextString(m) }
func (*HelloVoicyReply) ProtoMessage()    {}
func (*HelloVoicyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99e77eda75db571, []int{1}
}

func (m *HelloVoicyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloVoicyReply.Unmarshal(m, b)
}
func (m *HelloVoicyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloVoicyReply.Marshal(b, m, deterministic)
}
func (m *HelloVoicyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloVoicyReply.Merge(m, src)
}
func (m *HelloVoicyReply) XXX_Size() int {
	return xxx_messageInfo_HelloVoicyReply.Size(m)
}
func (m *HelloVoicyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloVoicyReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloVoicyReply proto.InternalMessageInfo

func (m *HelloVoicyReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloVoicyRequest)(nil), "hellovoicy.HelloVoicyRequest")
	proto.RegisterType((*HelloVoicyReply)(nil), "hellovoicy.HelloVoicyReply")
}

func init() { proto.RegisterFile("hellovoicy.proto", fileDescriptor_d99e77eda75db571) }

var fileDescriptor_d99e77eda75db571 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcb, 0xcf, 0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xe9, 0x72, 0x09, 0x7a, 0x80, 0x78, 0x61, 0x20, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xcd, 0xc5, 0x8f, 0xac, 0xbc, 0x20, 0xa7, 0x12, 0xb7, 0x62,
	0xa3, 0x60, 0x2e, 0x2e, 0x84, 0x62, 0x21, 0x57, 0x2e, 0x56, 0xb0, 0xbd, 0x42, 0xb2, 0x7a, 0x48,
	0x2e, 0xc2, 0xb0, 0x5c, 0x4a, 0x1a, 0x97, 0x74, 0x41, 0x4e, 0xa5, 0x12, 0x43, 0x12, 0x1b, 0xd8,
	0x0f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x5c, 0xc6, 0x0a, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloVoicyClient is the client API for HelloVoicy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloVoicyClient interface {
	Hello(ctx context.Context, in *HelloVoicyRequest, opts ...grpc.CallOption) (*HelloVoicyReply, error)
}

type helloVoicyClient struct {
	cc *grpc.ClientConn
}

func NewHelloVoicyClient(cc *grpc.ClientConn) HelloVoicyClient {
	return &helloVoicyClient{cc}
}

func (c *helloVoicyClient) Hello(ctx context.Context, in *HelloVoicyRequest, opts ...grpc.CallOption) (*HelloVoicyReply, error) {
	out := new(HelloVoicyReply)
	err := c.cc.Invoke(ctx, "/hellovoicy.HelloVoicy/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloVoicyServer is the server API for HelloVoicy service.
type HelloVoicyServer interface {
	Hello(context.Context, *HelloVoicyRequest) (*HelloVoicyReply, error)
}

func RegisterHelloVoicyServer(s *grpc.Server, srv HelloVoicyServer) {
	s.RegisterService(&_HelloVoicy_serviceDesc, srv)
}

func _HelloVoicy_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloVoicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloVoicyServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellovoicy.HelloVoicy/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloVoicyServer).Hello(ctx, req.(*HelloVoicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloVoicy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hellovoicy.HelloVoicy",
	HandlerType: (*HelloVoicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hello",
			Handler:    _HelloVoicy_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hellovoicy.proto",
}
