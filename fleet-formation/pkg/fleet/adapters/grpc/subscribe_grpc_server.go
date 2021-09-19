package grpc

import (
	"fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	fsvc := NewAssignAssetsToFleetServiceServer(app)

	skysign_proto.RegisterAssignAssetsToFleetServiceServer(s, fsvc)
}
