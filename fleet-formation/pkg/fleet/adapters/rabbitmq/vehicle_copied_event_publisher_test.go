package rabbitmq

import (
	"errors"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishVehicleCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetVehicleID + "-original"
		NewID      = DefaultFleetVehicleID + "-new"
	)

	event := fleet.VehicleCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishVehicleCopiedEvent(chMock, event)

	expectPb := skysign_proto.VehicleCopiedEvent{
		FleetId:           string(DefaultFleetID),
		OriginalVehicleId: string(OriginalID),
		NewVehicleId:      string(NewID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishVehicleCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetVehicleID + "-original"
		NewID      = DefaultFleetVehicleID + "-new"
	)

	event := fleet.VehicleCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishVehicleCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishVehicleCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetVehicleID + "-original"
		NewID      = DefaultFleetVehicleID + "-new"
	)

	event := fleet.VehicleCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishVehicleCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
