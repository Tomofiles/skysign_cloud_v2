package grpc

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/adapters/proto"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/app"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// manageMissionServiceServer .
type manageMissionServiceServer struct {
	skysign_proto.UnimplementedManageMissionServiceServer
	app app.Application
}

// NewManageMissionServiceServer .
func NewManageMissionServiceServer(application app.Application) skysign_proto.ManageMissionServiceServer {
	return &manageMissionServiceServer{app: application}
}

// ListMissions .
func (s *manageMissionServiceServer) ListMissions(
	ctx context.Context,
	request *skysign_proto.Empty,
) (*skysign_proto.ListMissionsResponses, error) {
	response := &skysign_proto.ListMissionsResponses{}
	if ret := s.app.Services.ManageMission.ListMissions(
		func(model service.MissionPresentationModel) {
			mission := MissionProtoTransformerFromModel(model)
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
func (s *manageMissionServiceServer) GetMission(
	ctx context.Context,
	request *skysign_proto.GetMissionRequest,
) (*skysign_proto.Mission, error) {
	var response *skysign_proto.Mission
	command := &missionIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageMission.GetMission(
		command,
		func(model service.MissionPresentationModel) {
			response = MissionProtoTransformerFromModel(model)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CreateMission .
func (s *manageMissionServiceServer) CreateMission(
	ctx context.Context,
	request *skysign_proto.Mission,
) (*skysign_proto.Mission, error) {
	if ret := proto.ValidateCreateMissionRequest(request); ret != nil {
		return nil, ret
	}
	response := &skysign_proto.Mission{}
	command := &createCommand{
		request: request,
	}
	if ret := s.app.Services.ManageMission.CreateMission(
		command,
		func(id string) {
			response.Id = id
		},
		func(uploadID string) {
			request.Navigation.UploadId = uploadID
		},
	); ret != nil {
		return nil, ret
	}
	response.Name = request.Name
	response.Navigation = request.Navigation
	return response, nil
}

// UpdateMission .
func (s *manageMissionServiceServer) UpdateMission(
	ctx context.Context,
	request *skysign_proto.Mission,
) (*skysign_proto.Mission, error) {
	response := &skysign_proto.Mission{}
	command := &updateCommand{
		request: request,
	}
	if ret := s.app.Services.ManageMission.UpdateMission(
		command,
		func(uploadID string) {
			request.Navigation.UploadId = uploadID
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.Name = request.Name
	response.Navigation = request.Navigation
	return response, nil
}

// DeleteMission .
func (s *manageMissionServiceServer) DeleteMission(
	ctx context.Context,
	request *skysign_proto.DeleteMissionRequest,
) (*skysign_proto.Empty, error) {
	response := &skysign_proto.Empty{}
	command := &missionIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageMission.DeleteMission(command); ret != nil {
		return nil, ret
	}
	return response, nil
}

type createCommand struct {
	request *skysign_proto.Mission
}

func (f *createCommand) GetMission() service.Mission {
	return &mission{
		request: f.request,
	}
}

type updateCommand struct {
	request *skysign_proto.Mission
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
	request *skysign_proto.Mission
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
				latitudeDegree:    w.Latitude,
				longitudeDegree:   w.Longitude,
				relativeAltitudeM: w.RelativeAltitude,
				speedMS:           w.Speed,
			},
		)
	}
	navigation := &navigation{
		takeoffPointGroundAltitudeM: f.request.Navigation.TakeoffPointGroundAltitude,
		waypoints:                   waypoints,
	}
	return navigation
}

type navigation struct {
	takeoffPointGroundAltitudeM float64
	waypoints                   []waypoint
	uploadID                    string
}

func (f *navigation) GetTakeoffPointGroundAltitudeM() float64 {
	return f.takeoffPointGroundAltitudeM
}

func (f *navigation) GetWaypoints() []service.Waypoint {
	waypoints := []service.Waypoint{}
	for _, w := range f.waypoints {
		waypoints = append(
			waypoints,
			&waypoint{
				latitudeDegree:    w.latitudeDegree,
				longitudeDegree:   w.longitudeDegree,
				relativeAltitudeM: w.relativeAltitudeM,
				speedMS:           w.speedMS,
			},
		)
	}
	return waypoints
}

func (f *navigation) GetUploadID() string {
	return f.uploadID
}

type waypoint struct {
	latitudeDegree    float64
	longitudeDegree   float64
	relativeAltitudeM float64
	speedMS           float64
}

func (f *waypoint) GetLatitudeDegree() float64 {
	return f.latitudeDegree
}

func (f *waypoint) GetLongitudeDegree() float64 {
	return f.longitudeDegree
}

func (f *waypoint) GetRelativeAltitudeM() float64 {
	return f.relativeAltitudeM
}

func (f *waypoint) GetSpeedMS() float64 {
	return f.speedMS
}

type missionIDCommand struct {
	id string
}

func (f *missionIDCommand) GetID() string {
	return f.id
}
