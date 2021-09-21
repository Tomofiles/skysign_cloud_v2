package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
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
