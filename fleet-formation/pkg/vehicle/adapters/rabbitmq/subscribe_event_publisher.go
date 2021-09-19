package rabbitmq

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/domain/vehicle"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventPublisher .
func SubscribeEventPublisher(psm crm.PubSubManagerSetter) {
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
