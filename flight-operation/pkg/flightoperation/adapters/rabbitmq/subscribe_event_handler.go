package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
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
