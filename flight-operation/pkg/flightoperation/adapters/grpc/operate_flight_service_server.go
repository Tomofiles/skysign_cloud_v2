package grpc

import (
	"context"

	"flight-operation/pkg/flightoperation/app"
	"flight-operation/pkg/flightoperation/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type operateFlightServiceServer struct {
	proto.UnimplementedOperateFlightServiceServer
	app app.Application
}

// NewOperateFlightServiceServer .
func NewOperateFlightServiceServer(application app.Application) proto.OperateFlightServiceServer {
	return &operateFlightServiceServer{app: application}
}

// ListFlightoperations .
func (s *operateFlightServiceServer) ListFlightoperations(
	ctx context.Context,
	request *proto.Empty,
) (*proto.ListFlightoperationsResponses, error) {
	response := &proto.ListFlightoperationsResponses{}
	if ret := s.app.Services.ManageFlightoperation.ListFlightoperations(
		func(model service.FlightoperationPresentationModel) {
			response.Flightoperations = append(
				response.Flightoperations,
				&proto.Flightoperation{
					Id:          model.GetFlightoperation().GetID(),
					Name:        model.GetFlightoperation().GetName(),
					Description: model.GetFlightoperation().GetDescription(),
					FleetId:     model.GetFlightoperation().GetFleetID(),
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetFlightoperation .
func (s *operateFlightServiceServer) GetFlightoperation(
	ctx context.Context,
	request *proto.GetFlightoperationRequest,
) (*proto.Flightoperation, error) {
	response := &proto.Flightoperation{}
	command := &flightoperationIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightoperation.GetFlightoperation(
		command,
		func(model service.FlightoperationPresentationModel) {
			response.Id = model.GetFlightoperation().GetID()
			response.Name = model.GetFlightoperation().GetName()
			response.Description = model.GetFlightoperation().GetDescription()
			response.FleetId = model.GetFlightoperation().GetFleetID()
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CompleteFlightoperation .
func (s *operateFlightServiceServer) CompleteFlightoperation(
	ctx context.Context,
	request *proto.CompleteFlightoperationRequest,
) (*proto.Empty, error) {
	response := &proto.Empty{}
	command := &flightoperationIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.OperateFlightoperation.CompleteFlightoperation(command); ret != nil {
		return nil, ret
	}
	return response, nil
}

type flightoperationIDCommand struct {
	id string
}

func (f *flightoperationIDCommand) GetID() string {
	return f.id
}
