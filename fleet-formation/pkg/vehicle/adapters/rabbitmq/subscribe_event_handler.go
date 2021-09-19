package rabbitmq

import (
	"context"
	"fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
	app app.Application,
) {
	evt := NewVehicleCopiedEventHandler(app)
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
