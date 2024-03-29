package rabbitmq

import (
	"fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const copiedVehicleCreatedEventExchangeName = "vehicle.copied_vehicle_created_event"

// PublishCopiedVehicleCreatedEvent .
func PublishCopiedVehicleCreatedEvent(
	ch crm.Channel,
	event vehicle.CopiedVehicleCreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		copiedVehicleCreatedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.CopiedVehicleCreatedEvent{
		VehicleId:       event.GetVehicleID(),
		CommunicationId: event.GetCommunicationID(),
		FleetId:         event.GetFleetID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		copiedVehicleCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", copiedVehicleCreatedEventExchangeName, eventPb.String())
	return nil
}
