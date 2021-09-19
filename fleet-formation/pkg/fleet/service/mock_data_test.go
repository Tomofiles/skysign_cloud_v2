package service

import (
	fl "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultFleetID = fl.ID("fleet-id")
const DefaultFleetVersion = fl.Version("version")
const DefaultFleetNumberOfVehicles = 3
const DefaultFleetAssignmentID = fl.AssignmentID("assignment-id")
const DefaultFleetEventID = fl.EventID("event-id")
const DefaultFleetVehicleID = fl.VehicleID("vehicle-id")
const DefaultFleetMissionID = fl.MissionID("mission-id")

type fleetRepositoryMock struct {
	mock.Mock
	fleet    *fl.Fleet
	deleteID fl.ID
}

func (r *fleetRepositoryMock) GetByID(
	tx txmanager.Tx,
	id fl.ID,
) (*fl.Fleet, error) {
	ret := r.Called(id)
	var f *fl.Fleet
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*fl.Fleet)
	}
	return f, ret.Error(1)
}

func (r *fleetRepositoryMock) Save(
	tx txmanager.Tx,
	fleet *fl.Fleet,
) error {
	ret := r.Called(fleet)
	r.fleet = fleet
	return ret.Error(0)
}

func (r *fleetRepositoryMock) Delete(
	tx txmanager.Tx,
	id fl.ID,
) error {
	ret := r.Called(id)
	r.deleteID = id
	return ret.Error(0)
}

type generatorMockFleet struct {
	fl.Generator
	assignmentIDs     []fl.AssignmentID
	assignmentIDIndex int
	eventIDs          []fl.EventID
	eventIDIndex      int
	vehicleIDs        []fl.VehicleID
	vehicleIDIndex    int
	missionIDs        []fl.MissionID
	missionIDIndex    int
	versions          []fl.Version
	versionIndex      int
}

func (gen *generatorMockFleet) NewAssignmentID() fl.AssignmentID {
	assignmentID := gen.assignmentIDs[gen.assignmentIDIndex]
	gen.assignmentIDIndex++
	return assignmentID
}
func (gen *generatorMockFleet) NewEventID() fl.EventID {
	eventID := gen.eventIDs[gen.eventIDIndex]
	gen.eventIDIndex++
	return eventID
}
func (gen *generatorMockFleet) NewVehicleID() fl.VehicleID {
	vehicleID := gen.vehicleIDs[gen.vehicleIDIndex]
	gen.vehicleIDIndex++
	return vehicleID
}
func (gen *generatorMockFleet) NewMissionID() fl.MissionID {
	missionID := gen.missionIDs[gen.missionIDIndex]
	gen.missionIDIndex++
	return missionID
}
func (gen *generatorMockFleet) NewVersion() fl.Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

type publisherMock struct {
	events  []interface{}
	isFlush bool
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	rm.isFlush = true
	return nil
}

type pubSubManagerMock struct {
	mock.Mock
}

func (psm *pubSubManagerMock) GetPublisher() (event.Publisher, event.ChannelClose, error) {
	ret := psm.Called()
	var pub event.Publisher
	var close func() error
	if ret.Get(0) == nil {
		pub = nil
	} else {
		pub = ret.Get(0).(event.Publisher)
		close = ret.Get(1).(func() error)
	}
	return pub, close, ret.Error(2)
}

type txManagerMock struct {
	isOpe, isEH error
}

func (txm *txManagerMock) Do(operation func(txmanager.Tx) error) error {
	txm.isOpe = operation(nil)
	return nil
}

func (txm *txManagerMock) DoAndEndHook(operation func(txmanager.Tx) error, endHook func() error) error {
	txm.isOpe = operation(nil)
	txm.isEH = endHook()
	return nil
}

type fleetComponentMock struct {
	ID           string
	Assignments  []*assignmentComponentMock
	Events       []*eventComponentMock
	IsCarbonCopy bool
	Version      string
}

func (f *fleetComponentMock) GetID() string {
	return f.ID
}

func (f *fleetComponentMock) GetIsCarbonCopy() bool {
	return f.IsCarbonCopy
}

func (f *fleetComponentMock) GetVersion() string {
	return f.Version
}

func (f *fleetComponentMock) GetAssignments() []fl.AssignmentComponent {
	var assignments []fl.AssignmentComponent
	for _, a := range f.Assignments {
		assignments = append(assignments, a)
	}
	return assignments
}

func (f *fleetComponentMock) GetEvents() []fl.EventComponent {
	var events []fl.EventComponent
	for _, e := range f.Events {
		events = append(events, e)
	}
	return events
}

type assignmentComponentMock struct {
	ID        string
	FleetID   string
	VehicleID string
}

func (a *assignmentComponentMock) GetID() string {
	return a.ID
}

func (a *assignmentComponentMock) GetFleetID() string {
	return a.FleetID
}

func (a *assignmentComponentMock) GetVehicleID() string {
	return a.VehicleID
}

type eventComponentMock struct {
	ID           string
	FleetID      string
	AssignmentID string
	MissionID    string
}

func (e *eventComponentMock) GetID() string {
	return e.ID
}

func (e *eventComponentMock) GetFleetID() string {
	return e.FleetID
}

func (e *eventComponentMock) GetAssignmentID() string {
	return e.AssignmentID
}

func (e *eventComponentMock) GetMissionID() string {
	return e.MissionID
}

type fleetIDCommandMock struct {
	FleetID string
}

func (f *fleetIDCommandMock) GetID() string {
	return f.FleetID
}

type carbonCopyCommandMock struct {
	OriginalID string
	NewID      string
}

func (f *carbonCopyCommandMock) GetOriginalID() string {
	return f.OriginalID
}

func (f *carbonCopyCommandMock) GetNewID() string {
	return f.NewID
}

type changeNumberOfVehiclesCommandFleetMock struct {
	FleetID          string
	NumberOfVehicles int
}

func (c *changeNumberOfVehiclesCommandFleetMock) GetID() string {
	return c.FleetID
}

func (c *changeNumberOfVehiclesCommandFleetMock) GetNumberOfVehicles() int {
	return c.NumberOfVehicles
}

type updateAssignmentCommandMock struct {
	ID           string
	EventID      string
	AssignmentID string
	VehicleID    string
	MissionID    string
}

func (u *updateAssignmentCommandMock) GetID() string {
	return u.ID
}

func (u *updateAssignmentCommandMock) GetEventID() string {
	return u.EventID
}

func (u *updateAssignmentCommandMock) GetAssignmentID() string {
	return u.AssignmentID
}

func (u *updateAssignmentCommandMock) GetVehicleID() string {
	return u.VehicleID
}

func (u *updateAssignmentCommandMock) GetMissionID() string {
	return u.MissionID
}
