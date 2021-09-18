package grpc

import (
	"context"
	"flight-operation/pkg/flightreport/app"
	"flight-operation/pkg/flightreport/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

type reportFlightServiceServer struct {
	proto.UnimplementedReportFlightServiceServer
	app app.Application
}

// NewReportFlightServiceServer .
func NewReportFlightServiceServer(application app.Application) proto.ReportFlightServiceServer {
	return &reportFlightServiceServer{app: application}
}

// ListFlightreports .
func (s *reportFlightServiceServer) ListFlightreports(
	ctx context.Context,
	request *proto.Empty,
) (*proto.ListFlightreportsResponses, error) {
	response := &proto.ListFlightreportsResponses{}
	if ret := s.app.Services.ManageFlightreport.ListFlightreports(
		func(model service.FlightreportPresentationModel) {
			response.Flightreports = append(
				response.Flightreports,
				&proto.Flightreport{
					Id:          model.GetFlightreport().GetID(),
					Name:        model.GetFlightreport().GetName(),
					Description: model.GetFlightreport().GetDescription(),
					FleetId:     model.GetFlightreport().GetFleetID(),
				},
			)
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// GetFlightreport .
func (s *reportFlightServiceServer) GetFlightreport(
	ctx context.Context,
	request *proto.GetFlightreportRequest,
) (*proto.Flightreport, error) {
	response := &proto.Flightreport{}
	command := &flightreportIDCommand{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightreport.GetFlightreport(
		command,
		func(model service.FlightreportPresentationModel) {
			response.Id = model.GetFlightreport().GetID()
			response.Name = model.GetFlightreport().GetName()
			response.Description = model.GetFlightreport().GetDescription()
			response.FleetId = model.GetFlightreport().GetFleetID()
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

type flightreportIDCommand struct {
	id string
}

func (f *flightreportIDCommand) GetID() string {
	return f.id
}
