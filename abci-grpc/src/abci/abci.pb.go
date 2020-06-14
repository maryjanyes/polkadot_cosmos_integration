// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: abci.proto

package abci

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TxResponse) Reset() {
	*x = TxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResponse) ProtoMessage() {}

func (x *TxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_abci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResponse.ProtoReflect.Descriptor instead.
func (*TxResponse) Descriptor() ([]byte, []int) {
	return file_abci_proto_rawDescGZIP(), []int{0}
}

type CheckTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *CheckTxRequest) Reset() {
	*x = CheckTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTxRequest) ProtoMessage() {}

func (x *CheckTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_abci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTxRequest.ProtoReflect.Descriptor instead.
func (*CheckTxRequest) Descriptor() ([]byte, []int) {
	return file_abci_proto_rawDescGZIP(), []int{1}
}

func (x *CheckTxRequest) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

type DeliverTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *DeliverTxRequest) Reset() {
	*x = DeliverTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverTxRequest) ProtoMessage() {}

func (x *DeliverTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_abci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverTxRequest.ProtoReflect.Descriptor instead.
func (*DeliverTxRequest) Descriptor() ([]byte, []int) {
	return file_abci_proto_rawDescGZIP(), []int{2}
}

func (x *DeliverTxRequest) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

type EchoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoMessage) Reset() {
	*x = EchoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abci_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoMessage) ProtoMessage() {}

func (x *EchoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_abci_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoMessage.ProtoReflect.Descriptor instead.
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return file_abci_proto_rawDescGZIP(), []int{3}
}

var File_abci_proto protoreflect.FileDescriptor

var file_abci_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x62,
	0x63, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0c, 0x0a, 0x0a, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x78,
	0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x78, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xf1, 0x01, 0x0a, 0x04, 0x41, 0x62, 0x63, 0x69, 0x12, 0x4e, 0x0a, 0x07,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x12, 0x14, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x61, 0x62, 0x63, 0x69, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x12, 0x16, 0x2e, 0x61, 0x62, 0x63, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x62,
	0x63, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x3a,
	0x01, 0x2a, 0x12, 0x43, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x61, 0x62, 0x63,
	0x69, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e,
	0x61, 0x62, 0x63, 0x69, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x45, 0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abci_proto_rawDescOnce sync.Once
	file_abci_proto_rawDescData = file_abci_proto_rawDesc
)

func file_abci_proto_rawDescGZIP() []byte {
	file_abci_proto_rawDescOnce.Do(func() {
		file_abci_proto_rawDescData = protoimpl.X.CompressGZIP(file_abci_proto_rawDescData)
	})
	return file_abci_proto_rawDescData
}

var file_abci_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_abci_proto_goTypes = []interface{}{
	(*TxResponse)(nil),       // 0: abci.TxResponse
	(*CheckTxRequest)(nil),   // 1: abci.CheckTxRequest
	(*DeliverTxRequest)(nil), // 2: abci.DeliverTxRequest
	(*EchoMessage)(nil),      // 3: abci.EchoMessage
}
var file_abci_proto_depIdxs = []int32{
	1, // 0: abci.Abci.CheckTx:input_type -> abci.CheckTxRequest
	2, // 1: abci.Abci.DeliverTx:input_type -> abci.DeliverTxRequest
	3, // 2: abci.Abci.Echo:input_type -> abci.EchoMessage
	0, // 3: abci.Abci.CheckTx:output_type -> abci.TxResponse
	0, // 4: abci.Abci.DeliverTx:output_type -> abci.TxResponse
	3, // 5: abci.Abci.Echo:output_type -> abci.EchoMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_abci_proto_init() }
func file_abci_proto_init() {
	if File_abci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_abci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResponse); i {
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
		file_abci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTxRequest); i {
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
		file_abci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverTxRequest); i {
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
		file_abci_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoMessage); i {
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
			RawDescriptor: file_abci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_abci_proto_goTypes,
		DependencyIndexes: file_abci_proto_depIdxs,
		MessageInfos:      file_abci_proto_msgTypes,
	}.Build()
	File_abci_proto = out.File
	file_abci_proto_rawDesc = nil
	file_abci_proto_goTypes = nil
	file_abci_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AbciClient is the client API for Abci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AbciClient interface {
	CheckTx(ctx context.Context, in *CheckTxRequest, opts ...grpc.CallOption) (*TxResponse, error)
	DeliverTx(ctx context.Context, in *DeliverTxRequest, opts ...grpc.CallOption) (*TxResponse, error)
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type abciClient struct {
	cc grpc.ClientConnInterface
}

func NewAbciClient(cc grpc.ClientConnInterface) AbciClient {
	return &abciClient{cc}
}

func (c *abciClient) CheckTx(ctx context.Context, in *CheckTxRequest, opts ...grpc.CallOption) (*TxResponse, error) {
	out := new(TxResponse)
	err := c.cc.Invoke(ctx, "/abci.Abci/CheckTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abciClient) DeliverTx(ctx context.Context, in *DeliverTxRequest, opts ...grpc.CallOption) (*TxResponse, error) {
	out := new(TxResponse)
	err := c.cc.Invoke(ctx, "/abci.Abci/DeliverTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abciClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/abci.Abci/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AbciServer is the server API for Abci service.
type AbciServer interface {
	CheckTx(context.Context, *CheckTxRequest) (*TxResponse, error)
	DeliverTx(context.Context, *DeliverTxRequest) (*TxResponse, error)
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
}

// UnimplementedAbciServer can be embedded to have forward compatible implementations.
type UnimplementedAbciServer struct {
}

func (*UnimplementedAbciServer) CheckTx(context.Context, *CheckTxRequest) (*TxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTx not implemented")
}
func (*UnimplementedAbciServer) DeliverTx(context.Context, *DeliverTxRequest) (*TxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTx not implemented")
}
func (*UnimplementedAbciServer) Echo(context.Context, *EchoMessage) (*EchoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterAbciServer(s *grpc.Server, srv AbciServer) {
	s.RegisterService(&_Abci_serviceDesc, srv)
}

func _Abci_CheckTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbciServer).CheckTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abci.Abci/CheckTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbciServer).CheckTx(ctx, req.(*CheckTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Abci_DeliverTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbciServer).DeliverTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abci.Abci/DeliverTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbciServer).DeliverTx(ctx, req.(*DeliverTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Abci_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbciServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abci.Abci/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbciServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Abci_serviceDesc = grpc.ServiceDesc{
	ServiceName: "abci.Abci",
	HandlerType: (*AbciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckTx",
			Handler:    _Abci_CheckTx_Handler,
		},
		{
			MethodName: "DeliverTx",
			Handler:    _Abci_DeliverTx_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Abci_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "abci.proto",
}
