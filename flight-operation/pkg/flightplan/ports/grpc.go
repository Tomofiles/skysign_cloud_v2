package ports

import (
	"context"

	"flight-operation/pkg/flightplan/app"
	"flight-operation/pkg/flightplan/service"
	proto "flight-operation/pkg/skysign_proto"

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
		func(model service.FlightplanPresentationModel) {
			response.Flightplans = append(
				response.Flightplans,
				&proto.Flightplan{
					Id:          model.GetFlightplan().GetID(),
					Name:        model.GetFlightplan().GetName(),
					Description: model.GetFlightplan().GetDescription(),
					FleetId:     model.GetFlightplan().GetFleetID(),
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
	command := &flightplanIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightplan.GetFlightplan(
		command,
		func(model service.FlightplanPresentationModel) {
			response.Id = model.GetFlightplan().GetID()
			response.Name = model.GetFlightplan().GetName()
			response.Description = model.GetFlightplan().GetDescription()
			response.FleetId = model.GetFlightplan().GetFleetID()
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
	command := &createCommand{
		request: request,
	}
	if ret := s.app.Services.ManageFlightplan.CreateFlightplan(
		command,
		func(id string) {
			response.Id = id
		},
		func(fleetID string) {
			response.FleetId = fleetID
		},
	); ret != nil {
		return nil, ret
	}
	response.Name = request.Name
	response.Description = request.Description
	return response, nil
}

// UpdateFlightplan .
func (s *GrpcServer) UpdateFlightplan(
	ctx context.Context,
	request *proto.Flightplan,
) (*proto.Flightplan, error) {
	response := &proto.Flightplan{}
	command := &updateCommand{
		request: request,
	}
	if ret := s.app.Services.ManageFlightplan.UpdateFlightplan(
		command,
		func(fleetID string) {
			response.FleetId = fleetID
		},
	); ret != nil {
		return nil, ret
	}
	response.Id = request.Id
	response.Name = request.Name
	response.Description = request.Description
	return response, nil
}

// DeleteFlightplan .
func (s *GrpcServer) DeleteFlightplan(
	ctx context.Context,
	request *proto.DeleteFlightplanRequest,
) (*proto.Empty, error) {
	response := &proto.Empty{}
	command := &flightplanIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightplan.DeleteFlightplan(command); ret != nil {
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

// ExecuteFlightplan .
func (s *GrpcServer) ExecuteFlightplan(
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

type createCommand struct {
	request *proto.Flightplan
}

func (f *createCommand) GetFlightplan() service.Flightplan {
	return &flightplan{
		request: f.request,
	}
}

type updateCommand struct {
	request *proto.Flightplan
}

func (f *updateCommand) GetID() string {
	return f.request.Id
}

func (f *updateCommand) GetFlightplan() service.Flightplan {
	return &flightplan{
		request: f.request,
	}
}

type flightplan struct {
	request *proto.Flightplan
}

func (f *flightplan) GetID() string {
	return f.request.Id
}

func (f *flightplan) GetName() string {
	return f.request.Name
}

func (f *flightplan) GetDescription() string {
	return f.request.Description
}

func (f *flightplan) GetFleetID() string {
	return f.request.FleetId
}

type flightplanIDCommand struct {
	id string
}

func (f *flightplanIDCommand) GetID() string {
	return f.id
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
