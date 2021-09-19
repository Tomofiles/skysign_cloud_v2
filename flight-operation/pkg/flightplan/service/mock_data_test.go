package service

import (
	fpl "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightplanID = fpl.ID("flightplan-id")
const DefaultFlightplanVersion = fpl.Version("version")
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFlightplanFleetID = fpl.FleetID("fleet-id")

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

type generatorMockFlightplan struct {
	fpl.Generator
	id           fpl.ID
	fleetID      fpl.FleetID
	versions     []fpl.Version
	versionIndex int
}

func (gen *generatorMockFlightplan) NewID() fpl.ID {
	return gen.id
}
func (gen *generatorMockFlightplan) NewFleetID() fpl.FleetID {
	return gen.fleetID
}
func (gen *generatorMockFlightplan) NewVersion() fpl.Version {
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

type flightplanComponentMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
	Version     string
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

func (f *flightplanComponentMock) GetFleetID() string {
	return f.FleetID
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.Version
}

type flightplanCommandMock struct {
	Flightplan flightplanMock
}

func (f *flightplanCommandMock) GetID() string {
	return f.Flightplan.ID
}

func (f *flightplanCommandMock) GetFlightplan() Flightplan {
	return &f.Flightplan
}

type flightplanMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
}

func (f *flightplanMock) GetID() string {
	return f.ID
}

func (f *flightplanMock) GetName() string {
	return f.Name
}

func (f *flightplanMock) GetDescription() string {
	return f.Description
}

func (f *flightplanMock) GetFleetID() string {
	return f.FleetID
}

type flightplanIDCommandMock struct {
	ID string
}

func (f *flightplanIDCommandMock) GetID() string {
	return f.ID
}

type fleetIDCommandMock struct {
	FleetID string
}

func (f *fleetIDCommandMock) GetID() string {
	return f.FleetID
}

type changeNumberOfVehiclesCommandFlightplanMock struct {
	ID               string
	NumberOfVehicles int
}

func (c *changeNumberOfVehiclesCommandFlightplanMock) GetID() string {
	return c.ID
}

func (c *changeNumberOfVehiclesCommandFlightplanMock) GetNumberOfVehicles() int {
	return c.NumberOfVehicles
}
