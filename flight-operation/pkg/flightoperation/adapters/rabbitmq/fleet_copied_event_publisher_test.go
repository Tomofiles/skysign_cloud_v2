package rabbitmq

import (
	"errors"
	"flight-operation/pkg/flightoperation/domain/flightoperation"
	"flight-operation/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishFleetCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightoperationFleetID + "-original"
		NewID      = DefaultFlightoperationFleetID + "-new"
	)

	event := flightoperation.FleetCopiedEvent{
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetCopiedEvent(chMock, event)

	expectPb := skysign_proto.FleetCopiedEvent{
		OriginalFleetId: string(OriginalID),
		NewFleetId:      string(NewID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishFleetCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightoperationFleetID + "-original"
		NewID      = DefaultFlightoperationFleetID + "-new"
	)

	event := flightoperation.FleetCopiedEvent{
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishFleetCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightoperationFleetID + "-original"
		NewID      = DefaultFlightoperationFleetID + "-new"
	)

	event := flightoperation.FleetCopiedEvent{
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishFleetCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
