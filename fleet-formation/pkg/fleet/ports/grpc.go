package ports

import (
	"context"

	"fleet-formation/pkg/fleet/app"
	"fleet-formation/pkg/fleet/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// GrpcServer .
type GrpcServer struct {
	proto.UnimplementedAssignAssetsToFleetServiceServer
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
