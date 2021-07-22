package ports

import (
	"collection-analysis/pkg/common/ports"
	"context"

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
		CopiedVehicleCreatedEventExchangeName,
		CopiedVehicleCreatedEventQueueName,
		func(event []byte) {
			if err := evt.HandleCopiedVehicleCreatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		FlightoperationCompletedEventExchangeName,
		FlightoperationCompletedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFlightoperationCompletedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		TelemetryUpdatedEventExchangeName,
		TelemetryUpdatedEventQueueName,
		func(event []byte) {
			if err := evt.HandleTelemetryUpdatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
}
