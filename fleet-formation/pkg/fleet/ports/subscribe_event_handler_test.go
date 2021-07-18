package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeEventHandlerFleetIdGaveEvent .
func TestSubscribeEventHandlerFleetIdGaveEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "flightplan.fleet_id_gave_event"
		QueueName    = "fleet.fleet_id_gave_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Equal(evt.events1, EventByte)
	a.Empty(evt.events2)
	a.Empty(evt.events3)
}

// TestSubscribeEventHandlerFleetIdRemovedEvent .
func TestSubscribeEventHandlerFleetIdRemovedEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "flightplan.fleet_id_removed_event"
		QueueName    = "fleet.fleet_id_removed_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Empty(evt.events1)
	a.Equal(evt.events2, EventByte)
	a.Empty(evt.events3)
}

// TestSubscribeEventHandlerFleetCopiedEvent .
func TestSubscribeEventHandlerFleetCopiedEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "flightoperation.fleet_copied_event"
		QueueName    = "fleet.fleet_copied_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Empty(evt.events1)
	a.Empty(evt.events2)
	a.Equal(evt.events3, EventByte)
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
	events1, events2, events3 []byte
}

func (h *eventHandlerMock) HandleFleetIDGaveEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events1 = append(h.events1, event...)
	return nil
}

func (h *eventHandlerMock) HandleFleetIDRemovedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events2 = append(h.events2, event...)
	return nil
}

func (h *eventHandlerMock) HandleFleetCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events3 = append(h.events3, event...)
	return nil
}
