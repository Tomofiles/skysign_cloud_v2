// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: flightplan_assignments.proto

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

type ChangeNumberOfVehiclesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumberOfVehicles int32  `protobuf:"varint,2,opt,name=number_of_vehicles,json=numberOfVehicles,proto3" json:"number_of_vehicles,omitempty"`
}

func (x *ChangeNumberOfVehiclesRequest) Reset() {
	*x = ChangeNumberOfVehiclesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNumberOfVehiclesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNumberOfVehiclesRequest) ProtoMessage() {}

func (x *ChangeNumberOfVehiclesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNumberOfVehiclesRequest.ProtoReflect.Descriptor instead.
func (*ChangeNumberOfVehiclesRequest) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeNumberOfVehiclesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeNumberOfVehiclesRequest) GetNumberOfVehicles() int32 {
	if x != nil {
		return x.NumberOfVehicles
	}
	return 0
}

type ChangeNumberOfVehiclesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumberOfVehicles int32  `protobuf:"varint,2,opt,name=number_of_vehicles,json=numberOfVehicles,proto3" json:"number_of_vehicles,omitempty"`
}

func (x *ChangeNumberOfVehiclesResponse) Reset() {
	*x = ChangeNumberOfVehiclesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNumberOfVehiclesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNumberOfVehiclesResponse) ProtoMessage() {}

func (x *ChangeNumberOfVehiclesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNumberOfVehiclesResponse.ProtoReflect.Descriptor instead.
func (*ChangeNumberOfVehiclesResponse) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeNumberOfVehiclesResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeNumberOfVehiclesResponse) GetNumberOfVehicles() int32 {
	if x != nil {
		return x.NumberOfVehicles
	}
	return 0
}

type GetAssignmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssignmentsRequest) Reset() {
	*x = GetAssignmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentsRequest) ProtoMessage() {}

func (x *GetAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssignmentsRequest.ProtoReflect.Descriptor instead.
func (*GetAssignmentsRequest) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{2}
}

func (x *GetAssignmentsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAssignmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Assignments []*Assignment `protobuf:"bytes,2,rep,name=assignments,proto3" json:"assignments,omitempty"`
}

func (x *GetAssignmentsResponse) Reset() {
	*x = GetAssignmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentsResponse) ProtoMessage() {}

func (x *GetAssignmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssignmentsResponse.ProtoReflect.Descriptor instead.
func (*GetAssignmentsResponse) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssignmentsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAssignmentsResponse) GetAssignments() []*Assignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

type UpdateAssignmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Assignments []*Assignment `protobuf:"bytes,2,rep,name=assignments,proto3" json:"assignments,omitempty"`
}

func (x *UpdateAssignmentsRequest) Reset() {
	*x = UpdateAssignmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentsRequest) ProtoMessage() {}

func (x *UpdateAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssignmentsRequest.ProtoReflect.Descriptor instead.
func (*UpdateAssignmentsRequest) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAssignmentsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAssignmentsRequest) GetAssignments() []*Assignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

type UpdateAssignmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Assignments []*Assignment `protobuf:"bytes,2,rep,name=assignments,proto3" json:"assignments,omitempty"`
}

func (x *UpdateAssignmentsResponse) Reset() {
	*x = UpdateAssignmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flightplan_assignments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentsResponse) ProtoMessage() {}

func (x *UpdateAssignmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flightplan_assignments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssignmentsResponse.ProtoReflect.Descriptor instead.
func (*UpdateAssignmentsResponse) Descriptor() ([]byte, []int) {
	return file_flightplan_assignments_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAssignmentsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAssignmentsResponse) GetAssignments() []*Assignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

var File_flightplan_assignments_proto protoreflect.FileDescriptor

var file_flightplan_assignments_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x1d, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x1e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x67, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x68, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b,
	0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xf7, 0x03, 0x0a, 0x1f,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x54, 0x6f, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xab, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x1a,
	0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x6f, 0x66, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x97, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x27, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x1a, 0x24, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_flightplan_assignments_proto_rawDescOnce sync.Once
	file_flightplan_assignments_proto_rawDescData = file_flightplan_assignments_proto_rawDesc
)

func file_flightplan_assignments_proto_rawDescGZIP() []byte {
	file_flightplan_assignments_proto_rawDescOnce.Do(func() {
		file_flightplan_assignments_proto_rawDescData = protoimpl.X.CompressGZIP(file_flightplan_assignments_proto_rawDescData)
	})
	return file_flightplan_assignments_proto_rawDescData
}

var file_flightplan_assignments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_flightplan_assignments_proto_goTypes = []interface{}{
	(*ChangeNumberOfVehiclesRequest)(nil),  // 0: skysign_proto.ChangeNumberOfVehiclesRequest
	(*ChangeNumberOfVehiclesResponse)(nil), // 1: skysign_proto.ChangeNumberOfVehiclesResponse
	(*GetAssignmentsRequest)(nil),          // 2: skysign_proto.GetAssignmentsRequest
	(*GetAssignmentsResponse)(nil),         // 3: skysign_proto.GetAssignmentsResponse
	(*UpdateAssignmentsRequest)(nil),       // 4: skysign_proto.UpdateAssignmentsRequest
	(*UpdateAssignmentsResponse)(nil),      // 5: skysign_proto.UpdateAssignmentsResponse
	(*Assignment)(nil),                     // 6: skysign_proto.Assignment
}
var file_flightplan_assignments_proto_depIdxs = []int32{
	6, // 0: skysign_proto.GetAssignmentsResponse.assignments:type_name -> skysign_proto.Assignment
	6, // 1: skysign_proto.UpdateAssignmentsRequest.assignments:type_name -> skysign_proto.Assignment
	6, // 2: skysign_proto.UpdateAssignmentsResponse.assignments:type_name -> skysign_proto.Assignment
	0, // 3: skysign_proto.AssignAssetsToFlightplanService.ChangeNumberOfVehicles:input_type -> skysign_proto.ChangeNumberOfVehiclesRequest
	2, // 4: skysign_proto.AssignAssetsToFlightplanService.GetAssignments:input_type -> skysign_proto.GetAssignmentsRequest
	4, // 5: skysign_proto.AssignAssetsToFlightplanService.UpdateAssignments:input_type -> skysign_proto.UpdateAssignmentsRequest
	1, // 6: skysign_proto.AssignAssetsToFlightplanService.ChangeNumberOfVehicles:output_type -> skysign_proto.ChangeNumberOfVehiclesResponse
	3, // 7: skysign_proto.AssignAssetsToFlightplanService.GetAssignments:output_type -> skysign_proto.GetAssignmentsResponse
	5, // 8: skysign_proto.AssignAssetsToFlightplanService.UpdateAssignments:output_type -> skysign_proto.UpdateAssignmentsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flightplan_assignments_proto_init() }
func file_flightplan_assignments_proto_init() {
	if File_flightplan_assignments_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flightplan_assignments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNumberOfVehiclesRequest); i {
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
		file_flightplan_assignments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNumberOfVehiclesResponse); i {
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
		file_flightplan_assignments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssignmentsRequest); i {
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
		file_flightplan_assignments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssignmentsResponse); i {
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
		file_flightplan_assignments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssignmentsRequest); i {
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
		file_flightplan_assignments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssignmentsResponse); i {
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
			RawDescriptor: file_flightplan_assignments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flightplan_assignments_proto_goTypes,
		DependencyIndexes: file_flightplan_assignments_proto_depIdxs,
		MessageInfos:      file_flightplan_assignments_proto_msgTypes,
	}.Build()
	File_flightplan_assignments_proto = out.File
	file_flightplan_assignments_proto_rawDesc = nil
	file_flightplan_assignments_proto_goTypes = nil
	file_flightplan_assignments_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssignAssetsToFlightplanServiceClient is the client API for AssignAssetsToFlightplanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssignAssetsToFlightplanServiceClient interface {
	ChangeNumberOfVehicles(ctx context.Context, in *ChangeNumberOfVehiclesRequest, opts ...grpc.CallOption) (*ChangeNumberOfVehiclesResponse, error)
	GetAssignments(ctx context.Context, in *GetAssignmentsRequest, opts ...grpc.CallOption) (*GetAssignmentsResponse, error)
	UpdateAssignments(ctx context.Context, in *UpdateAssignmentsRequest, opts ...grpc.CallOption) (*UpdateAssignmentsResponse, error)
}

type assignAssetsToFlightplanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignAssetsToFlightplanServiceClient(cc grpc.ClientConnInterface) AssignAssetsToFlightplanServiceClient {
	return &assignAssetsToFlightplanServiceClient{cc}
}

func (c *assignAssetsToFlightplanServiceClient) ChangeNumberOfVehicles(ctx context.Context, in *ChangeNumberOfVehiclesRequest, opts ...grpc.CallOption) (*ChangeNumberOfVehiclesResponse, error) {
	out := new(ChangeNumberOfVehiclesResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.AssignAssetsToFlightplanService/ChangeNumberOfVehicles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignAssetsToFlightplanServiceClient) GetAssignments(ctx context.Context, in *GetAssignmentsRequest, opts ...grpc.CallOption) (*GetAssignmentsResponse, error) {
	out := new(GetAssignmentsResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.AssignAssetsToFlightplanService/GetAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignAssetsToFlightplanServiceClient) UpdateAssignments(ctx context.Context, in *UpdateAssignmentsRequest, opts ...grpc.CallOption) (*UpdateAssignmentsResponse, error) {
	out := new(UpdateAssignmentsResponse)
	err := c.cc.Invoke(ctx, "/skysign_proto.AssignAssetsToFlightplanService/UpdateAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignAssetsToFlightplanServiceServer is the server API for AssignAssetsToFlightplanService service.
type AssignAssetsToFlightplanServiceServer interface {
	ChangeNumberOfVehicles(context.Context, *ChangeNumberOfVehiclesRequest) (*ChangeNumberOfVehiclesResponse, error)
	GetAssignments(context.Context, *GetAssignmentsRequest) (*GetAssignmentsResponse, error)
	UpdateAssignments(context.Context, *UpdateAssignmentsRequest) (*UpdateAssignmentsResponse, error)
}

// UnimplementedAssignAssetsToFlightplanServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssignAssetsToFlightplanServiceServer struct {
}

func (*UnimplementedAssignAssetsToFlightplanServiceServer) ChangeNumberOfVehicles(context.Context, *ChangeNumberOfVehiclesRequest) (*ChangeNumberOfVehiclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNumberOfVehicles not implemented")
}
func (*UnimplementedAssignAssetsToFlightplanServiceServer) GetAssignments(context.Context, *GetAssignmentsRequest) (*GetAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignments not implemented")
}
func (*UnimplementedAssignAssetsToFlightplanServiceServer) UpdateAssignments(context.Context, *UpdateAssignmentsRequest) (*UpdateAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssignments not implemented")
}

func RegisterAssignAssetsToFlightplanServiceServer(s *grpc.Server, srv AssignAssetsToFlightplanServiceServer) {
	s.RegisterService(&_AssignAssetsToFlightplanService_serviceDesc, srv)
}

func _AssignAssetsToFlightplanService_ChangeNumberOfVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNumberOfVehiclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignAssetsToFlightplanServiceServer).ChangeNumberOfVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.AssignAssetsToFlightplanService/ChangeNumberOfVehicles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignAssetsToFlightplanServiceServer).ChangeNumberOfVehicles(ctx, req.(*ChangeNumberOfVehiclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignAssetsToFlightplanService_GetAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignAssetsToFlightplanServiceServer).GetAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.AssignAssetsToFlightplanService/GetAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignAssetsToFlightplanServiceServer).GetAssignments(ctx, req.(*GetAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignAssetsToFlightplanService_UpdateAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignAssetsToFlightplanServiceServer).UpdateAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.AssignAssetsToFlightplanService/UpdateAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignAssetsToFlightplanServiceServer).UpdateAssignments(ctx, req.(*UpdateAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssignAssetsToFlightplanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.AssignAssetsToFlightplanService",
	HandlerType: (*AssignAssetsToFlightplanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeNumberOfVehicles",
			Handler:    _AssignAssetsToFlightplanService_ChangeNumberOfVehicles_Handler,
		},
		{
			MethodName: "GetAssignments",
			Handler:    _AssignAssetsToFlightplanService_GetAssignments_Handler,
		},
		{
			MethodName: "UpdateAssignments",
			Handler:    _AssignAssetsToFlightplanService_UpdateAssignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flightplan_assignments.proto",
}
