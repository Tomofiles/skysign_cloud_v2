package bridge

import (
	"context"
	"flightplan/pkg/flightplan/api"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/infra"
)

// Bind .
func Bind(eventHandler api.EventHandler, app app.Application) {
	if pub, ok := app.Pub.(*infra.PublisherDirect); ok {
		pub.CreatedEventHandler = func(event flightplan.CreatedEvent) {
			ctx := context.Background()
			eventHandler.HandleCreatedEvent(ctx, event)
		}
		pub.DeletedEventHandler = func(event flightplan.DeletedEvent) {
			ctx := context.Background()
			eventHandler.HandleDeletedEvent(ctx, event)
		}
	}
}
