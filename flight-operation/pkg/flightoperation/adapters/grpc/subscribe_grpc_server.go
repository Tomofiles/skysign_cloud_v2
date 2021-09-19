package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	osvc := NewOperateFlightServiceServer(app)

	skysign_proto.RegisterOperateFlightServiceServer(s, osvc)
}
