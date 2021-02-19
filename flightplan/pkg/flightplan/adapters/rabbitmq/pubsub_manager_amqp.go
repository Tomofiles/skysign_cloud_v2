package rabbitmq

import (
	"context"
	"flightplan/pkg/flightplan/domain/event"

	"github.com/streadway/amqp"
)

// PubSubManager .
type PubSubManager struct {
	conn *amqp.Connection
}

// NewPubSubManager .
func NewPubSubManager(conn *amqp.Connection) *PubSubManager {
	return &PubSubManager{
		conn: conn,
	}
}

// GetPublisher .
func (psm *PubSubManager) GetPublisher() (event.Publisher, event.ConnectionClose, error) {
	ch, err := psm.conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	var connClose = func() {
		ch.Close()
	}

	return NewPublisher(ch), connClose, nil
}

// SetConsumer .
func (psm *PubSubManager) SetConsumer(ctx context.Context, exchangeName string, handler event.Handler) error {
	ch, err := psm.conn.Channel()
	if err != nil {
		return err
	}
	var connClose = func() {
		ch.Close()
	}

	if err := ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		connClose()
		return err
	}

	q, err := ch.QueueDeclare(
		exchangeName,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		connClose()
		return err
	}

	if err := ch.QueueBind(
		q.Name,
		"",
		exchangeName,
		false,
		nil,
	); err != nil {
		connClose()
		return err
	}

	eventCh, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		connClose()
		return err
	}

	go func() {
		defer connClose()
		for {
			select {
			case <-ctx.Done():
				break
			case amqpEvent, ok := <-eventCh:
				if !ok {
					continue
				}

				handler(amqpEvent.Body)
			}
		}
	}()

	return nil
}
