package ports

import (
	"context"

	"remote-communication/pkg/communication/app"
	"remote-communication/pkg/communication/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// GrpcServer .
type GrpcServer struct {
	proto.UnimplementedCommunicationUserServiceServer
	proto.UnimplementedCommunicationEdgeServiceServer
	app app.Application
}

// NewGrpcServer .
func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

// PushCommand .
func (s *GrpcServer) PushCommand(
	ctx context.Context,
	request *proto.PushCommandRequest,
) (*proto.PushCommandResponse, error) {
	response := &proto.PushCommandResponse{}
	command := &pushCommandCommand{
		id:    request.Id,
		cType: request.Type.String(),
	}
	if ret := s.app.Services.UserCommunication.PushCommand(
		command,
		func(commandID string) {
			response.CommandId = commandID
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.Type = request.Type
	return response, nil
}

// PullCommand .
func (s *GrpcServer) PullCommand(
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

// PushUploadMission .
func (s *GrpcServer) PushUploadMission(
	ctx context.Context,
	request *proto.PushUploadMissionRequest,
) (*proto.PushUploadMissionResponse, error) {
	response := &proto.PushUploadMissionResponse{}
	command := &pushUploadMissionCommand{
		id:        request.Id,
		missionID: request.MissionId,
	}
	if ret := s.app.Services.UserCommunication.PushUploadMission(
		command,
		func(commandID string) {
			response.CommandId = commandID
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.MissionId = request.MissionId
	return response, nil
}

// PullUploadMission .
func (s *GrpcServer) PullUploadMission(
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
func (s *GrpcServer) PushTelemetry(
	ctx context.Context,
	request *proto.PushTelemetryRequest,
) (*proto.PushTelemetryResponse, error) {
	response := &proto.PushTelemetryResponse{}
	command := &pushTelemetryCommand{
		id: request.Id,
		telemetry: &telemetry{
			latitude:         request.Telemetry.Latitude,
			longitude:        request.Telemetry.Longitude,
			altitude:         request.Telemetry.Altitude,
			relativeAltitude: request.Telemetry.RelativeAltitude,
			speed:            request.Telemetry.Speed,
			armed:            request.Telemetry.Armed,
			flightMode:       request.Telemetry.FlightMode,
			x:                request.Telemetry.OrientationX,
			y:                request.Telemetry.OrientationY,
			z:                request.Telemetry.OrientationZ,
			w:                request.Telemetry.OrientationW,
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

// PullTelemetry .
func (s *GrpcServer) PullTelemetry(
	ctx context.Context,
	request *proto.PullTelemetryRequest,
) (*proto.PullTelemetryResponse, error) {
	response := &proto.PullTelemetryResponse{}
	command := &communicationIDCommand{
		id: request.Id,
	}
	response.Telemetry = &proto.Telemetry{}
	if ret := s.app.Services.UserCommunication.PullTelemetry(
		command,
		func(telemetry service.UserTelemetry) {
			response.Telemetry.Latitude = telemetry.GetLatitude()
			response.Telemetry.Longitude = telemetry.GetLongitude()
			response.Telemetry.Altitude = telemetry.GetAltitude()
			response.Telemetry.RelativeAltitude = telemetry.GetRelativeAltitude()
			response.Telemetry.Speed = telemetry.GetSpeed()
			response.Telemetry.Armed = telemetry.GetArmed()
			response.Telemetry.FlightMode = telemetry.GetFlightMode()
			response.Telemetry.OrientationX = telemetry.GetX()
			response.Telemetry.OrientationY = telemetry.GetY()
			response.Telemetry.OrientationZ = telemetry.GetZ()
			response.Telemetry.OrientationW = telemetry.GetW()
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	return response, nil
}

type communicationIDCommand struct {
	id string
}

func (f *communicationIDCommand) GetID() string {
	return f.id
}

type pushCommandCommand struct {
	id, cType string
}

func (f *pushCommandCommand) GetID() string {
	return f.id
}

func (f *pushCommandCommand) GetType() string {
	return f.cType
}

type pushUploadMissionCommand struct {
	id, missionID string
}

func (f *pushUploadMissionCommand) GetID() string {
	return f.id
}

func (f *pushUploadMissionCommand) GetMissionID() string {
	return f.missionID
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
	latitude         float64
	longitude        float64
	altitude         float64
	relativeAltitude float64
	speed            float64
	armed            bool
	flightMode       string
	x                float64
	y                float64
	z                float64
	w                float64
}

func (t *telemetry) GetLatitude() float64 {
	return t.latitude
}

func (t *telemetry) GetLongitude() float64 {
	return t.longitude
}

func (t *telemetry) GetAltitude() float64 {
	return t.altitude
}

func (t *telemetry) GetRelativeAltitude() float64 {
	return t.relativeAltitude
}

func (t *telemetry) GetSpeed() float64 {
	return t.speed
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
