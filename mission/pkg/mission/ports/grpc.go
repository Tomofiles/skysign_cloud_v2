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

// // ListVehicles .
// func (s *GrpcServer) ListVehicles(
// 	ctx context.Context,
// 	request *proto.Empty,
// ) (*proto.ListVehiclesResponses, error) {
// 	response := &proto.ListVehiclesResponses{}
// 	if ret := s.app.Services.ManageVehicle.ListVehicles(
// 		func(id, name, communicationID string) {
// 			response.Vehicles = append(
// 				response.Vehicles,
// 				&proto.Vehicle{
// 					Id:              id,
// 					Name:            name,
// 					CommunicationId: communicationID,
// 				},
// 			)
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// // GetVehicle .
// func (s *GrpcServer) GetVehicle(
// 	ctx context.Context,
// 	request *proto.GetVehicleRequest,
// ) (*proto.Vehicle, error) {
// 	response := &proto.Vehicle{}
// 	requestDpo := &vehicleIDRequestDpo{
// 		id: request.Id,
// 	}
// 	if ret := s.app.Services.ManageVehicle.GetVehicle(
// 		requestDpo,
// 		func(id, name, communicationID string) {
// 			response.Id = id
// 			response.Name = name
// 			response.CommunicationId = communicationID
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// CreateMission .
func (s *GrpcServer) CreateMission(
	ctx context.Context,
	request *proto.Mission,
) (*proto.Mission, error) {
	response := &proto.Mission{}
	command := &missionCommand{
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
	command := &missionCommand{
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

type missionCommand struct {
	request *proto.Mission
}

func (f *missionCommand) GetID() string {
	return f.request.Id
}

func (f *missionCommand) GetName() string {
	return f.request.Name
}

func (f *missionCommand) GetNavigation() service.Navigation {
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
