package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeEventHandlerMissionCopiedEvent .
func TestSubscribeEventHandlerMissionCopiedEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "fleet.mission_copied_event"
		QueueName    = "mission.mission_copied_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Equal(evt.events, EventByte)
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
	events []byte
}

func (h *eventHandlerMock) HandleMissionCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events = append(h.events, event...)
	return nil
}
