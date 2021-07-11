package rabbitmq

import "mission/pkg/mission/domain/mission"

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
		if event, ok := e.(mission.CopiedMissionCreatedEvent); ok {
			if err := PublishCopiedMissionCreatedEvent(p.ch, event); err != nil {
				return err
			}
		}
	}
	return nil
}
