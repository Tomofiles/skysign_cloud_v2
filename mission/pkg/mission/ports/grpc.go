package ports

// import (
// 	"context"

// 	proto "vehicle/pkg/skysign_proto"
// 	"vehicle/pkg/vehicle/app"

// 	"github.com/golang/glog"
// 	"google.golang.org/grpc"
// )

// // LogBodyInterceptor .
// func LogBodyInterceptor() grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
// 		glog.Infof("REQUEST , API: %s, Message: %+v", info.FullMethod, req)
// 		defer func() {
// 			if err != nil {
// 				glog.Errorf("RESPONSE, API: %s, Error: %+v", info.FullMethod, err)
// 			} else {
// 				glog.Infof("RESPONSE, API: %s, Message: %+v", info.FullMethod, resp)
// 			}
// 		}()

// 		resp, err = handler(ctx, req)
// 		return
// 	}
// }

// // GrpcServer .
// type GrpcServer struct {
// 	app app.Application
// }

// // NewGrpcServer .
// func NewGrpcServer(application app.Application) GrpcServer {
// 	return GrpcServer{app: application}
// }

// // ListVehicles .
// func (s *GrpcServer) ListVehicles(
// 	ctx context.Context,
// 	request *proto.Empty,
// ) (*proto.ListVehiclesResponses, error) {
// 	response := &proto.ListVehiclesResponses{}
// 	if ret := s.app.Services.ManageVehicle.ListVehicles(
// 		func(id, name, communicationID string) {
// 			response.Vehicles = append(
// 				response.Vehicles,
// 				&proto.Vehicle{
// 					Id:              id,
// 					Name:            name,
// 					CommunicationId: communicationID,
// 				},
// 			)
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// // GetVehicle .
// func (s *GrpcServer) GetVehicle(
// 	ctx context.Context,
// 	request *proto.GetVehicleRequest,
// ) (*proto.Vehicle, error) {
// 	response := &proto.Vehicle{}
// 	requestDpo := &vehicleIDRequestDpo{
// 		id: request.Id,
// 	}
// 	if ret := s.app.Services.ManageVehicle.GetVehicle(
// 		requestDpo,
// 		func(id, name, communicationID string) {
// 			response.Id = id
// 			response.Name = name
// 			response.CommunicationId = communicationID
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// // CreateVehicle .
// func (s *GrpcServer) CreateVehicle(
// 	ctx context.Context,
// 	request *proto.Vehicle,
// ) (*proto.Vehicle, error) {
// 	response := &proto.Vehicle{}
// 	requestDpo := &vehicleRequestDpo{
// 		name:            request.Name,
// 		communicationID: request.CommunicationId,
// 	}
// 	if ret := s.app.Services.ManageVehicle.CreateVehicle(
// 		requestDpo,
// 		func(id, name, communicationID string) {
// 			response.Id = id
// 			response.Name = name
// 			response.CommunicationId = communicationID
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// // UpdateVehicle .
// func (s *GrpcServer) UpdateVehicle(
// 	ctx context.Context,
// 	request *proto.Vehicle,
// ) (*proto.Vehicle, error) {
// 	response := &proto.Vehicle{}
// 	requestDpo := &vehicleRequestDpo{
// 		id:              request.Id,
// 		name:            request.Name,
// 		communicationID: request.CommunicationId,
// 	}
// 	if ret := s.app.Services.ManageVehicle.UpdateVehicle(
// 		requestDpo,
// 		func(id, name, communicationID string) {
// 			response.Id = id
// 			response.Name = name
// 			response.CommunicationId = communicationID
// 		},
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// // DeleteVehicle .
// func (s *GrpcServer) DeleteVehicle(
// 	ctx context.Context,
// 	request *proto.DeleteVehicleRequest,
// ) (*proto.Empty, error) {
// 	response := &proto.Empty{}
// 	requestDpo := &vehicleIDRequestDpo{
// 		id: request.Id,
// 	}
// 	if ret := s.app.Services.ManageVehicle.DeleteVehicle(
// 		requestDpo,
// 	); ret != nil {
// 		return nil, ret
// 	}
// 	return response, nil
// }

// type vehicleRequestDpo struct {
// 	id              string
// 	name            string
// 	communicationID string
// }

// func (f *vehicleRequestDpo) GetID() string {
// 	return f.id
// }

// func (f *vehicleRequestDpo) GetName() string {
// 	return f.name
// }

// func (f *vehicleRequestDpo) GetCommunicationID() string {
// 	return f.communicationID
// }

// type vehicleIDRequestDpo struct {
// 	id string
// }

// func (f *vehicleIDRequestDpo) GetID() string {
// 	return f.id
// }
