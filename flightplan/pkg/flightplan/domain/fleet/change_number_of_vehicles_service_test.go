package fleet

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles int32 = 3
		DefaultAssignmentID1          = DefaultAssignmentID + "-1"
		DefaultAssignmentID2          = DefaultAssignmentID + "-2"
		DefaultAssignmentID3          = DefaultAssignmentID + "-3"
		DefaultEventID1               = DefaultEventID + "-1"
		DefaultEventID2               = DefaultEventID + "-2"
		DefaultEventID3               = DefaultEventID + "-3"
		DefaultVersion1               = DefaultVersion + "-1"
		DefaultVersion2               = DefaultVersion + "-2"
		DefaultVersion3               = DefaultVersion + "-3"
		DefaultVersion4               = DefaultVersion + "-4"
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			flightplanID: string(DefaultFlightplanID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}

	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultFlightplanID, DefaultNumberOfVehicles)

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

func TestGetErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles int32 = 3
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(nil, ErrGet)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultFlightplanID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrGet)
}

func TestCannotChangeErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles int32 = 3
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			flightplanID: string(DefaultFlightplanID),
			isCarbonCopy: CarbonCopy,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultFlightplanID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrCannotChange)
}

func TestDeleteErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles int32 = 3
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			flightplanID: string(DefaultFlightplanID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{}

	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(ErrDelete)
	repo.On("Save", mock.Anything).Return(nil)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultFlightplanID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrDelete)
}

func TestSaveErrorWhenChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultNumberOfVehicles int32 = 3
		DefaultAssignmentID1          = DefaultAssignmentID + "-1"
		DefaultAssignmentID2          = DefaultAssignmentID + "-2"
		DefaultAssignmentID3          = DefaultAssignmentID + "-3"
		DefaultEventID1               = DefaultEventID + "-1"
		DefaultEventID2               = DefaultEventID + "-2"
		DefaultEventID3               = DefaultEventID + "-3"
		DefaultVersion1               = DefaultVersion + "-1"
		DefaultVersion2               = DefaultVersion + "-2"
		DefaultVersion3               = DefaultVersion + "-3"
		DefaultVersion4               = DefaultVersion + "-4"
	)

	fleet := AssembleFrom(
		nil,
		&fleetComponentMock{
			id:           string(DefaultID),
			flightplanID: string(DefaultFlightplanID),
			isCarbonCopy: Original,
			version:      string(DefaultVersion),
		},
	)

	gen := &generatorMock{
		id:            DefaultID,
		assignmentIDs: []AssignmentID{DefaultAssignmentID1, DefaultAssignmentID2, DefaultAssignmentID3},
		eventIDs:      []EventID{DefaultEventID1, DefaultEventID2, DefaultEventID3},
		versions:      []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}

	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultFlightplanID).Return(fleet, nil)
	repo.On("DeleteByFlightplanID", DefaultFlightplanID).Return(nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	ret := ChangeNumberOfVehicles(nil, gen, repo, DefaultFlightplanID, DefaultNumberOfVehicles)

	a.Equal(ret, ErrSave)
}

type repositoryMock struct {
	mock.Mock
	fleet    *Fleet
	deleteID flightplan.ID
}

func (r *repositoryMock) GetByFlightplanID(
	tx txmanager.Tx,
	id flightplan.ID,
) (*Fleet, error) {
	ret := r.Called(id)
	var f *Fleet
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*Fleet)
	}
	return f, ret.Error(1)
}

func (r *repositoryMock) Save(
	tx txmanager.Tx,
	fleet *Fleet,
) error {
	ret := r.Called(fleet)
	r.fleet = fleet
	return ret.Error(0)
}

func (r *repositoryMock) DeleteByFlightplanID(
	tx txmanager.Tx,
	id flightplan.ID,
) error {
	ret := r.Called(id)
	r.deleteID = id
	return ret.Error(0)
}
