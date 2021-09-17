package grpc

import (
	"remote-communication/pkg/mission/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectUsvc := NewUploadMissionEdgeServiceServer(app)

	a.Equal("skysign_proto.UploadMissionEdgeService", sMock.descs[0].ServiceName)
	a.Equal(expectUsvc, sMock.impls[0])
}
