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
		MissionCopiedEventExchangeName,
		MissionCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleMissionCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
