package api

import (
	"context"
	"errors"

	"flightplan/pkg/flightplan/app"
	proto "flightplan/pkg/skysign_proto"
)

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
	ret := s.app.Services.ManageFlightplan.ListFlightplans(
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
	)
	if ret != nil {
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
	ret := s.app.Services.ManageFlightplan.GetFlightplan(
		request,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	)
	if ret != nil {
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
	ret := s.app.Services.ManageFlightplan.CreateFlightplan(
		request,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	)
	if ret != nil {
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
	ret := s.app.Services.ManageFlightplan.UpdateFlightplan(
		request,
		func(id, name, description string) {
			response.Id = id
			response.Name = name
			response.Description = description
		},
	)
	if ret != nil {
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
	ret := s.app.Services.ManageFlightplan.DeleteFlightplan(
		request,
	)
	if ret != nil {
		return nil, ret
	}
	return response, nil
}

// ChangeNumberOfVehicles .
func (s *GrpcServer) ChangeNumberOfVehicles(
	ctx context.Context,
	request *proto.ChangeNumberOfVehiclesRequest,
) (*proto.ChangeNumberOfVehiclesResponse, error) {
	return nil, errors.New("")
}

// GetAssignments .
func (s *GrpcServer) GetAssignments(
	ctx context.Context,
	request *proto.GetAssignmentsRequest,
) (*proto.GetAssignmentsResponse, error) {
	return nil, errors.New("")
}

// UpdateAssignments .
func (s *GrpcServer) UpdateAssignments(
	ctx context.Context,
	request *proto.UpdateAssignmentsRequest,
) (*proto.UpdateAssignmentsResponse, error) {
	return nil, errors.New("")
}
