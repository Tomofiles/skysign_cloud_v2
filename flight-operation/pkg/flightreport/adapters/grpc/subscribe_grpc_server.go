package grpc

import (
	"flight-operation/pkg/flightreport/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	rsvc := NewReportFlightServiceServer(app)

	skysign_proto.RegisterReportFlightServiceServer(s, rsvc)
}
