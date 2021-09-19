package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribeEventPublisher(t *testing.T) {
	a := assert.New(t)

	event1 := communication.TelemetryUpdatedEvent{}

	psm := &pubSubManagerMock{}
	SubscribeEventPublisher(psm)

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	for _, h := range psm.publishHandlers {
		h(chMock, event1)
	}

	a.Equal(chMock.messageCallCount, 1)
}
