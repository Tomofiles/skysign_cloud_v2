package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeEventHandlerCommunicationIdGaveEvent .
func TestSubscribeEventHandlerCommunicationIdGaveEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "vehicle.communication_id_gave_event"
		QueueName    = "communication.communication_id_gave_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Equal(evt.events1, EventByte)
	a.Empty(evt.events2)
}

// TestSubscribeEventHandlerCommunicationIdRemovedEvent .
func TestSubscribeEventHandlerCommunicationIdRemovedEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "vehicle.communication_id_removed_event"
		QueueName    = "communication.communication_id_removed_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Empty(evt.events1)
	a.Equal(evt.events2, EventByte)
}

type publishHandlerMock struct {
	consumers []consumer
}

func (h *publishHandlerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	h.consumers = append(
		h.consumers,
		consumer{
			exchangeName: exchangeName,
			queueName:    queueName,
			handler:      handler,
		})
	return nil
}

type consumer struct {
	exchangeName, queueName string
	handler                 func([]byte)
}

type eventHandlerMock struct {
	events1, events2 []byte
}

func (h *eventHandlerMock) HandleCommunicationIDGaveEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events1 = append(h.events1, event...)
	return nil
}

func (h *eventHandlerMock) HandleCommunicationIDRemovedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events2 = append(h.events2, event...)
	return nil
}
