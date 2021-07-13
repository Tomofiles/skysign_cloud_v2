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
		DefaultAssignmentID1    = DefaultFleetAssignmentID + "-1"
		DefaultAssignmentID2    = DefaultFleetAssignmentID + "-2"
		DefaultAssignmentID3    = DefaultFleetAssignmentID + "-3"
		DefaultFleetVersion1    = DefaultFleetVersion + "-1"
		DefaultFleetVersion2    = DefaultFleetVersion + "-2"
		DefaultFleetVersion3    = DefaultFleetVersion + "-3"
		DefaultNumberOfVehicles = 3
	)

	gen := &generatorMockFleet{
		assignmentIDs: []fl.AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3},
	}

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	command := &changeNumberOfVehiclesCommandFleetMock{
		FleetID:          string(DefaultFleetID),
		NumberOfVehicles: DefaultNumberOfVehicles,
	}
	ret := service.CreateFleet(command)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestCreateFleetOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1    = DefaultFleetAssignmentID + "-1"
		DefaultAssignmentID2    = DefaultFleetAssignmentID + "-2"
		DefaultAssignmentID3    = DefaultFleetAssignmentID + "-3"
		DefaultFleetVersion1    = DefaultFleetVersion + "-1"
		DefaultFleetVersion2    = DefaultFleetVersion + "-2"
		DefaultFleetVersion3    = DefaultFleetVersion + "-3"
		DefaultNumberOfVehicles = 3
	)

	gen := &generatorMockFleet{
		assignmentIDs: []fl.AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3},
	}

	repo := &fleetRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	command := &changeNumberOfVehiclesCommandFleetMock{
		FleetID:          string(DefaultFleetID),
		NumberOfVehicles: DefaultNumberOfVehicles,
	}
	ret := service.createFleetOperation(nil, command)

	a.Nil(ret)
	a.Equal(repo.fleet.GetID(), DefaultFleetID)
	a.Len(repo.fleet.GetAllAssignmentID(), 3)
}

func TestDeleteFleetTransaction(t *testing.T) {
	a := assert.New(t)

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("Delete", DefaultFleetID).Return(nil)

	service := &manageFleetService{
		gen:  nil,
		repo: repo,
		txm:  txm,
	}

	command := &fleetIDCommandMock{
		FleetID: string(DefaultFleetID),
	}
	ret := service.DeleteFleet(command)

	a.Nil(ret)
	a.Nil(txm.isOpe)
}

func TestDeleteFleetOperation(t *testing.T) {
	a := assert.New(t)

	repo := &fleetRepositoryMock{}
	repo.On("Delete", DefaultFleetID).Return(nil)

	service := &manageFleetService{
		gen:  nil,
		repo: repo,
		txm:  nil,
	}

	command := &fleetIDCommandMock{
		FleetID: string(DefaultFleetID),
	}
	ret := service.deleteFleetOperation(nil, command)

	a.Nil(ret)
	a.Equal(repo.deleteID, DefaultFleetID)
}

func TestCarbonCopyFleetTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetOriginalID = DefaultFleetID + "-new"
		DefaultFleetNewID      = DefaultFleetID + "-new"
	)

	fleet := fl.AssembleFrom(
		nil,
		&fleetComponentMock{
			ID:           string(DefaultFleetOriginalID),
			IsCarbonCopy: fl.Original,
			Version:      string(DefaultFleetVersion),
		},
	)

	gen := &generatorMockFleet{}
	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultFleetOriginalID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &carbonCopyRequestMock{
		OriginalID: string(DefaultFleetOriginalID),
		NewID:      string(DefaultFleetNewID),
	}
	ret := service.CarbonCopyFleet(req)

	a.Nil(ret)
	a.Nil(ret)
	a.Len(pub.events, 0)
	a.True(pub.isFlush)
	a.True(isClose)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCarbonCopyFleetOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetOriginalID = DefaultFleetID + "-new"
		DefaultFleetNewID      = DefaultFleetID + "-new"
	)

	fleet := fl.AssembleFrom(
		nil,
		&fleetComponentMock{
			ID:           string(DefaultFleetOriginalID),
			IsCarbonCopy: fl.Original,
			Version:      string(DefaultFleetVersion),
		},
	)

	gen := &generatorMockFleet{}

	repo := &fleetRepositoryMock{}
	repo.On("GetByID", DefaultFleetOriginalID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &carbonCopyRequestMock{
		OriginalID: string(DefaultFleetOriginalID),
		NewID:      string(DefaultFleetNewID),
	}
	ret := service.carbonCopyFleetOperation(
		nil,
		pub,
		req,
	)

	a.Nil(ret)
	a.Len(pub.events, 0)
}
