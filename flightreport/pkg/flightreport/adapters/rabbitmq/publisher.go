package rabbitmq

import frep "flightreport/pkg/flightreport/domain/flightreport"

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
		if event, ok := e.(frep.CreatedEvent); ok {
			if err := PublishFlightreportCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
		if event, ok := e.(frep.FlightoperationCopiedWhenCreatedEvent); ok {
			if err := PublishFlightoperationCopiedWhenFlightreportCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
