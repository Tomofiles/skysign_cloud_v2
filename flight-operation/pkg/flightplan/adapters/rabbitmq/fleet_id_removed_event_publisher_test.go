package rabbitmq

import (
	"errors"
	"flight-operation/pkg/flightplan/domain/flightplan"
	"flight-operation/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishFleetIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDRemovedEvent{
		FleetID: DefaultFlightplanFleetID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetIDRemovedEvent(chMock, event)

	expectPb := skysign_proto.FleetIDRemovedEvent{
		FleetId: string(DefaultFlightplanFleetID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishFleetIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDRemovedEvent{
		FleetID: DefaultFlightplanFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetIDRemovedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishFleetIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDRemovedEvent{
		FleetID: DefaultFlightplanFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishFleetIDRemovedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
