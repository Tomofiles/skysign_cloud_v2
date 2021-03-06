// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: communication_edge.proto

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

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Telemetry *Telemetry `protobuf:"bytes,2,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
}

func (x *PushTelemetryRequest) Reset() {
	*x = PushTelemetryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTelemetryRequest) ProtoMessage() {}

func (x *PushTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[0]
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
	return file_communication_edge_proto_rawDescGZIP(), []int{0}
}

func (x *PushTelemetryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PushTelemetryRequest) GetTelemetry() *Telemetry {
	if x != nil {
		return x.Telemetry
	}
	return nil
}

type PushTelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommIds []string `protobuf:"bytes,2,rep,name=commIds,proto3" json:"commIds,omitempty"`
}

func (x *PushTelemetryResponse) Reset() {
	*x = PushTelemetryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTelemetryResponse) ProtoMessage() {}

func (x *PushTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[1]
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
	return file_communication_edge_proto_rawDescGZIP(), []int{1}
}

func (x *PushTelemetryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PushTelemetryResponse) GetCommIds() []string {
	if x != nil {
		return x.CommIds
	}
	return nil
}

type PullCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandId string `protobuf:"bytes,2,opt,name=commandId,proto3" json:"commandId,omitempty"`
}

func (x *PullCommandRequest) Reset() {
	*x = PullCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCommandRequest) ProtoMessage() {}

func (x *PullCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[2]
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
	return file_communication_edge_proto_rawDescGZIP(), []int{2}
}

func (x *PullCommandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullCommandRequest) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

type PullCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandId string      `protobuf:"bytes,2,opt,name=commandId,proto3" json:"commandId,omitempty"`
	Type      CommandType `protobuf:"varint,3,opt,name=type,proto3,enum=skysign_proto.CommandType" json:"type,omitempty"`
}

func (x *PullCommandResponse) Reset() {
	*x = PullCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCommandResponse) ProtoMessage() {}

func (x *PullCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[3]
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
	return file_communication_edge_proto_rawDescGZIP(), []int{3}
}

func (x *PullCommandResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullCommandResponse) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *PullCommandResponse) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_ARM
}

type PullUploadMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandId string `protobuf:"bytes,2,opt,name=commandId,proto3" json:"commandId,omitempty"`
}

func (x *PullUploadMissionRequest) Reset() {
	*x = PullUploadMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUploadMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUploadMissionRequest) ProtoMessage() {}

func (x *PullUploadMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUploadMissionRequest.ProtoReflect.Descriptor instead.
func (*PullUploadMissionRequest) Descriptor() ([]byte, []int) {
	return file_communication_edge_proto_rawDescGZIP(), []int{4}
}

func (x *PullUploadMissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullUploadMissionRequest) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

type PullUploadMissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandId string `protobuf:"bytes,2,opt,name=commandId,proto3" json:"commandId,omitempty"`
	MissionId string `protobuf:"bytes,3,opt,name=missionId,proto3" json:"missionId,omitempty"`
}

func (x *PullUploadMissionResponse) Reset() {
	*x = PullUploadMissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_edge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUploadMissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUploadMissionResponse) ProtoMessage() {}

func (x *PullUploadMissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_edge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUploadMissionResponse.ProtoReflect.Descriptor instead.
func (*PullUploadMissionResponse) Descriptor() ([]byte, []int) {
	return file_communication_edge_proto_rawDescGZIP(), []int{5}
}

func (x *PullUploadMissionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PullUploadMissionResponse) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *PullUploadMissionResponse) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

var File_communication_edge_proto protoreflect.FileDescriptor

var file_communication_edge_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a,
	0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x41, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x13,
	0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x48, 0x0a, 0x18, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x19, 0x50,
	0x75, 0x6c, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x32, 0xe9, 0x03, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x8c, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x3a, 0x01, 0x2a,
	0x12, 0x91, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22,
	0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xa9, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x22, 0x36, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x7d, 0x3a, 0x01, 0x2a,
	0x42, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_edge_proto_rawDescOnce sync.Once
	file_communication_edge_proto_rawDescData = file_communication_edge_proto_rawDesc
)

func file_communication_edge_proto_rawDescGZIP() []byte {
	file_communication_edge_proto_rawDescOnce.Do(func() {
		file_communication_edge_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_edge_proto_rawDescData)
	})
	return file_communication_edge_proto_rawDescData
}

var file_communication_edge_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_communication_edge_proto_goTypes = []interface{}{
	(*PushTelemetryRequest)(nil),      // 0: skysign_proto.PushTelemetryRequest
	(*PushTelemetryResponse)(nil),     // 1: skysign_proto.PushTelemetryResponse
	(*PullCommandRequest)(nil),        // 2: skysign_proto.PullCommandRequest
	(*PullCommandResponse)(nil),       // 3: skysign_proto.PullCommandResponse
	(*PullUploadMissionRequest)(nil),  // 4: skysign_proto.PullUploadMissionRequest
	(*PullUploadMissionResponse)(nil), // 5: skysign_proto.PullUploadMissionResponse
	(*Telemetry)(nil),                 // 6: skysign_proto.Telemetry
	(CommandType)(0),                  // 7: skysign_proto.CommandType
}
var file_communication_edge_proto_depIdxs = []int32{
	6, // 0: skysign_proto.PushTelemetryRequest.telemetry:type_name -> skysign_proto.Telemetry
	7, // 1: skysign_proto.PullCommandResponse.type:type_name -> skysign_proto.CommandType
	0, // 2: skysign_proto.CommunicationEdgeService.PushTelemetry:input_type -> skysign_proto.PushTelemetryRequest
	2, // 3: skysign_proto.CommunicationEdgeService.PullCommand:input_type -> skysign_proto.PullCommandRequest
	4, // 4: skysign_proto.CommunicationEdgeService.PullUploadMission:input_type -> skysign_proto.PullUploadMissionRequest
	1, // 5: skysign_proto.CommunicationEdgeService.PushTelemetry:output_type -> skysign_proto.PushTelemetryResponse
	3, // 6: skysign_proto.CommunicationEdgeService.PullCommand:output_type -> skysign_proto.PullCommandResponse
	5, // 7: skysign_proto.CommunicationEdgeService.PullUploadMission:output_type -> skysign_proto.PullUploadMissionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_communication_edge_proto_init() }
func file_communication_edge_proto_init() {
	if File_communication_edge_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_communication_edge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_communication_edge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_communication_edge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_communication_edge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_communication_edge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUploadMissionRequest); i {
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
		file_communication_edge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUploadMissionResponse); i {
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
			RawDescriptor: file_communication_edge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_edge_proto_goTypes,
		DependencyIndexes: file_communication_edge_proto_depIdxs,
		MessageInfos:      file_communication_edge_proto_msgTypes,
	}.Build()
	File_communication_edge_proto = out.File
	file_communication_edge_proto_rawDesc = nil
	file_communication_edge_proto_goTypes = nil
	file_communication_edge_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommunicationEdgeServiceClient is the client API for CommunicationEdgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationEdgeServiceClient interface {
	PushTelemetry(ctx context.Context, in *PushTelemetryRequest, opts ...grpc.CallOption) (*PushTelemetryResponse, error)
	PullCommand(ctx context.Context, in *PullCommandRequest, opts ...grpc.CallOption) (*PullCommandResponse, error)
	PullUploadMission(ctx context.Context, in *PullUploadMissionRequest, opts ...grpc.CallOption) (*PullUploadMissionResponse, error)
}

type communicationEdgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationEdgeServiceClient(cc grpc.ClientConnInterface) CommunicationEdgeServiceClient {
	return &communicationEdgeServiceClient{cc}
}

func (c *communicationEdgeServiceClient) PushTelemetry(ctx context.Context, in *PushTelemetryRequest, opts ...grpc.CallOption) (*PushTelemetryResponse, error) {
	out := new(PushTelemetryResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationEdgeService/PushTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationEdgeServiceClient) PullCommand(ctx context.Context, in *PullCommandRequest, opts ...grpc.CallOption) (*PullCommandResponse, error) {
	out := new(PullCommandResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationEdgeService/PullCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationEdgeServiceClient) PullUploadMission(ctx context.Context, in *PullUploadMissionRequest, opts ...grpc.CallOption) (*PullUploadMissionResponse, error) {
	out := new(PullUploadMissionResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.CommunicationEdgeService/PullUploadMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationEdgeServiceServer is the server API for CommunicationEdgeService service.
type CommunicationEdgeServiceServer interface {
	PushTelemetry(context.Context, *PushTelemetryRequest) (*PushTelemetryResponse, error)
	PullCommand(context.Context, *PullCommandRequest) (*PullCommandResponse, error)
	PullUploadMission(context.Context, *PullUploadMissionRequest) (*PullUploadMissionResponse, error)
}

// UnimplementedCommunicationEdgeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationEdgeServiceServer struct {
}

func (*UnimplementedCommunicationEdgeServiceServer) PushTelemetry(context.Context, *PushTelemetryRequest) (*PushTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTelemetry not implemented")
}
func (*UnimplementedCommunicationEdgeServiceServer) PullCommand(context.Context, *PullCommandRequest) (*PullCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullCommand not implemented")
}
func (*UnimplementedCommunicationEdgeServiceServer) PullUploadMission(context.Context, *PullUploadMissionRequest) (*PullUploadMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullUploadMission not implemented")
}

func RegisterCommunicationEdgeServiceServer(s *grpc.Server, srv CommunicationEdgeServiceServer) {
	s.RegisterService(&_CommunicationEdgeService_serviceDesc, srv)
}

func _CommunicationEdgeService_PushTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationEdgeServiceServer).PushTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationEdgeService/PushTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationEdgeServiceServer).PushTelemetry(ctx, req.(*PushTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationEdgeService_PullCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationEdgeServiceServer).PullCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationEdgeService/PullCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationEdgeServiceServer).PullCommand(ctx, req.(*PullCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationEdgeService_PullUploadMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullUploadMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationEdgeServiceServer).PullUploadMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.CommunicationEdgeService/PullUploadMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationEdgeServiceServer).PullUploadMission(ctx, req.(*PullUploadMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationEdgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.CommunicationEdgeService",
	HandlerType: (*CommunicationEdgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushTelemetry",
			Handler:    _CommunicationEdgeService_PushTelemetry_Handler,
		},
		{
			MethodName: "PullCommand",
			Handler:    _CommunicationEdgeService_PullCommand_Handler,
		},
		{
			MethodName: "PullUploadMission",
			Handler:    _CommunicationEdgeService_PullUploadMission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication_edge.proto",
}
