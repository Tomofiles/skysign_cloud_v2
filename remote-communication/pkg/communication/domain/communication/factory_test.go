package communication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Communicationを一つ新しく作成し、初期状態を検証する。
func TestCreateNewCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	expectTelemetry := NewTelemetry()
	var expectCommands []*Command
	var expectUploadMissions []*UploadMission

	a.Equal(communication.GetID(), DefaultID)
	a.Equal(communication.telemetry, expectTelemetry)
	a.Equal(communication.commands, expectCommands)
	a.Equal(communication.uploadMissions, expectUploadMissions)
	a.Equal(communication.gen, gen)
}

// // Fleetを構成オブジェクトから組み立て直し、
// // 内部状態を検証する。
// func TestFleetAssembleFromComponent(t *testing.T) {
// 	a := assert.New(t)

// 	var (
// 		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
// 		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
// 		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
// 		DefaultEventID1      = DefaultEventID + "-1"
// 		DefaultEventID2      = DefaultEventID + "-2"
// 		DefaultEventID3      = DefaultEventID + "-3"
// 		DefaultVehicleID1    = DefaultVehicleID + "-1"
// 		DefaultVehicleID2    = DefaultVehicleID + "-2"
// 		DefaultVehicleID3    = DefaultVehicleID + "-3"
// 		DefaultMissionID1    = DefaultMissionID + "-1"
// 		DefaultMissionID2    = DefaultMissionID + "-2"
// 		DefaultMissionID3    = DefaultMissionID + "-3"
// 	)

// 	gen := &generatorMock{}

// 	assignmentComps := []assignmentComponentMock{
// 		{
// 			id:        string(DefaultAssignmentID1),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID1),
// 		},
// 		{
// 			id:        string(DefaultAssignmentID2),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID2),
// 		},
// 		{
// 			id:        string(DefaultAssignmentID3),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID3),
// 		},
// 	}
// 	eventComps := []eventComponentMock{
// 		{
// 			id:           string(DefaultEventID1),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID1),
// 			missionID:    string(DefaultMissionID1),
// 		},
// 		{
// 			id:           string(DefaultEventID2),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID2),
// 			missionID:    string(DefaultMissionID2),
// 		},
// 		{
// 			id:           string(DefaultEventID3),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID3),
// 			missionID:    string(DefaultMissionID3),
// 		},
// 	}
// 	fleetComp := fleetComponentMock{
// 		id:           string(DefaultID),
// 		isCarbonCopy: CarbonCopy,
// 		assignments:  assignmentComps,
// 		events:       eventComps,
// 		version:      string(DefaultVersion),
// 	}
// 	fleet := AssembleFrom(gen, &fleetComp)

// 	expectAssignment1 := &VehicleAssignment{
// 		assignmentID: DefaultAssignmentID1,
// 		vehicleID:    DefaultVehicleID1,
// 	}
// 	expectAssignment2 := &VehicleAssignment{
// 		assignmentID: DefaultAssignmentID2,
// 		vehicleID:    DefaultVehicleID2,
// 	}
// 	expectAssignment3 := &VehicleAssignment{
// 		assignmentID: DefaultAssignmentID3,
// 		vehicleID:    DefaultVehicleID3,
// 	}

// 	expectEvent1 := &EventPlanning{
// 		eventID:      DefaultEventID1,
// 		assignmentID: DefaultAssignmentID1,
// 		missionID:    DefaultMissionID1,
// 	}
// 	expectEvent2 := &EventPlanning{
// 		eventID:      DefaultEventID2,
// 		assignmentID: DefaultAssignmentID2,
// 		missionID:    DefaultMissionID2,
// 	}
// 	expectEvent3 := &EventPlanning{
// 		eventID:      DefaultEventID3,
// 		assignmentID: DefaultAssignmentID3,
// 		missionID:    DefaultMissionID3,
// 	}

// 	a.Equal(fleet.GetID(), DefaultID)
// 	a.Equal(fleet.GetNumberOfVehicles(), 3)
// 	a.Equal(fleet.GetAllAssignmentID(), []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3})
// 	a.Equal(fleet.isCarbonCopy, CarbonCopy)
// 	a.Equal(fleet.GetVersion(), DefaultVersion)
// 	a.Equal(fleet.GetNewVersion(), DefaultVersion)
// 	a.Len(fleet.vehicleAssignments, 3)
// 	a.Equal(fleet.vehicleAssignments[0], expectAssignment1)
// 	a.Equal(fleet.vehicleAssignments[1], expectAssignment2)
// 	a.Equal(fleet.vehicleAssignments[2], expectAssignment3)
// 	a.Len(fleet.eventPlannings, 3)
// 	a.Equal(fleet.eventPlannings[0], expectEvent1)
// 	a.Equal(fleet.eventPlannings[1], expectEvent2)
// 	a.Equal(fleet.eventPlannings[2], expectEvent3)
// }

// // Fleetを構成オブジェクトに分解し、
// // 内部状態を検証する。
// func TestTakeApartFleet(t *testing.T) {
// 	a := assert.New(t)

// 	var (
// 		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
// 		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
// 		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
// 		DefaultEventID1      = DefaultEventID + "-1"
// 		DefaultEventID2      = DefaultEventID + "-2"
// 		DefaultEventID3      = DefaultEventID + "-3"
// 		DefaultVehicleID1    = DefaultVehicleID + "-1"
// 		DefaultVehicleID2    = DefaultVehicleID + "-2"
// 		DefaultVehicleID3    = DefaultVehicleID + "-3"
// 		DefaultMissionID1    = DefaultMissionID + "-1"
// 		DefaultMissionID2    = DefaultMissionID + "-2"
// 		DefaultMissionID3    = DefaultMissionID + "-3"
// 		DefaultVersion1      = DefaultVersion + "-1"
// 		DefaultVersion2      = DefaultVersion + "-2"
// 	)

// 	gen := &generatorMock{
// 		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
// 		versions:      []Version{DefaultVersion1, DefaultVersion2},
// 	}
// 	fleet := NewInstance(gen, DefaultID, 3)
// 	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID1
// 	fleet.vehicleAssignments[1].vehicleID = DefaultVehicleID2
// 	fleet.vehicleAssignments[2].vehicleID = DefaultVehicleID3
// 	fleet.eventPlannings = append(
// 		fleet.eventPlannings,
// 		&EventPlanning{
// 			eventID:      DefaultEventID1,
// 			assignmentID: DefaultAssignmentID1,
// 			missionID:    DefaultMissionID1,
// 		},
// 	)
// 	fleet.eventPlannings = append(
// 		fleet.eventPlannings,
// 		&EventPlanning{
// 			eventID:      DefaultEventID2,
// 			assignmentID: DefaultAssignmentID2,
// 			missionID:    DefaultMissionID2,
// 		},
// 	)
// 	fleet.eventPlannings = append(
// 		fleet.eventPlannings,
// 		&EventPlanning{
// 			eventID:      DefaultEventID3,
// 			assignmentID: DefaultAssignmentID3,
// 			missionID:    DefaultMissionID3,
// 		},
// 	)
// 	fleet.isCarbonCopy = CarbonCopy

// 	var fleetComp fleetComponentMock
// 	var assignmentComps []assignmentComponentMock
// 	var eventComps []eventComponentMock

// 	TakeApart(
// 		fleet,
// 		func(id, version string, isCarbonCopy bool) {
// 			fleetComp.id = id
// 			fleetComp.version = version
// 			fleetComp.isCarbonCopy = isCarbonCopy
// 		},
// 		func(id, fleetID, vehicleID string) {
// 			assignmentComps = append(
// 				assignmentComps,
// 				assignmentComponentMock{
// 					id:        id,
// 					fleetID:   fleetID,
// 					vehicleID: vehicleID,
// 				},
// 			)
// 		},
// 		func(id, fleetID, assignmentID, missionID string) {
// 			eventComps = append(
// 				eventComps,
// 				eventComponentMock{
// 					id:           id,
// 					fleetID:      fleetID,
// 					assignmentID: assignmentID,
// 					missionID:    missionID,
// 				},
// 			)
// 		},
// 	)

// 	expectFleet := fleetComponentMock{
// 		id:           string(DefaultID),
// 		version:      string(DefaultVersion1),
// 		isCarbonCopy: CarbonCopy,
// 	}
// 	expectAssignments := []assignmentComponentMock{
// 		{
// 			id:        string(DefaultAssignmentID1),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID1),
// 		},
// 		{
// 			id:        string(DefaultAssignmentID2),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID2),
// 		},
// 		{
// 			id:        string(DefaultAssignmentID3),
// 			fleetID:   string(DefaultID),
// 			vehicleID: string(DefaultVehicleID3),
// 		},
// 	}
// 	expectEvents := []eventComponentMock{
// 		{
// 			id:           string(DefaultEventID1),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID1),
// 			missionID:    string(DefaultMissionID1),
// 		},
// 		{
// 			id:           string(DefaultEventID2),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID2),
// 			missionID:    string(DefaultMissionID2),
// 		},
// 		{
// 			id:           string(DefaultEventID3),
// 			fleetID:      string(DefaultID),
// 			assignmentID: string(DefaultAssignmentID3),
// 			missionID:    string(DefaultMissionID3),
// 		},
// 	}

// 	a.Equal(fleetComp, expectFleet)
// 	a.Equal(assignmentComps, expectAssignments)
// 	a.Equal(eventComps, expectEvents)
// }
