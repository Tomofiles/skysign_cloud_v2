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
	return nil, errors.New("")
}

// GetFlightplan .
func (s *GrpcServer) GetFlightplan(
	ctx context.Context,
	request *proto.GetFlightplanRequest,
) (*proto.Flightplan, error) {
	return &proto.Flightplan{
		Id: request.Id,
	}, nil
}

// CreateFlightplan .
func (s *GrpcServer) CreateFlightplan(
	ctx context.Context,
	request *proto.Flightplan,
) (*proto.Flightplan, error) {
	return nil, errors.New("")
}

// UpdateFlightplan .
func (s *GrpcServer) UpdateFlightplan(
	ctx context.Context,
	request *proto.Flightplan,
) (*proto.Flightplan, error) {
	return nil, errors.New("")
}

// DeleteFlightplan .
func (s *GrpcServer) DeleteFlightplan(
	ctx context.Context,
	request *proto.DeleteFlightplanRequest,
) (*proto.Empty, error) {
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
