package ports

import (
	"context"

	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"
	proto "remote-communication/pkg/skysign_proto"

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

// GetUploadMission .
func (s *GrpcServer) GetUploadMission(
	ctx context.Context,
	request *proto.GetUploadMissionRequest,
) (*proto.UploadMission, error) {
	response := &proto.UploadMission{}
	command := &missionIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.EdgeMission.PullMission(
		command,
		func(id string, waypoints []service.Waypoint) {
			response.Id = request.Id
			for _, w := range waypoints {
				response.Waypoints = append(
					response.Waypoints,
					&proto.Waypoint{
						Latitude:       w.GetLatitudeDegree(),
						Longitude:      w.GetLongitudeDegree(),
						RelativeHeight: w.GetRelativeHeightM(),
						Speed:          w.GetSpeedMS(),
					},
				)
			}
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

type missionIDCommand struct {
	id string
}

func (f *missionIDCommand) GetID() string {
	return f.id
}
