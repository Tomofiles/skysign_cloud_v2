package grpc

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/adapters/proto"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// manageVehicleServiceServer .
type manageVehicleServiceServer struct {
	skysign_proto.UnimplementedManageVehicleServiceServer
	app app.Application
}

// NewManageVehicleServiceServer .
func NewManageVehicleServiceServer(application app.Application) skysign_proto.ManageVehicleServiceServer {
	return &manageVehicleServiceServer{app: application}
}

// ListVehicles .
func (s *manageVehicleServiceServer) ListVehicles(
	ctx context.Context,
	request *skysign_proto.Empty,
) (*skysign_proto.ListVehiclesResponses, error) {
	response := &skysign_proto.ListVehiclesResponses{}
	if ret := s.app.Services.ManageVehicle.ListVehicles(
		func(model service.VehiclePresentationModel) {
			response.Vehicles = append(
				response.Vehicles,
				&skysign_proto.Vehicle{
					Id:              model.GetVehicle().GetID(),
					Name:            model.GetVehicle().GetName(),
					CommunicationId: model.GetVehicle().GetCommunicationID(),
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetVehicle .
func (s *manageVehicleServiceServer) GetVehicle(
	ctx context.Context,
	request *skysign_proto.GetVehicleRequest,
) (*skysign_proto.Vehicle, error) {
	if ret := proto.ValidateGetVehicleRequest(request); ret != nil {
		return nil, ret
	}
	response := &skysign_proto.Vehicle{}
	command := &vehicleIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageVehicle.GetVehicle(
		command,
		func(model service.VehiclePresentationModel) {
			response.Id = model.GetVehicle().GetID()
			response.Name = model.GetVehicle().GetName()
			response.CommunicationId = model.GetVehicle().GetCommunicationID()
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CreateVehicle .
func (s *manageVehicleServiceServer) CreateVehicle(
	ctx context.Context,
	request *skysign_proto.Vehicle,
) (*skysign_proto.Vehicle, error) {
	if ret := proto.ValidateCreateVehicleRequest(request); ret != nil {
		return nil, ret
	}
	response := &skysign_proto.Vehicle{}
	command := &createCommand{
		request: request,
	}
	if ret := s.app.Services.ManageVehicle.CreateVehicle(
		command,
		func(id string) {
			response.Id = id
		},
	); ret != nil {
		return nil, ret
	}
	response.Name = request.Name
	response.CommunicationId = request.CommunicationId
	return response, nil
}

// UpdateVehicle .
func (s *manageVehicleServiceServer) UpdateVehicle(
	ctx context.Context,
	request *skysign_proto.Vehicle,
) (*skysign_proto.Vehicle, error) {
	if ret := proto.ValidateUpdateVehicleRequest(request); ret != nil {
		return nil, ret
	}
	response := &skysign_proto.Vehicle{
		Id:              request.Id,
		Name:            request.Name,
		CommunicationId: request.CommunicationId,
	}
	command := &updateCommand{
		request: request,
	}
	if ret := s.app.Services.ManageVehicle.UpdateVehicle(command); ret != nil {
		return nil, ret
	}
	return response, nil
}

// DeleteVehicle .
func (s *manageVehicleServiceServer) DeleteVehicle(
	ctx context.Context,
	request *skysign_proto.DeleteVehicleRequest,
) (*skysign_proto.Empty, error) {
	if ret := proto.ValidateDeleteVehicleRequest(request); ret != nil {
		return nil, ret
	}
	response := &skysign_proto.Empty{}
	command := &vehicleIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageVehicle.DeleteVehicle(command); ret != nil {
		return nil, ret
	}
	return response, nil
}

type vehicleIDCommand struct {
	id string
}

func (f *vehicleIDCommand) GetID() string {
	return f.id
}

type createCommand struct {
	request *skysign_proto.Vehicle
}

func (f *createCommand) GetVehicle() service.Vehicle {
	return &vehicle{
		request: f.request,
	}
}

type updateCommand struct {
	request *skysign_proto.Vehicle
}

func (f *updateCommand) GetID() string {
	return f.request.Id
}

func (f *updateCommand) GetVehicle() service.Vehicle {
	return &vehicle{
		request: f.request,
	}
}

type vehicle struct {
	request *skysign_proto.Vehicle
}

func (f *vehicle) GetID() string {
	return f.request.Id
}

func (f *vehicle) GetName() string {
	return f.request.Name
}

func (f *vehicle) GetCommunicationID() string {
	return f.request.CommunicationId
}
