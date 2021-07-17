package rabbitmq

import fope "flightoperation/pkg/flightoperation/domain/flightoperation"

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
		if event, ok := e.(fope.FlightoperationCompletedEvent); ok {
			if err := PublishFlightoperationCompletedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fope.FleetCopiedEvent); ok {
			if err := PublishFleetCopiedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
