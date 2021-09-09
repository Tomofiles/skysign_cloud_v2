package ports

import (
	"collection-analysis/pkg/action/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

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
		FleetId:         string(DefaultActionFleetID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCopiedVehicleCreatedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.GetVehicleId(), service.command.GetID())
	a.Equal(requestPb.GetCommunicationId(), service.command.GetCommunicationID())
	a.Equal(requestPb.GetFleetId(), service.command.GetFleetID())
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
		Flightoperation: &skysign_proto.Flightoperation{
			FleetId: string(DefaultActionFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightoperationCompletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.Flightoperation.GetFleetId(), service.completeCommand.GetFleetID())
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
	a.Equal(requestPb.GetCommunicationId(), service.telemetryCommand.GetCommunicationID())
	a.Equal(requestPb.GetTelemetry().Latitude, service.telemetryCommand.GetLatitude())
	a.Equal(requestPb.GetTelemetry().Longitude, service.telemetryCommand.GetLongitude())
	a.Equal(requestPb.GetTelemetry().Altitude, service.telemetryCommand.GetAltitude())
	a.Equal(requestPb.GetTelemetry().RelativeAltitude, service.telemetryCommand.GetRelativeAltitude())
	a.Equal(requestPb.GetTelemetry().Speed, service.telemetryCommand.GetSpeed())
	a.Equal(requestPb.GetTelemetry().Armed, service.telemetryCommand.GetArmed())
	a.Equal(requestPb.GetTelemetry().FlightMode, service.telemetryCommand.GetFlightMode())
	a.Equal(requestPb.GetTelemetry().OrientationX, service.telemetryCommand.GetOrientationX())
	a.Equal(requestPb.GetTelemetry().OrientationY, service.telemetryCommand.GetOrientationY())
	a.Equal(requestPb.GetTelemetry().OrientationZ, service.telemetryCommand.GetOrientationZ())
	a.Equal(requestPb.GetTelemetry().OrientationW, service.telemetryCommand.GetOrientationW())
}
