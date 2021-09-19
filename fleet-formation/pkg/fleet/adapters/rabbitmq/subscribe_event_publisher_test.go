package rabbitmq

import (
	"fleet-formation/pkg/fleet/domain/fleet"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribeEventPublisher(t *testing.T) {
	a := assert.New(t)

	event1 := fleet.MissionCopiedEvent{}
	event2 := fleet.VehicleCopiedEvent{}

	psm := &pubSubManagerMock{}
	SubscribeEventPublisher(psm)

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	for _, h := range psm.publishHandlers {
		h(chMock, event1)
	}
	for _, h := range psm.publishHandlers {
		h(chMock, event2)
	}

	a.Equal(chMock.messageCallCount, 2)
}
