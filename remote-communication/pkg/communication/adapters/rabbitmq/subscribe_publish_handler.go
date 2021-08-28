package rabbitmq

import (
	crm "remote-communication/pkg/common/adapters/rabbitmq"
	c "remote-communication/pkg/communication/domain/communication"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
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
