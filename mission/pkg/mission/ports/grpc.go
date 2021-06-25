package ports

import (
	"context"

	"mission/pkg/mission/app"
	"mission/pkg/mission/service"
	proto "mission/pkg/skysign_proto"

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

// ListMissions .
func (s *GrpcServer) ListMissions(
	ctx context.Context,
	request *proto.Empty,
) (*proto.ListMissionsResponses, error) {
	response := &proto.ListMissionsResponses{}
	if ret := s.app.Services.ManageMission.ListMissions(
		func(model service.MissionPresentationModel) {
			mission := MissionProtoTransfomerFromModel(model)
			response.Missions = append(
				response.Missions,
				mission,
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetMission .
func (s *GrpcServer) GetMission(
	ctx context.Context,
	request *proto.GetMissionRequest,
) (*proto.Mission, error) {
	var response *proto.Mission
	command := &missionIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageMission.GetMission(
		command,
		func(model service.MissionPresentationModel) {
			response = MissionProtoTransfomerFromModel(model)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CreateMission .
func (s *GrpcServer) CreateMission(
	ctx context.Context,
	request *proto.Mission,
) (*proto.Mission, error) {
	response := &proto.Mission{}
	command := &createCommand{
		request: request,
	}
	if ret := s.app.Services.ManageMission.CreateMission(
		command,
		func(id string) {
			response.Id = id
		},
	); ret != nil {
		return nil, ret
	}
	response.Name = request.Name
	response.Navigation = request.Navigation
	return response, nil
}

// UpdateMission .
func (s *GrpcServer) UpdateMission(
	ctx context.Context,
	request *proto.Mission,
) (*proto.Mission, error) {
	response := &proto.Mission{}
	command := &updateCommand{
		request: request,
	}
	if ret := s.app.Services.ManageMission.UpdateMission(command); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.Name = request.Name
	response.Navigation = request.Navigation
	return response, nil
}

// DeleteMission .
func (s *GrpcServer) DeleteMission(
	ctx context.Context,
	request *proto.DeleteMissionRequest,
) (*proto.Empty, error) {
	response := &proto.Empty{}
	command := &missionIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageMission.DeleteMission(command); ret != nil {
		return nil, ret
	}
	return response, nil
}

type createCommand struct {
	request *proto.Mission
}

func (f *createCommand) GetMission() service.Mission {
	return &mission{
		request: f.request,
	}
}

type updateCommand struct {
	request *proto.Mission
}

func (f *updateCommand) GetID() string {
	return f.request.Id
}

func (f *updateCommand) GetMission() service.Mission {
	return &mission{
		request: f.request,
	}
}

type mission struct {
	request *proto.Mission
}

func (f *mission) GetID() string {
	return f.request.Id
}

func (f *mission) GetName() string {
	return f.request.Name
}

func (f *mission) GetNavigation() service.Navigation {
	waypoints := []waypoint{}
	for _, w := range f.request.Navigation.Waypoints {
		waypoints = append(
			waypoints,
			waypoint{
				latitude:       w.Latitude,
				longitude:      w.Longitude,
				relativeHeight: w.RelativeHeight,
				speed:          w.Speed,
			},
		)
	}
	navigation := &navigation{
		takeoffPointGroundHeight: f.request.Navigation.TakeoffPointGroundHeight,
		waypoints:                waypoints,
	}
	return navigation
}

type navigation struct {
	takeoffPointGroundHeight float64
	waypoints                []waypoint
}

func (f *navigation) GetTakeoffPointGroundHeight() float64 {
	return f.takeoffPointGroundHeight
}

func (f *navigation) GetWaypoints() []service.Waypoint {
	waypoints := []service.Waypoint{}
	for _, w := range f.waypoints {
		waypoints = append(
			waypoints,
			&waypoint{
				latitude:       w.latitude,
				longitude:      w.longitude,
				relativeHeight: w.relativeHeight,
				speed:          w.speed,
			},
		)
	}
	return waypoints
}

type waypoint struct {
	latitude       float64
	longitude      float64
	relativeHeight float64
	speed          float64
}

func (f *waypoint) GetLatitude() float64 {
	return f.latitude
}

func (f *waypoint) GetLongitude() float64 {
	return f.longitude
}

func (f *waypoint) GetRelativeHeight() float64 {
	return f.relativeHeight
}

func (f *waypoint) GetSpeed() float64 {
	return f.speed
}

type missionIDCommand struct {
	id string
}

func (f *missionIDCommand) GetID() string {
	return f.id
}
