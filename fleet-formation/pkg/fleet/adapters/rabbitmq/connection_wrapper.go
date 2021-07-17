package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
)

type connection struct {
	conn *amqp.Connection
}

// NewConnection .
func NewConnection(conn *amqp.Connection) Connection {
	return &connection{
		conn: conn,
	}
}

func (conn *connection) GetChannel() (Channel, error) {
	ch, err := conn.conn.Channel()
	if err != nil {
		return nil, err
	}
	var channel = &channel{
		ch: ch,
	}
	return channel, nil
}

func (conn *connection) Close() error {
	return conn.conn.Close()
}

type channel struct {
	ch *amqp.Channel
}

func (ch *channel) FanoutExchangeDeclare(exchange string) error {
	return ch.ch.ExchangeDeclare(
		exchange,
		"fanout",
		false,
		true,
		false,
		false,
		nil,
	)
}

func (ch *channel) QueueDeclareAndBind(exchange, queue string) error {
	if _, err := ch.ch.QueueDeclare(
		queue,
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	if err := ch.ch.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (ch *channel) Publish(queue string, message Message) error {
	return ch.ch.Publish(
		queue,
		"",
		false,
		false,
		amqp.Publishing{
			Body: message,
		},
	)
}

func (ch *channel) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	deliveryCh, err := ch.ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	messageCh := deliveryToMessageCh(ctx, deliveryCh)

	return messageCh, nil
}

func (ch *channel) Close() error {
	return ch.ch.Close()
}

func deliveryToMessageCh(ctx context.Context, deliveryCh <-chan amqp.Delivery) <-chan Message {
	var messageCh = make(chan Message)

	go func() {
		defer close(messageCh)
		for {
			select {
			case <-ctx.Done():
				return
			case delivery, ok := <-deliveryCh:
				if !ok {
					return
				}
				messageCh <- delivery.Body
			}
		}
	}()

	return messageCh
}
