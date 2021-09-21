package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
	app app.Application,
) {
	evt := NewFlightoperationCompletedEventHandler(app)
	psm.SetConsumer(
		ctx,
		FlightoperationCompletedEventExchangeName,
		FlightoperationCompletedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFlightoperationCompletedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
