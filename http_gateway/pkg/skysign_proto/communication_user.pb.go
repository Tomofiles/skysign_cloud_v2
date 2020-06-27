// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: communication_user.proto

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

type CommandType int32

const (
	CommandType_UPLOAD CommandType = 0
	CommandType_NONE   CommandType = 99
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0:  "UPLOAD",
		99: "NONE",
	}
	CommandType_value = map[string]int32{
		"UPLOAD": 0,
		"NONE":   99,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_communication_user_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_communication_user_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{0}
}

type PushCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type CommandType `protobuf:"varint,2,opt,name=type,proto3,enum=skysign_proto.CommandType" json:"type,omitempty"`
}

func (x *PushCommandRequest) Reset() {
	*x = PushCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushCommandRequest) ProtoMessage() {}

func (x *PushCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushCommandRequest.ProtoReflect.Descriptor instead.
func (*PushCommandRequest) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{0}
}

func (x *PushCommandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PushCommandRequest) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_UPLOAD
}

type PushCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type CommandType `protobuf:"varint,2,opt,name=type,proto3,enum=skysign_proto.CommandType" json:"type,omitempty"`
}

func (x *PushCommandResponse) Reset() {
	*x = PushCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushCommandResponse) ProtoMessage() {}

func (x *PushCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushCommandResponse.ProtoReflect.Descriptor instead.
func (*PushCommandResponse) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{1}
}

func (x *PushCommandResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PushCommandResponse) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_UPLOAD
}

type PullTelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PullTelemetryRequest) Reset() {
	*x = PullTelemetryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTelemetryRequest) ProtoMessage() {}

func (x *PullTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTelemetryRequest.ProtoReflect.Descriptor instead.
func (*PullTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{2}
}

func (x *PullTelemetryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PullTelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude     float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude    float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude     float64 `protobuf:"fixed64,4,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Speed        float64 `protobuf:"fixed64,5,opt,name=speed,proto3" json:"speed,omitempty"`
	Armed        bool    `protobuf:"varint,6,opt,name=armed,proto3" json:"armed,omitempty"`
	FlightMode   string  `protobuf:"bytes,7,opt,name=flightMode,proto3" json:"flightMode,omitempty"`
	OrientationX float64 `protobuf:"fixed64,8,opt,name=orientationX,proto3" json:"orientationX,omitempty"`
	OrientationY float64 `protobuf:"fixed64,9,opt,name=orientationY,proto3" json:"orientationY,omitempty"`
	OrientationZ float64 `protobuf:"fixed64,10,opt,name=orientationZ,proto3" json:"orientationZ,omitempty"`
	OrientationW float64 `protobuf:"fixed64,11,opt,name=orientationW,proto3" json:"orientationW,omitempty"`
}

func (x *PullTelemetryResponse) Reset() {
	*x = PullTelemetryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTelemetryResponse) ProtoMessage() {}

func (x *PullTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTelemetryResponse.ProtoReflect.Descriptor instead.
func (*PullTelemetryResponse) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{3}
}

func (x *PullTelemetryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullTelemetryResponse) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PullTelemetryResponse) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PullTelemetryResponse) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *PullTelemetryResponse) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *PullTelemetryResponse) GetArmed() bool {
	if x != nil {
		return x.Armed
	}
	return false
}

func (x *PullTelemetryResponse) GetFlightMode() string {
	if x != nil {
		return x.FlightMode
	}
	return ""
}

func (x *PullTelemetryResponse) GetOrientationX() float64 {
	if x != nil {
		return x.OrientationX
	}
	return 0
}

func (x *PullTelemetryResponse) GetOrientationY() float64 {
	if x != nil {
		return x.OrientationY
	}
	return 0
}

func (x *PullTelemetryResponse) GetOrientationZ() float64 {
	if x != nil {
		return x.OrientationZ
	}
	return 0
}

func (x *PullTelemetryResponse) GetOrientationW() float64 {
	if x != nil {
		return x.OrientationW
	}
	return 0
}

type StandByRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MissionId string `protobuf:"bytes,2,opt,name=missionId,proto3" json:"missionId,omitempty"`
}

func (x *StandByRequest) Reset() {
	*x = StandByRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandByRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandByRequest) ProtoMessage() {}

func (x *StandByRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandByRequest.ProtoReflect.Descriptor instead.
func (*StandByRequest) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{4}
}

func (x *StandByRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StandByRequest) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

type StandByResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MissionId string `protobuf:"bytes,2,opt,name=missionId,proto3" json:"missionId,omitempty"`
}

func (x *StandByResponse) Reset() {
	*x = StandByResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandByResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandByResponse) ProtoMessage() {}

func (x *StandByResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandByResponse.ProtoReflect.Descriptor instead.
func (*StandByResponse) Descriptor() ([]byte, []int) {
	return file_communication_user_proto_rawDescGZIP(), []int{5}
}

func (x *StandByResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StandByResponse) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

var File_communication_user_proto protoreflect.FileDescriptor

var file_communication_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x12, 0x50, 0x75, 0x73, 0x68, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x6b,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a,
	0x13, 0x50, 0x75, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd9, 0x02, 0x0a,
	0x15, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59,
	0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5a, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x42, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x23, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x4c, 0x4f,
	0x41, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x63, 0x32, 0x97,
	0x03, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x0b, 0x50,
	0x75, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a,
	0x0d, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x23,
	0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x12,
	0x1d, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x62, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_communication_user_proto_rawDescOnce sync.Once
	file_communication_user_proto_rawDescData = file_communication_user_proto_rawDesc
)

func file_communication_user_proto_rawDescGZIP() []byte {
	file_communication_user_proto_rawDescOnce.Do(func() {
		file_communication_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_user_proto_rawDescData)
	})
	return file_communication_user_proto_rawDescData
}

var file_communication_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_communication_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_communication_user_proto_goTypes = []interface{}{
	(CommandType)(0),              // 0: skysign_proto.CommandType
	(*PushCommandRequest)(nil),    // 1: skysign_proto.PushCommandRequest
	(*PushCommandResponse)(nil),   // 2: skysign_proto.PushCommandResponse
	(*PullTelemetryRequest)(nil),  // 3: skysign_proto.PullTelemetryRequest
	(*PullTelemetryResponse)(nil), // 4: skysign_proto.PullTelemetryResponse
	(*StandByRequest)(nil),        // 5: skysign_proto.StandByRequest
	(*StandByResponse)(nil),       // 6: skysign_proto.StandByResponse
}
var file_communication_user_proto_depIdxs = []int32{
	0, // 0: skysign_proto.PushCommandRequest.type:type_name -> skysign_proto.CommandType
	0, // 1: skysign_proto.PushCommandResponse.type:type_name -> skysign_proto.CommandType
	1, // 2: skysign_proto.CommunicationUserService.PushCommand:input_type -> skysign_proto.PushCommandRequest
	3, // 3: skysign_proto.CommunicationUserService.PullTelemetry:input_type -> skysign_proto.PullTelemetryRequest
	5, // 4: skysign_proto.CommunicationUserService.StandBy:input_type -> skysign_proto.StandByRequest
	2, // 5: skysign_proto.CommunicationUserService.PushCommand:output_type -> skysign_proto.PushCommandResponse
	4, // 6: skysign_proto.CommunicationUserService.PullTelemetry:output_type -> skysign_proto.PullTelemetryResponse
	6, // 7: skysign_proto.CommunicationUserService.StandBy:output_type -> skysign_proto.StandByResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_communication_user_proto_init() }
func file_communication_user_proto_init() {
	if File_communication_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushCommandRequest); i {
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
		file_communication_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushCommandResponse); i {
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
		file_communication_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTelemetryRequest); i {
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
		file_communication_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTelemetryResponse); i {
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
		file_communication_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandByRequest); i {
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
		file_communication_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandByResponse); i {
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
			RawDescriptor: file_communication_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_user_proto_goTypes,
		DependencyIndexes: file_communication_user_proto_depIdxs,
		EnumInfos:         file_communication_user_proto_enumTypes,
		MessageInfos:      file_communication_user_proto_msgTypes,
	}.Build()
	File_communication_user_proto = out.File
	file_communication_user_proto_rawDesc = nil
	file_communication_user_proto_goTypes = nil
	file_communication_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommunicationUserServiceClient is the client API for CommunicationUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationUserServiceClient interface {
	PushCommand(ctx context.Context, in *PushCommandRequest, opts ...grpc.CallOption) (*PushCommandResponse, error)
	PullTelemetry(ctx context.Context, in *PullTelemetryRequest, opts ...grpc.CallOption) (*PullTelemetryResponse, error)
	StandBy(ctx context.Context, in *StandByRequest, opts ...grpc.CallOption) (*StandByResponse, error)
}

type communicationUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationUserServiceClient(cc grpc.ClientConnInterface) CommunicationUserServiceClient {
	return &communicationUserServiceClient{cc}
}

func (c *communicationUserServiceClient) PushCommand(ctx context.Context, in *PushCommandRequest, opts ...grpc.CallOption) (*PushCommandResponse, error) {
	out := new(PushCommandResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationUserService/PushCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationUserServiceClient) PullTelemetry(ctx context.Context, in *PullTelemetryRequest, opts ...grpc.CallOption) (*PullTelemetryResponse, error) {
	out := new(PullTelemetryResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationUserService/PullTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationUserServiceClient) StandBy(ctx context.Context, in *StandByRequest, opts ...grpc.CallOption) (*StandByResponse, error) {
	out := new(StandByResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationUserService/StandBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationUserServiceServer is the server API for CommunicationUserService service.
type CommunicationUserServiceServer interface {
	PushCommand(context.Context, *PushCommandRequest) (*PushCommandResponse, error)
	PullTelemetry(context.Context, *PullTelemetryRequest) (*PullTelemetryResponse, error)
	StandBy(context.Context, *StandByRequest) (*StandByResponse, error)
}

// UnimplementedCommunicationUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationUserServiceServer struct {
}

func (*UnimplementedCommunicationUserServiceServer) PushCommand(context.Context, *PushCommandRequest) (*PushCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCommand not implemented")
}
func (*UnimplementedCommunicationUserServiceServer) PullTelemetry(context.Context, *PullTelemetryRequest) (*PullTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullTelemetry not implemented")
}
func (*UnimplementedCommunicationUserServiceServer) StandBy(context.Context, *StandByRequest) (*StandByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StandBy not implemented")
}

func RegisterCommunicationUserServiceServer(s *grpc.Server, srv CommunicationUserServiceServer) {
	s.RegisterService(&_CommunicationUserService_serviceDesc, srv)
}

func _CommunicationUserService_PushCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationUserServiceServer).PushCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationUserService/PushCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationUserServiceServer).PushCommand(ctx, req.(*PushCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationUserService_PullTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationUserServiceServer).PullTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationUserService/PullTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationUserServiceServer).PullTelemetry(ctx, req.(*PullTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationUserService_StandBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationUserServiceServer).StandBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationUserService/StandBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationUserServiceServer).StandBy(ctx, req.(*StandByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.CommunicationUserService",
	HandlerType: (*CommunicationUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushCommand",
			Handler:    _CommunicationUserService_PushCommand_Handler,
		},
		{
			MethodName: "PullTelemetry",
			Handler:    _CommunicationUserService_PullTelemetry_Handler,
		},
		{
			MethodName: "StandBy",
			Handler:    _CommunicationUserService_StandBy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication_user.proto",
}
