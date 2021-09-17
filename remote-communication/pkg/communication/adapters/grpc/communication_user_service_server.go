package grpc

import (
	"context"

	"remote-communication/pkg/communication/app"
	"remote-communication/pkg/communication/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// communicationUserServiceServer .
type communicationUserServiceServer struct {
	proto.UnimplementedCommunicationUserServiceServer
	app app.Application
}

// NewCommunicationUserServiceServer .
func NewCommunicationUserServiceServer(application app.Application) *communicationUserServiceServer {
	return &communicationUserServiceServer{app: application}
}

// PushCommand .
func (s *communicationUserServiceServer) PushCommand(
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

// PushUploadMission .
func (s *communicationUserServiceServer) PushUploadMission(
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

// PullTelemetry .
func (s *communicationUserServiceServer) PullTelemetry(
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
			response.Telemetry.Latitude = telemetry.GetLatitudeDegree()
			response.Telemetry.Longitude = telemetry.GetLongitudeDegree()
			response.Telemetry.Altitude = telemetry.GetAltitudeM()
			response.Telemetry.RelativeAltitude = telemetry.GetRelativeAltitudeM()
			response.Telemetry.Speed = telemetry.GetSpeedMS()
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
