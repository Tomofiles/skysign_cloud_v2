package grpc

import (
	"remote-communication/pkg/communication/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	usvc := NewCommunicationUserServiceServer(app)
	esvc := NewCommunicationEdgeServiceServer(app)

	skysign_proto.RegisterCommunicationUserServiceServer(s, usvc)
	skysign_proto.RegisterCommunicationEdgeServiceServer(s, esvc)
}
