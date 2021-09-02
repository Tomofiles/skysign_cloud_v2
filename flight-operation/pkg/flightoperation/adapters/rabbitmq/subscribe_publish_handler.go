package rabbitmq

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribePublishHandler .
func SubscribePublishHandler(psm crm.PubSubManagerSetter) {
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(fope.FleetCopiedEvent); ok {
				if err := PublishFleetCopiedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
	psm.SetPublishHandler(
		func(ch crm.Channel, e interface{}) {
			if event, ok := e.(fope.FlightoperationCompletedEvent); ok {
				if err := PublishFlightoperationCompletedEvent(ch, event); err != nil {
					glog.Error(err)
				}
			}
		},
	)
}
