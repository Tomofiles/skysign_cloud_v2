package rabbitmq

import (
	c "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventPublisher .
func SubscribeEventPublisher(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(c.TelemetryUpdatedEvent); ok {
				if err := PublishTelemetryUpdatedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
