package grpc

import (
	"remote-communication/pkg/communication/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectUsvc := NewCommunicationUserServiceServer(app)
	expectEsvc := NewCommunicationEdgeServiceServer(app)

	a.Equal("skysign_proto.CommunicationUserService", sMock.descs[0].ServiceName)
	a.Equal("skysign_proto.CommunicationEdgeService", sMock.descs[1].ServiceName)
	a.Equal(expectUsvc, sMock.impls[0])
	a.Equal(expectEsvc, sMock.impls[1])
}
