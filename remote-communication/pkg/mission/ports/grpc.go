package ports

import (
	"context"

	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// GrpcServer .
type GrpcServer struct {
	proto.UnimplementedUploadMissionEdgeServiceServer
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
						Latitude:       w.GetLatitude(),
						Longitude:      w.GetLongitude(),
						RelativeHeight: w.GetRelativeAltitude(),
						Speed:          w.GetSpeed(),
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
