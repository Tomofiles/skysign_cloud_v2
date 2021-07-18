package ports

import (
	"context"
	"fleet-formation/pkg/common/ports"

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
		VehicleCopiedEventExchangeName,
		VehicleCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleVehicleCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
