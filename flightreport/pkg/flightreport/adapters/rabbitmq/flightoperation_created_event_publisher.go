package rabbitmq

import (
	fope "flightreport/pkg/flightreport/domain/flightoperation"
	"flightreport/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightoperationCreatedEventExchangeName = "flightoperation.flightoperation_created_event"

// PublishFlightoperationCreatedEvent .
func PublishFlightoperationCreatedEvent(
	ch Channel,
	event fope.CreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightoperationCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightoperationCreatedEvent{
		FlightoperationId: event.GetID(),
		FlightplanId:      event.GetFlightplanID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightoperationCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightoperationCreatedEventExchangeName, eventPb.String())
	return nil
}
