package grpc

import (
	"flight-operation/pkg/flightreport/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeGrpcServer .
func TestSubscribeGrpcServer(t *testing.T) {
	a := assert.New(t)

	app := app.Application{}

	sMock := &serviceRegistrarMock{}

	SubscribeGrpcServer(sMock, app)

	expectMsvc := NewReportFlightServiceServer(app)

	a.Equal("skysign_proto.ReportFlightService", sMock.descs[0].ServiceName)
	a.Equal(expectMsvc, sMock.impls[0])
}
