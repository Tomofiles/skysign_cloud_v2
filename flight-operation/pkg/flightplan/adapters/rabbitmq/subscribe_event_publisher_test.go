package rabbitmq

import (
	"flight-operation/pkg/flightplan/domain/flightplan"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscribeEventPublisher(t *testing.T) {
	a := assert.New(t)

	event1 := flightplan.FleetIDGaveEvent{}
	event2 := flightplan.FleetIDRemovedEvent{}
	event3 := flightplan.FlightplanExecutedEvent{}

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
	for _, h := range psm.publishHandlers {
		h(chMock, event3)
	}

	a.Equal(chMock.messageCallCount, 3)
}
