package service

import (
	"testing"

	m "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/domain/mission"

	"github.com/stretchr/testify/assert"
)

func TestPullMissionTransaction(t *testing.T) {
	a := assert.New(t)

	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	testMission := m.NewInstance(DefaultMissionID)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultMissionID).Return(testMission, nil)

	service := &edgeMissionService{
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	var resCall bool
	ret := service.PullMission(
		command,
		func(id string, waypoints []Waypoint) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPullMissionOperation(t *testing.T) {
	a := assert.New(t)

	testMission := m.NewInstance(DefaultMissionID)
	testMission.PushWaypoint(11.0, 21.0, 31.0, 41.0)
	testMission.PushWaypoint(12.0, 22.0, 32.0, 42.0)
	testMission.PushWaypoint(13.0, 23.0, 33.0, 43.0)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultMissionID).Return(testMission, nil)
	pub := &publisherMock{}

	service := &edgeMissionService{
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	var resID string
	var resWaypoints []Waypoint
	ret := service.pullMissionOperation(
		nil,
		pub,
		command,
		func(id string, waypoints []Waypoint) {
			resID = id
			resWaypoints = waypoints
		},
	)

	expectWaypoints := []Waypoint{
		&waypoint{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		&waypoint{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		&waypoint{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}

	a.Nil(ret)
	a.Equal(resID, string(DefaultMissionID))
	a.Equal(resWaypoints, expectWaypoints)
}
