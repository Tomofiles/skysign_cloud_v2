package rabbitmq

import (
	"context"
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestDeliveryToMessageCh(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	dlvCh := make(chan amqp.Delivery)

	msgCh := deliveryToMessageCh(ctx, dlvCh)

	dlvBody := "delivery-body"
	dlv := amqp.Delivery{
		Body: []byte(dlvBody),
	}

	dlvCh <- dlv

	msg := <-msgCh

	a.Equal(msg, []byte(dlvBody))
}

func TestMessageChCloseWhenDeliveryChClose(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	dlvCh := make(chan amqp.Delivery)

	msgCh := deliveryToMessageCh(ctx, dlvCh)

	close(dlvCh)

	_, ok := <-msgCh

	a.False(ok)
}

func TestMessageChCloseWhenContextDone(t *testing.T) {
	a := assert.New(t)

	ctx, cancel := context.WithCancel(context.Background())

	dlvCh := make(chan amqp.Delivery)

	msgCh := deliveryToMessageCh(ctx, dlvCh)

	cancel()

	_, ok := <-msgCh

	a.False(ok)
}
