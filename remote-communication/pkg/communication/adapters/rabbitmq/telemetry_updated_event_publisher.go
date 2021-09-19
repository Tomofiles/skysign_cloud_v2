package rabbitmq

import (
	c "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

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
			Latitude:         event.GetTelemetry().LatitudeDegree,
			Longitude:        event.GetTelemetry().LongitudeDegree,
			Altitude:         event.GetTelemetry().AltitudeM,
			RelativeAltitude: event.GetTelemetry().RelativeAltitudeM,
			Speed:            event.GetTelemetry().SpeedMS,
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
