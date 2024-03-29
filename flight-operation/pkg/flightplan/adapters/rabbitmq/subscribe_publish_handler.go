package rabbitmq

import (
	"flight-operation/pkg/flightplan/domain/flightplan"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(flightplan.FleetIDGaveEvent); ok {
				if err := PublishFleetIDGaveEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(flightplan.FleetIDRemovedEvent); ok {
				if err := PublishFleetIDRemovedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(flightplan.FlightplanExecutedEvent); ok {
				if err := PublishFlightplanExecutedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
