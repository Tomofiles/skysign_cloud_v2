package flightplan

import "errors"

const DefaultID = ID("flightplan-id")
const DefaultVersion = Version("version")
const DefaultName = "flightplan-name"
const DefaultDescription = "flightplan-description"

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

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	id           string
	name         string
	description  string
	isCarbonCopy bool
	version      string
}

func (f *flightplanComponentMock) GetID() string {
	return f.id
}

func (f *flightplanComponentMock) GetName() string {
	return f.name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.description
}

func (f *flightplanComponentMock) GetIsCarbonCopy() bool {
	return f.isCarbonCopy
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.version
}
