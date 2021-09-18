package grpc

import (
	"context"

	"remote-communication/pkg/communication/app"
	"remote-communication/pkg/communication/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// communicationEdgeServiceServer .
type communicationEdgeServiceServer struct {
	proto.UnimplementedCommunicationEdgeServiceServer
	app app.Application
}

// NewCommunicationEdgeServiceServer .
func NewCommunicationEdgeServiceServer(application app.Application) proto.CommunicationEdgeServiceServer {
	return &communicationEdgeServiceServer{app: application}
}

// PullCommand .
func (s *communicationEdgeServiceServer) PullCommand(
	ctx context.Context,
	request *proto.PullCommandRequest,
) (*proto.PullCommandResponse, error) {
	response := &proto.PullCommandResponse{}
	command := &pullCommand{
		id:        request.Id,
		commandID: request.CommandId,
	}
	if ret := s.app.Services.EdgeCommunication.PullCommand(
		command,
		func(cType string) {
			response.Type = proto.CommandType(proto.CommandType_value[cType])
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.CommandId = request.CommandId
	return response, nil
}

// PullUploadMission .
func (s *communicationEdgeServiceServer) PullUploadMission(
	ctx context.Context,
	request *proto.PullUploadMissionRequest,
) (*proto.PullUploadMissionResponse, error) {
	response := &proto.PullUploadMissionResponse{}
	command := &pullCommand{
		id:        request.Id,
		commandID: request.CommandId,
	}
	if ret := s.app.Services.EdgeCommunication.PullUploadMission(
		command,
		func(missionID string) {
			response.MissionId = missionID
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.CommandId = request.CommandId
	return response, nil
}

// PushTelemetry .
func (s *communicationEdgeServiceServer) PushTelemetry(
	ctx context.Context,
	request *proto.PushTelemetryRequest,
) (*proto.PushTelemetryResponse, error) {
	response := &proto.PushTelemetryResponse{}
	command := &pushTelemetryCommand{
		id: request.Id,
		telemetry: &telemetry{
			latitudeDegree:    request.Telemetry.Latitude,
			longitudeDegree:   request.Telemetry.Longitude,
			altitudeM:         request.Telemetry.Altitude,
			relativeAltitudeM: request.Telemetry.RelativeAltitude,
			speedMS:           request.Telemetry.Speed,
			armed:             request.Telemetry.Armed,
			flightMode:        request.Telemetry.FlightMode,
			x:                 request.Telemetry.OrientationX,
			y:                 request.Telemetry.OrientationY,
			z:                 request.Telemetry.OrientationZ,
			w:                 request.Telemetry.OrientationW,
		},
	}
	if ret := s.app.Services.EdgeCommunication.PushTelemetry(
		command,
		func(commandIDs []string) {
			response.CommandIds = append(response.CommandIds, commandIDs...)
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	return response, nil
}

type pullCommand struct {
	id, commandID string
}

func (f *pullCommand) GetID() string {
	return f.id
}

func (f *pullCommand) GetCommandID() string {
	return f.commandID
}

type pushTelemetryCommand struct {
	id        string
	telemetry *telemetry
}

func (f *pushTelemetryCommand) GetID() string {
	return f.id
}

func (f *pushTelemetryCommand) GetTelemetry() service.EdgeTelemetry {
	return f.telemetry
}

type telemetry struct {
	latitudeDegree    float64
	longitudeDegree   float64
	altitudeM         float64
	relativeAltitudeM float64
	speedMS           float64
	armed             bool
	flightMode        string
	x                 float64
	y                 float64
	z                 float64
	w                 float64
}

func (t *telemetry) GetLatitudeDegree() float64 {
	return t.latitudeDegree
}

func (t *telemetry) GetLongitudeDegree() float64 {
	return t.longitudeDegree
}

func (t *telemetry) GetAltitudeM() float64 {
	return t.altitudeM
}

func (t *telemetry) GetRelativeAltitudeM() float64 {
	return t.relativeAltitudeM
}

func (t *telemetry) GetSpeedMS() float64 {
	return t.speedMS
}

func (t *telemetry) GetArmed() bool {
	return t.armed
}

func (t *telemetry) GetFlightMode() string {
	return t.flightMode
}

func (t *telemetry) GetX() float64 {
	return t.x
}

func (t *telemetry) GetY() float64 {
	return t.y
}

func (t *telemetry) GetZ() float64 {
	return t.z
}

func (t *telemetry) GetW() float64 {
	return t.w
}
