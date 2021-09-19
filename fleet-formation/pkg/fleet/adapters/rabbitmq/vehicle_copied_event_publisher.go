package rabbitmq

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const vehicleCopiedEventExchangeName = "fleet.vehicle_copied_event"

// PublishVehicleCopiedEvent .
func PublishVehicleCopiedEvent(
	ch crm.Channel,
	event fleet.VehicleCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		vehicleCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.VehicleCopiedEvent{
		FleetId:           event.GetFleetID(),
		OriginalVehicleId: event.GetOriginalID(),
		NewVehicleId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		vehicleCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", vehicleCopiedEventExchangeName, eventPb.String())
	return nil
}
