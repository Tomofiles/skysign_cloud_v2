package rabbitmq

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"flightoperation/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const fleetCopiedEventExchangeName = "flightoperation.fleet_copied_event"

// PublishFleetCopiedEvent .
func PublishFleetCopiedEvent(
	ch Channel,
	event fope.FleetCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		fleetCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FleetCopiedEvent{
		OriginalFleetId: event.GetOriginalID(),
		NewFleetId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		fleetCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", fleetCopiedEventExchangeName, eventPb.String())
	return nil
}
