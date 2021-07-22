package fleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FleetのAssignementにVehicleを割り当てる。
// 割り当てられた後の内部状態と、バージョンが更新されることを検証する。
func TestAssignVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)

	ret := fleet.AssignVehicle(DefaultAssignmentID, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID,
		vehicleID:    DefaultVehicleID,
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのAssignementにVehicleを割り当てる。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenAssignVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	ret := fleet.AssignVehicle(DefaultAssignmentID, DefaultVehicleID)

	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementにVehicleを割り当てる。
// 割り当ての際、同Vehicleが別のAssignmentに割り当てられている場合、
// エラーとなり、割り当てが失敗となることを検証する。
func TestVehicleHasAlreadyAssignedWhenAssignVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 3)
	fleet.vehicleAssignments[2].vehicleID = DefaultVehicleID

	ret := fleet.AssignVehicle(DefaultAssignmentID1, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 3)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, ErrVehicleHasAlreadyAssigned)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementにVehicleを割り当てる。
// 割り当ての際、指定されたAssingmentが存在しない場合
// エラーとなり、割り当てが失敗となることを検証する。
func TestNotFoundErrorWhenAssignVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)

	ret := fleet.AssignVehicle(DefaultAssignmentID2, DefaultVehicleID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, ErrAssignmentNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのVehicle割り当てをキャンセルする。
// キャンセル後の内部状態と、バージョンが更新されることを検証する。
func TestCancelVehiclesAssignment(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID

	ret := fleet.CancelVehiclesAssignment(DefaultAssignmentID)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID,
		vehicleID:    "",
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのAssignementのVehicle割り当てをキャンセルする。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenCancelVehiclesAssignment(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	ret := fleet.CancelVehiclesAssignment(DefaultAssignmentID)

	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのVehicle割り当てをキャンセルする。
// キャンセルの際、指定されたAssingmentが存在しない場合
// エラーとなり、割り当てのキャンセルが失敗となることを検証する。
func TestNotFoundErrorWhenCancelVehiclesAssignment(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.vehicleAssignments[0].vehicleID = DefaultVehicleID

	ret := fleet.CancelVehiclesAssignment(DefaultAssignmentID2)

	expectAssignment := &VehicleAssignment{
		assignmentID: DefaultAssignmentID1,
		vehicleID:    DefaultVehicleID,
	}

	a.Len(fleet.vehicleAssignments, 1)
	a.Equal(fleet.vehicleAssignments[0], expectAssignment)
	a.Equal(ret, ErrAssignmentNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのEventを作成して追加する。
// Event作成後の内部状態と、バージョンが更新されることを検証する。
func TestAddNewEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		eventIDs:      []EventID{DefaultEventID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)

	eventID, ret := fleet.AddNewEvent(DefaultAssignmentID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID,
		assignmentID: DefaultAssignmentID,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(eventID, DefaultEventID)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのAssignementのEventを作成して追加する。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenAddNewEvent(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	eventID, ret := fleet.AddNewEvent(DefaultAssignmentID)

	a.Empty(eventID)
	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのEventを作成して追加する。
// 指定されたAssignmentが存在しない場合、エラーが発生し、
// 作成・追加が失敗することを検証する。
func TestNotAssignedErrorWhenAddNewEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1},
		eventIDs:      []EventID{DefaultEventID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)

	eventID, ret := fleet.AddNewEvent(DefaultAssignmentID2)

	a.Len(fleet.eventPlannings, 0)
	a.Empty(eventID)
	a.Equal(ret, ErrAssignmentNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのEventを削除する。
// Event削除後の内部状態と、バージョンが更新されることを検証する。
func TestRemoveEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID,
			assignmentID: DefaultAssignmentID,
			missionID:    "",
		},
	)

	ret := fleet.RemoveEvent(DefaultEventID)

	a.Len(fleet.eventPlannings, 0)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのAssignementのEventを削除する。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenRemoveEvent(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	ret := fleet.RemoveEvent(DefaultEventID)

	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのAssignementのEventを削除する。
// 指定されたAssignmentが存在しない場合、エラーが発生し、
// 削除が失敗することを検証する。
func TestNotFoundWhenRemoveEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultEventID1 = DefaultEventID + "-1"
		DefaultEventID2 = DefaultEventID + "-2"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID,
			missionID:    "",
		},
	)

	ret := fleet.RemoveEvent(DefaultEventID2)

	a.Len(fleet.eventPlannings, 1)
	a.Equal(ret, ErrEventNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのEventにMissionを割り当てる。
// 割り当てられた後の内部状態と、バージョンが更新されることを検証する。
func TestAssignMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID,
			assignmentID: DefaultAssignmentID,
			missionID:    "",
		},
	)

	ret := fleet.AssignMission(DefaultEventID, DefaultMissionID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID,
		assignmentID: DefaultAssignmentID,
		missionID:    DefaultMissionID,
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのEventにMissionを割り当てる。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenAssignMission(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	ret := fleet.AssignMission(DefaultEventID, DefaultMissionID)

	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのEventにMissionを割り当てる。
// 割り当ての際、同Missionが別のEventに割り当てられている場合、
// エラーとなり、割り当てが失敗となることを検証する。
func TestMissionHasAlreadyAssignedWhenAssignMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultAssignmentID1 = DefaultAssignmentID + "-1"
		DefaultAssignmentID2 = DefaultAssignmentID + "-2"
		DefaultAssignmentID3 = DefaultAssignmentID + "-3"
		DefaultEventID1      = DefaultEventID + "-1"
		DefaultEventID2      = DefaultEventID + "-2"
		DefaultEventID3      = DefaultEventID + "-3"
		DefaultVersion1      = DefaultVersion + "-1"
		DefaultVersion2      = DefaultVersion + "-2"
		DefaultVersion3      = DefaultVersion + "-3"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	fleet := NewInstance(gen, DefaultID, 3)
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
	a.Equal(ret, ErrMissionHasAlreadyAssigned)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのEventにMissionを割り当てる。
// 割り当ての際、指定されたEventが存在しない場合
// エラーとなり、割り当てが失敗となることを検証する。
func TestNotFoundErrorWhenAssignMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultEventID1 = DefaultEventID + "-1"
		DefaultEventID2 = DefaultEventID + "-2"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID,
			missionID:    "",
		},
	)

	ret := fleet.AssignMission(DefaultEventID2, DefaultMissionID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(ret, ErrEventNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのEventにMission割り当てをキャンセルする。
// キャンセル後の内部状態と、バージョンが更新されることを検証する。
func TestCancelMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID,
			assignmentID: DefaultAssignmentID,
			missionID:    DefaultMissionID,
		},
	)

	ret := fleet.CancelMission(DefaultEventID)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID,
		assignmentID: DefaultAssignmentID,
		missionID:    "",
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Nil(ret)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion2)
}

// FleetのEventにMission割り当てをキャンセルする。
// カーボンコピーされたFleetの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenCancelMission(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		versions: []Version{DefaultVersion2},
	}
	original := &Fleet{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	fleet := Copy(gen, nil, CopiedID, original)

	ret := fleet.CancelMission(DefaultEventID)

	a.Equal(ret, ErrCannotChange)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FleetのEventにMission割り当てをキャンセルする。
// キャンセルの際、指定されたEventが存在しない場合
// エラーとなり、割り当てのキャンセルが失敗となることを検証する。
func TestNotFoundErrorWhenCancelMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultEventID1 = DefaultEventID + "-1"
		DefaultEventID2 = DefaultEventID + "-2"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 1)
	fleet.eventPlannings = append(
		fleet.eventPlannings,
		&EventPlanning{
			eventID:      DefaultEventID1,
			assignmentID: DefaultAssignmentID,
			missionID:    DefaultMissionID,
		},
	)

	ret := fleet.CancelMission(DefaultEventID2)

	expectEvent := &EventPlanning{
		eventID:      DefaultEventID1,
		assignmentID: DefaultAssignmentID,
		missionID:    DefaultMissionID,
	}

	a.Len(fleet.eventPlannings, 1)
	a.Equal(fleet.eventPlannings[0], expectEvent)
	a.Equal(ret, ErrEventNotFound)
	a.Equal(fleet.GetVersion(), DefaultVersion1)
	a.Equal(fleet.GetNewVersion(), DefaultVersion1)
}

// FlightplanのAssignの内部状態をキャプチャする。
// ダブルディスパッチで公開された内部状態を検証する。
func TestProvideAssignmentsInterest(t *testing.T) {
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
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2},
	}
	fleet := NewInstance(gen, DefaultID, 3)
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

	var assignments []assignmentComponentMock
	var events []eventComponentMock
	fleet.ProvideAssignmentsInterest(
		func(assignmentID, vehicleID string) {
			assignments = append(
				assignments,
				assignmentComponentMock{
					id:        assignmentID,
					vehicleID: vehicleID,
				},
			)
		},
		func(eventID, assignmentID, missionID string) {
			events = append(
				events,
				eventComponentMock{
					id:           eventID,
					assignmentID: assignmentID,
					missionID:    missionID,
				},
			)
		},
	)

	expectAssignments := []assignmentComponentMock{
		{
			id:        string(DefaultAssignmentID1),
			vehicleID: string(DefaultVehicleID1),
		},
		{
			id:        string(DefaultAssignmentID2),
			vehicleID: string(DefaultVehicleID2),
		},
		{
			id:        string(DefaultAssignmentID3),
			vehicleID: string(DefaultVehicleID3),
		},
	}
	expectEvents := []eventComponentMock{
		{
			id:           string(DefaultEventID1),
			assignmentID: string(DefaultAssignmentID1),
			missionID:    string(DefaultMissionID1),
		},
		{
			id:           string(DefaultEventID2),
			assignmentID: string(DefaultAssignmentID2),
			missionID:    string(DefaultMissionID2),
		},
		{
			id:           string(DefaultEventID3),
			assignmentID: string(DefaultAssignmentID3),
			missionID:    string(DefaultMissionID3),
		},
	}

	a.Equal(assignments, expectAssignments)
	a.Equal(events, expectEvents)
}
