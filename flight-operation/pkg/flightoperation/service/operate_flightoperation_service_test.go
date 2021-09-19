package service

import (
	"testing"

	fope "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompleteFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
			IsCompleted: fope.Operating,
			Version:     string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	pub := &publisherMock{}
	txm := &txManagerMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &operateFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &flightoperationIDCommandMock{
		ID: string(DefaultID),
	}
	ret := service.CompleteFlightoperation(command)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCompleteFlightoperationOperation(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
			IsCompleted: fope.Operating,
			Version:     string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &operateFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightoperationIDCommandMock{
		ID: string(DefaultID),
	}
	ret := service.completeFlightoperationOperation(
		nil,
		pub,
		command,
	)

	expectEvent := fope.FlightoperationCompletedEvent{
		ID:          DefaultID,
		Name:        DefaultName,
		Description: DefaultDescription,
		FleetID:     DefaultFleetID,
	}
	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
}
