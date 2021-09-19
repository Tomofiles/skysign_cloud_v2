package rabbitmq

import (
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const fleetIDGaveEventExchangeName = "flightplan.fleet_id_gave_event"

// PublishFleetIDGaveEvent .
func PublishFleetIDGaveEvent(
	ch crm.Channel,
	event flightplan.FleetIDGaveEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		fleetIDGaveEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.FleetIdGaveEvent{
		FleetId:          string(event.GetFleetID()),
		NumberOfVehicles: int32(event.NumberOfVehicles),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		fleetIDGaveEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", fleetIDGaveEventExchangeName, eventPb.String())
	return nil
}
