package rabbitmq

import (
	"fleet-formation/pkg/vehicle/domain/vehicle"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribePublishHandler(t *testing.T) {
	a := assert.New(t)

	event1 := vehicle.CommunicationIDGaveEvent{}
	event2 := vehicle.CommunicationIDRemovedEvent{}
	event3 := vehicle.CopiedVehicleCreatedEvent{}

	psm := &pubSubManagerMock{}
	SubscribePublishHandler(psm)

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	for _, h := range psm.publishHandlers {
		h(chMock, event1)
	}
	for _, h := range psm.publishHandlers {
		h(chMock, event2)
	}
	for _, h := range psm.publishHandlers {
		h(chMock, event3)
	}

	a.Equal(chMock.messageCallCount, 3)
}
