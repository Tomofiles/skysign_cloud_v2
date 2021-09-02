package rabbitmq

import (
	"fleet-formation/pkg/fleet/domain/fleet"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

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
