package grpc

import (
	"context"

	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// uploadMissionEdgeServiceServer .
type uploadMissionEdgeServiceServer struct {
	proto.UnimplementedUploadMissionEdgeServiceServer
	app app.Application
}

// NewUploadMissionEdgeServiceServer .
func NewUploadMissionEdgeServiceServer(application app.Application) *uploadMissionEdgeServiceServer {
	return &uploadMissionEdgeServiceServer{app: application}
}

// GetUploadMission .
func (s *uploadMissionEdgeServiceServer) GetUploadMission(
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
						Latitude:         w.GetLatitudeDegree(),
						Longitude:        w.GetLongitudeDegree(),
						RelativeAltitude: w.GetRelativeAltitudeM(),
						Speed:            w.GetSpeedMS(),
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

type waypoint struct {
	LatitudeDegree, LongitudeDegree, RelativeAltitudeM, SpeedMS float64
}

func (v *waypoint) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

func (v *waypoint) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

func (v *waypoint) GetRelativeAltitudeM() float64 {
	return v.RelativeAltitudeM
}

func (v *waypoint) GetSpeedMS() float64 {
	return v.SpeedMS
}
