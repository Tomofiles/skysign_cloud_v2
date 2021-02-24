package ports

import (
	"context"

	"flightplan/pkg/flightplan/app"
	proto "flightplan/pkg/skysign_proto"

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

// ListFlightplans .
func (s *GrpcServer) ListFlightplans(
	ctx context.Context,
	request *proto.Empty,
) (*proto.ListFlightplansResponses, error) {
	response := &proto.ListFlightplansResponses{}
	if ret := s.app.Services.ManageFlightplan.ListFlightplans(
		func(id, name, description string) {
			response.Flightplans = append(
				response.Flightplans,
				&proto.Flightplan{
					Id:          id,
					Name:        name,
					Description: description,
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetFlightplan .
func (s *GrpcServer) GetFlightplan(
	ctx context.Context,
	request *proto.GetFlightplanRequest,
) (*proto.Flightplan, error) {
	response := &proto.Flightplan{}
	requestDpo := &flightplanIDRequestDpo{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightplan.GetFlightplan(
		requestDpo,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CreateFlightplan .
func (s *GrpcServer) CreateFlightplan(
	ctx context.Context,
	request *proto.Flightplan,
) (*proto.Flightplan, error) {
	response := &proto.Flightplan{}
	requestDpo := &flightplanRequestDpo{
		name:        request.Name,
		description: request.Description,
	}
	if ret := s.app.Services.ManageFlightplan.CreateFlightplan(
		requestDpo,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// UpdateFlightplan .
func (s *GrpcServer) UpdateFlightplan(
	ctx context.Context,
	request *proto.Flightplan,
) (*proto.Flightplan, error) {
	response := &proto.Flightplan{}
	requestDpo := &flightplanRequestDpo{
		id:          request.Id,
		name:        request.Name,
		description: request.Description,
	}
	if ret := s.app.Services.ManageFlightplan.UpdateFlightplan(
		requestDpo,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// DeleteFlightplan .
func (s *GrpcServer) DeleteFlightplan(
	ctx context.Context,
	request *proto.DeleteFlightplanRequest,
) (*proto.Empty, error) {
	response := &proto.Empty{}
	requestDpo := &flightplanIDRequestDpo{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightplan.DeleteFlightplan(
		requestDpo,
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// ChangeNumberOfVehicles .
func (s *GrpcServer) ChangeNumberOfVehicles(
	ctx context.Context,
	request *proto.ChangeNumberOfVehiclesRequest,
) (*proto.ChangeNumberOfVehiclesResponse, error) {
	response := &proto.ChangeNumberOfVehiclesResponse{}
	requestDpo := &changeNumberOfVehiclesRequestDpo{
		flightplanID:     request.Id,
		numberOfVehicles: request.NumberOfVehicles,
	}
	if ret := s.app.Services.AssignFleet.ChangeNumberOfVehicles(
		requestDpo,
		func(id string, numberOfVehicles int32) {
			response.Id = id
			response.NumberOfVehicles = numberOfVehicles
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetAssignments .
func (s *GrpcServer) GetAssignments(
	ctx context.Context,
	request *proto.GetAssignmentsRequest,
) (*proto.GetAssignmentsResponse, error) {
	response := &proto.GetAssignmentsResponse{
		Id: request.Id,
	}
	requestDpo := &fleetIDRequestDpo{
		flightplanID: request.Id,
	}
	if ret := s.app.Services.AssignFleet.GetAssignments(
		requestDpo,
		func(id, assignmentID, vehicleID, missionID string) {
			response.Assignments = append(
				response.Assignments,
				&proto.Assignment{
					Id:           id,
					AssignmentId: assignmentID,
					VehicleId:    vehicleID,
					MissionId:    missionID,
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
		requestDpo := &updateAssignmentsRequestDpo{
			flightplanID: request.Id,
			assignment:   assignment,
		}
		if ret := s.app.Services.AssignFleet.UpdateAssignment(
			requestDpo,
			func(id, assignmentID, vehicleID, missionID string) {
				response.Assignments = append(
					response.Assignments,
					&proto.Assignment{
						Id:           id,
						AssignmentId: assignmentID,
						VehicleId:    vehicleID,
						MissionId:    missionID,
					},
				)
			},
		); ret != nil {
			return nil, ret
		}
	}
	return response, nil
}

type flightplanRequestDpo struct {
	id          string
	name        string
	description string
}

func (f *flightplanRequestDpo) GetID() string {
	return f.id
}

func (f *flightplanRequestDpo) GetName() string {
	return f.name
}

func (f *flightplanRequestDpo) GetDescription() string {
	return f.description
}

type flightplanIDRequestDpo struct {
	id string
}

func (f *flightplanIDRequestDpo) GetID() string {
	return f.id
}

type fleetIDRequestDpo struct {
	flightplanID string
}

func (f *fleetIDRequestDpo) GetFlightplanID() string {
	return f.flightplanID
}

type changeNumberOfVehiclesRequestDpo struct {
	flightplanID     string
	numberOfVehicles int32
}

func (c *changeNumberOfVehiclesRequestDpo) GetFlightplanID() string {
	return c.flightplanID
}

func (c *changeNumberOfVehiclesRequestDpo) GetNumberOfVehicles() int32 {
	return c.numberOfVehicles
}

type updateAssignmentsRequestDpo struct {
	flightplanID string
	assignment   *proto.Assignment
}

func (r *updateAssignmentsRequestDpo) GetFlightplanID() string {
	return r.flightplanID
}
func (r *updateAssignmentsRequestDpo) GetEventID() string {
	return r.assignment.Id
}
func (r *updateAssignmentsRequestDpo) GetAssignmentID() string {
	return r.assignment.AssignmentId
}
func (r *updateAssignmentsRequestDpo) GetVehicleID() string {
	return r.assignment.VehicleId
}
func (r *updateAssignmentsRequestDpo) GetMissionID() string {
	return r.assignment.MissionId
}
