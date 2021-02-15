package fleet

import (
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

type testGenerator struct {
	Generator
	id                ID
	assignmentIDs     []AssignmentID
	assignmentIDIndex int
	eventIDs          []EventID
	eventIDIndex      int
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

func TestCreateSingleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		eventIDs:      []EventID{DefaultEventID1},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetNumberOfVehicles(), 1)
	a.Len(fleet.GetVehicleAssignments(), 1)
	a.Equal(fleet.GetVehicleAssignments()[0], expectAssignment)
	a.Len(fleet.GetEventPlannings(), 0)
}

func TestCreateMultipleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
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
	a.Equal(fleet.GetNumberOfVehicles(), 3)
	a.Len(fleet.GetVehicleAssignments(), 3)
	a.Equal(fleet.GetVehicleAssignments()[0], expectAssignment1)
	a.Equal(fleet.GetVehicleAssignments()[1], expectAssignment2)
	a.Equal(fleet.GetVehicleAssignments()[2], expectAssignment3)
	a.Len(fleet.GetEventPlannings(), 0)
}
