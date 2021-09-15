// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: proto/event.proto

package skysign_proto

import (
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

type CommunicationIdGaveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunicationId string `protobuf:"bytes,1,opt,name=communication_id,json=communicationId,proto3" json:"communication_id,omitempty"`
}

func (x *CommunicationIdGaveEvent) Reset() {
	*x = CommunicationIdGaveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicationIdGaveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicationIdGaveEvent) ProtoMessage() {}

func (x *CommunicationIdGaveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicationIdGaveEvent.ProtoReflect.Descriptor instead.
func (*CommunicationIdGaveEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *CommunicationIdGaveEvent) GetCommunicationId() string {
	if x != nil {
		return x.CommunicationId
	}
	return ""
}

type CommunicationIdRemovedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunicationId string `protobuf:"bytes,1,opt,name=communication_id,json=communicationId,proto3" json:"communication_id,omitempty"`
}

func (x *CommunicationIdRemovedEvent) Reset() {
	*x = CommunicationIdRemovedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicationIdRemovedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicationIdRemovedEvent) ProtoMessage() {}

func (x *CommunicationIdRemovedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicationIdRemovedEvent.ProtoReflect.Descriptor instead.
func (*CommunicationIdRemovedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{1}
}

func (x *CommunicationIdRemovedEvent) GetCommunicationId() string {
	if x != nil {
		return x.CommunicationId
	}
	return ""
}

type FleetIdGaveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FleetId          string `protobuf:"bytes,1,opt,name=fleet_id,json=fleetId,proto3" json:"fleet_id,omitempty"`
	NumberOfVehicles int32  `protobuf:"varint,2,opt,name=number_of_vehicles,json=numberOfVehicles,proto3" json:"number_of_vehicles,omitempty"`
}

func (x *FleetIdGaveEvent) Reset() {
	*x = FleetIdGaveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleetIdGaveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleetIdGaveEvent) ProtoMessage() {}

func (x *FleetIdGaveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleetIdGaveEvent.ProtoReflect.Descriptor instead.
func (*FleetIdGaveEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{2}
}

func (x *FleetIdGaveEvent) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

func (x *FleetIdGaveEvent) GetNumberOfVehicles() int32 {
	if x != nil {
		return x.NumberOfVehicles
	}
	return 0
}

type FleetIdRemovedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FleetId string `protobuf:"bytes,1,opt,name=fleet_id,json=fleetId,proto3" json:"fleet_id,omitempty"`
}

func (x *FleetIdRemovedEvent) Reset() {
	*x = FleetIdRemovedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleetIdRemovedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleetIdRemovedEvent) ProtoMessage() {}

func (x *FleetIdRemovedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleetIdRemovedEvent.ProtoReflect.Descriptor instead.
func (*FleetIdRemovedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{3}
}

func (x *FleetIdRemovedEvent) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

type FlightplanExecutedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightplanId string      `protobuf:"bytes,1,opt,name=flightplan_id,json=flightplanId,proto3" json:"flightplan_id,omitempty"`
	Flightplan   *Flightplan `protobuf:"bytes,2,opt,name=flightplan,proto3" json:"flightplan,omitempty"`
}

func (x *FlightplanExecutedEvent) Reset() {
	*x = FlightplanExecutedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightplanExecutedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightplanExecutedEvent) ProtoMessage() {}

func (x *FlightplanExecutedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightplanExecutedEvent.ProtoReflect.Descriptor instead.
func (*FlightplanExecutedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{4}
}

func (x *FlightplanExecutedEvent) GetFlightplanId() string {
	if x != nil {
		return x.FlightplanId
	}
	return ""
}

func (x *FlightplanExecutedEvent) GetFlightplan() *Flightplan {
	if x != nil {
		return x.Flightplan
	}
	return nil
}

type FleetCopiedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalFleetId string `protobuf:"bytes,1,opt,name=original_fleet_id,json=originalFleetId,proto3" json:"original_fleet_id,omitempty"`
	NewFleetId      string `protobuf:"bytes,2,opt,name=new_fleet_id,json=newFleetId,proto3" json:"new_fleet_id,omitempty"`
}

func (x *FleetCopiedEvent) Reset() {
	*x = FleetCopiedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleetCopiedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleetCopiedEvent) ProtoMessage() {}

func (x *FleetCopiedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleetCopiedEvent.ProtoReflect.Descriptor instead.
func (*FleetCopiedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{5}
}

func (x *FleetCopiedEvent) GetOriginalFleetId() string {
	if x != nil {
		return x.OriginalFleetId
	}
	return ""
}

func (x *FleetCopiedEvent) GetNewFleetId() string {
	if x != nil {
		return x.NewFleetId
	}
	return ""
}

type VehicleCopiedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FleetId           string `protobuf:"bytes,1,opt,name=fleet_id,json=fleetId,proto3" json:"fleet_id,omitempty"`
	OriginalVehicleId string `protobuf:"bytes,2,opt,name=original_vehicle_id,json=originalVehicleId,proto3" json:"original_vehicle_id,omitempty"`
	NewVehicleId      string `protobuf:"bytes,3,opt,name=new_vehicle_id,json=newVehicleId,proto3" json:"new_vehicle_id,omitempty"`
}

func (x *VehicleCopiedEvent) Reset() {
	*x = VehicleCopiedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleCopiedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleCopiedEvent) ProtoMessage() {}

func (x *VehicleCopiedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleCopiedEvent.ProtoReflect.Descriptor instead.
func (*VehicleCopiedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6}
}

func (x *VehicleCopiedEvent) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

func (x *VehicleCopiedEvent) GetOriginalVehicleId() string {
	if x != nil {
		return x.OriginalVehicleId
	}
	return ""
}

func (x *VehicleCopiedEvent) GetNewVehicleId() string {
	if x != nil {
		return x.NewVehicleId
	}
	return ""
}

type MissionCopiedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FleetId           string `protobuf:"bytes,1,opt,name=fleet_id,json=fleetId,proto3" json:"fleet_id,omitempty"`
	OriginalMissionId string `protobuf:"bytes,2,opt,name=original_mission_id,json=originalMissionId,proto3" json:"original_mission_id,omitempty"`
	NewMissionId      string `protobuf:"bytes,3,opt,name=new_mission_id,json=newMissionId,proto3" json:"new_mission_id,omitempty"`
}

func (x *MissionCopiedEvent) Reset() {
	*x = MissionCopiedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissionCopiedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissionCopiedEvent) ProtoMessage() {}

func (x *MissionCopiedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissionCopiedEvent.ProtoReflect.Descriptor instead.
func (*MissionCopiedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{7}
}

func (x *MissionCopiedEvent) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

func (x *MissionCopiedEvent) GetOriginalMissionId() string {
	if x != nil {
		return x.OriginalMissionId
	}
	return ""
}

func (x *MissionCopiedEvent) GetNewMissionId() string {
	if x != nil {
		return x.NewMissionId
	}
	return ""
}

type FlightoperationCompletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightoperationId string           `protobuf:"bytes,1,opt,name=flightoperation_id,json=flightoperationId,proto3" json:"flightoperation_id,omitempty"`
	Flightoperation   *Flightoperation `protobuf:"bytes,2,opt,name=flightoperation,proto3" json:"flightoperation,omitempty"`
}

func (x *FlightoperationCompletedEvent) Reset() {
	*x = FlightoperationCompletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightoperationCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightoperationCompletedEvent) ProtoMessage() {}

func (x *FlightoperationCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightoperationCompletedEvent.ProtoReflect.Descriptor instead.
func (*FlightoperationCompletedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{8}
}

func (x *FlightoperationCompletedEvent) GetFlightoperationId() string {
	if x != nil {
		return x.FlightoperationId
	}
	return ""
}

func (x *FlightoperationCompletedEvent) GetFlightoperation() *Flightoperation {
	if x != nil {
		return x.Flightoperation
	}
	return nil
}

type CopiedVehicleCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId       string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	CommunicationId string `protobuf:"bytes,2,opt,name=communication_id,json=communicationId,proto3" json:"communication_id,omitempty"`
	FleetId         string `protobuf:"bytes,3,opt,name=fleet_id,json=fleetId,proto3" json:"fleet_id,omitempty"`
}

func (x *CopiedVehicleCreatedEvent) Reset() {
	*x = CopiedVehicleCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopiedVehicleCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopiedVehicleCreatedEvent) ProtoMessage() {}

func (x *CopiedVehicleCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopiedVehicleCreatedEvent.ProtoReflect.Descriptor instead.
func (*CopiedVehicleCreatedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{9}
}

func (x *CopiedVehicleCreatedEvent) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *CopiedVehicleCreatedEvent) GetCommunicationId() string {
	if x != nil {
		return x.CommunicationId
	}
	return ""
}

func (x *CopiedVehicleCreatedEvent) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

type TelemetryUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunicationId string     `protobuf:"bytes,1,opt,name=communication_id,json=communicationId,proto3" json:"communication_id,omitempty"`
	Telemetry       *Telemetry `protobuf:"bytes,2,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
}

func (x *TelemetryUpdatedEvent) Reset() {
	*x = TelemetryUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryUpdatedEvent) ProtoMessage() {}

func (x *TelemetryUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryUpdatedEvent.ProtoReflect.Descriptor instead.
func (*TelemetryUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{10}
}

func (x *TelemetryUpdatedEvent) GetCommunicationId() string {
	if x != nil {
		return x.CommunicationId
	}
	return ""
}

func (x *TelemetryUpdatedEvent) GetTelemetry() *Telemetry {
	if x != nil {
		return x.Telemetry
	}
	return nil
}

type CopiedMissionCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissionId string   `protobuf:"bytes,1,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	Mission   *Mission `protobuf:"bytes,2,opt,name=mission,proto3" json:"mission,omitempty"`
}

func (x *CopiedMissionCreatedEvent) Reset() {
	*x = CopiedMissionCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopiedMissionCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopiedMissionCreatedEvent) ProtoMessage() {}

func (x *CopiedMissionCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopiedMissionCreatedEvent.ProtoReflect.Descriptor instead.
func (*CopiedMissionCreatedEvent) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{11}
}

func (x *CopiedMissionCreatedEvent) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *CopiedMissionCreatedEvent) GetMission() *Mission {
	if x != nil {
		return x.Mission
	}
	return nil
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x47, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x10, 0x46, 0x6c, 0x65, 0x65, 0x74,
	0x49, 0x64, 0x47, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x17, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x0a, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x22, 0x60, 0x0a, 0x10, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x43, 0x6f, 0x70, 0x69, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x46, 0x6c, 0x65, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43,
	0x6f, 0x70, 0x69, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x65, 0x77, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x12,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x70, 0x69, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x1d, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80,
	0x01, 0x0a, 0x19, 0x43, 0x6f, 0x70, 0x69, 0x65, 0x64, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x49,
	0x64, 0x22, 0x7a, 0x0a, 0x15, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x6c, 0x0a,
	0x19, 0x43, 0x6f, 0x70, 0x69, 0x65, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x0a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x6b, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_event_proto_goTypes = []interface{}{
	(*CommunicationIdGaveEvent)(nil),      // 0: skysign_proto.CommunicationIdGaveEvent
	(*CommunicationIdRemovedEvent)(nil),   // 1: skysign_proto.CommunicationIdRemovedEvent
	(*FleetIdGaveEvent)(nil),              // 2: skysign_proto.FleetIdGaveEvent
	(*FleetIdRemovedEvent)(nil),           // 3: skysign_proto.FleetIdRemovedEvent
	(*FlightplanExecutedEvent)(nil),       // 4: skysign_proto.FlightplanExecutedEvent
	(*FleetCopiedEvent)(nil),              // 5: skysign_proto.FleetCopiedEvent
	(*VehicleCopiedEvent)(nil),            // 6: skysign_proto.VehicleCopiedEvent
	(*MissionCopiedEvent)(nil),            // 7: skysign_proto.MissionCopiedEvent
	(*FlightoperationCompletedEvent)(nil), // 8: skysign_proto.FlightoperationCompletedEvent
	(*CopiedVehicleCreatedEvent)(nil),     // 9: skysign_proto.CopiedVehicleCreatedEvent
	(*TelemetryUpdatedEvent)(nil),         // 10: skysign_proto.TelemetryUpdatedEvent
	(*CopiedMissionCreatedEvent)(nil),     // 11: skysign_proto.CopiedMissionCreatedEvent
	(*Flightplan)(nil),                    // 12: skysign_proto.Flightplan
	(*Flightoperation)(nil),               // 13: skysign_proto.Flightoperation
	(*Telemetry)(nil),                     // 14: skysign_proto.Telemetry
	(*Mission)(nil),                       // 15: skysign_proto.Mission
}
var file_proto_event_proto_depIdxs = []int32{
	12, // 0: skysign_proto.FlightplanExecutedEvent.flightplan:type_name -> skysign_proto.Flightplan
	13, // 1: skysign_proto.FlightoperationCompletedEvent.flightoperation:type_name -> skysign_proto.Flightoperation
	14, // 2: skysign_proto.TelemetryUpdatedEvent.telemetry:type_name -> skysign_proto.Telemetry
	15, // 3: skysign_proto.CopiedMissionCreatedEvent.mission:type_name -> skysign_proto.Mission
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicationIdGaveEvent); i {
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
		file_proto_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicationIdRemovedEvent); i {
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
		file_proto_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleetIdGaveEvent); i {
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
		file_proto_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleetIdRemovedEvent); i {
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
		file_proto_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightplanExecutedEvent); i {
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
		file_proto_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleetCopiedEvent); i {
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
		file_proto_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleCopiedEvent); i {
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
		file_proto_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissionCopiedEvent); i {
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
		file_proto_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightoperationCompletedEvent); i {
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
		file_proto_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopiedVehicleCreatedEvent); i {
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
		file_proto_event_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryUpdatedEvent); i {
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
		file_proto_event_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopiedMissionCreatedEvent); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}
