package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	msvc := NewManageMissionServiceServer(app)

	skysign_proto.RegisterManageMissionServiceServer(s, msvc)
}
