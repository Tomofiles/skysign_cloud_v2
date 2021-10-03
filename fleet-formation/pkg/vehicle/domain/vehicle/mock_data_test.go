package vehicle

import (
	"errors"

	"github.com/google/uuid"
)

var NewID = func() ID {
	id, _ := uuid.NewRandom()
	return ID(id.String())
}
var NewVersion = func() Version {
	id, _ := uuid.NewRandom()
	return Version(id.String())
}
var NewCommunicationID = func() CommunicationID {
	id, _ := uuid.NewRandom()
	return CommunicationID(id.String())
}
var NewFleetID = func() FleetID {
	id, _ := uuid.NewRandom()
	return FleetID(id.String())
}

const DefaultName = "vehicle-name"

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

// Flightplan用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id           ID
	versions     []Version
	versionIndex int
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewVersion() Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

// Flightplan用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

// Vehicle構成オブジェクトモック
type vehicleComponentMock struct {
	id              string
	name            string
	communicationID string
	isCarbonCopy    bool
	version         string
}

func (v *vehicleComponentMock) GetID() string {
	return v.id
}

func (v *vehicleComponentMock) GetName() string {
	return v.name
}

func (v *vehicleComponentMock) GetCommunicationID() string {
	return v.communicationID
}

func (v *vehicleComponentMock) GetIsCarbonCopy() bool {
	return v.isCarbonCopy
}

func (v *vehicleComponentMock) GetVersion() string {
	return v.version
}
