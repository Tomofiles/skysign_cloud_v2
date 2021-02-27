package service

import (
	fl "flightplan/pkg/flightplan/domain/fleet"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFleetTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetVersion1 = DefaultFleetVersion + "-1"
		DefaultFleetVersion2 = DefaultFleetVersion + "-2"
		DefaultFleetVersion3 = DefaultFleetVersion + "-3"
	)

	gen := &generatorMockFleet{
		id:       DefaultFleetID,
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3},
	}

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	ret := service.CreateFleet(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestCreateFleetOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetVersion1 = DefaultFleetVersion + "-1"
		DefaultFleetVersion2 = DefaultFleetVersion + "-2"
		DefaultFleetVersion3 = DefaultFleetVersion + "-3"
	)

	gen := &generatorMockFleet{
		id:       DefaultFleetID,
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3},
	}

	repo := &fleetRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	ret := service.createFleetOperation(nil, req)

	a.Nil(ret)
	a.Equal(repo.fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Len(repo.fleet.GetAllAssignmentID(), 0)
}

func TestDeleteFleetTransaction(t *testing.T) {
	a := assert.New(t)

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)

	service := &manageFleetService{
		gen:  nil,
		repo: repo,
		txm:  txm,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	ret := service.DeleteFleet(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestDeleteFleetOperation(t *testing.T) {
	a := assert.New(t)

	repo := &fleetRepositoryMock{}
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)

	service := &manageFleetService{
		gen:  nil,
		repo: repo,
		txm:  nil,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	ret := service.deleteFleetOperation(nil, req)

	a.Nil(ret)
	a.Equal(repo.deleteID, DefaultFlightplanID)
}
