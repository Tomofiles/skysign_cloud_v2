package service

import (
	"fleet-formation/pkg/common/domain/event"
	"fleet-formation/pkg/common/domain/txmanager"
	v "fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/stretchr/testify/mock"
)

const DefaultVehicleID = v.ID("vehicle-id")
const DefaultVehicleVersion = v.Version("version")
const DefaultVehicleName = "vehicle-name"
const DefaultVehicleCommunicationID = v.CommunicationID("communication-id")
const DefaultFleetID = v.FleetID("fleet-id")

type vehicleRepositoryMock struct {
	mock.Mock
}

func (r *vehicleRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*v.Vehicle, error) {
	ret := r.Called()
	var f []*v.Vehicle
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*v.Vehicle)
	}
	return f, ret.Error(1)
}

func (r *vehicleRepositoryMock) GetAllOrigin(
	tx txmanager.Tx,
) ([]*v.Vehicle, error) {
	ret := r.Called()
	var f []*v.Vehicle
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*v.Vehicle)
	}
	return f, ret.Error(1)
}

func (r *vehicleRepositoryMock) GetByID(
	tx txmanager.Tx,
	id v.ID,
) (*v.Vehicle, error) {
	ret := r.Called(id)
	var f *v.Vehicle
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*v.Vehicle)
	}
	return f, ret.Error(1)
}

func (r *vehicleRepositoryMock) Save(
	tx txmanager.Tx,
	flightplan *v.Vehicle,
) error {
	ret := r.Called(flightplan)
	return ret.Error(0)
}

func (r *vehicleRepositoryMock) Delete(
	tx txmanager.Tx,
	id v.ID,
) error {
	ret := r.Called(id)
	return ret.Error(0)
}

type generatorMock struct {
	v.Generator
	id           v.ID
	versions     []v.Version
	versionIndex int
}

func (gen *generatorMock) NewID() v.ID {
	return gen.id
}
func (gen *generatorMock) NewVersion() v.Version {
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

type vehicleComponentMock struct {
	ID              string
	Name            string
	CommunicationID string
	IsCarbonCopy    bool
	Version         string
}

func (v *vehicleComponentMock) GetID() string {
	return v.ID
}

func (v *vehicleComponentMock) GetName() string {
	return v.Name
}

func (v *vehicleComponentMock) GetCommunicationID() string {
	return v.CommunicationID
}

func (v *vehicleComponentMock) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

func (v *vehicleComponentMock) GetVersion() string {
	return v.Version
}

type vehicleCommandMock struct {
	vehicle *vehicleMock
}

func (f *vehicleCommandMock) GetID() string {
	return f.vehicle.ID
}

func (f *vehicleCommandMock) GetVehicle() Vehicle {
	return f.vehicle
}

type vehicleMock struct {
	ID              string
	Name            string
	CommunicationID string
}

func (f *vehicleMock) GetID() string {
	return f.ID
}

func (f *vehicleMock) GetName() string {
	return f.Name
}

func (f *vehicleMock) GetCommunicationID() string {
	return f.CommunicationID
}

type vehicleIDCommandMock struct {
	ID string
}

func (f *vehicleIDCommandMock) GetID() string {
	return f.ID
}

type carbonCopyCommandMock struct {
	OriginalID string
	NewID      string
	FleetID    string
}

func (f *carbonCopyCommandMock) GetOriginalID() string {
	return f.OriginalID
}

func (f *carbonCopyCommandMock) GetNewID() string {
	return f.NewID
}

func (f *carbonCopyCommandMock) GetFleetID() string {
	return f.FleetID
}
