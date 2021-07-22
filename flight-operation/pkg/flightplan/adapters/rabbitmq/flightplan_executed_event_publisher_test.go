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

func TestPublishFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FlightplanExecutedEvent{
		ID:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetID:     DefaultFlightplanFleetID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFlightplanExecutedEvent(chMock, event)

	expectPb := skysign_proto.FlightplanExecutedEvent{
		FlightplanId: string(DefaultFlightplanID),
		Flightplan: &skysign_proto.Flightplan{
			Id:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
			FleetId:     string(DefaultFlightplanFleetID),
		},
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FlightplanExecutedEvent{
		ID:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetID:     DefaultFlightplanFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFlightplanExecutedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FlightplanExecutedEvent{
		ID:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetID:     DefaultFlightplanFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishFlightplanExecutedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
