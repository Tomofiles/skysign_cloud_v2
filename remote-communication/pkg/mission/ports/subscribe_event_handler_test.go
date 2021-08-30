package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribeEventHandleCopiedMissionCreatedEvent .
func TestSubscribeEventHandleCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	psm := &publishHandlerMock{}
	evt := &eventHandlerMock{}
	SubscribeEventHandler(nil, psm, evt)

	var (
		ExchangeName = "mission.copied_mission_created_event"
		QueueName    = "uploadmission.copied_mission_created_event"
		EventByte    = []byte{0x20, 0x21}
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(EventByte)
		}
	}

	a.Equal(evt.events1, EventByte)
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
	events1 []byte
}

func (h *eventHandlerMock) HandleCopiedMissionCreatedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events1 = append(h.events1, event...)
	return nil
}
