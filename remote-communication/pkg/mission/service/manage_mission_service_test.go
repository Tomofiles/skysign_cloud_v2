package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMissionTransaction(t *testing.T) {
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

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageMissionService{
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &missionCommandMock{
		ID: string(DefaultMissionID),
		Waypoints: []Waypoint{
			&waypoint{
				PointOrder:       1,
				Latitude:         1.0,
				Longitude:        2.0,
				RelativeAltitude: 3.0,
				Speed:            4.0,
			},
		},
	}
	ret := service.CreateMission(command)

	a.Nil(ret)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateMissionOperation(t *testing.T) {
	a := assert.New(t)

	repo := &repositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageMissionService{
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionCommandMock{
		ID: string(DefaultMissionID),
		Waypoints: []Waypoint{
			&waypoint{
				PointOrder:       1,
				Latitude:         1.0,
				Longitude:        2.0,
				RelativeAltitude: 3.0,
				Speed:            4.0,
			},
		},
	}
	ret := service.createMissionOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
}
