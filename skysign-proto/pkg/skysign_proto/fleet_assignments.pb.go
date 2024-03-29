// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: proto/fleet_assignments.proto

package skysign_proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAssignmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssignmentsRequest) Reset() {
	*x = GetAssignmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fleet_assignments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentsRequest) ProtoMessage() {}

func (x *GetAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fleet_assignments_proto_msgTypes[0]
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
	return file_proto_fleet_assignments_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_fleet_assignments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentsResponse) ProtoMessage() {}

func (x *GetAssignmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fleet_assignments_proto_msgTypes[1]
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
	return file_proto_fleet_assignments_proto_rawDescGZIP(), []int{1}
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
		mi := &file_proto_fleet_assignments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentsRequest) ProtoMessage() {}

func (x *UpdateAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fleet_assignments_proto_msgTypes[2]
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
	return file_proto_fleet_assignments_proto_rawDescGZIP(), []int{2}
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
		mi := &file_proto_fleet_assignments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentsResponse) ProtoMessage() {}

func (x *UpdateAssignmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fleet_assignments_proto_msgTypes[3]
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
	return file_proto_fleet_assignments_proto_rawDescGZIP(), []int{3}
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

var File_proto_fleet_assignments_proto protoreflect.FileDescriptor

var file_proto_fleet_assignments_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x67, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x68, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xba, 0x02, 0x0a, 0x1a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x54, 0x6f, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6b,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x27, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x1a, 0x1f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fleet_assignments_proto_rawDescOnce sync.Once
	file_proto_fleet_assignments_proto_rawDescData = file_proto_fleet_assignments_proto_rawDesc
)

func file_proto_fleet_assignments_proto_rawDescGZIP() []byte {
	file_proto_fleet_assignments_proto_rawDescOnce.Do(func() {
		file_proto_fleet_assignments_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fleet_assignments_proto_rawDescData)
	})
	return file_proto_fleet_assignments_proto_rawDescData
}

var file_proto_fleet_assignments_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_fleet_assignments_proto_goTypes = []interface{}{
	(*GetAssignmentsRequest)(nil),     // 0: skysign_proto.GetAssignmentsRequest
	(*GetAssignmentsResponse)(nil),    // 1: skysign_proto.GetAssignmentsResponse
	(*UpdateAssignmentsRequest)(nil),  // 2: skysign_proto.UpdateAssignmentsRequest
	(*UpdateAssignmentsResponse)(nil), // 3: skysign_proto.UpdateAssignmentsResponse
	(*Assignment)(nil),                // 4: skysign_proto.Assignment
}
var file_proto_fleet_assignments_proto_depIdxs = []int32{
	4, // 0: skysign_proto.GetAssignmentsResponse.assignments:type_name -> skysign_proto.Assignment
	4, // 1: skysign_proto.UpdateAssignmentsRequest.assignments:type_name -> skysign_proto.Assignment
	4, // 2: skysign_proto.UpdateAssignmentsResponse.assignments:type_name -> skysign_proto.Assignment
	0, // 3: skysign_proto.AssignAssetsToFleetService.GetAssignments:input_type -> skysign_proto.GetAssignmentsRequest
	2, // 4: skysign_proto.AssignAssetsToFleetService.UpdateAssignments:input_type -> skysign_proto.UpdateAssignmentsRequest
	1, // 5: skysign_proto.AssignAssetsToFleetService.GetAssignments:output_type -> skysign_proto.GetAssignmentsResponse
	3, // 6: skysign_proto.AssignAssetsToFleetService.UpdateAssignments:output_type -> skysign_proto.UpdateAssignmentsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_fleet_assignments_proto_init() }
func file_proto_fleet_assignments_proto_init() {
	if File_proto_fleet_assignments_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_fleet_assignments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_fleet_assignments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_fleet_assignments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_fleet_assignments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_fleet_assignments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fleet_assignments_proto_goTypes,
		DependencyIndexes: file_proto_fleet_assignments_proto_depIdxs,
		MessageInfos:      file_proto_fleet_assignments_proto_msgTypes,
	}.Build()
	File_proto_fleet_assignments_proto = out.File
	file_proto_fleet_assignments_proto_rawDesc = nil
	file_proto_fleet_assignments_proto_goTypes = nil
	file_proto_fleet_assignments_proto_depIdxs = nil
}
