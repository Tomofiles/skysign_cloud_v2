package grpc

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/app"
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type manageFlightplanServiceServer struct {
	proto.UnimplementedManageFlightplanServiceServer
	app app.Application
}

// NewManageFlightplanServiceServer .
func NewManageFlightplanServiceServer(application app.Application) proto.ManageFlightplanServiceServer {
	return &manageFlightplanServiceServer{app: application}
}

// ListFlightplans .
func (s *manageFlightplanServiceServer) ListFlightplans(
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
func (s *manageFlightplanServiceServer) GetFlightplan(
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
func (s *manageFlightplanServiceServer) CreateFlightplan(
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
func (s *manageFlightplanServiceServer) UpdateFlightplan(
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
func (s *manageFlightplanServiceServer) DeleteFlightplan(
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
