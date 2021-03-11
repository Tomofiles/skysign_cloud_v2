package service

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompleteFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	const (
		OperatingIsCompleted = fope.Operating
		NewVersion           = DefaultVersion + "-new"
	)

	gen := &generatorMockFlightoperation{
		version: NewVersion,
	}

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			ID:           string(DefaultFlightoperationID),
			FlightplanID: string(DefaultFlightplanID),
			IsCompleted:  OperatingIsCompleted,
			Version:      string(DefaultVersion),
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
	repo.On("GetByID", DefaultFlightoperationID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &operateFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &flightoperationIDRequestMock{
		ID: string(DefaultFlightoperationID),
	}
	ret := service.CompleteFlightoperation(req)

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
		OperatingIsCompleted = fope.Operating
		NewVersion           = DefaultVersion + "-new"
	)

	gen := &generatorMockFlightoperation{
		version: NewVersion,
	}

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			ID:           string(DefaultFlightoperationID),
			FlightplanID: string(DefaultFlightplanID),
			IsCompleted:  OperatingIsCompleted,
			Version:      string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultFlightoperationID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &operateFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		ID: string(DefaultFlightoperationID),
	}
	ret := service.completeFlightoperationOperation(
		nil,
		pub,
		req,
	)

	expectEvent := fope.CompletedEvent{
		ID: DefaultFlightoperationID,
	}
	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
}

func TestCannotChangeErrorWhenCompleteFlightoperationOperation(t *testing.T) {
	a := assert.New(t)

	const (
		CompletedIsCompleted = fope.Completed
		NewVersion           = DefaultVersion + "-new"
	)

	gen := &generatorMockFlightoperation{
		version: NewVersion,
	}

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			ID:           string(DefaultFlightoperationID),
			FlightplanID: string(DefaultFlightplanID),
			IsCompleted:  CompletedIsCompleted,
			Version:      string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultFlightoperationID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &operateFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		ID: string(DefaultFlightoperationID),
	}
	ret := service.completeFlightoperationOperation(
		nil,
		pub,
		req,
	)

	a.Len(pub.events, 0)
	a.Equal(ret, fope.ErrCannotChange)
}
