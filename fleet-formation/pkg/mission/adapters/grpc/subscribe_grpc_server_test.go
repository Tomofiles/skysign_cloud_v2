package grpc

import (
	"fleet-formation/pkg/mission/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectVsvc := NewManageMissionServiceServer(app)

	a.Equal("skysign_proto.ManageMissionService", sMock.descs[0].ServiceName)
	a.Equal(expectVsvc, sMock.impls[0])
}
