package rabbitmq

import (
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/mission/domain/mission"

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
