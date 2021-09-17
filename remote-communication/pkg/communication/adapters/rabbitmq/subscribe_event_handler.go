package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

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
		CommunicationIDGaveEventExchangeName,
		CommunicationIDGaveEventQueueName,
		func(event []byte) {
			if err := evt.HandleCommunicationIDGaveEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		CommunicationIDRemovedEventExchangeName,
		CommunicationIDRemovedEventQueueName,
		func(event []byte) {
			if err := evt.HandleCommunicationIDRemovedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
