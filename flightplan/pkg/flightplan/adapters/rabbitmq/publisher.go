package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
)

// Publisher .
type Publisher struct {
	events []interface{}
	ch     Channel
}

// NewPublisher .
func NewPublisher(ch Channel) *Publisher {
	return &Publisher{
		ch: ch,
	}
}

// Publish .
func (p *Publisher) Publish(event interface{}) {
	p.events = append(p.events, event)
}

// Flush .
func (p *Publisher) Flush() error {
	for _, e := range p.events {
		if event, ok := e.(flightplan.CreatedEvent); ok {
			if err := PublishFlightplanCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(flightplan.DeletedEvent); ok {
			if err := PublishFlightplanDeletedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(flightplan.CopiedEvent); ok {
			if err := PublishFlightplanCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fleet.VehicleCopiedWhenCopiedEvent); ok {
			if err := PublishVehicleCopiedWhenCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fleet.MissionCopiedWhenCopiedEvent); ok {
			if err := PublishMissionCopiedWhenCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
