package service

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFlightreportTransaction(t *testing.T) {
	a := assert.New(t)

	flightreport := frep.AssembleFrom(
		nil,
		&flightreportComponentMock{
			ID:                string(DefaultFlightreportID),
			FlightoperationID: string(DefaultFlightoperationID),
		},
	)

	repo := &flightreportRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultFlightreportID).Return(flightreport, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	req := &flightreportIDRequestMock{
		ID: string(DefaultFlightreportID),
	}
	var resCall bool
	ret := service.GetFlightreport(
		req,
		func(id, flightoperationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetFlightreportOperation(t *testing.T) {
	a := assert.New(t)

	flightreport := frep.AssembleFrom(
		nil,
		&flightreportComponentMock{
			ID:                string(DefaultFlightreportID),
			FlightoperationID: string(DefaultFlightoperationID),
		},
	)

	repo := &flightreportRepositoryMock{}
	repo.On("GetByID", DefaultFlightreportID).Return(flightreport, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightreportIDRequestMock{
		ID: string(DefaultFlightreportID),
	}
	var resID, resFlightoperationID string
	ret := service.getFlightreportOperation(
		nil,
		req,
		func(id, flightoperationID string) {
			resID = id
			resFlightoperationID = flightoperationID
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightreportID))
	a.Equal(resFlightoperationID, string(DefaultFlightoperationID))
}

func TestListFlightreportsTransaction(t *testing.T) {
	a := assert.New(t)

	flightreports := []*frep.Flightreport{
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:                string(DefaultFlightreportID),
				FlightoperationID: string(DefaultFlightoperationID),
			},
		),
	}

	repo := &flightreportRepositoryMock{}
	txm := &txManagerMock{}
	repo.On("GetAll").Return(flightreports, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListFlightreports(
		func(id, flightoperationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestListFlightreportsOperation(t *testing.T) {
	a := assert.New(t)

	const (
		DefaultFlightreportID1    = DefaultFlightreportID + "-1"
		DefaultFlightreportID2    = DefaultFlightreportID + "-2"
		DefaultFlightreportID3    = DefaultFlightreportID + "-3"
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
	)

	flightreports := []*frep.Flightreport{
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:                string(DefaultFlightreportID1),
				FlightoperationID: string(DefaultFlightoperationID1),
			},
		),
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:                string(DefaultFlightreportID2),
				FlightoperationID: string(DefaultFlightoperationID2),
			},
		),
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:                string(DefaultFlightreportID3),
				FlightoperationID: string(DefaultFlightoperationID3),
			},
		),
	}

	repo := &flightreportRepositoryMock{}
	repo.On("GetAll").Return(flightreports, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resID, resFlightoperationID []string
	ret := service.listFlightreportsOperation(
		nil,
		func(id, flightoperationID string) {
			resID = append(resID, id)
			resFlightoperationID = append(resFlightoperationID, flightoperationID)
		},
	)

	a.Nil(ret)
	a.Equal(resID[0], string(DefaultFlightreportID1))
	a.Equal(resFlightoperationID[0], string(DefaultFlightoperationID1))
	a.Equal(resID[1], string(DefaultFlightreportID2))
	a.Equal(resFlightoperationID[1], string(DefaultFlightoperationID2))
	a.Equal(resID[2], string(DefaultFlightreportID3))
	a.Equal(resFlightoperationID[2], string(DefaultFlightoperationID3))
}

func TestCreateFlightreportTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightoperationID + "-original"
		NewID      = DefaultFlightoperationID + "-new"
	)

	gen := &generatorMockFlightreport{
		id:                DefaultFlightreportID,
		flightoperationID: NewID,
	}
	repo := &flightreportRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	txm := &txManagerMock{}

	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFlightreportService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		FlightoperationID: string(OriginalID),
	}
	ret := service.CreateFlightreport(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestCreateFlightreportOperation(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFlightoperationID + "-original"
		NewID      = DefaultFlightoperationID + "-new"
	)

	gen := &generatorMockFlightreport{
		id:                DefaultFlightreportID,
		flightoperationID: NewID,
	}
	repo := &flightreportRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFlightreportService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightoperationIDRequestMock{
		FlightoperationID: string(OriginalID),
	}
	ret := service.createFlightreportOperation(
		nil,
		req,
	)

	a.Nil(ret)
}
