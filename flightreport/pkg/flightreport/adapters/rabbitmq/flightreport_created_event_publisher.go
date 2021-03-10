package rabbitmq

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightreportCreatedEventExchangeName = "flightreport.flightreport_created_event"

// PublishFlightreportCreatedEvent .
func PublishFlightreportCreatedEvent(
	ch Channel,
	event frep.CreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightreportCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightreportCreatedEvent{
		FlightreportId:    event.GetID(),
		FlightoperationId: event.GetFlightoperationID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightreportCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightreportCreatedEventExchangeName, eventPb.String())
	return nil
}
