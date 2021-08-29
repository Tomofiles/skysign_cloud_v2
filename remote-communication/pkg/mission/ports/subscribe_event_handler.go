package ports

import (
	"context"
	"remote-communication/pkg/common/ports"

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
		CopiedMissionCreatedEventExchangeName,
		CopiedMissionCreatedEventQueueName,
		func(event []byte) {
			if err := evt.HandleCopiedMissionCreatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
