package rabbitmq

import (
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(vehicle.CommunicationIDGaveEvent); ok {
				if err := PublishCommunicationIDGaveEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(vehicle.CommunicationIDRemovedEvent); ok {
				if err := PublishCommunicationIDRemovedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(vehicle.CopiedVehicleCreatedEvent); ok {
				if err := PublishCopiedVehicleCreatedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
