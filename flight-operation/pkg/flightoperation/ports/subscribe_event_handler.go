package ports

import (
	"context"
	"flight-operation/pkg/common/ports"

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
		FlightplanExecutedEventExchangeName,
		FlightplanExecutedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFlightplanExecutedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
