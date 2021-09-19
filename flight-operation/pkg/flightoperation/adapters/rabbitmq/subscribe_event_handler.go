package rabbitmq

import (
	"context"
	"flight-operation/pkg/flightoperation/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
	app app.Application,
) {
	evt := NewFlightplanExecutedEventHandler(app)
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
