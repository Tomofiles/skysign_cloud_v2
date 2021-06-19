package rabbitmq

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
	// for _, e := range p.events {
	// 	if event, ok := e.(vehicle.CommunicationIDGaveEvent); ok {
	// 		if err := PublishCommunicationIDGaveEvent(p.ch, event); err != nil {
	// 			return err
	// 		}
	// 	}
	// 	if event, ok := e.(vehicle.CommunicationIDRemovedEvent); ok {
	// 		if err := PublishCommunicationIDRemovedEvent(p.ch, event); err != nil {
	// 			return err
	// 		}
	// 	}
	// 	if event, ok := e.(vehicle.CopiedVehicleCreatedEvent); ok {
	// 		if err := PublishCopiedVehicleCreatedEvent(p.ch, event); err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	return nil
}
