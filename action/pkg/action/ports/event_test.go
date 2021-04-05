package ports

import (
	"action/pkg/action/app"
	"action/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestHandleCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}

	service.On("CreateAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.CopiedVehicleCreatedEvent{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FlightplanId:    string(DefaultActionFlightplanID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCopiedVehicleCreatedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.GetVehicleId(), service.requestDpo.GetID())
	a.Equal(requestPb.GetCommunicationId(), service.requestDpo.GetCommunicationID())
	a.Equal(requestPb.GetFlightplanId(), service.requestDpo.GetFlightplanID())
}

func TestHandleFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}

	service.On("CompleteAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FlightoperationCompletedEvent{
		FlightplanId: string(DefaultActionFlightplanID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightoperationCompletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.GetFlightplanId(), service.completeRequestDpo.GetFlightplanID())
}

func TestHandleTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}

	service.On("PushTelemetry", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	handler := NewEventHandler(app)

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
	a.Equal(requestPb.GetCommunicationId(), service.telemetryRequestDpo.GetCommunicationID())
	a.Equal(requestPb.GetTelemetry().Latitude, service.telemetryRequestDpo.GetLatitude())
	a.Equal(requestPb.GetTelemetry().Longitude, service.telemetryRequestDpo.GetLongitude())
	a.Equal(requestPb.GetTelemetry().Altitude, service.telemetryRequestDpo.GetAltitude())
	a.Equal(requestPb.GetTelemetry().RelativeAltitude, service.telemetryRequestDpo.GetRelativeAltitude())
	a.Equal(requestPb.GetTelemetry().Speed, service.telemetryRequestDpo.GetSpeed())
	a.Equal(requestPb.GetTelemetry().Armed, service.telemetryRequestDpo.GetArmed())
	a.Equal(requestPb.GetTelemetry().FlightMode, service.telemetryRequestDpo.GetFlightMode())
	a.Equal(requestPb.GetTelemetry().OrientationX, service.telemetryRequestDpo.GetOrientationX())
	a.Equal(requestPb.GetTelemetry().OrientationY, service.telemetryRequestDpo.GetOrientationY())
	a.Equal(requestPb.GetTelemetry().OrientationZ, service.telemetryRequestDpo.GetOrientationZ())
	a.Equal(requestPb.GetTelemetry().OrientationW, service.telemetryRequestDpo.GetOrientationW())
}
