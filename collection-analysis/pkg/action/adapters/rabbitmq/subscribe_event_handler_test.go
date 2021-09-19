package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestSubscribeEventHandleCopiedVehicleCreatedEvent .
func TestSubscribeEventHandleCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}
	service.On("CreateAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.CopiedVehicleCreatedEvent{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FleetId:         string(DefaultActionFleetID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "vehicle.copied_vehicle_created_event"
		QueueName    = "action.copied_vehicle_created_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(requestPb.GetVehicleId(), service.command.GetID())
	a.Equal(requestPb.GetCommunicationId(), service.command.GetCommunicationID())
	a.Equal(requestPb.GetFleetId(), service.command.GetFleetID())
}

// TestSubscribeEventHandleFlightoperationCompletedEvent .
func TestSubscribeEventHandleFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}
	service.On("CompleteAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.FlightoperationCompletedEvent{
		Flightoperation: &skysign_proto.Flightoperation{
			FleetId: string(DefaultActionFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "flightoperation.flightoperation_completed_event"
		QueueName    = "action.flightoperation_completed_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(requestPb.Flightoperation.GetFleetId(), service.completeCommand.GetFleetID())
}

// TestSubscribeEventHandleTelemetryUpdatedEvent .
func TestSubscribeEventHandleTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}
	service.On("PushTelemetry", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.TelemetryUpdatedEvent{
		CommunicationId: string(DefaultActionCommunicationID),
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "communication.telemetry_updated_event"
		QueueName    = "action.telemetry_updated_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

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
