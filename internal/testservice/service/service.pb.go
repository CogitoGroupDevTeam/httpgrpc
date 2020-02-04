// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

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

type FibonacciRequest struct {
	// The 1-indexed point in the Fibonacci sequence
	N                    uint64   `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetN() uint64 {
	if m != nil {
		return m.N
	}
	return 0
}

type FibonacciResponse struct {
	// The number found in the nth place of the sequence
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type RandomRequest struct {
	// The lowest inclusive integer for the resulting random number
	LowerBound int64 `protobuf:"varint,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	// The highest inclusive integer for the resulting random number
	UpperBound           int64    `protobuf:"varint,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomRequest) Reset()         { *m = RandomRequest{} }
func (m *RandomRequest) String() string { return proto.CompactTextString(m) }
func (*RandomRequest) ProtoMessage()    {}
func (*RandomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RandomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomRequest.Unmarshal(m, b)
}
func (m *RandomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomRequest.Marshal(b, m, deterministic)
}
func (m *RandomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomRequest.Merge(m, src)
}
func (m *RandomRequest) XXX_Size() int {
	return xxx_messageInfo_RandomRequest.Size(m)
}
func (m *RandomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RandomRequest proto.InternalMessageInfo

func (m *RandomRequest) GetLowerBound() int64 {
	if m != nil {
		return m.LowerBound
	}
	return 0
}

func (m *RandomRequest) GetUpperBound() int64 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

type RandomResponse struct {
	// The generated number
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomResponse) Reset()         { *m = RandomResponse{} }
func (m *RandomResponse) String() string { return proto.CompactTextString(m) }
func (*RandomResponse) ProtoMessage()    {}
func (*RandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RandomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomResponse.Unmarshal(m, b)
}
func (m *RandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomResponse.Marshal(b, m, deterministic)
}
func (m *RandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomResponse.Merge(m, src)
}
func (m *RandomResponse) XXX_Size() int {
	return xxx_messageInfo_RandomResponse.Size(m)
}
func (m *RandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RandomResponse proto.InternalMessageInfo

func (m *RandomResponse) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type UploadPhotoRequest struct {
	// The raw bytes of the photo
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoRequest) Reset()         { *m = UploadPhotoRequest{} }
func (m *UploadPhotoRequest) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoRequest) ProtoMessage()    {}
func (*UploadPhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *UploadPhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoRequest.Unmarshal(m, b)
}
func (m *UploadPhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoRequest.Marshal(b, m, deterministic)
}
func (m *UploadPhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoRequest.Merge(m, src)
}
func (m *UploadPhotoRequest) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoRequest.Size(m)
}
func (m *UploadPhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoRequest proto.InternalMessageInfo

func (m *UploadPhotoRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadPhotoResponse struct {
	// The uuid generated to identify and retreive this photo
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoResponse) Reset()         { *m = UploadPhotoResponse{} }
func (m *UploadPhotoResponse) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoResponse) ProtoMessage()    {}
func (*UploadPhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UploadPhotoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoResponse.Unmarshal(m, b)
}
func (m *UploadPhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoResponse.Marshal(b, m, deterministic)
}
func (m *UploadPhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoResponse.Merge(m, src)
}
func (m *UploadPhotoResponse) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoResponse.Size(m)
}
func (m *UploadPhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoResponse proto.InternalMessageInfo

func (m *UploadPhotoResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*FibonacciRequest)(nil), "service.FibonacciRequest")
	proto.RegisterType((*FibonacciResponse)(nil), "service.FibonacciResponse")
	proto.RegisterType((*RandomRequest)(nil), "service.RandomRequest")
	proto.RegisterType((*RandomResponse)(nil), "service.RandomResponse")
	proto.RegisterType((*UploadPhotoRequest)(nil), "service.UploadPhotoRequest")
	proto.RegisterType((*UploadPhotoResponse)(nil), "service.UploadPhotoResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0xff, 0xf9, 0xb7, 0x54, 0xe6, 0xb6, 0xf5, 0x23, 0x42, 0xd5, 0x51, 0xb0, 0x64, 0x35,
	0x22, 0x74, 0x51, 0x97, 0x82, 0xa0, 0xf8, 0x01, 0x2e, 0xa4, 0x06, 0x5c, 0xcb, 0x7c, 0x04, 0x1c,
	0x68, 0x73, 0xe3, 0x24, 0x51, 0x1f, 0xc6, 0xc7, 0xf2, 0x81, 0x64, 0x32, 0x33, 0xa1, 0xb5, 0xb3,
	0x12, 0x77, 0xc9, 0xc9, 0x6f, 0xce, 0x3d, 0xf7, 0x30, 0x30, 0xd4, 0xa2, 0x78, 0xcb, 0x53, 0x31,
	0x51, 0x05, 0x1a, 0xa4, 0x1b, 0xf5, 0x95, 0x8d, 0x61, 0xfb, 0x36, 0x4f, 0x50, 0xc6, 0x69, 0x9a,
	0x73, 0xf1, 0x6a, 0x85, 0x36, 0x74, 0x00, 0x44, 0xee, 0x93, 0x31, 0x89, 0xba, 0x9c, 0x48, 0x76,
	0x0a, 0x3b, 0x4b, 0x84, 0x56, 0x28, 0xb5, 0xa0, 0x23, 0xe8, 0x49, 0xbb, 0x48, 0x44, 0x51, 0x73,
	0xf5, 0x8d, 0x3d, 0xc2, 0x90, 0xc7, 0x32, 0xc3, 0x45, 0xe3, 0x75, 0x0c, 0xfd, 0x39, 0xbe, 0x8b,
	0xe2, 0x39, 0x41, 0x2b, 0x33, 0x47, 0x77, 0x38, 0x38, 0xe9, 0xaa, 0x54, 0x4a, 0xc0, 0x2a, 0xe5,
	0x81, 0xff, 0x15, 0xe0, 0x24, 0x07, 0xb0, 0x08, 0x36, 0x1b, 0xcb, 0xd6, 0xe1, 0x1d, 0x3f, 0x3c,
	0x02, 0xfa, 0xa4, 0xe6, 0x18, 0x67, 0xb3, 0x17, 0x34, 0xd8, 0x24, 0xa0, 0xd0, 0xcd, 0x62, 0x13,
	0x3b, 0x76, 0xc0, 0xdd, 0x99, 0x9d, 0xc0, 0xee, 0x0a, 0x59, 0x1b, 0x53, 0xe8, 0x5a, 0x9b, 0x57,
	0x29, 0x03, 0xee, 0xce, 0xd3, 0x2f, 0x02, 0x9d, 0x4b, 0xa5, 0xe8, 0x35, 0x04, 0xbe, 0x06, 0x7a,
	0x30, 0x69, 0xea, 0xfc, 0x59, 0x5e, 0x18, 0xb6, 0x3d, 0x55, 0xfe, 0xec, 0x1f, 0x3d, 0x87, 0x5e,
	0xb5, 0x0c, 0x1d, 0x79, 0x6e, 0xa5, 0xb0, 0x70, 0x6f, 0x4d, 0xf7, 0x1f, 0xdf, 0x43, 0x7f, 0x29,
	0x35, 0x3d, 0xf4, 0xe4, 0xfa, 0xd6, 0xe1, 0x51, 0xfb, 0x63, 0xe3, 0x35, 0xfd, 0x24, 0x00, 0x37,
	0x1f, 0x0a, 0xb5, 0xc8, 0xca, 0xed, 0x2e, 0x20, 0xb8, 0x13, 0xe6, 0xf7, 0xd1, 0x1e, 0x60, 0x6b,
	0x86, 0xda, 0xfc, 0x55, 0xbc, 0xa4, 0xe7, 0x7e, 0xd3, 0xb3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x68, 0xf6, 0xa7, 0xf4, 0xb7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error)
	// Random returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
	// UploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	UploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error) {
	out := new(FibonacciResponse)
	err := c.cc.Invoke(ctx, "/service.App/Fibonacci", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/service.App/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error) {
	out := new(UploadPhotoResponse)
	err := c.cc.Invoke(ctx, "/service.App/UploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(context.Context, *FibonacciRequest) (*FibonacciResponse, error)
	// Random returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	Random(context.Context, *RandomRequest) (*RandomResponse, error)
	// UploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	UploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) Fibonacci(ctx context.Context, req *FibonacciRequest) (*FibonacciResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (*UnimplementedAppServer) Random(ctx context.Context, req *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (*UnimplementedAppServer) UploadPhoto(ctx context.Context, req *UploadPhotoRequest) (*UploadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPhoto not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Fibonacci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibonacciRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Fibonacci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/Fibonacci",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Fibonacci(ctx, req.(*FibonacciRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Random(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/UploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UploadPhoto(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibonacci",
			Handler:    _App_Fibonacci_Handler,
		},
		{
			MethodName: "Random",
			Handler:    _App_Random_Handler,
		},
		{
			MethodName: "UploadPhoto",
			Handler:    _App_UploadPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// ExposedAppClient is the client API for ExposedApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExposedAppClient interface {
	GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
	PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error)
}

type exposedAppClient struct {
	cc *grpc.ClientConn
}

func NewExposedAppClient(cc *grpc.ClientConn) ExposedAppClient {
	return &exposedAppClient{cc}
}

func (c *exposedAppClient) GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/service.ExposedApp/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposedAppClient) PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error) {
	out := new(UploadPhotoResponse)
	err := c.cc.Invoke(ctx, "/service.ExposedApp/PostUploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExposedAppServer is the server API for ExposedApp service.
type ExposedAppServer interface {
	GetRandom(context.Context, *RandomRequest) (*RandomResponse, error)
	PostUploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error)
}

// UnimplementedExposedAppServer can be embedded to have forward compatible implementations.
type UnimplementedExposedAppServer struct {
}

func (*UnimplementedExposedAppServer) GetRandom(ctx context.Context, req *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}
func (*UnimplementedExposedAppServer) PostUploadPhoto(ctx context.Context, req *UploadPhotoRequest) (*UploadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUploadPhoto not implemented")
}

func RegisterExposedAppServer(s *grpc.Server, srv ExposedAppServer) {
	s.RegisterService(&_ExposedApp_serviceDesc, srv)
}

func _ExposedApp_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposedAppServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExposedApp/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposedAppServer).GetRandom(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExposedApp_PostUploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposedAppServer).PostUploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExposedApp/PostUploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposedAppServer).PostUploadPhoto(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExposedApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ExposedApp",
	HandlerType: (*ExposedAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandom",
			Handler:    _ExposedApp_GetRandom_Handler,
		},
		{
			MethodName: "PostUploadPhoto",
			Handler:    _ExposedApp_PostUploadPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
