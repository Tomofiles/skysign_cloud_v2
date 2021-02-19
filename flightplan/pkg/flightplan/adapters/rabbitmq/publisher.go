package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/flightplan"

	"github.com/streadway/amqp"
)

// Publisher .
type Publisher struct {
	events []interface{}
	ch     *amqp.Channel
}

// NewPublisher .
func NewPublisher(ch *amqp.Channel) *Publisher {
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
	}
	return nil
}
