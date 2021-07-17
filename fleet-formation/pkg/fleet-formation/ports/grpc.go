package ports

import (
	"context"

	"fleet-formation/pkg/fleet-formation/app"
	"fleet-formation/pkg/fleet-formation/service"
	proto "fleet-formation/pkg/skysign_proto"

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

// GetAssignments .
func (s *GrpcServer) GetAssignments(
	ctx context.Context,
	request *proto.GetAssignmentsRequest,
) (*proto.GetAssignmentsResponse, error) {
	response := &proto.GetAssignmentsResponse{
		Id: request.Id,
	}
	command := &fleetIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.AssignFleet.GetAssignments(
		command,
		func(model service.AssignmentPresentationModel) {
			response.Assignments = append(
				response.Assignments,
				&proto.Assignment{
					Id:           model.GetAssignment().GetEventID(),
					AssignmentId: model.GetAssignment().GetAssignmentID(),
					VehicleId:    model.GetAssignment().GetVehicleID(),
					MissionId:    model.GetAssignment().GetMissionID(),
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// UpdateAssignments .
func (s *GrpcServer) UpdateAssignments(
	ctx context.Context,
	request *proto.UpdateAssignmentsRequest,
) (*proto.UpdateAssignmentsResponse, error) {
	response := &proto.UpdateAssignmentsResponse{
		Id: request.Id,
	}
	for _, assignment := range request.Assignments {
		command := &updateAssignmentsCommand{
			id:         request.Id,
			assignment: assignment,
		}
		if ret := s.app.Services.AssignFleet.UpdateAssignment(
			command,
			func(model service.AssignmentPresentationModel) {
				response.Assignments = append(
					response.Assignments,
					&proto.Assignment{
						Id:           model.GetAssignment().GetEventID(),
						AssignmentId: model.GetAssignment().GetAssignmentID(),
						VehicleId:    model.GetAssignment().GetVehicleID(),
						MissionId:    model.GetAssignment().GetMissionID(),
					},
				)
			},
		); ret != nil {
			return nil, ret
		}
	}
	return response, nil
}

type fleetIDCommand struct {
	id string
}

func (f *fleetIDCommand) GetID() string {
	return f.id
}

type updateAssignmentsCommand struct {
	id         string
	assignment *proto.Assignment
}

func (r *updateAssignmentsCommand) GetID() string {
	return r.id
}
func (r *updateAssignmentsCommand) GetEventID() string {
	return r.assignment.Id
}
func (r *updateAssignmentsCommand) GetAssignmentID() string {
	return r.assignment.AssignmentId
}
func (r *updateAssignmentsCommand) GetVehicleID() string {
	return r.assignment.VehicleId
}
func (r *updateAssignmentsCommand) GetMissionID() string {
	return r.assignment.MissionId
}
