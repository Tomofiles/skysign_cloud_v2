package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
	app app.Application,
) {
	gevt := NewCommunicationIDGaveEventHandler(app)
	revt := NewCommunicationIDRemovedEventHandler(app)
	psm.SetConsumer(
		ctx,
		CommunicationIDGaveEventExchangeName,
		CommunicationIDGaveEventQueueName,
		func(event []byte) {
			if err := gevt.HandleCommunicationIDGaveEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		CommunicationIDRemovedEventExchangeName,
		CommunicationIDRemovedEventQueueName,
		func(event []byte) {
			if err := revt.HandleCommunicationIDRemovedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
