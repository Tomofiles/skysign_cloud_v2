package rabbitmq

// Publisher .
type Publisher struct {
	events   []interface{}
	ch       Channel
	handlers []PublishHandler
}

// NewPublisher .
func NewPublisher(ch Channel, handlers []PublishHandler) *Publisher {
	return &Publisher{
		ch:       ch,
		handlers: handlers,
	}
}

// Publish .
func (p *Publisher) Publish(event interface{}) {
	p.events = append(p.events, event)
}

// Flush .
func (p *Publisher) Flush() error {
	for _, e := range p.events {
		for _, handler := range p.handlers {
			handler(p.ch, e)
		}
	}
	return nil
}
