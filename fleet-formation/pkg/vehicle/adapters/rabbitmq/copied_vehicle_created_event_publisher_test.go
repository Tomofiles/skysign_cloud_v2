package rabbitmq

import (
	"errors"
	"fleet-formation/pkg/skysign_proto"
	"fleet-formation/pkg/vehicle/domain/vehicle"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CopiedVehicleCreatedEvent{
		ID:              DefaultVehicleID,
		CommunicationID: DefaultVehicleCommunicationID,
		FleetID:         DefaultFleetID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCopiedVehicleCreatedEvent(chMock, event)

	expectPb := skysign_proto.CopiedVehicleCreatedEvent{
		VehicleId:       string(DefaultVehicleID),
		CommunicationId: string(DefaultVehicleCommunicationID),
		FleetId:         string(DefaultFleetID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CopiedVehicleCreatedEvent{
		ID:              DefaultVehicleID,
		CommunicationID: DefaultVehicleCommunicationID,
		FleetID:         DefaultFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCopiedVehicleCreatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	event := vehicle.CopiedVehicleCreatedEvent{
		ID:              DefaultVehicleID,
		CommunicationID: DefaultVehicleCommunicationID,
		FleetID:         DefaultFleetID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishCopiedVehicleCreatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
