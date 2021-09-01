package rabbitmq

import (
	"errors"
	"fleet-formation/pkg/fleet/domain/fleet"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishMissionCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetMissionID + "-original"
		NewID      = DefaultFleetMissionID + "-new"
	)

	event := fleet.MissionCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishMissionCopiedEvent(chMock, event)

	expectPb := skysign_proto.MissionCopiedEvent{
		FleetId:           string(DefaultFleetID),
		OriginalMissionId: string(OriginalID),
		NewMissionId:      string(NewID),
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishMissionCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetMissionID + "-original"
		NewID      = DefaultFleetMissionID + "-new"
	)

	event := fleet.MissionCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishMissionCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishMissionCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetMissionID + "-original"
		NewID      = DefaultFleetMissionID + "-new"
	)

	event := fleet.MissionCopiedEvent{
		FleetID:    DefaultFleetID,
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishMissionCopiedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
