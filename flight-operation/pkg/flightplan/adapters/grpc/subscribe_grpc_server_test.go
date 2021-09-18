package grpc

import (
	"flight-operation/pkg/flightplan/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectMsvc := NewManageFlightplanServiceServer(app)
	expectCsvc := NewChangeFlightplanServiceServer(app)
	expectEsvc := NewExecuteFlightplanServiceServer(app)

	a.Equal("skysign_proto.ManageFlightplanService", sMock.descs[0].ServiceName)
	a.Equal(expectMsvc, sMock.impls[0])
	a.Equal("skysign_proto.ChangeFlightplanService", sMock.descs[1].ServiceName)
	a.Equal(expectCsvc, sMock.impls[1])
	a.Equal("skysign_proto.ExecuteFlightplanService", sMock.descs[2].ServiceName)
	a.Equal(expectEsvc, sMock.impls[2])
}
