package ports

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
	evt EventHandler,
) {
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
