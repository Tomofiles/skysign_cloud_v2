package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
)

const DefaultID = ID("fleet-id")
const DefaultFlightplanID = flightplan.ID("flightplan-id")
const DefaultAssignmentID = AssignmentID("assignment-id")
const DefaultEventID = EventID("event-id")
const DefaultVehicleID = VehicleID("vehicle-id")
const DefaultMissionID = MissionID("mission-id")
const DefaultVersion = Version("version")

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

// Fleet用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id                ID
	assignmentIDs     []AssignmentID
	assignmentIDIndex int
	eventIDs          []EventID
	eventIDIndex      int
	versions          []Version
	versionIndex      int
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewAssignmentID() AssignmentID {
	assignmentID := gen.assignmentIDs[gen.assignmentIDIndex]
	gen.assignmentIDIndex++
	return assignmentID
}
func (gen *generatorMock) NewEventID() EventID {
	eventID := gen.eventIDs[gen.eventIDIndex]
	gen.eventIDIndex++
	return eventID
}
func (gen *generatorMock) NewVersion() Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

// Fleet構成オブジェクトモック
type fleetComponentMock struct {
	id           string
	flightplanID string
	assignments  []assignmentComponentMock
	events       []eventComponentMock
	isCarbonCopy bool
	version      string
}

func (f *fleetComponentMock) GetID() string {
	return f.id
}

func (f *fleetComponentMock) GetFlightplanID() string {
	return f.flightplanID
}

func (f *fleetComponentMock) GetIsCarbonCopy() bool {
	return f.isCarbonCopy
}

func (f *fleetComponentMock) GetVersion() string {
	return f.version
}

func (f *fleetComponentMock) GetAssignments() []AssignmentComponent {
	var assignments []AssignmentComponent
	for _, a := range f.assignments {
		assignments = append(
			assignments,
			&assignmentComponentMock{
				id:        a.id,
				fleetID:   a.fleetID,
				vehicleID: a.vehicleID,
			},
		)
	}
	return assignments
}

func (f *fleetComponentMock) GetEvents() []EventComponent {
	var events []EventComponent
	for _, e := range f.events {
		events = append(
			events,
			&eventComponentMock{
				id:           e.id,
				fleetID:      e.fleetID,
				assignmentID: e.assignmentID,
				missionID:    e.missionID,
			},
		)
	}
	return events
}

// Assignment構成オブジェクトモック
type assignmentComponentMock struct {
	id        string
	fleetID   string
	vehicleID string
}

func (a *assignmentComponentMock) GetID() string {
	return a.id
}

func (a *assignmentComponentMock) GetFleetID() string {
	return a.fleetID
}

func (a *assignmentComponentMock) GetVehicleID() string {
	return a.vehicleID
}

// Event構成オブジェクトモック
type eventComponentMock struct {
	id           string
	fleetID      string
	assignmentID string
	missionID    string
}

func (e *eventComponentMock) GetID() string {
	return e.id
}

func (e *eventComponentMock) GetFleetID() string {
	return e.fleetID
}

func (e *eventComponentMock) GetAssignmentID() string {
	return e.assignmentID
}

func (e *eventComponentMock) GetMissionID() string {
	return e.missionID
}
