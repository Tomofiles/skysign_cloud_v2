package rabbitmq

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"flightoperation/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightplanCopiedWhenFlightoperationCreatedEventExchangeName = "flightoperation.flightplan_copied_when_flightoperation_created_event"

// PublishFlightplanCopiedWhenFlightoperationCreatedEvent .
func PublishFlightplanCopiedWhenFlightoperationCreatedEvent(
	ch Channel,
	event fope.FlightplanCopiedWhenCreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightplanCopiedWhenFlightoperationCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightplanCopiedWhenFlightoperationCreatedEvent{
		FlightoperationId:    event.GetFlightoperationID(),
		OriginalFlightplanId: event.GetOriginalID(),
		NewFlightplanId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightplanCopiedWhenFlightoperationCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightplanCopiedWhenFlightoperationCreatedEventExchangeName, eventPb.String())
	return nil
}
