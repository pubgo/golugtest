// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: proto/hello/transport.proto

package hello

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_hello_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_proto_hello_transport_proto protoreflect.FileDescriptor

var file_proto_hello_transport_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0xd6, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x32, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x31, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x32, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x33, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_transport_proto_rawDescOnce sync.Once
	file_proto_hello_transport_proto_rawDescData = file_proto_hello_transport_proto_rawDesc
)

func file_proto_hello_transport_proto_rawDescGZIP() []byte {
	file_proto_hello_transport_proto_rawDescOnce.Do(func() {
		file_proto_hello_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_transport_proto_rawDescData)
	})
	return file_proto_hello_transport_proto_rawDescData
}

var file_proto_hello_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_hello_transport_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: hello.Message
	nil,             // 1: hello.Message.HeaderEntry
}
var file_proto_hello_transport_proto_depIdxs = []int32{
	1, // 0: hello.Message.header:type_name -> hello.Message.HeaderEntry
	0, // 1: hello.Transport.TestStream:input_type -> hello.Message
	0, // 2: hello.Transport.TestStream1:input_type -> hello.Message
	0, // 3: hello.Transport.TestStream2:input_type -> hello.Message
	0, // 4: hello.Transport.TestStream3:input_type -> hello.Message
	0, // 5: hello.Transport.TestStream:output_type -> hello.Message
	0, // 6: hello.Transport.TestStream1:output_type -> hello.Message
	0, // 7: hello.Transport.TestStream2:output_type -> hello.Message
	0, // 8: hello.Transport.TestStream3:output_type -> hello.Message
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_hello_transport_proto_init() }
func file_proto_hello_transport_proto_init() {
	if File_proto_hello_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_proto_hello_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_transport_proto_goTypes,
		DependencyIndexes: file_proto_hello_transport_proto_depIdxs,
		MessageInfos:      file_proto_hello_transport_proto_msgTypes,
	}.Build()
	File_proto_hello_transport_proto = out.File
	file_proto_hello_transport_proto_rawDesc = nil
	file_proto_hello_transport_proto_goTypes = nil
	file_proto_hello_transport_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransportClient is the client API for Transport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransportClient interface {
	TestStream(ctx context.Context, opts ...grpc.CallOption) (Transport_TestStreamClient, error)
	TestStream1(ctx context.Context, opts ...grpc.CallOption) (Transport_TestStream1Client, error)
	TestStream2(ctx context.Context, in *Message, opts ...grpc.CallOption) (Transport_TestStream2Client, error)
	TestStream3(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type transportClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportClient(cc grpc.ClientConnInterface) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) TestStream(ctx context.Context, opts ...grpc.CallOption) (Transport_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transport_serviceDesc.Streams[0], "/hello.Transport/TestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportTestStreamClient{stream}
	return x, nil
}

type Transport_TestStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type transportTestStreamClient struct {
	grpc.ClientStream
}

func (x *transportTestStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transportTestStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transportClient) TestStream1(ctx context.Context, opts ...grpc.CallOption) (Transport_TestStream1Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Transport_serviceDesc.Streams[1], "/hello.Transport/TestStream1", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportTestStream1Client{stream}
	return x, nil
}

type Transport_TestStream1Client interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type transportTestStream1Client struct {
	grpc.ClientStream
}

func (x *transportTestStream1Client) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transportTestStream1Client) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transportClient) TestStream2(ctx context.Context, in *Message, opts ...grpc.CallOption) (Transport_TestStream2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Transport_serviceDesc.Streams[2], "/hello.Transport/TestStream2", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportTestStream2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transport_TestStream2Client interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type transportTestStream2Client struct {
	grpc.ClientStream
}

func (x *transportTestStream2Client) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transportClient) TestStream3(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/hello.Transport/TestStream3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServer is the server API for Transport service.
type TransportServer interface {
	TestStream(Transport_TestStreamServer) error
	TestStream1(Transport_TestStream1Server) error
	TestStream2(*Message, Transport_TestStream2Server) error
	TestStream3(context.Context, *Message) (*Message, error)
}

// UnimplementedTransportServer can be embedded to have forward compatible implementations.
type UnimplementedTransportServer struct {
}

func (*UnimplementedTransportServer) TestStream(Transport_TestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}
func (*UnimplementedTransportServer) TestStream1(Transport_TestStream1Server) error {
	return status.Errorf(codes.Unimplemented, "method TestStream1 not implemented")
}
func (*UnimplementedTransportServer) TestStream2(*Message, Transport_TestStream2Server) error {
	return status.Errorf(codes.Unimplemented, "method TestStream2 not implemented")
}
func (*UnimplementedTransportServer) TestStream3(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestStream3 not implemented")
}

func RegisterTransportServer(s *grpc.Server, srv TransportServer) {
	s.RegisterService(&_Transport_serviceDesc, srv)
}

func _Transport_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransportServer).TestStream(&transportTestStreamServer{stream})
}

type Transport_TestStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type transportTestStreamServer struct {
	grpc.ServerStream
}

func (x *transportTestStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transportTestStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Transport_TestStream1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransportServer).TestStream1(&transportTestStream1Server{stream})
}

type Transport_TestStream1Server interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type transportTestStream1Server struct {
	grpc.ServerStream
}

func (x *transportTestStream1Server) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transportTestStream1Server) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Transport_TestStream2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransportServer).TestStream2(m, &transportTestStream2Server{stream})
}

type Transport_TestStream2Server interface {
	Send(*Message) error
	grpc.ServerStream
}

type transportTestStream2Server struct {
	grpc.ServerStream
}

func (x *transportTestStream2Server) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Transport_TestStream3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServer).TestStream3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Transport/TestStream3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServer).TestStream3(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Transport",
	HandlerType: (*TransportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestStream3",
			Handler:    _Transport_TestStream3_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStream",
			Handler:       _Transport_TestStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TestStream1",
			Handler:       _Transport_TestStream1_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TestStream2",
			Handler:       _Transport_TestStream2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/hello/transport.proto",
}
