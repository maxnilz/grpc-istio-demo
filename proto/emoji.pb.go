// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emoji.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
	return fileDescriptor_emoji_2c56afabb313f538, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
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

type HelloResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_2c56afabb313f538, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EmojiRequest struct {
	InputText            string   `protobuf:"bytes,1,opt,name=input_text,json=inputText,proto3" json:"input_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojiRequest) Reset()         { *m = EmojiRequest{} }
func (m *EmojiRequest) String() string { return proto.CompactTextString(m) }
func (*EmojiRequest) ProtoMessage()    {}
func (*EmojiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_2c56afabb313f538, []int{2}
}
func (m *EmojiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojiRequest.Unmarshal(m, b)
}
func (m *EmojiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojiRequest.Marshal(b, m, deterministic)
}
func (dst *EmojiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojiRequest.Merge(dst, src)
}
func (m *EmojiRequest) XXX_Size() int {
	return xxx_messageInfo_EmojiRequest.Size(m)
}
func (m *EmojiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmojiRequest proto.InternalMessageInfo

func (m *EmojiRequest) GetInputText() string {
	if m != nil {
		return m.InputText
	}
	return ""
}

type EmojiResponse struct {
	OutputText           string   `protobuf:"bytes,1,opt,name=output_text,json=outputText,proto3" json:"output_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojiResponse) Reset()         { *m = EmojiResponse{} }
func (m *EmojiResponse) String() string { return proto.CompactTextString(m) }
func (*EmojiResponse) ProtoMessage()    {}
func (*EmojiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_emoji_2c56afabb313f538, []int{3}
}
func (m *EmojiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojiResponse.Unmarshal(m, b)
}
func (m *EmojiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojiResponse.Marshal(b, m, deterministic)
}
func (dst *EmojiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojiResponse.Merge(dst, src)
}
func (m *EmojiResponse) XXX_Size() int {
	return xxx_messageInfo_EmojiResponse.Size(m)
}
func (m *EmojiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmojiResponse proto.InternalMessageInfo

func (m *EmojiResponse) GetOutputText() string {
	if m != nil {
		return m.OutputText
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
	proto.RegisterType((*EmojiRequest)(nil), "proto.EmojiRequest")
	proto.RegisterType((*EmojiResponse)(nil), "proto.EmojiResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmojiServiceClient is the client API for EmojiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmojiServiceClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error)
}

type emojiServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmojiServiceClient(cc *grpc.ClientConn) EmojiServiceClient {
	return &emojiServiceClient{cc}
}

func (c *emojiServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.EmojiService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emojiServiceClient) InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error) {
	out := new(EmojiResponse)
	err := c.cc.Invoke(ctx, "/proto.EmojiService/InsertEmojis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmojiServiceServer is the server API for EmojiService service.
type EmojiServiceServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	InsertEmojis(context.Context, *EmojiRequest) (*EmojiResponse, error)
}

func RegisterEmojiServiceServer(s *grpc.Server, srv EmojiServiceServer) {
	s.RegisterService(&_EmojiService_serviceDesc, srv)
}

func _EmojiService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmojiService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmojiService_InsertEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmojiService/InsertEmojis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, req.(*EmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmojiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmojiService",
	HandlerType: (*EmojiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _EmojiService_SayHello_Handler,
		},
		{
			MethodName: "InsertEmojis",
			Handler:    _EmojiService_InsertEmojis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emoji.proto",
}

func init() { proto.RegisterFile("emoji.proto", fileDescriptor_emoji_2c56afabb313f538) }

var fileDescriptor_emoji_2c56afabb313f538 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0xcd, 0xcf,
	0xca, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x32, 0xe9, 0xf9, 0xf9,
	0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99,
	0xf9, 0x79, 0xc5, 0x10, 0x45, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x32, 0x17, 0x2f, 0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0x2a, 0x56, 0x45, 0xba, 0x5c, 0x3c, 0xae, 0x20, 0xcb, 0x61, 0x06, 0xc9, 0x72, 0x71, 0x65,
	0xe6, 0x15, 0x94, 0x96, 0xc4, 0x97, 0xa4, 0x56, 0x94, 0x40, 0x55, 0x72, 0x82, 0x45, 0x42, 0x52,
	0x2b, 0x4a, 0x94, 0x0c, 0xb8, 0x78, 0xa1, 0xca, 0xa1, 0x66, 0xca, 0x73, 0x71, 0xe7, 0x97, 0x96,
	0xa0, 0x69, 0xe0, 0x82, 0x08, 0x81, 0x74, 0x18, 0xad, 0x61, 0x84, 0xda, 0x10, 0x9c, 0x5a, 0x54,
	0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc5, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0x76, 0x99, 0x90, 0x30, 0xc4,
	0x3b, 0x7a, 0xc8, 0x7e, 0x91, 0x12, 0x41, 0x15, 0x84, 0x58, 0xa4, 0x24, 0xdc, 0x74, 0xf9, 0xc9,
	0x64, 0x26, 0x5e, 0x25, 0x0e, 0x70, 0x88, 0x14, 0x27, 0x56, 0x5a, 0x31, 0x6a, 0x09, 0x05, 0x70,
	0xf1, 0x78, 0xe6, 0x15, 0xa7, 0x16, 0x95, 0x80, 0x6d, 0x28, 0x86, 0x9b, 0x87, 0xec, 0x25, 0xb8,
	0x79, 0x28, 0x0e, 0x57, 0x12, 0x05, 0x9b, 0xc7, 0xaf, 0xc4, 0x05, 0x36, 0x0f, 0x1c, 0x01, 0x56,
	0x8c, 0x5a, 0x49, 0x6c, 0x60, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc4, 0x4a,
	0x5d, 0x93, 0x01, 0x00, 0x00,
}