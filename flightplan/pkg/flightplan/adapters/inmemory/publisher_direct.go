package inmemory

import (
	"flightplan/pkg/flightplan/domain/flightplan"
)

// PublisherDirect .
type PublisherDirect struct {
	CreatedEventHandler func(flightplan.CreatedEvent)
	DeletedEventHandler func(flightplan.DeletedEvent)
}

// Publish .
func (p *PublisherDirect) Publish(event interface{}) {
	if createdEvent, ok := event.(flightplan.CreatedEvent); ok {
		p.CreatedEventHandler(createdEvent)
	}
	if deletedEvent, ok := event.(flightplan.DeletedEvent); ok {
		p.DeletedEventHandler(deletedEvent)
	}
}

// Flush .
func (p *PublisherDirect) Flush() error {
	return nil
}
