package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const flightplanDeletedEventExchangeName = "flightplan.flightplan_deleted_event"

// PublishFlightplanDeletedEvent .
func PublishFlightplanDeletedEvent(
	ch Channel,
	event flightplan.DeletedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		flightplanDeletedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FlightplanDeletedEvent{
		FlightplanId: event.GetFlightplanID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		flightplanDeletedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", flightplanDeletedEventExchangeName, eventPb.String())
	return nil
}
