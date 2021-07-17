package flightoperation

import (
	"errors"
	"flightoperation/pkg/flightoperation/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultID = ID("flightoperation-id")
const DefaultName = "flightoperation-name"
const DefaultDescription = "flightoperation-description"
const DefaultFleetID = FleetID("fleet-id")
const DefaultIsCompleted = Completed
const DefaultVersion = Version("version")

var (
	ErrSave = errors.New("save error")
	ErrGet  = errors.New("get error")
)

// Flightoperation用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id      ID
	fleetID FleetID
	version Version
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewFleetID() FleetID {
	return gen.fleetID
}
func (gen *generatorMock) NewVersion() Version {
	return gen.version
}

// Flightoperation用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

type flightoperationRepositoryMock struct {
	mock.Mock

	saveFlightoperations []*Flightoperation
}

func (r *flightoperationRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*Flightoperation, error) {
	panic("implement me")
}

func (r *flightoperationRepositoryMock) GetAllOperating(
	tx txmanager.Tx,
) ([]*Flightoperation, error) {
	panic("implement me")
}

func (r *flightoperationRepositoryMock) GetByID(
	tx txmanager.Tx,
	id ID,
) (*Flightoperation, error) {
	ret := r.Called(id)
	var f *Flightoperation
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*Flightoperation)
	}
	return f, ret.Error(1)
}

func (r *flightoperationRepositoryMock) Save(
	tx txmanager.Tx,
	flightoperation *Flightoperation,
) error {
	ret := r.Called(flightoperation)
	if ret.Error(0) == nil {
		r.saveFlightoperations = append(r.saveFlightoperations, flightoperation)
	}
	return ret.Error(0)
}

// Flightoperation構成オブジェクトモック
type flightoperationComponentMock struct {
	id          string
	name        string
	description string
	fleetID     string
	isCompleted bool
	version     string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.id
}

func (f *flightoperationComponentMock) GetName() string {
	return f.name
}

func (f *flightoperationComponentMock) GetDescription() string {
	return f.description
}

func (f *flightoperationComponentMock) GetFleetID() string {
	return f.fleetID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.isCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.version
}
