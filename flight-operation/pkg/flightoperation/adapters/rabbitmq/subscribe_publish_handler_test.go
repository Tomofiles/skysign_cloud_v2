package rabbitmq

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribePublishHandler(t *testing.T) {
	a := assert.New(t)

	event1 := fope.FleetCopiedEvent{}
	event2 := fope.FlightoperationCompletedEvent{}

	psm := &publishHandlerMock{}
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

	a.Equal(chMock.messageCallCount, 2)
}
