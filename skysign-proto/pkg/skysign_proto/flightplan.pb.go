// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: proto/flightplan.proto

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

type GetFlightplanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFlightplanRequest) Reset() {
	*x = GetFlightplanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flightplan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightplanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightplanRequest) ProtoMessage() {}

func (x *GetFlightplanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flightplan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightplanRequest.ProtoReflect.Descriptor instead.
func (*GetFlightplanRequest) Descriptor() ([]byte, []int) {
	return file_proto_flightplan_proto_rawDescGZIP(), []int{0}
}

func (x *GetFlightplanRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFlightplanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFlightplanRequest) Reset() {
	*x = DeleteFlightplanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flightplan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlightplanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlightplanRequest) ProtoMessage() {}

func (x *DeleteFlightplanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flightplan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlightplanRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlightplanRequest) Descriptor() ([]byte, []int) {
	return file_proto_flightplan_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFlightplanRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFlightplansResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flightplans []*Flightplan `protobuf:"bytes,1,rep,name=flightplans,proto3" json:"flightplans,omitempty"`
}

func (x *ListFlightplansResponses) Reset() {
	*x = ListFlightplansResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flightplan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightplansResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightplansResponses) ProtoMessage() {}

func (x *ListFlightplansResponses) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flightplan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightplansResponses.ProtoReflect.Descriptor instead.
func (*ListFlightplansResponses) Descriptor() ([]byte, []int) {
	return file_proto_flightplan_proto_rawDescGZIP(), []int{2}
}

func (x *ListFlightplansResponses) GetFlightplans() []*Flightplan {
	if x != nil {
		return x.Flightplans
	}
	return nil
}

var File_proto_flightplan_proto protoreflect.FileDescriptor

var file_proto_flightplan_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x0b, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x73, 0x32, 0xc8, 0x04, 0x0a, 0x17, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73,
	0x12, 0x71, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x12, 0x23, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x1a, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x12, 0x19, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x19, 0x2e, 0x73,
	0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x1a,
	0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x12, 0x26, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_flightplan_proto_rawDescOnce sync.Once
	file_proto_flightplan_proto_rawDescData = file_proto_flightplan_proto_rawDesc
)

func file_proto_flightplan_proto_rawDescGZIP() []byte {
	file_proto_flightplan_proto_rawDescOnce.Do(func() {
		file_proto_flightplan_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_flightplan_proto_rawDescData)
	})
	return file_proto_flightplan_proto_rawDescData
}

var file_proto_flightplan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_flightplan_proto_goTypes = []interface{}{
	(*GetFlightplanRequest)(nil),     // 0: skysign_proto.GetFlightplanRequest
	(*DeleteFlightplanRequest)(nil),  // 1: skysign_proto.DeleteFlightplanRequest
	(*ListFlightplansResponses)(nil), // 2: skysign_proto.ListFlightplansResponses
	(*Flightplan)(nil),               // 3: skysign_proto.Flightplan
	(*Empty)(nil),                    // 4: skysign_proto.Empty
}
var file_proto_flightplan_proto_depIdxs = []int32{
	3, // 0: skysign_proto.ListFlightplansResponses.flightplans:type_name -> skysign_proto.Flightplan
	4, // 1: skysign_proto.ManageFlightplanService.ListFlightplans:input_type -> skysign_proto.Empty
	0, // 2: skysign_proto.ManageFlightplanService.GetFlightplan:input_type -> skysign_proto.GetFlightplanRequest
	3, // 3: skysign_proto.ManageFlightplanService.CreateFlightplan:input_type -> skysign_proto.Flightplan
	3, // 4: skysign_proto.ManageFlightplanService.UpdateFlightplan:input_type -> skysign_proto.Flightplan
	1, // 5: skysign_proto.ManageFlightplanService.DeleteFlightplan:input_type -> skysign_proto.DeleteFlightplanRequest
	2, // 6: skysign_proto.ManageFlightplanService.ListFlightplans:output_type -> skysign_proto.ListFlightplansResponses
	3, // 7: skysign_proto.ManageFlightplanService.GetFlightplan:output_type -> skysign_proto.Flightplan
	3, // 8: skysign_proto.ManageFlightplanService.CreateFlightplan:output_type -> skysign_proto.Flightplan
	3, // 9: skysign_proto.ManageFlightplanService.UpdateFlightplan:output_type -> skysign_proto.Flightplan
	4, // 10: skysign_proto.ManageFlightplanService.DeleteFlightplan:output_type -> skysign_proto.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_flightplan_proto_init() }
func file_proto_flightplan_proto_init() {
	if File_proto_flightplan_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_flightplan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightplanRequest); i {
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
		file_proto_flightplan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlightplanRequest); i {
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
		file_proto_flightplan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightplansResponses); i {
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
			RawDescriptor: file_proto_flightplan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_flightplan_proto_goTypes,
		DependencyIndexes: file_proto_flightplan_proto_depIdxs,
		MessageInfos:      file_proto_flightplan_proto_msgTypes,
	}.Build()
	File_proto_flightplan_proto = out.File
	file_proto_flightplan_proto_rawDesc = nil
	file_proto_flightplan_proto_goTypes = nil
	file_proto_flightplan_proto_depIdxs = nil
}
