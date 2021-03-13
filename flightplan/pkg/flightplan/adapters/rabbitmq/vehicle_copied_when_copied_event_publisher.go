package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const vehicleCopiedWhenFlightplanCopiedEventExchangeName = "flightplan.vehicle_copied_when_flightplan_copied_event"

// PublishVehicleCopiedWhenFlightplanCopiedEvent .
func PublishVehicleCopiedWhenFlightplanCopiedEvent(
	ch Channel,
	event fleet.VehicleCopiedWhenFlightplanCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		vehicleCopiedWhenFlightplanCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.VehicleCopiedWhenFlightplanCopiedEvent{
		FlightplanId:      event.GetFlightplanID(),
		OriginalVehicleId: event.GetOriginalID(),
		NewVehicleId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		vehicleCopiedWhenFlightplanCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", vehicleCopiedWhenFlightplanCopiedEventExchangeName, eventPb.String())
	return nil
}
