// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: proto/flightreport.proto

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

type GetFlightreportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFlightreportRequest) Reset() {
	*x = GetFlightreportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flightreport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightreportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightreportRequest) ProtoMessage() {}

func (x *GetFlightreportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flightreport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightreportRequest.ProtoReflect.Descriptor instead.
func (*GetFlightreportRequest) Descriptor() ([]byte, []int) {
	return file_proto_flightreport_proto_rawDescGZIP(), []int{0}
}

func (x *GetFlightreportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFlightreportsResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flightreports []*Flightreport `protobuf:"bytes,1,rep,name=flightreports,proto3" json:"flightreports,omitempty"`
}

func (x *ListFlightreportsResponses) Reset() {
	*x = ListFlightreportsResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flightreport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightreportsResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightreportsResponses) ProtoMessage() {}

func (x *ListFlightreportsResponses) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flightreport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightreportsResponses.ProtoReflect.Descriptor instead.
func (*ListFlightreportsResponses) Descriptor() ([]byte, []int) {
	return file_proto_flightreport_proto_rawDescGZIP(), []int{1}
}

func (x *ListFlightreportsResponses) GetFlightreports() []*Flightreport {
	if x != nil {
		return x.Flightreports
	}
	return nil
}

var File_proto_flightreport_proto protoreflect.FileDescriptor

var file_proto_flightreport_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x85, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x73, 0x6b, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50,
	0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_flightreport_proto_rawDescOnce sync.Once
	file_proto_flightreport_proto_rawDescData = file_proto_flightreport_proto_rawDesc
)

func file_proto_flightreport_proto_rawDescGZIP() []byte {
	file_proto_flightreport_proto_rawDescOnce.Do(func() {
		file_proto_flightreport_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_flightreport_proto_rawDescData)
	})
	return file_proto_flightreport_proto_rawDescData
}

var file_proto_flightreport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_flightreport_proto_goTypes = []interface{}{
	(*GetFlightreportRequest)(nil),     // 0: skysign_proto.GetFlightreportRequest
	(*ListFlightreportsResponses)(nil), // 1: skysign_proto.ListFlightreportsResponses
	(*Flightreport)(nil),               // 2: skysign_proto.Flightreport
	(*Empty)(nil),                      // 3: skysign_proto.Empty
}
var file_proto_flightreport_proto_depIdxs = []int32{
	2, // 0: skysign_proto.ListFlightreportsResponses.flightreports:type_name -> skysign_proto.Flightreport
	3, // 1: skysign_proto.ReportFlightService.ListFlightreports:input_type -> skysign_proto.Empty
	0, // 2: skysign_proto.ReportFlightService.GetFlightreport:input_type -> skysign_proto.GetFlightreportRequest
	1, // 3: skysign_proto.ReportFlightService.ListFlightreports:output_type -> skysign_proto.ListFlightreportsResponses
	2, // 4: skysign_proto.ReportFlightService.GetFlightreport:output_type -> skysign_proto.Flightreport
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_flightreport_proto_init() }
func file_proto_flightreport_proto_init() {
	if File_proto_flightreport_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_flightreport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightreportRequest); i {
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
		file_proto_flightreport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightreportsResponses); i {
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
			RawDescriptor: file_proto_flightreport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_flightreport_proto_goTypes,
		DependencyIndexes: file_proto_flightreport_proto_depIdxs,
		MessageInfos:      file_proto_flightreport_proto_msgTypes,
	}.Build()
	File_proto_flightreport_proto = out.File
	file_proto_flightreport_proto_rawDesc = nil
	file_proto_flightreport_proto_goTypes = nil
	file_proto_flightreport_proto_depIdxs = nil
}
