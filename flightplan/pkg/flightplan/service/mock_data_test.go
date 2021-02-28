package service

import (
	"context"
	"flightplan/pkg/flightplan/domain/event"
	fl "flightplan/pkg/flightplan/domain/fleet"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightplanID = fpl.ID("flightplan-id")
const DefaultFlightplanVersion = fpl.Version("version")
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFleetID = fl.ID("fleet-id")
const DefaultFleetVersion = fl.Version("version")
const DefaultFleetNumberOfVehicles int32 = 3
const DefaultFleetAssignmentID = fl.AssignmentID("assignment-id")
const DefaultFleetEventID = fl.EventID("event-id")
const DefaultFleetVehicleID = fl.VehicleID("vehicle-id")
const DefaultFleetMissionID = fl.MissionID("mission-id")

type flightplanRepositoryMock struct {
	mock.Mock
}

func (r *flightplanRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*fpl.Flightplan, error) {
	ret := r.Called()
	var f []*fpl.Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*fpl.Flightplan)
	}
	return f, ret.Error(1)
}

func (r *flightplanRepositoryMock) GetAllOrigin(
	tx txmanager.Tx,
) ([]*fpl.Flightplan, error) {
	ret := r.Called()
	var f []*fpl.Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*fpl.Flightplan)
	}
	return f, ret.Error(1)
}

func (r *flightplanRepositoryMock) GetByID(
	tx txmanager.Tx,
	id fpl.ID,
) (*fpl.Flightplan, error) {
	ret := r.Called(id)
	var f *fpl.Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*fpl.Flightplan)
	}
	return f, ret.Error(1)
}

func (r *flightplanRepositoryMock) Save(
	tx txmanager.Tx,
	flightplan *fpl.Flightplan,
) error {
	ret := r.Called(flightplan)
	return ret.Error(0)
}

func (r *flightplanRepositoryMock) Delete(
	tx txmanager.Tx,
	id fpl.ID,
) error {
	ret := r.Called(id)
	return ret.Error(0)
}

type fleetRepositoryMock struct {
	mock.Mock
	fleet    *fl.Fleet
	deleteID fpl.ID
}

func (r *fleetRepositoryMock) GetByFlightplanID(
	tx txmanager.Tx,
	id fpl.ID,
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

func (r *fleetRepositoryMock) DeleteByFlightplanID(
	tx txmanager.Tx,
	id fpl.ID,
) error {
	ret := r.Called(id)
	r.deleteID = id
	return ret.Error(0)
}

type generatorMockFlightplan struct {
	fpl.Generator
	id           fpl.ID
	versions     []fpl.Version
	versionIndex int
}

func (gen *generatorMockFlightplan) NewID() fpl.ID {
	return gen.id
}
func (gen *generatorMockFlightplan) NewVersion() fpl.Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

type generatorMockFleet struct {
	fpl.Generator
	id                fl.ID
	assignmentIDs     []fl.AssignmentID
	assignmentIDIndex int
	eventIDs          []fl.EventID
	eventIDIndex      int
	versions          []fl.Version
	versionIndex      int
}

func (gen *generatorMockFleet) NewID() fl.ID {
	return gen.id
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

func (psm *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName string, handler event.Handler) error {
	return nil
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

type flightplanComponentMock struct {
	ID           string
	Name         string
	Description  string
	IsCarbonCopy bool
	Version      string
}

func (f *flightplanComponentMock) GetID() string {
	return f.ID
}

func (f *flightplanComponentMock) GetName() string {
	return f.Name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.Description
}

func (f *flightplanComponentMock) GetIsCarbonCopy() bool {
	return f.IsCarbonCopy
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.Version
}

type fleetComponentMock struct {
	ID           string
	FlightplanID string
	Assignments  []*assignmentComponentMock
	Events       []*eventComponentMock
	Version      string
}

func (f *fleetComponentMock) GetID() string {
	return f.ID
}

func (f *fleetComponentMock) GetFlightplanID() string {
	return f.FlightplanID
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

type flightplanRequestMock struct {
	ID          string
	Name        string
	Description string
}

func (f *flightplanRequestMock) GetID() string {
	return f.ID
}

func (f *flightplanRequestMock) GetName() string {
	return f.Name
}

func (f *flightplanRequestMock) GetDescription() string {
	return f.Description
}

type flightplanIDRequestMock struct {
	ID string
}

func (f *flightplanIDRequestMock) GetID() string {
	return f.ID
}

type fleetIDRequestMock struct {
	FlightplanID string
}

func (f *fleetIDRequestMock) GetFlightplanID() string {
	return f.FlightplanID
}

type carbonCopyRequestMock struct {
	OriginalID string
	NewID      string
}

func (f *carbonCopyRequestMock) GetOriginalID() string {
	return f.OriginalID
}

func (f *carbonCopyRequestMock) GetNewID() string {
	return f.NewID
}

type changeNumberOfVehiclesRequestMock struct {
	FlightplanID     string
	NumberOfVehicles int32
}

func (c *changeNumberOfVehiclesRequestMock) GetFlightplanID() string {
	return c.FlightplanID
}

func (c *changeNumberOfVehiclesRequestMock) GetNumberOfVehicles() int32 {
	return c.NumberOfVehicles
}

// UpdateAssignmentRequestDpo .
type updateAssignmentRequestMock struct {
	FlightplanID string
	EventID      string
	AssignmentID string
	VehicleID    string
	MissionID    string
}

func (u *updateAssignmentRequestMock) GetFlightplanID() string {
	return u.FlightplanID
}

func (u *updateAssignmentRequestMock) GetEventID() string {
	return u.EventID
}

func (u *updateAssignmentRequestMock) GetAssignmentID() string {
	return u.AssignmentID
}

func (u *updateAssignmentRequestMock) GetVehicleID() string {
	return u.VehicleID
}

func (u *updateAssignmentRequestMock) GetMissionID() string {
	return u.MissionID
}
