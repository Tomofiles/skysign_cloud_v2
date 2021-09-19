package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	usvc := NewUploadMissionEdgeServiceServer(app)

	skysign_proto.RegisterUploadMissionEdgeServiceServer(s, usvc)
}
