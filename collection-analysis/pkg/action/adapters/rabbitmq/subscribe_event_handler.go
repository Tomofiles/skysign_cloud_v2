package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
)

// SubscribeEventHandler .
func SubscribeEventHandler(
	ctx context.Context,
	psm crm.PubSubManagerSetter,
	app app.Application,
) {
	vevt := NewCopiedVehicleCreatedEventHandler(app)
	fevt := NewFlightoperationCompletedEventHandler(app)
	tevt := NewTelemetryUpdatedEventHandler(app)
	psm.SetConsumer(
		ctx,
		CopiedVehicleCreatedEventExchangeName,
		CopiedVehicleCreatedEventQueueName,
		func(event []byte) {
			if err := vevt.HandleCopiedVehicleCreatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FlightoperationCompletedEventExchangeName,
		FlightoperationCompletedEventQueueName,
		func(event []byte) {
			if err := fevt.HandleFlightoperationCompletedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		TelemetryUpdatedEventExchangeName,
		TelemetryUpdatedEventQueueName,
		func(event []byte) {
			if err := tevt.HandleTelemetryUpdatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
