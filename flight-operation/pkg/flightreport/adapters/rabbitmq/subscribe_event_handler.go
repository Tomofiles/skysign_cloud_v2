package rabbitmq

import (
	"context"
	"flight-operation/pkg/flightreport/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"
	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
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
