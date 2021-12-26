package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
	app app.Application,
) {
	gevt := NewFleetIDGaveEventHandler(app)
	revt := NewFleetIDRemovedEventHandler(app)
	cevt := NewFleetCopiedEventHandler(app)
	psm.SetConsumer(
		ctx,
		FleetIDGaveEventExchangeName,
		FleetIDGaveEventQueueName,
		func(event []byte) {
			if err := gevt.HandleFleetIDGaveEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FleetIDRemovedEventExchangeName,
		FleetIDRemovedEventQueueName,
		func(event []byte) {
			if err := revt.HandleFleetIDRemovedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FleetCopiedEventExchangeName,
		FleetCopiedEventQueueName,
		func(event []byte) {
			if err := cevt.HandleFleetCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}