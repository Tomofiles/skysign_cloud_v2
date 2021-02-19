package bridge

import (
	"context"
	"flightplan/pkg/flightplan/adapters/inmemory"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/ports"
)

// Bind .
func Bind(eventHandler ports.EventHandler, app app.Application) {
	if pub, ok := app.Pub.(*inmemory.PublisherDirect); ok {
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
