// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello2.proto

package helloservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import "net/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello2_4d4f3d4a9b4e2a63, []int{0}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "helloservice.String")
}

type HelloServiceInterface interface {
	Hello(*String, *String) error
}

func RegisterHelloService(
	srv *rpc.Server, x HelloServiceInterface,
) error {
	if err := srv.RegisterName("HelloService", x); err != nil {
		return err
	}
	return nil
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (
	*HelloServiceClient, error,
) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(
	in *String, out *String,
) error {
	return p.Client.Call("HelloService.Hello", in, out)
}

func init() { proto.RegisterFile("hello2.proto", fileDescriptor_hello2_4d4f3d4a9b4e2a63) }

var fileDescriptor_hello2_4d4f3d4a9b4e2a63 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0xcd, 0xc9,
	0xc9, 0x37, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0xf0, 0x8a, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x95, 0xe4, 0xb8, 0xd8, 0x82, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x85, 0x44, 0xb8, 0x58,
	0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x23, 0x67,
	0x2e, 0x1e, 0x0f, 0x90, 0xfa, 0x60, 0x88, 0x7a, 0x21, 0x63, 0x2e, 0x56, 0x30, 0x5f, 0x48, 0x44,
	0x0f, 0xd9, 0x1c, 0x3d, 0x88, 0x21, 0x52, 0x58, 0x45, 0x93, 0xd8, 0xc0, 0x36, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0x11, 0x08, 0x20, 0x89, 0x00, 0x00, 0x00,
}
