package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestHandleTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}

	service.On("PushTelemetry", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	handler := NewTelemetryUpdatedEventHandler(app)

	requestPb := &skysign_proto.TelemetryUpdatedEvent{
		CommunicationId: string(DefaultActionCommunicationID),
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            true,
			FlightMode:       "state",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleTelemetryUpdatedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.GetCommunicationId(), service.telemetryCommand.GetCommunicationID())
	a.Equal(requestPb.GetTelemetry().Latitude, service.telemetryCommand.GetLatitudeDegree())
	a.Equal(requestPb.GetTelemetry().Longitude, service.telemetryCommand.GetLongitudeDegree())
	a.Equal(requestPb.GetTelemetry().Altitude, service.telemetryCommand.GetAltitudeM())
	a.Equal(requestPb.GetTelemetry().RelativeAltitude, service.telemetryCommand.GetRelativeAltitudeM())
	a.Equal(requestPb.GetTelemetry().Speed, service.telemetryCommand.GetSpeedMS())
	a.Equal(requestPb.GetTelemetry().Armed, service.telemetryCommand.GetArmed())
	a.Equal(requestPb.GetTelemetry().FlightMode, service.telemetryCommand.GetFlightMode())
	a.Equal(requestPb.GetTelemetry().OrientationX, service.telemetryCommand.GetOrientationX())
	a.Equal(requestPb.GetTelemetry().OrientationY, service.telemetryCommand.GetOrientationY())
	a.Equal(requestPb.GetTelemetry().OrientationZ, service.telemetryCommand.GetOrientationZ())
	a.Equal(requestPb.GetTelemetry().OrientationW, service.telemetryCommand.GetOrientationW())
}
