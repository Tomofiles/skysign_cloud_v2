package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const vehicleCopiedWhenCopiedEventExchangeName = "flightplan.vehicle_copied_when_copied_event"

// PublishVehicleCopiedWhenCopiedEvent .
func PublishVehicleCopiedWhenCopiedEvent(
	ch Channel,
	event fleet.VehicleCopiedWhenCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		vehicleCopiedWhenCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.VehicleCopiedWhenCopiedEvent{
		OriginalVehicleId: event.GetOriginalID(),
		NewVehicleId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		vehicleCopiedWhenCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", vehicleCopiedWhenCopiedEventExchangeName, eventPb.String())
	return nil
}
