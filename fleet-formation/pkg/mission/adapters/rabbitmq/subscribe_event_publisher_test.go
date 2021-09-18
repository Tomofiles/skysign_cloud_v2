package rabbitmq

import (
	"fleet-formation/pkg/mission/domain/mission"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribeEventPublisher(t *testing.T) {
	a := assert.New(t)

	event1 := mission.CopiedMissionCreatedEvent{
		Mission: &mission.Mission{},
	}

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
