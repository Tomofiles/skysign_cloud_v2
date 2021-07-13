package fleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 機体数の変更によりFleetを再作成する。
func TestChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles = 3
		DefaultAssignmentID1    = DefaultAssignmentID + "-1"
		DefaultAssignmentID2    = DefaultAssignmentID + "-2"
		DefaultAssignmentID3    = DefaultAssignmentID + "-3"
		DefaultEventID1         = DefaultEventID + "-1"
		DefaultEventID2         = DefaultEventID + "-2"
		DefaultEventID3         = DefaultEventID + "-3"
		DefaultVersion1         = DefaultVersion + "-1"
		DefaultVersion2         = DefaultVersion + "-2"
		DefaultVersion3         = DefaultVersion + "-3"
		DefaultVersion4         = DefaultVersion + "-4"
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(fleet, nil)
	repo.On("Delete", DefaultID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultID, DefaultNumberOfVehicles)

	var actualAssignments []assignmentComponentMock
	var actualEvents []eventComponentMock
	repo.fleet.ProvideAssignmentsInterest(
		func(assignmentID string, vehicleID string) {
			actualAssignments = append(
				actualAssignments,
				assignmentComponentMock{
					id:        assignmentID,
					vehicleID: vehicleID,
				},
			)
		},
		func(eventID string, assignmentID string, missionID string) {
			actualEvents = append(
				actualEvents,
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
			id: string(DefaultAssignmentID1),
		},
		{
			id: string(DefaultAssignmentID2),
		},
		{
			id: string(DefaultAssignmentID3),
		},
	}
	expectEvents := []eventComponentMock{
		{
			id:           string(DefaultEventID1),
			assignmentID: string(DefaultAssignmentID1),
		},
		{
			id:           string(DefaultEventID2),
			assignmentID: string(DefaultAssignmentID2),
		},
		{
			id:           string(DefaultEventID3),
			assignmentID: string(DefaultAssignmentID3),
		},
	}

	a.Nil(ret)
	a.Equal(actualAssignments, expectAssignments)
	a.Equal(actualEvents, expectEvents)
}

// 機体数の変更によりFleetを再作成する。
// 指定されたIDのFleetの取得がエラーとなった場合、
// 再作成が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles = 3
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Delete", DefaultID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrGet)
}

// 機体数の変更によりFleetを再作成する。
// カーボンコピーされたFleetの再作成となった場合、
// 再作成が失敗し、エラーが返却されることを検証する。
func TestCannotChangeErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles = 3
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			isCarbonCopy: CarbonCopy,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(fleet, nil)
	repo.On("Delete", DefaultID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrCannotChange)
}

// 機体数の変更によりFleetを再作成する。
// Fleetの削除がエラーとなった場合、
// 再作成が失敗し、エラーが返却されることを検証する。
func TestDeleteErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles = 3
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(fleet, nil)
	repo.On("Delete", DefaultID).Return(ErrDelete)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrDelete)
}

// 機体数の変更によりFleetを再作成する。
// 再作成されたFleetの保存がエラーとなった場合、
// 再作成が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles = 3
		DefaultAssignmentID1    = DefaultAssignmentID + "-1"
		DefaultAssignmentID2    = DefaultAssignmentID + "-2"
		DefaultAssignmentID3    = DefaultAssignmentID + "-3"
		DefaultEventID1         = DefaultEventID + "-1"
		DefaultEventID2         = DefaultEventID + "-2"
		DefaultEventID3         = DefaultEventID + "-3"
		DefaultVersion1         = DefaultVersion + "-1"
		DefaultVersion2         = DefaultVersion + "-2"
		DefaultVersion3         = DefaultVersion + "-3"
		DefaultVersion4         = DefaultVersion + "-4"
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(fleet, nil)
	repo.On("Delete", DefaultID).Return(nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrSave)
}
