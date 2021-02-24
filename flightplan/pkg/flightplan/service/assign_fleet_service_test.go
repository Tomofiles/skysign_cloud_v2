package service

import (
	fl "flightplan/pkg/flightplan/domain/fleet"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeNumberOfVehiclesTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetAssignmentID1 = DefaultFleetAssignmentID + "-1"
		DefaultFleetAssignmentID2 = DefaultFleetAssignmentID + "-2"
		DefaultFleetAssignmentID3 = DefaultFleetAssignmentID + "-3"
		DefaultFleetEventID1      = DefaultFleetEventID + "-1"
		DefaultFleetEventID2      = DefaultFleetEventID + "-2"
		DefaultFleetEventID3      = DefaultFleetEventID + "-3"
		DefaultFleetVersion1      = DefaultFleetVersion + "-1"
		DefaultFleetVersion2      = DefaultFleetVersion + "-2"
		DefaultFleetVersion3      = DefaultFleetVersion + "-3"
		DefaultFleetVersion4      = DefaultFleetVersion + "-4"
	)

	gen := &generatorMockFleet{
		id:            DefaultFleetID,
		assignmentIDs: []fl.AssignmentID{DefaultFleetAssignmentID1, DefaultFleetAssignmentID2, DefaultFleetAssignmentID3},
		eventIDs:      []fl.EventID{DefaultFleetEventID1, DefaultFleetEventID2, DefaultFleetEventID3},
		versions:      []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3, DefaultFleetVersion4},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
			Version:      string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	req := &changeNumberOfVehiclesRequestMock{
		FlightplanID:     string(DefaultFlightplanID),
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	var resCall bool
	ret := service.ChangeNumberOfVehicles(
		req,
		func(id string, numberOfVehicles int32) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestChangeNumberOfVehiclesOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetAssignmentID1 = DefaultFleetAssignmentID + "-1"
		DefaultFleetAssignmentID2 = DefaultFleetAssignmentID + "-2"
		DefaultFleetAssignmentID3 = DefaultFleetAssignmentID + "-3"
		DefaultFleetEventID1      = DefaultFleetEventID + "-1"
		DefaultFleetEventID2      = DefaultFleetEventID + "-2"
		DefaultFleetEventID3      = DefaultFleetEventID + "-3"
		DefaultFleetVersion1      = DefaultFleetVersion + "-1"
		DefaultFleetVersion2      = DefaultFleetVersion + "-2"
		DefaultFleetVersion3      = DefaultFleetVersion + "-3"
		DefaultFleetVersion4      = DefaultFleetVersion + "-4"
	)

	gen := &generatorMockFleet{
		id:            DefaultFleetID,
		assignmentIDs: []fl.AssignmentID{DefaultFleetAssignmentID1, DefaultFleetAssignmentID2, DefaultFleetAssignmentID3},
		eventIDs:      []fl.EventID{DefaultFleetEventID1, DefaultFleetEventID2, DefaultFleetEventID3},
		versions:      []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3, DefaultFleetVersion4},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
			Version:      string(DefaultFleetVersion),
		},
	)

	repo := &fleetRepositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	req := &changeNumberOfVehiclesRequestMock{
		FlightplanID:     string(DefaultFlightplanID),
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	var resFlightplanID string
	var resNumberOfVehicles int32
	ret := service.changeNumberOfVehiclesOperation(
		nil,
		req,
		func(id string, numberOfVehicles int32) {
			resFlightplanID = id
			resNumberOfVehicles = numberOfVehicles
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
			ID: string(DefaultFleetAssignmentID1),
		},
		{
			ID: string(DefaultFleetAssignmentID2),
		},
		{
			ID: string(DefaultFleetAssignmentID3),
		},
	}
	expectEvents := []eventComponentMock{
		{
			ID:           string(DefaultFleetEventID1),
			AssignmentID: string(DefaultFleetAssignmentID1),
		},
		{
			ID:           string(DefaultFleetEventID2),
			AssignmentID: string(DefaultFleetAssignmentID2),
		},
		{
			ID:           string(DefaultFleetEventID3),
			AssignmentID: string(DefaultFleetAssignmentID3),
		},
	}

	a.Nil(ret)
	a.Equal(resFlightplanID, string(DefaultFlightplanID))
	a.Equal(resNumberOfVehicles, DefaultFleetNumberOfVehicles)
	a.Equal(actualAssignments, expectAssignments)
	a.Equal(actualEvents, expectEvents)
}

func TestGetAssignmentsTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMockFleet{}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
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

	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	var resCall bool
	ret := service.GetAssignments(
		req,
		func(id, assignmentID, vehicleID, missionID string) {
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
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
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
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	req := &fleetIDRequestMock{
		FlightplanID: string(DefaultFlightplanID),
	}
	var resIDs []string
	var resAssignmentIDs []string
	var resVehicleIDs []string
	var resMissionIDs []string
	ret := service.getAssignmentsOperation(
		nil,
		req,
		func(id, assignmentID, vehicleID, missionID string) {
			resIDs = append(resIDs, id)
			resAssignmentIDs = append(resAssignmentIDs, assignmentID)
			resVehicleIDs = append(resVehicleIDs, vehicleID)
			resMissionIDs = append(resMissionIDs, missionID)
		},
	)

	expectIDs := []string{string(DefaultFleetEventID1), string(DefaultFleetEventID2), string(DefaultFleetEventID3)}
	expectAssignmentIDs := []string{string(DefaultFleetAssignmentID1), string(DefaultFleetAssignmentID2), string(DefaultFleetAssignmentID3)}
	expectVehicleIDs := []string{string(DefaultFleetVehicleID1), string(DefaultFleetVehicleID2), string(DefaultFleetVehicleID3)}
	expectMissionIDs := []string{string(DefaultFleetMissionID1), string(DefaultFleetMissionID2), string(DefaultFleetMissionID3)}

	a.Nil(ret)
	a.Equal(resIDs, expectIDs)
	a.Equal(resAssignmentIDs, expectAssignmentIDs)
	a.Equal(resVehicleIDs, expectVehicleIDs)
	a.Equal(resMissionIDs, expectMissionIDs)
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
		id:       DefaultFleetID,
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2, DefaultFleetVersion3, DefaultFleetVersion4},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
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

	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}

	req := &updateAssignmentRequestMock{
		FlightplanID: string(DefaultFlightplanID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    string(AfterFleetVehicleID),
		MissionID:    string(AfterFleetMissionID),
	}
	var resCall bool
	ret := service.UpdateAssignment(
		req,
		func(id, assignmentID, vehicleID, missionID string) {
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
		id:       DefaultFleetID,
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
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
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	req := &updateAssignmentRequestMock{
		FlightplanID: string(DefaultFlightplanID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    string(AfterFleetVehicleID),
		MissionID:    string(AfterFleetMissionID),
	}
	var resID string
	var resAssignmentID string
	var resVehicleID string
	var resMissionID string
	ret := service.updateAssignmentOperation(
		nil,
		req,
		func(id, assignmentID, vehicleID, missionID string) {
			resID = id
			resAssignmentID = assignmentID
			resVehicleID = vehicleID
			resMissionID = missionID
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
	a.Equal(resID, string(DefaultFleetEventID))
	a.Equal(resAssignmentID, string(DefaultFleetAssignmentID))
	a.Equal(resVehicleID, string(AfterFleetVehicleID))
	a.Equal(resMissionID, string(AfterFleetMissionID))
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
		id:       DefaultFleetID,
		versions: []fl.Version{DefaultFleetVersion1, DefaultFleetVersion2},
	}

	fleet := fl.AssembleFrom(
		gen,
		&fleetComponentMock{
			ID:           string(DefaultFleetID),
			FlightplanID: string(DefaultFlightplanID),
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
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  nil,
	}

	req := &updateAssignmentRequestMock{
		FlightplanID: string(DefaultFlightplanID),
		EventID:      string(DefaultFleetEventID),
		AssignmentID: string(DefaultFleetAssignmentID),
		VehicleID:    "",
		MissionID:    "",
	}
	var resID string
	var resAssignmentID string
	var resVehicleID string
	var resMissionID string
	ret := service.updateAssignmentOperation(
		nil,
		req,
		func(id, assignmentID, vehicleID, missionID string) {
			resID = id
			resAssignmentID = assignmentID
			resVehicleID = vehicleID
			resMissionID = missionID
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
	a.Equal(resID, string(DefaultFleetEventID))
	a.Equal(resAssignmentID, string(DefaultFleetAssignmentID))
	a.Empty(resVehicleID)
	a.Empty(resMissionID)
	a.Equal(actualAssignments, expectAssignments)
	a.Equal(actualEvents, expectEvents)
}
