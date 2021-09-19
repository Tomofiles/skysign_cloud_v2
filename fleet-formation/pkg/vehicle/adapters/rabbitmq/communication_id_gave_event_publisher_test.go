package rabbitmq

import (
	"errors"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishCommunicationIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CommunicationIDGaveEvent{
		CommunicationID: DefaultVehicleCommunicationID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCommunicationIDGaveEvent(chMock, event)

	expectPb := skysign_proto.CommunicationIdGaveEvent{
		CommunicationId: string(DefaultVehicleCommunicationID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishCommunicationIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CommunicationIDGaveEvent{
		CommunicationID: DefaultVehicleCommunicationID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCommunicationIDGaveEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishCommunicationIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CommunicationIDGaveEvent{
		CommunicationID: DefaultVehicleCommunicationID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishCommunicationIDGaveEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
