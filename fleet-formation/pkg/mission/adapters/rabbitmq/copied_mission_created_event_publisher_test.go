package rabbitmq

import (
	"errors"
	m "fleet-formation/pkg/mission/domain/mission"
	"fleet-formation/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
		nil,
		&missionComponentMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationComponentMock{
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints: []waypointComponentMock{
					{
						1,
						11.0,
						21.0,
						31.0,
						41.0,
					},
					{
						2,
						12.0,
						22.0,
						32.0,
						42.0,
					},
					{
						3,
						13.0,
						23.0,
						33.0,
						43.0,
					},
				},
				UploadID: string(DefaultMissionUploadID),
			},
			Version: string(DefaultMissionVersion),
		},
	)
	event := m.CopiedMissionCreatedEvent{
		ID:      DefaultMissionID,
		Mission: mission,
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCopiedMissionCreatedEvent(chMock, event)

	expectPb := skysign_proto.CopiedMissionCreatedEvent{
		MissionId: string(DefaultMissionID),
		Mission: &skysign_proto.Mission{
			Id:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: &skysign_proto.Navigation{
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints: []*skysign_proto.Waypoint{
					{
						Latitude:       11.0,
						Longitude:      21.0,
						RelativeHeight: 31.0,
						Speed:          41.0,
					},
					{
						Latitude:       12.0,
						Longitude:      22.0,
						RelativeHeight: 32.0,
						Speed:          42.0,
					},
					{
						Latitude:       13.0,
						Longitude:      23.0,
						RelativeHeight: 33.0,
						Speed:          43.0,
					},
				},
				UploadId: string(DefaultMissionUploadID),
			},
		},
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
		nil,
		&missionComponentMock{
			ID:         string(DefaultMissionID),
			Name:       DefaultMissionName,
			Navigation: navigationComponentMock{},
		},
	)
	event := m.CopiedMissionCreatedEvent{
		ID:      DefaultMissionID,
		Mission: mission,
	}
	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishCopiedMissionCreatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
		nil,
		&missionComponentMock{
			ID:         string(DefaultMissionID),
			Name:       DefaultMissionName,
			Navigation: navigationComponentMock{},
		},
	)
	event := m.CopiedMissionCreatedEvent{
		ID:      DefaultMissionID,
		Mission: mission,
	}
	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishCopiedMissionCreatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
