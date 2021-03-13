package rabbitmq

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"flightoperation/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"

// PublishFlightoperationCompletedEvent .
func PublishFlightoperationCompletedEvent(
	ch Channel,
	event fope.CompletedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightoperationCompletedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightoperationCompletedEvent{
		FlightoperationId: event.GetID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightoperationCompletedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightoperationCompletedEventExchangeName, eventPb.String())
	return nil
}
