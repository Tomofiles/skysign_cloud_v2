package communication

import (
	"errors"
	"time"
)

const DefaultID = ID("communication-id")
const DefaultCommandID = CommandID("command-id")
const DefaultMissionID = MissionID("mission-id")

var DefaultTime = time.Now()

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

// Communication用汎用ジェネレータモック
type generatorMock struct {
	Generator
	commandIDs     []CommandID
	commandIDIndex int
	times          []time.Time
	timeIndex      int
}

func (gen *generatorMock) NewCommandID() CommandID {
	commandID := gen.commandIDs[gen.commandIDIndex]
	gen.commandIDIndex++
	return commandID
}
func (gen *generatorMock) NewTime() time.Time {
	time := gen.times[gen.timeIndex]
	gen.timeIndex++
	return time
}

// Communication用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

// // Fleet用リポジトリモック
// type repositoryMock struct {
// 	mock.Mock
// 	fleet    *Fleet
// 	deleteID ID
// }

// func (r *repositoryMock) GetByID(
// 	tx txmanager.Tx,
// 	id ID,
// ) (*Fleet, error) {
// 	ret := r.Called(id)
// 	var f *Fleet
// 	if ret.Get(0) == nil {
// 		f = nil
// 	} else {
// 		f = ret.Get(0).(*Fleet)
// 	}
// 	return f, ret.Error(1)
// }

// func (r *repositoryMock) Save(
// 	tx txmanager.Tx,
// 	fleet *Fleet,
// ) error {
// 	ret := r.Called(fleet)
// 	r.fleet = fleet
// 	return ret.Error(0)
// }

// func (r *repositoryMock) Delete(
// 	tx txmanager.Tx,
// 	id ID,
// ) error {
// 	ret := r.Called(id)
// 	r.deleteID = id
// 	return ret.Error(0)
// }

// // Fleet構成オブジェクトモック
// type fleetComponentMock struct {
// 	id           string
// 	assignments  []assignmentComponentMock
// 	events       []eventComponentMock
// 	isCarbonCopy bool
// 	version      string
// }

// func (f *fleetComponentMock) GetID() string {
// 	return f.id
// }

// func (f *fleetComponentMock) GetIsCarbonCopy() bool {
// 	return f.isCarbonCopy
// }

// func (f *fleetComponentMock) GetVersion() string {
// 	return f.version
// }

// func (f *fleetComponentMock) GetAssignments() []AssignmentComponent {
// 	var assignments []AssignmentComponent
// 	for _, a := range f.assignments {
// 		assignments = append(
// 			assignments,
// 			&assignmentComponentMock{
// 				id:        a.id,
// 				fleetID:   a.fleetID,
// 				vehicleID: a.vehicleID,
// 			},
// 		)
// 	}
// 	return assignments
// }

// func (f *fleetComponentMock) GetEvents() []EventComponent {
// 	var events []EventComponent
// 	for _, e := range f.events {
// 		events = append(
// 			events,
// 			&eventComponentMock{
// 				id:           e.id,
// 				fleetID:      e.fleetID,
// 				assignmentID: e.assignmentID,
// 				missionID:    e.missionID,
// 			},
// 		)
// 	}
// 	return events
// }

// // Assignment構成オブジェクトモック
// type assignmentComponentMock struct {
// 	id        string
// 	fleetID   string
// 	vehicleID string
// }

// func (a *assignmentComponentMock) GetID() string {
// 	return a.id
// }

// func (a *assignmentComponentMock) GetFleetID() string {
// 	return a.fleetID
// }

// func (a *assignmentComponentMock) GetVehicleID() string {
// 	return a.vehicleID
// }

// // Event構成オブジェクトモック
// type eventComponentMock struct {
// 	id           string
// 	fleetID      string
// 	assignmentID string
// 	missionID    string
// }

// func (e *eventComponentMock) GetID() string {
// 	return e.id
// }

// func (e *eventComponentMock) GetFleetID() string {
// 	return e.fleetID
// }

// func (e *eventComponentMock) GetAssignmentID() string {
// 	return e.assignmentID
// }

// func (e *eventComponentMock) GetMissionID() string {
// 	return e.missionID
// }
