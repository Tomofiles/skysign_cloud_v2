package rabbitmq

import (
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
