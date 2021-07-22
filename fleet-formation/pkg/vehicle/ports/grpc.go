package ports

import (
	"context"

	proto "fleet-formation/pkg/skysign_proto"
	"fleet-formation/pkg/vehicle/app"
	"fleet-formation/pkg/vehicle/service"

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

// ListVehicles .
func (s *GrpcServer) ListVehicles(
	ctx context.Context,
	request *proto.Empty,
) (*proto.ListVehiclesResponses, error) {
	response := &proto.ListVehiclesResponses{}
	if ret := s.app.Services.ManageVehicle.ListVehicles(
		func(model service.VehiclePresentationModel) {
			response.Vehicles = append(
				response.Vehicles,
				&proto.Vehicle{
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
func (s *GrpcServer) GetVehicle(
	ctx context.Context,
	request *proto.GetVehicleRequest,
) (*proto.Vehicle, error) {
	response := &proto.Vehicle{}
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
func (s *GrpcServer) CreateVehicle(
	ctx context.Context,
	request *proto.Vehicle,
) (*proto.Vehicle, error) {
	response := &proto.Vehicle{}
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
func (s *GrpcServer) UpdateVehicle(
	ctx context.Context,
	request *proto.Vehicle,
) (*proto.Vehicle, error) {
	response := &proto.Vehicle{
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
func (s *GrpcServer) DeleteVehicle(
	ctx context.Context,
	request *proto.DeleteVehicleRequest,
) (*proto.Empty, error) {
	response := &proto.Empty{}
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
	request *proto.Vehicle
}

func (f *createCommand) GetVehicle() service.Vehicle {
	return &vehicle{
		request: f.request,
	}
}

type updateCommand struct {
	request *proto.Vehicle
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
	request *proto.Vehicle
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
