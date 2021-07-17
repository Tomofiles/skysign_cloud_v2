package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const fleetIDRemovedEventExchangeName = "flightplan.fleet_id_removed_event"

// PublishFleetIDRemovedEvent .
func PublishFleetIDRemovedEvent(
	ch Channel,
	event flightplan.FleetIDRemovedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		fleetIDRemovedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FleetIDRemovedEvent{
		FleetId: string(event.GetFleetID()),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		fleetIDRemovedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", fleetIDRemovedEventExchangeName, eventPb.String())
	return nil
}