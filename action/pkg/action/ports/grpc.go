package ports

import (
	"context"

	"action/pkg/action/app"
	"action/pkg/action/domain/action"
	proto "action/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// LogBodyInterceptor .
func LogBodyInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		glog.Infof("REQUEST , API: %s, Message: %+v", info.FullMethod, req)
		defer func() {
			if err != nil {
				glog.Errorf("RESPONSE, API: %s, Error: %+v", info.FullMethod, err)
			} else {
				glog.Infof("RESPONSE, API: %s, Message: %+v", info.FullMethod, resp)
			}
		}()

		resp, err = handler(ctx, req)
		return
	}
}

// GrpcServer .
type GrpcServer struct {
	app app.Application
}

// NewGrpcServer .
func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

// GetFlightplan .
func (s *GrpcServer) GetTrajectory(
	ctx context.Context,
	request *proto.GetTrajectoryRequest,
) (*proto.GetTrajectoryResponse, error) {
	response := &proto.GetTrajectoryResponse{}
	requestDpo := &actionIDRequestDpo{
		id: request.VehicleId,
	}
	if ret := s.app.Services.ManageAction.GetTrajectory(
		requestDpo,
		func(snapshot action.TelemetrySnapshot) {
			response.Telemetries = append(
				response.Telemetries,
				&proto.Telemetry{
					Latitude:         snapshot.Latitude,
					Longitude:        snapshot.Longitude,
					Altitude:         snapshot.Altitude,
					RelativeAltitude: snapshot.RelativeAltitude,
					Speed:            snapshot.Speed,
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

type actionIDRequestDpo struct {
	id string
}

func (f *actionIDRequestDpo) GetID() string {
	return f.id
}
