package service

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	flightoperation := fope.AssembleFrom(
		nil,
		&flightoperationComponentMock{
			ID:           string(DefaultFlightoperationID),
			FlightplanID: string(DefaultFlightplanID),
			IsCompleted:  DefaultIsCompleted,
			Version:      string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultFlightoperationID).Return(flightoperation, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		ID: string(DefaultFlightoperationID),
	}
	var resCall bool
	ret := service.GetFlightoperation(
		req,
		func(id, flightplanID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetFlightoperationOperation(t *testing.T) {
	a := assert.New(t)

	flightoperation := fope.AssembleFrom(
		nil,
		&flightoperationComponentMock{
			ID:           string(DefaultFlightoperationID),
			FlightplanID: string(DefaultFlightplanID),
			IsCompleted:  DefaultIsCompleted,
			Version:      string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultFlightoperationID).Return(flightoperation, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		ID: string(DefaultFlightoperationID),
	}
	var resID, resFlightplanID string
	ret := service.getFlightoperationOperation(
		nil,
		req,
		func(id, flightplanID string) {
			resID = id
			resFlightplanID = flightplanID
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightoperationID))
	a.Equal(resFlightplanID, string(DefaultFlightplanID))
}

func TestListFlightoperationsTransaction(t *testing.T) {
	a := assert.New(t)

	flightoperations := []*fope.Flightoperation{
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:           string(DefaultFlightoperationID),
				FlightplanID: string(DefaultFlightplanID),
				IsCompleted:  DefaultIsCompleted,
				Version:      string(DefaultVersion),
			},
		),
	}

	repo := &flightoperationRepositoryMock{}
	txm := &txManagerMock{}
	repo.On("GetAll").Return(flightoperations, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListFlightoperations(
		func(id, flightplanID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestListFlightoperationsOperation(t *testing.T) {
	a := assert.New(t)

	const (
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
		DefaultFlightplanID1      = DefaultFlightplanID + "-1"
		DefaultFlightplanID2      = DefaultFlightplanID + "-2"
		DefaultFlightplanID3      = DefaultFlightplanID + "-3"
		DefaultVersion1           = DefaultVersion + "-1"
		DefaultVersion2           = DefaultVersion + "-2"
		DefaultVersion3           = DefaultVersion + "-3"
	)

	flightoperations := []*fope.Flightoperation{
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:           string(DefaultFlightoperationID1),
				FlightplanID: string(DefaultFlightplanID1),
				IsCompleted:  DefaultIsCompleted,
				Version:      string(DefaultVersion1),
			},
		),
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:           string(DefaultFlightoperationID2),
				FlightplanID: string(DefaultFlightplanID2),
				IsCompleted:  DefaultIsCompleted,
				Version:      string(DefaultVersion2),
			},
		),
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:           string(DefaultFlightoperationID3),
				FlightplanID: string(DefaultFlightplanID3),
				IsCompleted:  DefaultIsCompleted,
				Version:      string(DefaultVersion3),
			},
		),
	}

	repo := &flightoperationRepositoryMock{}
	repo.On("GetAll").Return(flightoperations, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resID, resFlightplanID []string
	ret := service.listFlightoperationsOperation(
		nil,
		func(id, flightplanID string) {
			resID = append(resID, id)
			resFlightplanID = append(resFlightplanID, flightplanID)
		},
	)

	a.Nil(ret)
	a.Equal(resID[0], string(DefaultFlightoperationID1))
	a.Equal(resFlightplanID[0], string(DefaultFlightplanID1))
	a.Equal(resID[1], string(DefaultFlightoperationID2))
	a.Equal(resFlightplanID[1], string(DefaultFlightplanID2))
	a.Equal(resID[2], string(DefaultFlightoperationID3))
	a.Equal(resFlightplanID[2], string(DefaultFlightplanID3))
}

func TestCreateFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightplanID + "-original"
		NewID      = DefaultFlightplanID + "-new"
	)

	gen := &generatorMockFlightoperation{
		id:           DefaultFlightoperationID,
		flightplanID: NewID,
	}
	repo := &flightoperationRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}
	txm := &txManagerMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFlightoperationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &flightplanIDRequestMock{
		FlightplanID: string(OriginalID),
	}
	ret := service.CreateFlightoperation(req)

	a.Nil(ret)
	a.Len(pub.events, 2)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateFlightoperationOperation(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightplanID + "-original"
		NewID      = DefaultFlightplanID + "-new"
	)

	gen := &generatorMockFlightoperation{
		id:           DefaultFlightoperationID,
		flightplanID: NewID,
		version:      DefaultVersion,
	}
	repo := &flightoperationRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFlightoperationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightplanIDRequestMock{
		FlightplanID: string(OriginalID),
	}
	ret := service.createFlightoperationOperation(
		nil,
		pub,
		req,
	)

	expectEvent1 := fope.CreatedEvent{
		ID:           DefaultFlightoperationID,
		FlightplanID: NewID,
	}
	expectEvent2 := fope.FlightplanCopiedWhenCreatedEvent{
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	a.Nil(ret)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2})
}
