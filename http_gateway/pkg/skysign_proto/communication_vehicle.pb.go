// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: communication_vehicle.proto

package skysign_proto

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

type PushTelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PushTelemetryRequest) Reset() {
	*x = PushTelemetryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTelemetryRequest) ProtoMessage() {}

func (x *PushTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTelemetryRequest.ProtoReflect.Descriptor instead.
func (*PushTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_communication_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *PushTelemetryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PushTelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PushTelemetryResponse) Reset() {
	*x = PushTelemetryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTelemetryResponse) ProtoMessage() {}

func (x *PushTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTelemetryResponse.ProtoReflect.Descriptor instead.
func (*PushTelemetryResponse) Descriptor() ([]byte, []int) {
	return file_communication_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *PushTelemetryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PullCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommId string `protobuf:"bytes,1,opt,name=commId,proto3" json:"commId,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PullCommandRequest) Reset() {
	*x = PullCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_vehicle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCommandRequest) ProtoMessage() {}

func (x *PullCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_vehicle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullCommandRequest.ProtoReflect.Descriptor instead.
func (*PullCommandRequest) Descriptor() ([]byte, []int) {
	return file_communication_vehicle_proto_rawDescGZIP(), []int{2}
}

func (x *PullCommandRequest) GetCommId() string {
	if x != nil {
		return x.CommId
	}
	return ""
}

func (x *PullCommandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PullCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommId string `protobuf:"bytes,1,opt,name=commId,proto3" json:"commId,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PullCommandResponse) Reset() {
	*x = PullCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_vehicle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCommandResponse) ProtoMessage() {}

func (x *PullCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_vehicle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullCommandResponse.ProtoReflect.Descriptor instead.
func (*PullCommandResponse) Descriptor() ([]byte, []int) {
	return file_communication_vehicle_proto_rawDescGZIP(), []int{3}
}

func (x *PullCommandResponse) GetCommId() string {
	if x != nil {
		return x.CommId
	}
	return ""
}

func (x *PullCommandResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_communication_vehicle_proto protoreflect.FileDescriptor

var file_communication_vehicle_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x75,
	0x73, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x50,
	0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x13, 0x50, 0x75, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbf, 0x02, 0x0a, 0x1b, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x73,
	0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x50, 0x75,
	0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_vehicle_proto_rawDescOnce sync.Once
	file_communication_vehicle_proto_rawDescData = file_communication_vehicle_proto_rawDesc
)

func file_communication_vehicle_proto_rawDescGZIP() []byte {
	file_communication_vehicle_proto_rawDescOnce.Do(func() {
		file_communication_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_vehicle_proto_rawDescData)
	})
	return file_communication_vehicle_proto_rawDescData
}

var file_communication_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communication_vehicle_proto_goTypes = []interface{}{
	(*PushTelemetryRequest)(nil),  // 0: skysign_proto.PushTelemetryRequest
	(*PushTelemetryResponse)(nil), // 1: skysign_proto.PushTelemetryResponse
	(*PullCommandRequest)(nil),    // 2: skysign_proto.PullCommandRequest
	(*PullCommandResponse)(nil),   // 3: skysign_proto.PullCommandResponse
}
var file_communication_vehicle_proto_depIdxs = []int32{
	0, // 0: skysign_proto.CommunicationVehicleService.PushTelemetry:input_type -> skysign_proto.PushTelemetryRequest
	2, // 1: skysign_proto.CommunicationVehicleService.PullCommand:input_type -> skysign_proto.PullCommandRequest
	1, // 2: skysign_proto.CommunicationVehicleService.PushTelemetry:output_type -> skysign_proto.PushTelemetryResponse
	3, // 3: skysign_proto.CommunicationVehicleService.PullCommand:output_type -> skysign_proto.PullCommandResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_communication_vehicle_proto_init() }
func file_communication_vehicle_proto_init() {
	if File_communication_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTelemetryRequest); i {
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
		file_communication_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTelemetryResponse); i {
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
		file_communication_vehicle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullCommandRequest); i {
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
		file_communication_vehicle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullCommandResponse); i {
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
			RawDescriptor: file_communication_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_vehicle_proto_goTypes,
		DependencyIndexes: file_communication_vehicle_proto_depIdxs,
		MessageInfos:      file_communication_vehicle_proto_msgTypes,
	}.Build()
	File_communication_vehicle_proto = out.File
	file_communication_vehicle_proto_rawDesc = nil
	file_communication_vehicle_proto_goTypes = nil
	file_communication_vehicle_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommunicationVehicleServiceClient is the client API for CommunicationVehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationVehicleServiceClient interface {
	PushTelemetry(ctx context.Context, in *PushTelemetryRequest, opts ...grpc.CallOption) (*PushTelemetryResponse, error)
	PullCommand(ctx context.Context, in *PullCommandRequest, opts ...grpc.CallOption) (*PullCommandResponse, error)
}

type communicationVehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationVehicleServiceClient(cc grpc.ClientConnInterface) CommunicationVehicleServiceClient {
	return &communicationVehicleServiceClient{cc}
}

func (c *communicationVehicleServiceClient) PushTelemetry(ctx context.Context, in *PushTelemetryRequest, opts ...grpc.CallOption) (*PushTelemetryResponse, error) {
	out := new(PushTelemetryResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationVehicleService/PushTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationVehicleServiceClient) PullCommand(ctx context.Context, in *PullCommandRequest, opts ...grpc.CallOption) (*PullCommandResponse, error) {
	out := new(PullCommandResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationVehicleService/PullCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationVehicleServiceServer is the server API for CommunicationVehicleService service.
type CommunicationVehicleServiceServer interface {
	PushTelemetry(context.Context, *PushTelemetryRequest) (*PushTelemetryResponse, error)
	PullCommand(context.Context, *PullCommandRequest) (*PullCommandResponse, error)
}

// UnimplementedCommunicationVehicleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationVehicleServiceServer struct {
}

func (*UnimplementedCommunicationVehicleServiceServer) PushTelemetry(context.Context, *PushTelemetryRequest) (*PushTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTelemetry not implemented")
}
func (*UnimplementedCommunicationVehicleServiceServer) PullCommand(context.Context, *PullCommandRequest) (*PullCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullCommand not implemented")
}

func RegisterCommunicationVehicleServiceServer(s *grpc.Server, srv CommunicationVehicleServiceServer) {
	s.RegisterService(&_CommunicationVehicleService_serviceDesc, srv)
}

func _CommunicationVehicleService_PushTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationVehicleServiceServer).PushTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationVehicleService/PushTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationVehicleServiceServer).PushTelemetry(ctx, req.(*PushTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationVehicleService_PullCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationVehicleServiceServer).PullCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationVehicleService/PullCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationVehicleServiceServer).PullCommand(ctx, req.(*PullCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationVehicleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.CommunicationVehicleService",
	HandlerType: (*CommunicationVehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushTelemetry",
			Handler:    _CommunicationVehicleService_PushTelemetry_Handler,
		},
		{
			MethodName: "PullCommand",
			Handler:    _CommunicationVehicleService_PullCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication_vehicle.proto",
}
