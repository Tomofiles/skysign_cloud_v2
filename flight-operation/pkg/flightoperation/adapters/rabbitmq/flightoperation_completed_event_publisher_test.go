package rabbitmq

import (
	"errors"
	"testing"

	fo "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	event := fo.FlightoperationCompletedEvent{
		ID:          DefaultFlightoperationID,
		Name:        DefaultFlightoperationName,
		Description: DefaultFlightoperationDescription,
		FleetID:     DefaultFlightoperationFleetID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFlightoperationCompletedEvent(chMock, event)

	expectPb := skysign_proto.FlightoperationCompletedEvent{
		FlightoperationId: string(DefaultFlightoperationID),
		Flightoperation: &skysign_proto.Flightoperation{
			Id:          string(DefaultFlightoperationID),
			Name:        DefaultFlightoperationName,
			Description: DefaultFlightoperationDescription,
			FleetId:     string(DefaultFlightoperationFleetID),
		},
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	event := fo.FlightoperationCompletedEvent{
		ID:          DefaultFlightoperationID,
		Name:        DefaultFlightoperationName,
		Description: DefaultFlightoperationDescription,
		FleetID:     DefaultFlightoperationFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFlightoperationCompletedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	event := fo.FlightoperationCompletedEvent{
		ID:          DefaultFlightoperationID,
		Name:        DefaultFlightoperationName,
		Description: DefaultFlightoperationDescription,
		FleetID:     DefaultFlightoperationFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishFlightoperationCompletedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
