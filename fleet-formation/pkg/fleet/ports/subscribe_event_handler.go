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
		FleetIDGaveEventExchangeName,
		FleetIDGaveEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetIDGaveEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FleetIDRemovedEventExchangeName,
		FleetIDRemovedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetIDRemovedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FleetCopiedEventExchangeName,
		FleetCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
