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
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		eventIDs:      []EventID{DefaultEventID},
		versions:      []Version{DefaultVersion},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 1)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID,
		vehicleID:    "",
	}

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Equal(fleet.GetNumberOfVehicles(), 1)
	a.Equal(fleet.GetAllAssignmentID(), []AssignmentID{DefaultAssignmentID})
	a.Equal(fleet.isCarbonCopy, Original)
	a.Equal(fleet.GetVersion(), DefaultVersion)
	a.Equal(fleet.GetNewVersion(), DefaultVersion)
	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Len(fleet.eventPlannings, 0)
}

// Fleetを一つ新しく作成し、初期状態を検証する。
// Vehicle数は複数指定し、初期状態のAssignmentが同数
// 作成されることを検証する。
func TestCreateMultipleFleetNewFleet(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultEventID1      = DefaultEventID + "-1"
		DefaultEventID2      = DefaultEventID + "-2"
		DefaultEventID3      = DefaultEventID + "-3"
	)

	gen := &generatorMock{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion},
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
	a.Equal(fleet.isCarbonCopy, Original)
	a.Equal(fleet.GetVersion(), DefaultVersion)
	a.Equal(fleet.GetNewVersion(), DefaultVersion)
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
		versions: []Version{DefaultVersion},
	}
	fleet := NewInstance(gen, DefaultFlightplanID, 0)

	a.Equal(fleet.GetID(), DefaultID)
	a.Equal(fleet.GetFlightplanID(), DefaultFlightplanID)
	a.Equal(fleet.GetNumberOfVehicles(), 0)
	a.Equal(fleet.GetAllAssignmentID(), *new([]AssignmentID))
	a.Equal(fleet.isCarbonCopy, Original)
	a.Equal(fleet.GetVersion(), DefaultVersion)
	a.Equal(fleet.GetNewVersion(), DefaultVersion)
	a.Len(fleet.vehicleAssignments, 0)
	a.Len(fleet.eventPlannings, 0)
}

// Fleetのカーボンコピーを作成し、初期状態を検証する。
func TestCopyFleet(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedFlightplanID  = ID(DefaultFlightplanID) + "-copied"
		CopiedAssignmentID1 = DefaultAssignmentID + "-copied-1"
		CopiedAssignmentID2 = DefaultAssignmentID + "-copied-2"
		CopiedAssignmentID3 = DefaultAssignmentID + "-copied-3"
		CopiedEventID1      = DefaultEventID + "-copied-1"
		CopiedEventID2      = DefaultEventID + "-copied-2"
		CopiedEventID3      = DefaultEventID + "-copied-3"
		CopiedVehicleID1    = DefaultVehicleID + "-copied-1"
		// CopiedVehicleID2     = DefaultVehicleID + "-copied-2"
		CopiedVehicleID3 = DefaultVehicleID + "-copied-3"
		CopiedMissionID1 = DefaultMissionID + "-copied-1"
		// CopiedMissionID2     = DefaultMissionID + "-copied-2"
		CopiedMissionID3     = DefaultMissionID + "-copied-3"
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultEventID1      = DefaultEventID + "-1"
		DefaultEventID2      = DefaultEventID + "-2"
		DefaultEventID3      = DefaultEventID + "-3"
		DefaultVehicleID1    = DefaultVehicleID + "-1"
		// DefaultVehicleID2    = DefaultVehicleID + "-2"
		DefaultVehicleID3 = DefaultVehicleID + "-3"
		DefaultMissionID1 = DefaultMissionID + "-1"
		// DefaultMissionID2    = DefaultMissionID + "-2"
		DefaultMissionID3 = DefaultMissionID + "-3"
	)

	gen := &generatorMock{
		id:            CopiedFlightplanID,
		assignmentIDs: []AssignmentID{CopiedAssignmentID1, CopiedAssignmentID2, CopiedAssignmentID3},
		eventIDs:      []EventID{CopiedEventID1, CopiedEventID2, CopiedEventID3},
		vehicleIDs:    []VehicleID{CopiedVehicleID1, CopiedVehicleID3},
		missionIDs:    []MissionID{CopiedMissionID1, CopiedMissionID3},
	}
	pub := &publisherMock{}
	id := DefaultFlightplanID + "-copied"
	original := &Fleet{
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
		vehicleAssignments: []*VehicleAssignment{
			{
				assignmentID: DefaultAssignmentID1,
				vehicleID:    DefaultVehicleID1,
			},
			{
				assignmentID: DefaultAssignmentID2,
				vehicleID:    "",
			},
			{
				assignmentID: DefaultAssignmentID3,
				vehicleID:    DefaultVehicleID3,
			},
		},
		eventPlannings: []*EventPlanning{
			{
				eventID:      DefaultEventID1,
				assignmentID: DefaultAssignmentID1,
				missionID:    DefaultMissionID1,
			},
			{
				eventID:      DefaultEventID2,
				assignmentID: DefaultAssignmentID2,
				missionID:    "",
			},
			{
				eventID:      DefaultEventID3,
				assignmentID: DefaultAssignmentID3,
				missionID:    DefaultMissionID3,
			},
		},
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
	}
	fleet := Copy(gen, pub, id, original)

	expectAssignment1 := &VehicleAssignment{
		assignmentID: CopiedAssignmentID1,
		vehicleID:    CopiedVehicleID1,
	}
	expectAssignment2 := &VehicleAssignment{
		assignmentID: CopiedAssignmentID2,
		vehicleID:    "",
	}
	expectAssignment3 := &VehicleAssignment{
		assignmentID: CopiedAssignmentID3,
		vehicleID:    CopiedVehicleID3,
	}

	expectEventPlanning1 := &EventPlanning{
		eventID:      CopiedEventID1,
		assignmentID: CopiedAssignmentID1,
		missionID:    CopiedMissionID1,
	}
	expectEventPlanning2 := &EventPlanning{
		eventID:      CopiedEventID2,
		assignmentID: CopiedAssignmentID2,
		missionID:    "",
	}
	expectEventPlanning3 := &EventPlanning{
		eventID:      CopiedEventID3,
		assignmentID: CopiedAssignmentID3,
		missionID:    CopiedMissionID3,
	}

	expectEvent1 := VehicleCopiedWhenCopiedEvent{
		OriginalID: DefaultVehicleID1,
		NewID:      CopiedVehicleID1,
	}
	expectEvent2 := VehicleCopiedWhenCopiedEvent{
		OriginalID: DefaultVehicleID3,
		NewID:      CopiedVehicleID3,
	}
	expectEvent3 := MissionCopiedWhenCopiedEvent{
		OriginalID: DefaultMissionID1,
		NewID:      CopiedMissionID1,
	}
	expectEvent4 := MissionCopiedWhenCopiedEvent{
		OriginalID: DefaultMissionID3,
		NewID:      CopiedMissionID3,
	}

	a.Equal(fleet.GetID(), CopiedFlightplanID)
	a.Equal(fleet.GetFlightplanID(), id)
	a.Equal(fleet.GetNumberOfVehicles(), 3)
	a.Equal(fleet.GetAllAssignmentID(), []AssignmentID{CopiedAssignmentID1, CopiedAssignmentID2, CopiedAssignmentID3})
	a.Equal(fleet.isCarbonCopy, CarbonCopy)
	a.Equal(fleet.GetVersion(), DefaultVersion)
	a.Equal(fleet.GetNewVersion(), DefaultVersion)
	a.Len(fleet.vehicleAssignments, 3)
	a.Equal(fleet.vehicleAssignments, []*VehicleAssignment{expectAssignment1, expectAssignment2, expectAssignment3})
	a.Len(fleet.eventPlannings, 3)
	a.Equal(fleet.eventPlannings, []*EventPlanning{expectEventPlanning1, expectEventPlanning2, expectEventPlanning3})
	a.Len(pub.events, 4)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2, expectEvent3, expectEvent4})
}

// Fleetを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFleetAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultEventID1      = DefaultEventID + "-1"
		DefaultEventID2      = DefaultEventID + "-2"
		DefaultEventID3      = DefaultEventID + "-3"
		DefaultVehicleID1    = DefaultVehicleID + "-1"
		DefaultVehicleID2    = DefaultVehicleID + "-2"
		DefaultVehicleID3    = DefaultVehicleID + "-3"
		DefaultMissionID1    = DefaultMissionID + "-1"
		DefaultMissionID2    = DefaultMissionID + "-2"
		DefaultMissionID3    = DefaultMissionID + "-3"
	)

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
		isCarbonCopy: CarbonCopy,
		assignments:  assignmentComps,
		events:       eventComps,
		version:      string(DefaultVersion),
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
	a.Equal(fleet.isCarbonCopy, CarbonCopy)
	a.Equal(fleet.GetVersion(), DefaultVersion)
	a.Equal(fleet.GetNewVersion(), DefaultVersion)
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

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultEventID1      = DefaultEventID + "-1"
		DefaultEventID2      = DefaultEventID + "-2"
		DefaultEventID3      = DefaultEventID + "-3"
		DefaultVehicleID1    = DefaultVehicleID + "-1"
		DefaultVehicleID2    = DefaultVehicleID + "-2"
		DefaultVehicleID3    = DefaultVehicleID + "-3"
		DefaultMissionID1    = DefaultMissionID + "-1"
		DefaultMissionID2    = DefaultMissionID + "-2"
		DefaultMissionID3    = DefaultMissionID + "-3"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
	)

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
	fleet.isCarbonCopy = CarbonCopy

	var fleetComp fleetComponentMock
	var assignmentComps []assignmentComponentMock
	var eventComps []eventComponentMock

	TakeApart(
		fleet,
		func(id, flightplanID, version string, isCarbonCopy bool) {
			fleetComp.id = id
			fleetComp.flightplanID = flightplanID
			fleetComp.version = version
			fleetComp.isCarbonCopy = isCarbonCopy
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
		isCarbonCopy: CarbonCopy,
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
