package service

import (
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFlightplanChangeNumberOfVehiclesTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		NewFlightplanFleetID      = DefaultFlightplanFleetID + "-new"
		DefaultNumberOfVehicles   = 10
	)

	gen := &generatorMockFlightplan{
		fleetID:  NewFlightplanFleetID,
		versions: []fpl.Version{DefaultFlightplanVersion2},
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

	service := &changeFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &changeNumberOfVehiclesCommandFlightplanMock{
		ID:               string(DefaultFlightplanID),
		NumberOfVehicles: DefaultNumberOfVehicles,
	}
	ret := service.ChangeNumberOfVehicles(command)

	a.Nil(ret)
	a.Len(pub.events, 2)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestFlightplanChangeNumberOfVehiclesOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanVersion2 = DefaultFlightplanVersion + "-2"
		NewFlightplanFleetID      = DefaultFlightplanFleetID + "-new"
		DefaultNumberOfVehicles   = 10
	)

	gen := &generatorMockFlightplan{
		fleetID:  NewFlightplanFleetID,
		versions: []fpl.Version{DefaultFlightplanVersion2},
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

	service := &changeFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &changeNumberOfVehiclesCommandFlightplanMock{
		ID:               string(DefaultFlightplanID),
		NumberOfVehicles: DefaultNumberOfVehicles,
	}
	ret := service.changeNumberOfVehiclesOperation(
		nil,
		pub,
		command,
	)

	expectEvent1 := fpl.FleetIDRemovedEvent{
		FleetID: DefaultFlightplanFleetID,
	}
	expectEvent2 := fpl.FleetIDGaveEvent{
		FleetID:          NewFlightplanFleetID,
		NumberOfVehicles: DefaultNumberOfVehicles,
	}

	a.Nil(ret)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2})
}
