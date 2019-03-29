// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HelloResp struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Salutation           string               `protobuf:"bytes,2,opt,name=salutation,proto3" json:"salutation,omitempty"`
	Subject              string               `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HelloResp) Reset()         { *m = HelloResp{} }
func (m *HelloResp) String() string { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()    {}
func (*HelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{0}
}
func (m *HelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResp.Unmarshal(m, b)
}
func (m *HelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResp.Marshal(b, m, deterministic)
}
func (dst *HelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResp.Merge(dst, src)
}
func (m *HelloResp) XXX_Size() int {
	return xxx_messageInfo_HelloResp.Size(m)
}
func (m *HelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResp proto.InternalMessageInfo

func (m *HelloResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HelloResp) GetSalutation() string {
	if m != nil {
		return m.Salutation
	}
	return ""
}

func (m *HelloResp) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *HelloResp) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *HelloResp) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *HelloResp) GetDeleted() *timestamp.Timestamp {
	if m != nil {
		return m.Deleted
	}
	return nil
}

type AddHelloReq struct {
	Salutation           string   `protobuf:"bytes,1,opt,name=salutation,proto3" json:"salutation,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHelloReq) Reset()         { *m = AddHelloReq{} }
func (m *AddHelloReq) String() string { return proto.CompactTextString(m) }
func (*AddHelloReq) ProtoMessage()    {}
func (*AddHelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{1}
}
func (m *AddHelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHelloReq.Unmarshal(m, b)
}
func (m *AddHelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHelloReq.Marshal(b, m, deterministic)
}
func (dst *AddHelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHelloReq.Merge(dst, src)
}
func (m *AddHelloReq) XXX_Size() int {
	return xxx_messageInfo_AddHelloReq.Size(m)
}
func (m *AddHelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddHelloReq proto.InternalMessageInfo

func (m *AddHelloReq) GetSalutation() string {
	if m != nil {
		return m.Salutation
	}
	return ""
}

func (m *AddHelloReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type OvrHelloReq struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Req                  *AddHelloReq `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OvrHelloReq) Reset()         { *m = OvrHelloReq{} }
func (m *OvrHelloReq) String() string { return proto.CompactTextString(m) }
func (*OvrHelloReq) ProtoMessage()    {}
func (*OvrHelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{2}
}
func (m *OvrHelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OvrHelloReq.Unmarshal(m, b)
}
func (m *OvrHelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OvrHelloReq.Marshal(b, m, deterministic)
}
func (dst *OvrHelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OvrHelloReq.Merge(dst, src)
}
func (m *OvrHelloReq) XXX_Size() int {
	return xxx_messageInfo_OvrHelloReq.Size(m)
}
func (m *OvrHelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OvrHelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_OvrHelloReq proto.InternalMessageInfo

func (m *OvrHelloReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OvrHelloReq) GetReq() *AddHelloReq {
	if m != nil {
		return m.Req
	}
	return nil
}

type RmvHelloResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmvHelloResp) Reset()         { *m = RmvHelloResp{} }
func (m *RmvHelloResp) String() string { return proto.CompactTextString(m) }
func (*RmvHelloResp) ProtoMessage()    {}
func (*RmvHelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{3}
}
func (m *RmvHelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmvHelloResp.Unmarshal(m, b)
}
func (m *RmvHelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmvHelloResp.Marshal(b, m, deterministic)
}
func (dst *RmvHelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmvHelloResp.Merge(dst, src)
}
func (m *RmvHelloResp) XXX_Size() int {
	return xxx_messageInfo_RmvHelloResp.Size(m)
}
func (m *RmvHelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RmvHelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_RmvHelloResp proto.InternalMessageInfo

type RmvHelloReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmvHelloReq) Reset()         { *m = RmvHelloReq{} }
func (m *RmvHelloReq) String() string { return proto.CompactTextString(m) }
func (*RmvHelloReq) ProtoMessage()    {}
func (*RmvHelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{4}
}
func (m *RmvHelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmvHelloReq.Unmarshal(m, b)
}
func (m *RmvHelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmvHelloReq.Marshal(b, m, deterministic)
}
func (dst *RmvHelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmvHelloReq.Merge(dst, src)
}
func (m *RmvHelloReq) XXX_Size() int {
	return xxx_messageInfo_RmvHelloReq.Size(m)
}
func (m *RmvHelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RmvHelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_RmvHelloReq proto.InternalMessageInfo

func (m *RmvHelloReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type HellosResp struct {
	Items                []*HelloResp `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HellosResp) Reset()         { *m = HellosResp{} }
func (m *HellosResp) String() string { return proto.CompactTextString(m) }
func (*HellosResp) ProtoMessage()    {}
func (*HellosResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{5}
}
func (m *HellosResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HellosResp.Unmarshal(m, b)
}
func (m *HellosResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HellosResp.Marshal(b, m, deterministic)
}
func (dst *HellosResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HellosResp.Merge(dst, src)
}
func (m *HellosResp) XXX_Size() int {
	return xxx_messageInfo_HellosResp.Size(m)
}
func (m *HellosResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HellosResp.DiscardUnknown(m)
}

var xxx_messageInfo_HellosResp proto.InternalMessageInfo

func (m *HellosResp) GetItems() []*HelloResp {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *HellosResp) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FndHellosReq struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Salutations          []string `protobuf:"bytes,2,rep,name=salutations,proto3" json:"salutations,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Lapse                uint32   `protobuf:"varint,4,opt,name=lapse,proto3" json:"lapse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FndHellosReq) Reset()         { *m = FndHellosReq{} }
func (m *FndHellosReq) String() string { return proto.CompactTextString(m) }
func (*FndHellosReq) ProtoMessage()    {}
func (*FndHellosReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_062e3cf3bf8b761a, []int{6}
}
func (m *FndHellosReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FndHellosReq.Unmarshal(m, b)
}
func (m *FndHellosReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FndHellosReq.Marshal(b, m, deterministic)
}
func (dst *FndHellosReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FndHellosReq.Merge(dst, src)
}
func (m *FndHellosReq) XXX_Size() int {
	return xxx_messageInfo_FndHellosReq.Size(m)
}
func (m *FndHellosReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FndHellosReq.DiscardUnknown(m)
}

var xxx_messageInfo_FndHellosReq proto.InternalMessageInfo

func (m *FndHellosReq) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FndHellosReq) GetSalutations() []string {
	if m != nil {
		return m.Salutations
	}
	return nil
}

func (m *FndHellosReq) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FndHellosReq) GetLapse() uint32 {
	if m != nil {
		return m.Lapse
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloResp)(nil), "pb.HelloResp")
	proto.RegisterType((*AddHelloReq)(nil), "pb.AddHelloReq")
	proto.RegisterType((*OvrHelloReq)(nil), "pb.OvrHelloReq")
	proto.RegisterType((*RmvHelloResp)(nil), "pb.RmvHelloResp")
	proto.RegisterType((*RmvHelloReq)(nil), "pb.RmvHelloReq")
	proto.RegisterType((*HellosResp)(nil), "pb.HellosResp")
	proto.RegisterType((*FndHellosReq)(nil), "pb.FndHellosReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	AddHello(ctx context.Context, in *AddHelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	OvrHello(ctx context.Context, in *OvrHelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	RmvHello(ctx context.Context, in *RmvHelloReq, opts ...grpc.CallOption) (*RmvHelloResp, error)
	FndHellos(ctx context.Context, in *FndHellosReq, opts ...grpc.CallOption) (*HellosResp, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) AddHello(ctx context.Context, in *AddHelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/pb.Hello/AddHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) OvrHello(ctx context.Context, in *OvrHelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/pb.Hello/OvrHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) RmvHello(ctx context.Context, in *RmvHelloReq, opts ...grpc.CallOption) (*RmvHelloResp, error) {
	out := new(RmvHelloResp)
	err := c.cc.Invoke(ctx, "/pb.Hello/RmvHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) FndHellos(ctx context.Context, in *FndHellosReq, opts ...grpc.CallOption) (*HellosResp, error) {
	out := new(HellosResp)
	err := c.cc.Invoke(ctx, "/pb.Hello/FndHellos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	AddHello(context.Context, *AddHelloReq) (*HelloResp, error)
	OvrHello(context.Context, *OvrHelloReq) (*HelloResp, error)
	RmvHello(context.Context, *RmvHelloReq) (*RmvHelloResp, error)
	FndHellos(context.Context, *FndHellosReq) (*HellosResp, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_AddHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).AddHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hello/AddHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).AddHello(ctx, req.(*AddHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_OvrHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OvrHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).OvrHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hello/OvrHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).OvrHello(ctx, req.(*OvrHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_RmvHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmvHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).RmvHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hello/RmvHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).RmvHello(ctx, req.(*RmvHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_FndHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FndHellosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).FndHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hello/FndHellos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).FndHellos(ctx, req.(*FndHellosReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHello",
			Handler:    _Hello_AddHello_Handler,
		},
		{
			MethodName: "OvrHello",
			Handler:    _Hello_OvrHello_Handler,
		},
		{
			MethodName: "RmvHello",
			Handler:    _Hello_RmvHello_Handler,
		},
		{
			MethodName: "FndHellos",
			Handler:    _Hello_FndHellos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_062e3cf3bf8b761a) }

var fileDescriptor_hello_062e3cf3bf8b761a = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x94, 0xb6, 0x9b, 0xf1, 0xb6, 0x5b, 0x0c, 0x42, 0x51, 0xc4, 0x47, 0x31, 0x97,
	0x6a, 0x0f, 0x89, 0x54, 0x38, 0x21, 0x81, 0x80, 0x03, 0xcb, 0x0d, 0xc9, 0xe2, 0x05, 0x92, 0xb5,
	0x59, 0x8c, 0xf2, 0xe1, 0xd6, 0xee, 0x5e, 0x10, 0x17, 0x1e, 0x80, 0x0b, 0x8f, 0xc6, 0x2b, 0xf0,
	0x0e, 0x5c, 0x91, 0xc7, 0x49, 0x36, 0x2d, 0xd2, 0xf6, 0x96, 0xf9, 0xf8, 0xff, 0x32, 0xf3, 0x1f,
	0x03, 0xf9, 0x22, 0xcb, 0xb2, 0x49, 0xf5, 0xb6, 0xb1, 0x0d, 0x0d, 0x75, 0x91, 0x3c, 0xbc, 0x6a,
	0x9a, 0xab, 0x52, 0x66, 0xb9, 0x56, 0x59, 0x5e, 0xd7, 0x8d, 0xcd, 0xad, 0x6a, 0x6a, 0xe3, 0x3b,
	0x92, 0x27, 0x6d, 0x15, 0xa3, 0x62, 0xf7, 0x39, 0xb3, 0xaa, 0x92, 0xc6, 0xe6, 0x95, 0xf6, 0x0d,
	0xec, 0x6f, 0x00, 0xd1, 0x07, 0x87, 0xe4, 0xd2, 0x68, 0x3a, 0x87, 0x50, 0x89, 0x38, 0x58, 0x06,
	0xab, 0x88, 0x87, 0x4a, 0xd0, 0xc7, 0x00, 0x26, 0x2f, 0x77, 0x9e, 0x19, 0x87, 0x98, 0x1f, 0x64,
	0x68, 0x0c, 0x53, 0xb3, 0x2b, 0xbe, 0xca, 0x4b, 0x1b, 0x8f, 0xb0, 0xd8, 0x85, 0xf4, 0x05, 0x4c,
	0x2f, 0xb7, 0x32, 0xb7, 0x52, 0xc4, 0x77, 0x96, 0xc1, 0x8a, 0xac, 0x93, 0xd4, 0x8f, 0x92, 0x76,
	0xa3, 0xa4, 0x9f, 0xba, 0x51, 0x78, 0xd7, 0xea, 0x54, 0x3b, 0x2d, 0x50, 0x35, 0x3e, 0xae, 0x6a,
	0x5b, 0x9d, 0x4a, 0xc8, 0x52, 0x3a, 0xd5, 0xe4, 0xb8, 0xaa, 0x6d, 0x65, 0x17, 0x40, 0xde, 0x0a,
	0xd1, 0xee, 0xbe, 0x39, 0x58, 0x35, 0xb8, 0x6d, 0xd5, 0x70, 0x6f, 0x55, 0xf6, 0x06, 0xc8, 0xc7,
	0xeb, 0x6d, 0x0f, 0x3a, 0xf4, 0xf0, 0x29, 0x8c, 0xb6, 0x72, 0x83, 0x22, 0xb2, 0x3e, 0x4b, 0x75,
	0x91, 0x0e, 0x7e, 0xcb, 0x5d, 0x8d, 0xcd, 0xe1, 0x94, 0x57, 0xd7, 0xfd, 0x19, 0xd8, 0x23, 0x20,
	0x37, 0xf1, 0x7f, 0x44, 0x76, 0x01, 0x80, 0x35, 0x83, 0x37, 0x7b, 0x06, 0x63, 0x65, 0x65, 0x65,
	0xe2, 0x60, 0x39, 0x5a, 0x91, 0xf5, 0xcc, 0xfd, 0xa1, 0x47, 0x71, 0x5f, 0xa3, 0xf7, 0x61, 0x6c,
	0x1b, 0x9b, 0x97, 0x38, 0xc6, 0x8c, 0xfb, 0x80, 0xd5, 0x70, 0xfa, 0xbe, 0x16, 0x1d, 0x6b, 0x43,
	0x17, 0x30, 0x52, 0xc2, 0x83, 0x22, 0xee, 0x3e, 0xe9, 0x12, 0xc8, 0x8d, 0x07, 0x26, 0x0e, 0xb1,
	0x32, 0x4c, 0x39, 0x72, 0xa9, 0x2a, 0xe5, 0x1f, 0xc0, 0x8c, 0xfb, 0x00, 0xb3, 0xb9, 0x36, 0x12,
	0x8f, 0xef, 0xb2, 0x2e, 0x58, 0xff, 0x0c, 0x61, 0x8c, 0x7f, 0xa3, 0xaf, 0xe0, 0xa4, 0x73, 0x81,
	0x1e, 0x7a, 0x92, 0xec, 0xaf, 0xc0, 0xee, 0xfe, 0xf8, 0xfd, 0xe7, 0x57, 0x48, 0xd8, 0x24, 0xc3,
	0xb7, 0xff, 0x32, 0x38, 0xa7, 0xef, 0xe0, 0xa4, 0xb3, 0xdc, 0xcb, 0x07, 0x07, 0x38, 0x94, 0x3f,
	0x40, 0xf9, 0x22, 0x21, 0x5e, 0x9e, 0x7d, 0x53, 0xe2, 0x7b, 0xcb, 0xe8, 0x4c, 0xf6, 0x8c, 0x81,
	0xe5, 0xc9, 0x62, 0x3f, 0x61, 0x34, 0xbb, 0x87, 0x98, 0xd9, 0xf9, 0x10, 0x43, 0x5f, 0x43, 0xd4,
	0x1b, 0x48, 0x51, 0x33, 0xf4, 0x33, 0x99, 0xf7, 0x93, 0xe0, 0xa9, 0xd8, 0x19, 0x32, 0x22, 0x3a,
	0xf5, 0x0c, 0x53, 0x4c, 0xf0, 0x81, 0x3e, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x50, 0xd8,
	0x2f, 0xd6, 0x03, 0x00, 0x00,
}