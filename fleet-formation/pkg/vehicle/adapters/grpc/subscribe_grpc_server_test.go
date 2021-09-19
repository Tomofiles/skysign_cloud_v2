package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectVsvc := NewManageVehicleServiceServer(app)

	a.Equal("skysign_proto.ManageVehicleService", sMock.descs[0].ServiceName)
	a.Equal(expectVsvc, sMock.impls[0])
}
