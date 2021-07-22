package ports

import (
	"context"

	"flight-operation/pkg/flightreport/app"
	"flight-operation/pkg/flightreport/service"
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

// ListFlightreports .
func (s *GrpcServer) ListFlightreports(
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
func (s *GrpcServer) GetFlightreport(
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
