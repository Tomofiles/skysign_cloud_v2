package grpc

import (
	"context"

	"flight-operation/pkg/flightplan/app"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type executeFlightplanServiceServer struct {
	proto.UnimplementedExecuteFlightplanServiceServer
	app app.Application
}

// NewExecuteFlightplanServiceServer .
func NewExecuteFlightplanServiceServer(application app.Application) proto.ExecuteFlightplanServiceServer {
	return &executeFlightplanServiceServer{app: application}
}

// ExecuteFlightplan .
func (s *executeFlightplanServiceServer) ExecuteFlightplan(
	ctx context.Context,
	request *proto.ExecuteFlightplanRequest,
) (*proto.ExecuteFlightplanResponse, error) {
	response := &proto.ExecuteFlightplanResponse{
		Id: request.Id,
	}
	command := &flightplanIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ExecuteFlightplan.ExecuteFlightplan(command); ret != nil {
		return nil, ret
	}
	return response, nil
}
