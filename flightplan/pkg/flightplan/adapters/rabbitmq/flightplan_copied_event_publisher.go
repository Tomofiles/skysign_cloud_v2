package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightplanCopiedEventExchangeName = "flightplan.flightplan_copied_event"

// PublishFlightplanCopiedEvent .
func PublishFlightplanCopiedEvent(
	ch Channel,
	event flightplan.CopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightplanCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightplanCopiedEvent{
		OriginalFlightplanId: event.GetOriginalID(),
		NewFlightplanId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightplanCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightplanCopiedEventExchangeName, eventPb.String())
	return nil
}
