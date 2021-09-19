package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectAsvc := NewActionServiceServer(app)

	a.Equal("skysign_proto.ActionService", sMock.descs[0].ServiceName)
	a.Equal(expectAsvc, sMock.impls[0])
}
