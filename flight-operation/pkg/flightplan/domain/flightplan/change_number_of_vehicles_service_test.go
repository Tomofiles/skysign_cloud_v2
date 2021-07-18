package flightplan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
		NewFleetID          = DefaultFleetID + "-new"
		NewNumberOfVehicles = 10
	)

	gen := &generatorMock{
		fleetID:  NewFleetID,
		versions: []Version{DefaultVersion2},
	}
	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := ChangeNumberOfVehicles(ctx, gen, repo, pub, DefaultID, NewNumberOfVehicles)

	expectFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     NewFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion2,
		gen:         gen,
		pub:         pub,
	}

	expectEvent1 := FleetIDRemovedEvent{
		FleetID: DefaultFleetID,
	}
	expectEvent2 := FleetIDGaveEvent{
		FleetID:          NewFleetID,
		NumberOfVehicles: NewNumberOfVehicles,
	}

	a.Len(repo.saveFlightplans, 1)
	a.Equal(repo.saveFlightplans[0], &expectFlightplan)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2})

	a.Nil(ret)
}

func TestGetErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewNumberOfVehicles = 10
	)

	gen := &generatorMock{}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := ChangeNumberOfVehicles(ctx, gen, repo, pub, DefaultID, NewNumberOfVehicles)

	a.Equal(ret, ErrGet)
}

func TestSaveErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
		NewFleetID          = DefaultFleetID + "-new"
		NewNumberOfVehicles = 10
	)

	gen := &generatorMock{
		fleetID:  NewFleetID,
		versions: []Version{DefaultVersion2},
	}
	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	ret := ChangeNumberOfVehicles(ctx, gen, repo, pub, DefaultID, NewNumberOfVehicles)

	a.Equal(ret, ErrSave)
}
