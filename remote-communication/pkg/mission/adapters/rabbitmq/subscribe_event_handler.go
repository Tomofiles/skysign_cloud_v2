package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm ports.PubSubManagerSetter,
	app app.Application,
) {
	mevt := NewCopiedMissionCreatedEventHandler(app)
	psm.SetConsumer(
		ctx,
		CopiedMissionCreatedEventExchangeName,
		CopiedMissionCreatedEventQueueName,
		func(event []byte) {
			if err := mevt.HandleCopiedMissionCreatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
