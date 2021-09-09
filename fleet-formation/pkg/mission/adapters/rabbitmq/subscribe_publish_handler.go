package rabbitmq

import (
	"fleet-formation/pkg/mission/domain/mission"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(mission.CopiedMissionCreatedEvent); ok {
				if err := PublishCopiedMissionCreatedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
