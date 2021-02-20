package fleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fleetを一つ新しく作成し、初期状態を検証する。
// Vehicle数は1つ指定し、初期状態のAssignmentが1つ
// 作成されることを検証する。
func TestCreateSingleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
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

// Fleetを一つ新しく作成し、初期状態を検証する。
// Vehicle数は複数指定し、初期状態のAssignmentが同数
// 作成されることを検証する。
func TestCreateMultipleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
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

// Fleetを一つ新しく作成し、初期状態を検証する。
// Vehicle数は0を指定し、初期状態のAssignmentが1つも
// 作成されないことを検証する。
func TestCreateNonFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 0)

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Equal(fleet.GetNumberOfVehicles(), 0)
	a.Equal(fleet.GetAllAssignmentID(), *new([]AssignmentID))
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
	a.Len(fleet.vehicleAssignments, 0)
	a.Len(fleet.eventPlannings, 0)
}

// Fleetを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFleetAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id: DefaultID,
	}

	assignmentComps := []assignmentComponentMock{
		{
			id:        string(DefaultAssignmentID1),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID1),
		},
		{
			id:        string(DefaultAssignmentID2),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID2),
		},
		{
			id:        string(DefaultAssignmentID3),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID3),
		},
	}
	eventComps := []eventComponentMock{
		{
			id:           string(DefaultEventID1),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID1),
			missionID:    string(DefaultMissionID1),
		},
		{
			id:           string(DefaultEventID2),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID2),
			missionID:    string(DefaultMissionID2),
		},
		{
			id:           string(DefaultEventID3),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID3),
			missionID:    string(DefaultMissionID3),
		},
	}
	fleetComp := fleetComponentMock{
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
		assignments:  assignmentComps,
		events:       eventComps,
		version:      string(DefaultVersion1),
	}
	fleet := AssembleFrom(gen, &fleetComp)

	expectAssignment1 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    DefaultVehicleID1,
	}
	expectAssignment2 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID2,
		vehicleID:    DefaultVehicleID2,
	}
	expectAssignment3 := &VehicleAssignment{
		assignmentID: DefaultAssignmentID3,
		vehicleID:    DefaultVehicleID3,
	}

	expectEvent1 := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID1,
		missionID:    DefaultMissionID1,
	}
	expectEvent2 := &EventPlanning{
		eventID:      DefaultEventID2,
		assignmentID: DefaultAssignmentID2,
		missionID:    DefaultMissionID2,
	}
	expectEvent3 := &EventPlanning{
		eventID:      DefaultEventID3,
		assignmentID: DefaultAssignmentID3,
		missionID:    DefaultMissionID3,
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
	a.Len(fleet.eventPlannings, 3)
	a.Equal(fleet.eventPlannings[0], expectEvent1)
	a.Equal(fleet.eventPlannings[1], expectEvent2)
	a.Equal(fleet.eventPlannings[2], expectEvent3)
}

// Fleetを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFleet(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 3)
	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID1
	fleet.vehicleAssignments[1].vehicleID = DefaultVehicleID2
	fleet.vehicleAssignments[2].vehicleID = DefaultVehicleID3
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID1,
			missionID:    DefaultMissionID1,
		},
	)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID2,
			assignmentID: DefaultAssignmentID2,
			missionID:    DefaultMissionID2,
		},
	)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID3,
			assignmentID: DefaultAssignmentID3,
			missionID:    DefaultMissionID3,
		},
	)

	var fleetComp fleetComponentMock
	var assignmentComps []assignmentComponentMock
	var eventComps []eventComponentMock

	TakeApart(
		fleet,
		func(id, flightplanID, version string) {
			fleetComp.id = id
			fleetComp.flightplanID = flightplanID
			fleetComp.version = version
		},
		func(id, fleetID, vehicleID string) {
			assignmentComps = append(
				assignmentComps,
				assignmentComponentMock{
					id:        id,
					fleetID:   fleetID,
					vehicleID: vehicleID,
				},
			)
		},
		func(id, fleetID, assignmentID, missionID string) {
			eventComps = append(
				eventComps,
				eventComponentMock{
					id:           id,
					fleetID:      fleetID,
					assignmentID: assignmentID,
					missionID:    missionID,
				},
			)
		},
	)

	expectFleet := fleetComponentMock{
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
		version:      string(DefaultVersion1),
	}
	expectAssignments := []assignmentComponentMock{
		{
			id:        string(DefaultAssignmentID1),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID1),
		},
		{
			id:        string(DefaultAssignmentID2),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID2),
		},
		{
			id:        string(DefaultAssignmentID3),
			fleetID:   string(DefaultID),
			vehicleID: string(DefaultVehicleID3),
		},
	}
	expectEvents := []eventComponentMock{
		{
			id:           string(DefaultEventID1),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID1),
			missionID:    string(DefaultMissionID1),
		},
		{
			id:           string(DefaultEventID2),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID2),
			missionID:    string(DefaultMissionID2),
		},
		{
			id:           string(DefaultEventID3),
			fleetID:      string(DefaultID),
			assignmentID: string(DefaultAssignmentID3),
			missionID:    string(DefaultMissionID3),
		},
	}

	a.Equal(fleetComp, expectFleet)
	a.Equal(assignmentComps, expectAssignments)
	a.Equal(eventComps, expectEvents)
}
