package rabbitmq

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightoperationCopiedWhenFlightreportCreatedEventExchangeName = "flightreport.flightoperation_copied_when_flightreport_created_event"

// PublishFlightoperationCopiedWhenFlightreportCreatedEvent .
func PublishFlightoperationCopiedWhenFlightreportCreatedEvent(
	ch Channel,
	event frep.FlightoperationCopiedWhenCreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightoperationCopiedWhenFlightreportCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightoperationCopiedWhenFlightreportCreatedEvent{
		OriginalFlightoperationId: event.GetOriginalID(),
		NewFlightoperationId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightoperationCopiedWhenFlightreportCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightoperationCopiedWhenFlightreportCreatedEventExchangeName, eventPb.String())
	return nil
}
