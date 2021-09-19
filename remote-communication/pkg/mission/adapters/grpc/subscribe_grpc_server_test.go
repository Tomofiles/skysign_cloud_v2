package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/app"

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
