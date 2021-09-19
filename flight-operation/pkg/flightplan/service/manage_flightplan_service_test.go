package service

import (
	"testing"

	fpl "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	flightplan := fpl.AssembleFrom(
		nil,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			FleetID:     string(DefaultFlightplanFleetID),
			Description: DefaultFlightplanDescription,
			Version:     string(DefaultFlightplanVersion),
		},
	)

	repo := &flightplanRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)

	service := &manageFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	var resCall bool
	ret := service.GetFlightplan(
		command,
		func(model FlightplanPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	flightplan := fpl.AssembleFrom(
		nil,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			FleetID:     string(DefaultFlightplanFleetID),
			Description: DefaultFlightplanDescription,
			Version:     string(DefaultFlightplanVersion),
		},
	)

	repo := &flightplanRepositoryMock{}
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)

	service := &manageFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	var resModel FlightplanPresentationModel
	ret := service.getFlightplanOperation(
		nil,
		command,
		func(model FlightplanPresentationModel) {
			resModel = model
		},
	)

	a.Nil(ret)
	a.Equal(resModel.GetFlightplan().GetID(), string(DefaultFlightplanID))
	a.Equal(resModel.GetFlightplan().GetName(), DefaultFlightplanName)
	a.Equal(resModel.GetFlightplan().GetDescription(), DefaultFlightplanDescription)
	a.Equal(resModel.GetFlightplan().GetFleetID(), string(DefaultFlightplanFleetID))
}

func TestListFlightplansTransaction(t *testing.T) {
	a := assert.New(t)

	flightplans := []*fpl.Flightplan{
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          string(DefaultFlightplanID),
				Name:        DefaultFlightplanName,
				Description: DefaultFlightplanDescription,
				FleetID:     string(DefaultFlightplanFleetID),
				Version:     string(DefaultFlightplanVersion),
			},
		),
	}

	repo := &flightplanRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetAll").Return(flightplans, nil)

	service := &manageFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListFlightplans(
		func(model FlightplanPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestListFlightplansOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanID1          = string(DefaultFlightplanID) + "-1"
		DefaultFlightplanName1        = DefaultFlightplanName + "-1"
		DefaultFlightplanDescription1 = DefaultFlightplanDescription + "-1"
		DefaultFlightplanFleetID1     = string(DefaultFlightplanFleetID) + "-1"
		DefaultFlightplanVersion1     = string(DefaultFlightplanVersion) + "-1"
		DefaultFlightplanID2          = string(DefaultFlightplanID) + "-2"
		DefaultFlightplanName2        = DefaultFlightplanName + "-2"
		DefaultFlightplanDescription2 = DefaultFlightplanDescription + "-2"
		DefaultFlightplanFleetID2     = string(DefaultFlightplanFleetID) + "-2"
		DefaultFlightplanVersion2     = string(DefaultFlightplanVersion) + "-2"
		DefaultFlightplanID3          = string(DefaultFlightplanID) + "-3"
		DefaultFlightplanName3        = DefaultFlightplanName + "-3"
		DefaultFlightplanDescription3 = DefaultFlightplanDescription + "-3"
		DefaultFlightplanFleetID3     = string(DefaultFlightplanFleetID) + "-3"
		DefaultFlightplanVersion3     = string(DefaultFlightplanVersion) + "-3"
	)

	flightplans := []*fpl.Flightplan{
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID1,
				Name:        DefaultFlightplanName1,
				Description: DefaultFlightplanDescription1,
				FleetID:     DefaultFlightplanFleetID1,
				Version:     DefaultFlightplanVersion1,
			},
		),
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID2,
				Name:        DefaultFlightplanName2,
				Description: DefaultFlightplanDescription2,
				FleetID:     DefaultFlightplanFleetID2,
				Version:     DefaultFlightplanVersion2,
			},
		),
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID3,
				Name:        DefaultFlightplanName3,
				Description: DefaultFlightplanDescription3,
				FleetID:     DefaultFlightplanFleetID3,
				Version:     DefaultFlightplanVersion3,
			},
		),
	}

	repo := &flightplanRepositoryMock{}
	repo.On("GetAll").Return(flightplans, nil)

	service := &manageFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resModels []FlightplanPresentationModel
	ret := service.listFlightplansOperation(
		nil,
		func(model FlightplanPresentationModel) {
			resModels = append(resModels, model)
		},
	)

	a.Nil(ret)
	a.Len(resModels, 3)
}

func TestCreateFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		DefaultFlightplanVersion3 = DefaultFlightplanVersion + "-3"
		DefaultFlightplanVersion4 = DefaultFlightplanVersion + "-4"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		fleetID:  DefaultFlightplanFleetID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2, DefaultFlightplanVersion3, DefaultFlightplanVersion4},
	}
	repo := &flightplanRepositoryMock{}
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

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &flightplanCommandMock{
		Flightplan: flightplanMock{
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
		},
	}
	var resCall1, resCall2 bool
	ret := service.CreateFlightplan(
		command,
		func(id string) {
			resCall1 = true
		},
		func(fleetID string) {
			resCall2 = true
		},
	)

	a.Nil(ret)
	a.True(resCall1)
	a.True(resCall2)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		DefaultFlightplanVersion3 = DefaultFlightplanVersion + "-3"
		DefaultFlightplanVersion4 = DefaultFlightplanVersion + "-4"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		fleetID:  DefaultFlightplanFleetID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2, DefaultFlightplanVersion3, DefaultFlightplanVersion4},
	}
	repo := &flightplanRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightplanCommandMock{
		Flightplan: flightplanMock{
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
		},
	}
	var resID, resFleetID string
	ret := service.createFlightplanOperation(
		nil,
		pub,
		command,
		func(id string) {
			resID = id
		},
		func(fleetID string) {
			resFleetID = fleetID
		},
	)

	expectEvent := fpl.FleetIDGaveEvent{
		FleetID:          DefaultFlightplanFleetID,
		NumberOfVehicles: 0,
	}

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightplanID))
	a.Equal(resFleetID, string(DefaultFlightplanFleetID))
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}

func TestUpdateFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		AfterFlightplanName        = DefaultFlightplanName + "-after"
		AfterFlightplanDescription = DefaultFlightplanDescription + "-after"
		DefaultFlightplanVersion1  = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2  = DefaultFlightplanVersion + "-2"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2},
	}

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
			Version:     string(DefaultFlightplanVersion),
		},
	)

	repo := &flightplanRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &flightplanCommandMock{
		Flightplan: flightplanMock{
			ID:          string(DefaultFlightplanID),
			Name:        AfterFlightplanName,
			Description: AfterFlightplanDescription,
		},
	}
	var resCall bool
	ret := service.UpdateFlightplan(
		command,
		func(fleetID string) {
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

func TestUpdateFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	var (
		AfterFlightplanName        = DefaultFlightplanName + "-after"
		AfterFlightplanDescription = DefaultFlightplanDescription + "-after"
		DefaultFlightplanVersion1  = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2  = DefaultFlightplanVersion + "-2"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2},
	}

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
			FleetID:     string(DefaultFlightplanFleetID),
			Version:     string(DefaultFlightplanVersion),
		},
	)

	repo := &flightplanRepositoryMock{}
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightplanCommandMock{
		Flightplan: flightplanMock{
			ID:          string(DefaultFlightplanID),
			Name:        AfterFlightplanName,
			Description: AfterFlightplanDescription,
		},
	}
	var resFleetID string
	ret := service.updateFlightplanOperation(
		nil,
		pub,
		command,
		func(fleetID string) {
			resFleetID = fleetID
		},
	)

	a.Nil(ret)
	a.Equal(resFleetID, string(DefaultFlightplanFleetID))
	a.Len(pub.events, 0)
}

func TestDeleteFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
	)

	gen := &generatorMockFlightplan{
		versions: []fpl.Version{DefaultFlightplanVersion2},
	}

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
			FleetID:     string(DefaultFlightplanFleetID),
			Version:     string(DefaultFlightplanVersion1),
		},
	)

	repo := &flightplanRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.DeleteFlightplan(command)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestDeleteFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
	)

	gen := &generatorMockFlightplan{
		versions: []fpl.Version{DefaultFlightplanVersion2},
	}
	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			ID:          string(DefaultFlightplanID),
			Name:        DefaultFlightplanName,
			Description: DefaultFlightplanDescription,
			FleetID:     string(DefaultFlightplanFleetID),
			Version:     string(DefaultFlightplanVersion1),
		},
	)

	repo := &flightplanRepositoryMock{}
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)
	repo.On("Delete", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.deleteFlightplanOperation(
		nil,
		pub,
		command,
	)

	expectEvent := fpl.FleetIDRemovedEvent{FleetID: DefaultFlightplanFleetID}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}
