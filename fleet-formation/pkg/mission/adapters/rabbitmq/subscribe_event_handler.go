package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
	app app.Application,
) {
	evt := NewMissionCopiedEventHandler(app)
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
