package rabbitmq

import (
	"errors"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishFleetIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDGaveEvent{
		FleetID:          DefaultFlightplanFleetID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetIDGaveEvent(chMock, event)

	expectPb := skysign_proto.FleetIdGaveEvent{
		FleetId:          string(DefaultFlightplanFleetID),
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishFleetIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDGaveEvent{
		FleetID:          DefaultFlightplanFleetID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishFleetIDGaveEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishFleetIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := flightplan.FleetIDGaveEvent{
		FleetID:          DefaultFlightplanFleetID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishFleetIDGaveEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
