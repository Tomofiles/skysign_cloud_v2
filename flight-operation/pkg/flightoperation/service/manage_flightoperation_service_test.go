package service

import (
	"testing"

	fope "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	flightoperation := fope.AssembleFrom(
		nil,
		&flightoperationComponentMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
			IsCompleted: DefaultIsCompleted,
			Version:     string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultID).Return(flightoperation, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	command := &flightoperationIDCommandMock{
		ID: string(DefaultID),
	}
	var resCall bool
	ret := service.GetFlightoperation(
		command,
		func(model FlightoperationPresentationModel) {
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
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(DefaultFleetID),
			IsCompleted: DefaultIsCompleted,
			Version:     string(DefaultVersion),
		},
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightoperationIDCommandMock{
		ID: string(DefaultID),
	}
	var resModel FlightoperationPresentationModel
	ret := service.getFlightoperationOperation(
		nil,
		command,
		func(model FlightoperationPresentationModel) {
			resModel = model
		},
	)

	a.Nil(ret)
	a.Equal(resModel.GetFlightoperation().GetID(), string(DefaultID))
	a.Equal(resModel.GetFlightoperation().GetName(), DefaultName)
	a.Equal(resModel.GetFlightoperation().GetDescription(), DefaultDescription)
	a.Equal(resModel.GetFlightoperation().GetFleetID(), string(DefaultFleetID))
}

func TestListFlightoperationsTransaction(t *testing.T) {
	a := assert.New(t)

	flightoperations := []*fope.Flightoperation{
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:          string(DefaultID),
				Name:        DefaultName,
				Description: DefaultDescription,
				FleetID:     string(DefaultFleetID),
				IsCompleted: DefaultIsCompleted,
				Version:     string(DefaultVersion),
			},
		),
	}

	repo := &flightoperationRepositoryMock{}
	txm := &txManagerMock{}
	repo.On("GetAllOperating").Return(flightoperations, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListFlightoperations(
		func(model FlightoperationPresentationModel) {
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
		DefaultID1          = DefaultID + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName1        = DefaultName + "-1"
		DefaultName2        = DefaultName + "-2"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID1     = DefaultFleetID + "-1"
		DefaultFleetID2     = DefaultFleetID + "-2"
		DefaultFleetID3     = DefaultFleetID + "-3"
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
		DefaultVersion3     = DefaultVersion + "-3"
	)

	flightoperations := []*fope.Flightoperation{
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:          string(DefaultID1),
				Name:        DefaultName1,
				Description: DefaultDescription1,
				FleetID:     string(DefaultFleetID1),
				IsCompleted: DefaultIsCompleted,
				Version:     string(DefaultVersion1),
			},
		),
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:          string(DefaultID2),
				Name:        DefaultName2,
				Description: DefaultDescription2,
				FleetID:     string(DefaultFleetID2),
				IsCompleted: DefaultIsCompleted,
				Version:     string(DefaultVersion2),
			},
		),
		fope.AssembleFrom(
			nil,
			&flightoperationComponentMock{
				ID:          string(DefaultID3),
				Name:        DefaultName3,
				Description: DefaultDescription3,
				FleetID:     string(DefaultFleetID3),
				IsCompleted: DefaultIsCompleted,
				Version:     string(DefaultVersion3),
			},
		),
	}

	repo := &flightoperationRepositoryMock{}
	repo.On("GetAllOperating").Return(flightoperations, nil)

	service := &manageFlightoperationService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resModels []FlightoperationPresentationModel
	ret := service.listFlightoperationsOperation(
		nil,
		func(model FlightoperationPresentationModel) {
			resModels = append(resModels, model)
		},
	)

	a.Nil(ret)
	a.Len(resModels, 3)
}

func TestCreateFlightoperationTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetID + "-original"
		NewID      = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		fleetID: NewID,
		version: DefaultVersion,
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

	command := &flightoperationCommandMock{
		Flightoperation: flightoperationMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(OriginalID),
		},
	}
	ret := service.CreateFlightoperation(command)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateFlightoperationOperation(t *testing.T) {
	a := assert.New(t)

	var (
		OriginalID = DefaultFleetID + "-original"
		NewID      = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		fleetID: NewID,
		version: DefaultVersion,
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

	command := &flightoperationCommandMock{
		Flightoperation: flightoperationMock{
			ID:          string(DefaultID),
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetID:     string(OriginalID),
		},
	}
	ret := service.createFlightoperationOperation(
		nil,
		pub,
		command,
	)

	expectEvent := fope.FleetCopiedEvent{
		OriginalID: OriginalID,
		NewID:      NewID,
	}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
}
