package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
	"testing"

	"github.com/stretchr/testify/assert"
)

const DefaultID = ID("fleet-id")
const DefaultFlightplanID = flightplan.ID("flightplan-id")
const DefaultAssignmentID1 = AssignmentID("assignment-id-1")
const DefaultAssignmentID2 = AssignmentID("assignment-id-2")
const DefaultAssignmentID3 = AssignmentID("assignment-id-3")
const DefaultEventID1 = EventID("event-id-1")
const DefaultEventID2 = EventID("event-id-2")
const DefaultEventID3 = EventID("event-id-3")
const DefaultVehicleID = VehicleID("vehicle-id")
const DefaultMissionID = MissionID("mission-id")
const DefaultVersion1 = Version("version-1")
const DefaultVersion2 = Version("version-2")
const DefaultVersion3 = Version("version-3")

type testGenerator struct {
	Generator
	id                ID
	assignmentIDs     []AssignmentID
	assignmentIDIndex int
	eventIDs          []EventID
	eventIDIndex      int
	versions          []Version
	versionIndex      int
}

func (gen *testGenerator) NewID() ID {
	return gen.id
}
func (gen *testGenerator) NewAssignmentID() AssignmentID {
	assignmentID := gen.assignmentIDs[gen.assignmentIDIndex]
	gen.assignmentIDIndex++
	return assignmentID
}
func (gen *testGenerator) NewEventID() EventID {
	eventID := gen.eventIDs[gen.eventIDIndex]
	gen.eventIDIndex++
	return eventID
}
func (gen *testGenerator) NewVersion() Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

func TestCreateSingleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		eventIDs:      []EventID{DefaultEventID1},
		versions:      []Version{DefaultVersion1},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Equal(fleet.GetNumberOfVehicles(), 1)
	a.Equal(fleet.GetAllAssignmentID(), []AssignmentID{DefaultAssignmentID1})
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Len(fleet.eventPlannings, 0)
}

func TestCreateMultipleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion1},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 3)

	expectAssignment1 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}
	expectAssignment2 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID2,
		vehicleID:    "",
	}
	expectAssignment3 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID3,
		vehicleID:    "",
	}

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Equal(fleet.GetNumberOfVehicles(), 3)
	a.Equal(fleet.GetAllAssignmentID(), []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3})
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
	a.Len(fleet.vehicleAssignments, 3)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment1)
	a.Equal(fleet.vehicleAssignments[1], expectAssignment2)
	a.Equal(fleet.vehicleAssignments[2], expectAssignment3)
	a.Len(fleet.eventPlannings, 0)
}

func TestAssignVehicle(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	ret := fleet.AssignVehicle(DefaultAssignmentID1, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    DefaultVehicleID,
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestVehicleHasAlreadyAssignedWhenAssignVehicle(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 3)
	fleet.vehicleAssignments[2].vehicleID = DefaultVehicleID

	ret := fleet.AssignVehicle(DefaultAssignmentID1, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 3)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, errors.New("this vehicle has already assigned"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestNotFoundErrorWhenAssignVehicle(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	ret := fleet.AssignVehicle(DefaultAssignmentID2, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, errors.New("assignment not found"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestCancelVehiclesAssignment(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID

	ret := fleet.CancelVehiclesAssignment(DefaultAssignmentID1)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestNotFoundErrorWhenCancelVehiclesAssignment(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID

	ret := fleet.CancelVehiclesAssignment(DefaultAssignmentID2)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    DefaultVehicleID,
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, errors.New("assignment not found"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestAddNewEvent(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		eventIDs:      []EventID{DefaultEventID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	eventID, ret := fleet.AddNewEvent(DefaultAssignmentID1)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(eventID, DefaultEventID1)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestNotAssignedErrorWhenAddNewEvent(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		eventIDs:      []EventID{DefaultEventID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	eventID, ret := fleet.AddNewEvent(DefaultAssignmentID2)

	a.Len(fleet.eventPlannings, 0)
	a.Empty(eventID)
	a.Equal(ret, errors.New("this id not assigned"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestRemoveEvent(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    "",
		},
	)

	ret := fleet.RemoveEvent(DefaultEventID1)

	a.Len(fleet.eventPlannings, 0)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestNotFoundWhenRemoveEvent(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    "",
		},
	)

	ret := fleet.RemoveEvent(DefaultEventID2)

	a.Len(fleet.eventPlannings, 1)
	a.Equal(ret, errors.New("event not found"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestAssignMission(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    "",
		},
	)

	ret := fleet.AssignMission(DefaultEventID1, DefaultMissionID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    DefaultMissionID,
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestMissionHasAlreadyAssignedWhenAssignMission(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 2)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    "",
		},
	)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID2,
			assignmentID: DefaultAssignmentID2,
			missionID:    "",
		},
	)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID3,
			assignmentID: DefaultAssignmentID3,
			missionID:    DefaultMissionID,
		},
	)

	ret := fleet.AssignMission(DefaultEventID1, DefaultMissionID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 3)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(ret, errors.New("this mission has already assigned"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestNotFoundErrorWhenAssignMission(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    "",
		},
	)

	ret := fleet.AssignMission(DefaultEventID2, DefaultMissionID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(ret, errors.New("event not found"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

func TestCancelMission(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    DefaultMissionID,
		},
	)

	ret := fleet.CancelMission(DefaultEventID1)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

func TestNotFoundErrorWhenCancelMission(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    DefaultMissionID,
		},
	)

	ret := fleet.CancelMission(DefaultEventID2)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    DefaultMissionID,
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(ret, errors.New("event not found"))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}