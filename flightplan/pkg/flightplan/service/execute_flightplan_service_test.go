package service

import (
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightplanExecuteFlightplanTransaction(t *testing.T) {
	a := assert.New(t)

	flightplan := fpl.AssembleFrom(
		nil,
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

	service := &executeFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.ExecuteFlightplan(command)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestFlightplanExecuteFlightplanOperation(t *testing.T) {
	a := assert.New(t)

	flightplan := fpl.AssembleFrom(
		nil,
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
	pub := &publisherMock{}

	service := &executeFlightplanService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &flightplanIDCommandMock{
		ID: string(DefaultFlightplanID),
	}
	ret := service.executeFlightplanOperation(
		nil,
		pub,
		command,
	)

	expectEvent := &fpl.FlightplanExecutedEvent{
		ID:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetID:     DefaultFlightplanFleetID,
	}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
}
