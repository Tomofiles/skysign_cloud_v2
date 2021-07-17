package service

import (
	fl "fleet-formation/pkg/fleet-formation/domain/fleet"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAssignmentsTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMockFleet{}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID: string(DefaultFleetID),
			Assignments: []*assignmentComponentMock{
				{
					ID:        string(DefaultFleetAssignmentID),
					VehicleID: string(DefaultFleetVehicleID),
				},
			},
			Events: []*eventComponentMock{
				{
					ID:           string(DefaultFleetEventID),
					AssignmentID: string(DefaultFleetAssignmentID),
					MissionID:    string(DefaultFleetMissionID),
				},
			},
			Version: string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultFleetID).Return(fleet, nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	command := &fleetIDCommandMock{
		FleetID: string(DefaultFleetID),
	}
	var resCall bool
	ret := service.GetAssignments(
		command,
		func(model AssignmentPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetAssignmentsOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetAssignmentID1 = DefaultFleetAssignmentID + "-1"
		DefaultFleetAssignmentID2 = DefaultFleetAssignmentID + "-2"
		DefaultFleetAssignmentID3 = DefaultFleetAssignmentID + "-3"
		DefaultFleetEventID1      = DefaultFleetEventID + "-1"
		DefaultFleetEventID2      = DefaultFleetEventID + "-2"
		DefaultFleetEventID3      = DefaultFleetEventID + "-3"
		DefaultFleetVehicleID1    = DefaultFleetVehicleID + "-1"
		DefaultFleetVehicleID2    = DefaultFleetVehicleID + "-2"
		DefaultFleetVehicleID3    = DefaultFleetVehicleID + "-3"
		DefaultFleetMissionID1    = DefaultFleetMissionID + "-1"
		DefaultFleetMissionID2    = DefaultFleetMissionID + "-2"
		DefaultFleetMissionID3    = DefaultFleetMissionID + "-3"
	)

	gen := &generatorMockFleet{}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID: string(DefaultFleetID),
			Assignments: []*assignmentComponentMock{
				{
					ID:        string(DefaultFleetAssignmentID1),
					VehicleID: string(DefaultFleetVehicleID1),
				},
				{
					ID:        string(DefaultFleetAssignmentID2),
					VehicleID: string(DefaultFleetVehicleID2),
				},
				{
					ID:        string(DefaultFleetAssignmentID3),
					VehicleID: string(DefaultFleetVehicleID3),
				},
			},
			Events: []*eventComponentMock{
				{
					ID:           string(DefaultFleetEventID1),
					AssignmentID: string(DefaultFleetAssignmentID1),
					MissionID:    string(DefaultFleetMissionID1),
				},
				{
					ID:           string(DefaultFleetEventID2),
					AssignmentID: string(DefaultFleetAssignmentID2),
					MissionID:    string(DefaultFleetMissionID2),
				},
				{
					ID:           string(DefaultFleetEventID3),
					AssignmentID: string(DefaultFleetAssignmentID3),
					MissionID:    string(DefaultFleetMissionID3),
				},
			},
			Version: string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	repo.On("GetByID", DefaultFleetID).Return(fleet, nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	command := &fleetIDCommandMock{
		FleetID: string(DefaultFleetID),
	}
	var resModels []AssignmentPresentationModel
	ret := service.getAssignmentsOperation(
		nil,
		command,
		func(model AssignmentPresentationModel) {
			resModels = append(resModels, model)
		},
	)

	expectModels := []AssignmentPresentationModel{
		&assignmentModel{
			assignment: &assignment{
				ID:           string(DefaultFleetID),
				EventID:      string(DefaultFleetEventID1),
				AssignmentID: string(DefaultFleetAssignmentID1),
				VehicleID:    string(DefaultFleetVehicleID1),
				MissionID:    string(DefaultFleetMissionID1),
			},
		},
		&assignmentModel{
			assignment: &assignment{
				ID:           string(DefaultFleetID),
				EventID:      string(DefaultFleetEventID2),
				AssignmentID: string(DefaultFleetAssignmentID2),
				VehicleID:    string(DefaultFleetVehicleID2),
				MissionID:    string(DefaultFleetMissionID2),
			},
		},
		&assignmentModel{
			assignment: &assignment{
				ID:           string(DefaultFleetID),
				EventID:      string(DefaultFleetEventID3),
				AssignmentID: string(DefaultFleetAssignmentID3),
				VehicleID:    string(DefaultFleetVehicleID3),
				MissionID:    string(DefaultFleetMissionID3),
			},
		},
	}

	a.Nil(ret)
	a.Equal(resModels, expectModels)
}

func TestUpdateAssignmentTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetVersion1 = DefaultFleetVersion + "-1"
		DefaultFleetVersion2 = DefaultFleetVersion + "-2"
		DefaultFleetVersion3 = DefaultFleetVersion + "-3"
		DefaultFleetVersion4 = DefaultFleetVersion + "-4"

		AfterFleetVehicleID = DefaultFleetVehicleID + "-after"
		AfterFleetMissionID = DefaultFleetMissionID + "-after"
	)

	gen := &generatorMockFleet{
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3, DefaultFleetVersion4},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID: string(DefaultFleetID),
			Assignments: []*assignmentComponentMock{
				{
					ID:        string(DefaultFleetAssignmentID),
					VehicleID: string(DefaultFleetVehicleID),
				},
			},
			Events: []*eventComponentMock{
				{
					ID:           string(DefaultFleetEventID),
					AssignmentID: string(DefaultFleetAssignmentID),
					MissionID:    string(DefaultFleetMissionID),
				},
			},
			Version: string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultFleetID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	command := &updateAssignmentCommandMock{
		ID:           string(DefaultFleetID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    string(AfterFleetVehicleID),
		MissionID:    string(AfterFleetMissionID),
	}
	var resCall bool
	ret := service.UpdateAssignment(
		command,
		func(model AssignmentPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestUpdateAssignmentOperationAssign(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetVersion1 = DefaultFleetVersion + "-1"
		DefaultFleetVersion2 = DefaultFleetVersion + "-2"

		AfterFleetVehicleID = DefaultFleetVehicleID + "-after"
		AfterFleetMissionID = DefaultFleetMissionID + "-after"
	)

	gen := &generatorMockFleet{
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID: string(DefaultFleetID),
			Assignments: []*assignmentComponentMock{
				{
					ID:        string(DefaultFleetAssignmentID),
					VehicleID: string(DefaultFleetVehicleID),
				},
			},
			Events: []*eventComponentMock{
				{
					ID:           string(DefaultFleetEventID),
					AssignmentID: string(DefaultFleetAssignmentID),
					MissionID:    string(DefaultFleetMissionID),
				},
			},
			Version: string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	repo.On("GetByID", DefaultFleetID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	command := &updateAssignmentCommandMock{
		ID:           string(DefaultFleetID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    string(AfterFleetVehicleID),
		MissionID:    string(AfterFleetMissionID),
	}
	var resModel AssignmentPresentationModel
	ret := service.updateAssignmentOperation(
		nil,
		command,
		func(model AssignmentPresentationModel) {
			resModel = model
		},
	)

	var actualAssignments []assignmentComponentMock
	var actualEvents []eventComponentMock
	repo.fleet.ProvideAssignmentsInterest(
		func(assignmentID string, vehicleID string) {
			actualAssignments = append(
				actualAssignments,
				assignmentComponentMock{
					ID:        assignmentID,
					VehicleID: vehicleID,
				},
			)
		},
		func(eventID string, assignmentID string, missionID string) {
			actualEvents = append(
				actualEvents,
				eventComponentMock{
					ID:           eventID,
					AssignmentID: assignmentID,
					MissionID:    missionID,
				},
			)
		},
	)

	expectAssignments := []assignmentComponentMock{
		{
			ID:        string(DefaultFleetAssignmentID),
			VehicleID: string(AfterFleetVehicleID),
		},
	}
	expectEvents := []eventComponentMock{
		{
			ID:           string(DefaultFleetEventID),
			AssignmentID: string(DefaultFleetAssignmentID),
			MissionID:    string(AfterFleetMissionID),
		},
	}

	a.Nil(ret)
	a.Equal(resModel.GetAssignment().GetID(), string(DefaultFleetID))
	a.Equal(resModel.GetAssignment().GetEventID(), string(DefaultFleetEventID))
	a.Equal(resModel.GetAssignment().GetAssignmentID(), string(DefaultFleetAssignmentID))
	a.Equal(resModel.GetAssignment().GetVehicleID(), string(AfterFleetVehicleID))
	a.Equal(resModel.GetAssignment().GetMissionID(), string(AfterFleetMissionID))
	a.Equal(actualAssignments, expectAssignments)
	a.Equal(actualEvents, expectEvents)
}

func TestUpdateAssignmentOperationCancel(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetVersion1 = DefaultFleetVersion + "-1"
		DefaultFleetVersion2 = DefaultFleetVersion + "-2"
	)

	gen := &generatorMockFleet{
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID: string(DefaultFleetID),
			Assignments: []*assignmentComponentMock{
				{
					ID:        string(DefaultFleetAssignmentID),
					VehicleID: string(DefaultFleetVehicleID),
				},
			},
			Events: []*eventComponentMock{
				{
					ID:           string(DefaultFleetEventID),
					AssignmentID: string(DefaultFleetAssignmentID),
					MissionID:    string(DefaultFleetMissionID),
				},
			},
			Version: string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	repo.On("GetByID", DefaultFleetID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	command := &updateAssignmentCommandMock{
		ID:           string(DefaultFleetID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    "",
		MissionID:    "",
	}
	var resModel AssignmentPresentationModel
	ret := service.updateAssignmentOperation(
		nil,
		command,
		func(model AssignmentPresentationModel) {
			resModel = model
		},
	)

	var actualAssignments []assignmentComponentMock
	var actualEvents []eventComponentMock
	repo.fleet.ProvideAssignmentsInterest(
		func(assignmentID string, vehicleID string) {
			actualAssignments = append(
				actualAssignments,
				assignmentComponentMock{
					ID:        assignmentID,
					VehicleID: vehicleID,
				},
			)
		},
		func(eventID string, assignmentID string, missionID string) {
			actualEvents = append(
				actualEvents,
				eventComponentMock{
					ID:           eventID,
					AssignmentID: assignmentID,
					MissionID:    missionID,
				},
			)
		},
	)

	expectAssignments := []assignmentComponentMock{
		{
			ID: string(DefaultFleetAssignmentID),
		},
	}
	expectEvents := []eventComponentMock{
		{
			ID:           string(DefaultFleetEventID),
			AssignmentID: string(DefaultFleetAssignmentID),
		},
	}

	a.Nil(ret)
	a.Equal(resModel.GetAssignment().GetID(), string(DefaultFleetID))
	a.Equal(resModel.GetAssignment().GetEventID(), string(DefaultFleetEventID))
	a.Equal(resModel.GetAssignment().GetAssignmentID(), string(DefaultFleetAssignmentID))
	a.Empty(resModel.GetAssignment().GetVehicleID())
	a.Empty(resModel.GetAssignment().GetMissionID())
	a.Equal(actualAssignments, expectAssignments)
	a.Equal(actualEvents, expectEvents)
}
