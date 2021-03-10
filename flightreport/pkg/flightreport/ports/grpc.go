package ports

import (
	"context"

	"flightreport/pkg/flightreport/app"
	proto "flightreport/pkg/skysign_proto"

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
	if ret := s.app.Services.ManageFlightoperation.ListFlightoperations(
		func(id, flightplanID string) {
			response.Flightreports = append(
				response.Flightreports,
				&proto.Flightreport{
					Id:                id,
					FlightoperationId: flightplanID,
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
	requestDpo := &flightoperationIDRequestDpo{
		id: request.Id,
	}
	if ret := s.app.Services.ManageFlightoperation.GetFlightoperation(
		requestDpo,
		func(id, flightplanID string) {
			response.Id = id
			response.FlightoperationId = flightplanID
		},
	); ret != nil {
		return nil, ret
	}
	return response, nil
}

// CreateFlightreport .
func (s *GrpcServer) CreateFlightreport(
	ctx context.Context,
	request *proto.CreateFlightreportRequest,
) (*proto.Empty, error) {
	requestDpo := &flightplanIDRequestDpo{
		flightplanID: request.FlightplanId,
	}
	if ret := s.app.Services.ManageFlightoperation.CreateFlightoperation(
		requestDpo,
	); ret != nil {
		return nil, ret
	}
	return &proto.Empty{}, nil
}

type flightoperationIDRequestDpo struct {
	id string
}

func (f *flightoperationIDRequestDpo) GetID() string {
	return f.id
}

type flightplanIDRequestDpo struct {
	flightplanID string
}

func (f *flightplanIDRequestDpo) GetFlightplanID() string {
	return f.flightplanID
}
