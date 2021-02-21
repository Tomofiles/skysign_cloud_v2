package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightplanCreatedEventExchangeName = "flightplan.flightplan_created_event"

// PublishFlightplanCreatedEvent .
func PublishFlightplanCreatedEvent(
	ch Channel,
	event flightplan.CreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightplanCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightplanCreatedEvent{
		FlightplanId: event.GetFlightplanID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightplanCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightplanCreatedEventExchangeName, eventPb.String())
	return nil
}
