package grpc

import (
	"fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	vsvc := NewManageVehicleServiceServer(app)

	skysign_proto.RegisterManageVehicleServiceServer(s, vsvc)
}
