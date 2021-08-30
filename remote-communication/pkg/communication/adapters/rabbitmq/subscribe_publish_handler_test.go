package rabbitmq

import (
	"remote-communication/pkg/communication/domain/communication"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribePublishHandler(t *testing.T) {
	a := assert.New(t)

	event1 := communication.TelemetryUpdatedEvent{}

	psm := &publishHandlerMock{}
	SubscribePublishHandler(psm)

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	for _, h := range psm.publishHandlers {
		h(chMock, event1)
	}

	a.Equal(chMock.messageCallCount, 1)
}
