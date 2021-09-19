package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/grpc"
)

// SubscribeGrpcServer .
func SubscribeGrpcServer(
	s grpc.ServiceRegistrar,
	app app.Application,
) {
	msvc := NewManageFlightplanServiceServer(app)
	csvc := NewChangeFlightplanServiceServer(app)
	esvc := NewExecuteFlightplanServiceServer(app)

	skysign_proto.RegisterManageFlightplanServiceServer(s, msvc)
	skysign_proto.RegisterChangeFlightplanServiceServer(s, csvc)
	skysign_proto.RegisterExecuteFlightplanServiceServer(s, esvc)
}
