package rabbitmq

import (
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/fleet/domain/fleet"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(fleet.VehicleCopiedEvent); ok {
				if err := PublishVehicleCopiedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(fleet.MissionCopiedEvent); ok {
				if err := PublishMissionCopiedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
