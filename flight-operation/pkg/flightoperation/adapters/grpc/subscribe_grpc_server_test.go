package grpc

import (
	"flight-operation/pkg/flightoperation/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectVsvc := NewOperateFlightServiceServer(app)

	a.Equal("skysign_proto.OperateFlightService", sMock.descs[0].ServiceName)
	a.Equal(expectVsvc, sMock.impls[0])
}
