package service

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultID = fope.ID("flightoperation-id")
const DefaultName = "flightoperation-name"
const DefaultDescription = "flightoperation-description"
const DefaultFleetID = fope.FleetID("fleet-id")
const DefaultIsCompleted = fope.Completed
const DefaultVersion = fope.Version("version")

type flightoperationRepositoryMock struct {
	mock.Mock
}

func (r *flightoperationRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*fope.Flightoperation, error) {
	ret := r.Called()
	var f []*fope.Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*fope.Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) GetAllOperating(
	tx txmanager.Tx,
) ([]*fope.Flightoperation, error) {
	ret := r.Called()
	var f []*fope.Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*fope.Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) GetByID(
	tx txmanager.Tx,
	id fope.ID,
) (*fope.Flightoperation, error) {
	ret := r.Called(id)
	var f *fope.Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*fope.Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) Save(
	tx txmanager.Tx,
	flightoperation *fope.Flightoperation,
) error {
	ret := r.Called(flightoperation)
	return ret.Error(0)
}

type generatorMock struct {
	fope.Generator
	id      fope.ID
	fleetID fope.FleetID
	version fope.Version
}

func (gen *generatorMock) NewID() fope.ID {
	return gen.id
}
func (gen *generatorMock) NewFleetID() fope.FleetID {
	return gen.fleetID
}
func (gen *generatorMock) NewVersion() fope.Version {
	return gen.version
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

type flightoperationComponentMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
	IsCompleted bool
	Version     string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.ID
}

func (f *flightoperationComponentMock) GetName() string {
	return f.Name
}

func (f *flightoperationComponentMock) GetDescription() string {
	return f.Description
}

func (f *flightoperationComponentMock) GetFleetID() string {
	return f.FleetID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.IsCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.Version
}

type flightoperationIDCommandMock struct {
	ID string
}

func (f *flightoperationIDCommandMock) GetID() string {
	return f.ID
}

type flightoperationCommandMock struct {
	Flightoperation flightoperationMock
}

func (f *flightoperationCommandMock) GetFlightoperation() Flightoperation {
	return &f.Flightoperation
}

type flightoperationModelMock struct {
	flightoperation *flightoperationMock
}

func (f *flightoperationModelMock) GetFlightoperation() Flightoperation {
	return &flightoperationMock{
		ID:          f.flightoperation.ID,
		Name:        f.flightoperation.Name,
		Description: f.flightoperation.Description,
		FleetID:     f.flightoperation.FleetID,
	}
}

type flightoperationMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
}

func (f *flightoperationMock) GetID() string {
	return f.ID
}

func (f *flightoperationMock) GetName() string {
	return f.Name
}

func (f *flightoperationMock) GetDescription() string {
	return f.Description
}

func (f *flightoperationMock) GetFleetID() string {
	return f.FleetID
}
