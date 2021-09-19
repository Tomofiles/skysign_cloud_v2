package service

import (
	"testing"

	frep "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/domain/flightreport"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFlightreportTransaction(t *testing.T) {
	a := assert.New(t)

	flightreport := frep.AssembleFrom(
		nil,
		&flightreportComponentMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
		},
	)

	repo := &flightreportRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultID).Return(flightreport, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	command := &flightreportIDCommandMock{
		ID: string(DefaultID),
	}
	var resCall bool
	ret := service.GetFlightreport(
		command,
		func(model FlightreportPresentationModel) {
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
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
		},
	)

	repo := &flightreportRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightreport, nil)

	service := &manageFlightreportService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightreportIDCommandMock{
		ID: string(DefaultID),
	}
	var resModel FlightreportPresentationModel
	ret := service.getFlightreportOperation(
		nil,
		command,
		func(model FlightreportPresentationModel) {
			resModel = model
		},
	)

	a.Nil(ret)
	a.Equal(resModel.GetFlightreport().GetID(), string(DefaultID))
	a.Equal(resModel.GetFlightreport().GetName(), DefaultName)
	a.Equal(resModel.GetFlightreport().GetDescription(), DefaultDescription)
	a.Equal(resModel.GetFlightreport().GetFleetID(), string(DefaultFleetID))
}

func TestListFlightreportsTransaction(t *testing.T) {
	a := assert.New(t)

	flightreports := []*frep.Flightreport{
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:          string(DefaultID),
				Name:        DefaultName,
				Description: DefaultDescription,
				FleetID:     string(DefaultFleetID),
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
		func(model FlightreportPresentationModel) {
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
		DefaultID1          = string(DefaultID) + "-1"
		DefaultName1        = DefaultName + "-1"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultFleetID1     = string(DefaultFleetID) + "-1"
		DefaultID2          = string(DefaultID) + "-2"
		DefaultName2        = DefaultName + "-2"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultFleetID2     = string(DefaultFleetID) + "-2"
		DefaultID3          = string(DefaultID) + "-3"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID3     = string(DefaultFleetID) + "-3"
	)

	flightreports := []*frep.Flightreport{
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:          DefaultID1,
				Name:        DefaultName1,
				Description: DefaultDescription1,
				FleetID:     DefaultFleetID1,
			},
		),
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:          DefaultID2,
				Name:        DefaultName2,
				Description: DefaultDescription2,
				FleetID:     DefaultFleetID2,
			},
		),
		frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:          DefaultID3,
				Name:        DefaultName3,
				Description: DefaultDescription3,
				FleetID:     DefaultFleetID3,
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

	var resModels []FlightreportPresentationModel
	ret := service.listFlightreportsOperation(
		nil,
		func(model FlightreportPresentationModel) {
			resModels = append(resModels, model)
		},
	)

	a.Nil(ret)
	a.Len(resModels, 3)
}

func TestCreateFlightreportTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMockFlightreport{
		id: DefaultID,
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

	command := &flightreportCommandMock{
		Flightreport: flightreportMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
		},
	}
	ret := service.CreateFlightreport(command)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestCreateFlightreportOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMockFlightreport{
		id: DefaultID,
	}
	repo := &flightreportRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFlightreportService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightreportCommandMock{
		Flightreport: flightreportMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
		},
	}
	ret := service.createFlightreportOperation(
		nil,
		command,
	)

	a.Nil(ret)
}
