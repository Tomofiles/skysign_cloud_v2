package grpc

import (
	"context"

	"flight-operation/pkg/flightplan/app"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type changeFlightplanServiceServer struct {
	proto.UnimplementedChangeFlightplanServiceServer
	app app.Application
}

// NewChangeFlightplanServiceServer .
func NewChangeFlightplanServiceServer(application app.Application) proto.ChangeFlightplanServiceServer {
	return &changeFlightplanServiceServer{app: application}
}

// ChangeNumberOfVehicles .
func (s *changeFlightplanServiceServer) ChangeNumberOfVehicles(
	ctx context.Context,
	request *proto.ChangeNumberOfVehiclesRequest,
) (*proto.ChangeNumberOfVehiclesResponse, error) {
	response := &proto.ChangeNumberOfVehiclesResponse{}
	command := &changeNumberOfVehiclesCommand{
		id:               request.Id,
		numberOfVehicles: int(request.NumberOfVehicles),
	}
	if ret := s.app.Services.ChangeFlightplan.ChangeNumberOfVehicles(command); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.NumberOfVehicles = request.NumberOfVehicles
	return response, nil
}

type changeNumberOfVehiclesCommand struct {
	id               string
	numberOfVehicles int
}

func (c *changeNumberOfVehiclesCommand) GetID() string {
	return c.id
}

func (c *changeNumberOfVehiclesCommand) GetNumberOfVehicles() int {
	return c.numberOfVehicles
}
