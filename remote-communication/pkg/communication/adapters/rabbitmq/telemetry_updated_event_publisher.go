package rabbitmq

import (
	c "remote-communication/pkg/communication/domain/communication"
	"remote-communication/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const telemetryUpdatedEventExchangeName = "communication.telemetry_updated_event"

// PublishTelemetryUpdatedEvent .
func PublishTelemetryUpdatedEvent(
	ch crm.Channel,
	event c.TelemetryUpdatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		telemetryUpdatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.TelemetryUpdatedEvent{
		CommunicationId: event.GetID(),
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         event.GetTelemetry().Latitude,
			Longitude:        event.GetTelemetry().Longitude,
			Altitude:         event.GetTelemetry().Altitude,
			RelativeAltitude: event.GetTelemetry().RelativeAltitude,
			Speed:            event.GetTelemetry().Speed,
			Armed:            event.GetTelemetry().Armed,
			FlightMode:       event.GetTelemetry().FlightMode,
			OrientationX:     event.GetTelemetry().X,
			OrientationY:     event.GetTelemetry().Y,
			OrientationZ:     event.GetTelemetry().Z,
			OrientationW:     event.GetTelemetry().W,
		},
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		telemetryUpdatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", telemetryUpdatedEventExchangeName, eventPb.String())
	return nil
}
