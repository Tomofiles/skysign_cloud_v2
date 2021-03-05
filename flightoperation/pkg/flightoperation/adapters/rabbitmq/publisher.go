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
		if event, ok := e.(fope.CreatedEvent); ok {
			if err := PublishFlightoperationCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(fope.FlightplanCopiedWhenCreatedEvent); ok {
			if err := PublishFlightplanCopiedWhenFlightoperationCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
