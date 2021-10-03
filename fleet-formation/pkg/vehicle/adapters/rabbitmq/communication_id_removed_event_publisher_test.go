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

func TestPublishCommunicationIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	event := vehicle.CommunicationIDRemovedEvent{
		CommunicationID: vehicle.CommunicationID(DefaultVehicleCommunicationID),
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCommunicationIDRemovedEvent(chMock, event)

	expectPb := skysign_proto.CommunicationIdRemovedEvent{
		CommunicationId: string(DefaultVehicleCommunicationID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishCommunicationIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	event := vehicle.CommunicationIDRemovedEvent{
		CommunicationID: vehicle.CommunicationID(DefaultVehicleCommunicationID),
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCommunicationIDRemovedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishCommunicationIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	event := vehicle.CommunicationIDRemovedEvent{
		CommunicationID: vehicle.CommunicationID(DefaultVehicleCommunicationID),
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishCommunicationIDRemovedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
