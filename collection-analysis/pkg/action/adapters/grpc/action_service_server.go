package grpc

import (
	"context"

	"collection-analysis/pkg/action/app"
	"collection-analysis/pkg/action/domain/action"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type actionServiceServer struct {
	proto.UnimplementedActionServiceServer
	app app.Application
}

// NewActionServiceServer .
func NewActionServiceServer(application app.Application) proto.ActionServiceServer {
	return &actionServiceServer{app: application}
}

// GetFlightplan .
func (s *actionServiceServer) GetTrajectory(
	ctx context.Context,
	request *proto.GetTrajectoryRequest,
) (*proto.GetTrajectoryResponse, error) {
	response := &proto.GetTrajectoryResponse{}
	command := &actionIDCommand{
		id: request.VehicleId,
	}
	if ret := s.app.Services.ManageAction.GetTrajectory(
		command,
		func(snapshot action.TelemetrySnapshot) {
			response.Telemetries = append(
				response.Telemetries,
				&proto.Telemetry{
					Latitude:         snapshot.LatitudeDegree,
					Longitude:        snapshot.LongitudeDegree,
					Altitude:         snapshot.AltitudeM,
					RelativeAltitude: snapshot.RelativeAltitudeM,
					Speed:            snapshot.SpeedMS,
					Armed:            snapshot.Armed,
					FlightMode:       snapshot.FlightMode,
					OrientationX:     snapshot.OrientationX,
					OrientationY:     snapshot.OrientationY,
					OrientationZ:     snapshot.OrientationZ,
					OrientationW:     snapshot.OrientationW,
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

type actionIDCommand struct {
	id string
}

func (f *actionIDCommand) GetID() string {
	return f.id
}
