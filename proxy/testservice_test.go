// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.11.4
// source: test.proto

package proxy

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainData   []byte `protobuf:"bytes,1,opt,name=main_data,json=mainData,proto3" json:"main_data,omitempty"`
	OtherThing int32  `protobuf:"varint,2,opt,name=other_thing,json=otherThing,proto3" json:"other_thing,omitempty"`
	Toggle     bool   `protobuf:"varint,3,opt,name=toggle,proto3" json:"toggle,omitempty"`
}

func (x *ExampleRequest) Reset() {
	*x = ExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleRequest) ProtoMessage() {}

func (x *ExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleRequest.ProtoReflect.Descriptor instead.
func (*ExampleRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleRequest) GetMainData() []byte {
	if x != nil {
		return x.MainData
	}
	return nil
}

func (x *ExampleRequest) GetOtherThing() int32 {
	if x != nil {
		return x.OtherThing
	}
	return 0
}

func (x *ExampleRequest) GetToggle() bool {
	if x != nil {
		return x.Toggle
	}
	return false
}

type ExampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullResponseData string `protobuf:"bytes,1,opt,name=full_response_data,json=fullResponseData,proto3" json:"full_response_data,omitempty"`
	Done             bool   `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	Output           []byte `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ExampleResponse) Reset() {
	*x = ExampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleResponse) ProtoMessage() {}

func (x *ExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleResponse.ProtoReflect.Descriptor instead.
func (*ExampleResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleResponse) GetFullResponseData() string {
	if x != nil {
		return x.FullResponseData
	}
	return ""
}

func (x *ExampleResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *ExampleResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x22, 0x66, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x22, 0x6b, 0x0a, 0x0f, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x43, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x4c, 0x4b, 0x65,
	0x6e, 0x6e, 0x65, 0x64, 0x79, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*ExampleRequest)(nil),  // 0: proxy.ExampleRequest
	(*ExampleResponse)(nil), // 1: proxy.ExampleResponse
}
var file_test_proto_depIdxs = []int32{
	0, // 0: proxy.Example.Example:input_type -> proxy.ExampleRequest
	1, // 1: proxy.Example.Example:output_type -> proxy.ExampleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	Example(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Example(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error) {
	out := new(ExampleResponse)
	err := c.cc.Invoke(ctx, "/proxy.Example/Example", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	Example(context.Context, *ExampleRequest) (*ExampleResponse, error)
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Example(context.Context, *ExampleRequest) (*ExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Example not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Example_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Example(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.Example/Example",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Example(ctx, req.(*ExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Example",
			Handler:    _Example_Example_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
