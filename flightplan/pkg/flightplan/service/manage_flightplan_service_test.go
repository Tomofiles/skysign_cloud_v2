package service

import (
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"testing"

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

	req := &flightplanIDRequestMock{
		ID: string(DefaultFlightplanID),
	}
	var resCall bool
	ret := service.GetFlightplan(
		req,
		func(id, name, description string) {
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

	req := &flightplanIDRequestMock{
		ID: string(DefaultFlightplanID),
	}
	var resID, resName, resDescription string
	ret := service.getFlightplanOperation(
		nil,
		req,
		func(id, name, description string) {
			resID = id
			resName = name
			resDescription = description
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightplanID))
	a.Equal(resName, DefaultFlightplanName)
	a.Equal(resDescription, DefaultFlightplanDescription)
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
		func(id, name, description string) {
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
		DefaultFlightplanVersion1     = string(DefaultFlightplanVersion) + "-1"
		DefaultFlightplanID2          = string(DefaultFlightplanID) + "-2"
		DefaultFlightplanName2        = DefaultFlightplanName + "-2"
		DefaultFlightplanDescription2 = DefaultFlightplanDescription + "-2"
		DefaultFlightplanVersion2     = string(DefaultFlightplanVersion) + "-2"
		DefaultFlightplanID3          = string(DefaultFlightplanID) + "-3"
		DefaultFlightplanName3        = DefaultFlightplanName + "-3"
		DefaultFlightplanDescription3 = DefaultFlightplanDescription + "-3"
		DefaultFlightplanVersion3     = string(DefaultFlightplanVersion) + "-3"
	)

	flightplans := []*fpl.Flightplan{
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID1,
				Name:        DefaultFlightplanName1,
				Description: DefaultFlightplanDescription1,
				Version:     DefaultFlightplanVersion1,
			},
		),
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID2,
				Name:        DefaultFlightplanName2,
				Description: DefaultFlightplanDescription2,
				Version:     DefaultFlightplanVersion2,
			},
		),
		fpl.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          DefaultFlightplanID3,
				Name:        DefaultFlightplanName3,
				Description: DefaultFlightplanDescription3,
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

	var resID, resName, resDescription []string
	ret := service.listFlightplansOperation(
		nil,
		func(id, name, description string) {
			resID = append(resID, id)
			resName = append(resName, name)
			resDescription = append(resDescription, description)
		},
	)

	a.Nil(ret)
	a.Equal(resID[0], DefaultFlightplanID1)
	a.Equal(resName[0], DefaultFlightplanName1)
	a.Equal(resDescription[0], DefaultFlightplanDescription1)
	a.Equal(resID[1], DefaultFlightplanID2)
	a.Equal(resName[1], DefaultFlightplanName2)
	a.Equal(resDescription[1], DefaultFlightplanDescription2)
	a.Equal(resID[2], DefaultFlightplanID3)
	a.Equal(resName[2], DefaultFlightplanName3)
	a.Equal(resDescription[2], DefaultFlightplanDescription3)
}

func TestCreateFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		DefaultFlightplanVersion3 = DefaultFlightplanVersion + "-3"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2, DefaultFlightplanVersion3},
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

	req := &flightplanRequestMock{
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
	}
	var resCall bool
	ret := service.CreateFlightplan(
		req,
		func(id, name, description string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
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
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2, DefaultFlightplanVersion3},
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

	req := &flightplanRequestMock{
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
	}
	var resID, resName, resDescription string
	ret := service.createFlightplanOperation(
		nil,
		pub,
		req,
		func(id, name, description string) {
			resID = id
			resName = name
			resDescription = description
		},
	)

	expectEvent := fpl.CreatedEvent{ID: DefaultFlightplanID}

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightplanID))
	a.Equal(resName, DefaultFlightplanName)
	a.Equal(resDescription, DefaultFlightplanDescription)
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

	req := &flightplanRequestMock{
		ID:          string(DefaultFlightplanID),
		Name:        AfterFlightplanName,
		Description: AfterFlightplanDescription,
	}
	var resCall bool
	ret := service.UpdateFlightplan(
		req,
		func(id, name, description string) {
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

	req := &flightplanRequestMock{
		ID:          string(DefaultFlightplanID),
		Name:        AfterFlightplanName,
		Description: AfterFlightplanDescription,
	}
	var resID, resName, resDescription string
	ret := service.updateFlightplanOperation(
		nil,
		pub,
		req,
		func(id, name, description string) {
			resID = id
			resName = name
			resDescription = description
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultFlightplanID))
	a.Equal(resName, AfterFlightplanName)
	a.Equal(resDescription, AfterFlightplanDescription)
	a.Len(pub.events, 0)
}

func TestDeleteFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion1 = DefaultFlightplanVersion + "-1"
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		DefaultFlightplanVersion3 = DefaultFlightplanVersion + "-3"
	)

	gen := &generatorMockFlightplan{
		id:       DefaultFlightplanID,
		versions: []fpl.Version{DefaultFlightplanVersion1, DefaultFlightplanVersion2, DefaultFlightplanVersion3},
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
	repo.On("Delete", mock.Anything).Return(nil)

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &flightplanIDRequestMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.DeleteFlightplan(req)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestDeleteFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMockFlightplan{}

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
	repo.On("GetByID", DefaultFlightplanID).Return(flightplan, nil)
	repo.On("Delete", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &flightplanIDRequestMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.deleteFlightplanOperation(
		nil,
		pub,
		req,
	)

	expectEvent := fpl.DeletedEvent{ID: DefaultFlightplanID}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}
