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
		if event, ok := e.(flightplan.FleetIDGaveEvent); ok {
			if err := PublishFleetIDGaveEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(flightplan.FleetIDRemovedEvent); ok {
			if err := PublishFleetIDRemovedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(flightplan.FlightplanExecutedEvent); ok {
			if err := PublishFlightplanExecutedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fleet.VehicleCopiedEvent); ok {
			if err := PublishVehicleCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fleet.MissionCopiedEvent); ok {
			if err := PublishMissionCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
