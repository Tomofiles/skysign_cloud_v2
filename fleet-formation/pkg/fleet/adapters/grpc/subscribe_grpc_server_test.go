package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectFsvc := NewAssignAssetsToFleetServiceServer(app)

	a.Equal("skysign_proto.AssignAssetsToFleetService", sMock.descs[0].ServiceName)
	a.Equal(expectFsvc, sMock.impls[0])
}