package flightoperation

import "errors"

const DefaultID = ID("flightoperation-id")
const DefaultFlightplanID = FlightplanID("flightplan-id")
const DefaultIsCompleted = Completed
const DefaultVersion = Version("version")

var (
	ErrSave = errors.New("save error")
)

// Flightoperation用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id           ID
	flightplanID FlightplanID
	version      Version
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewFlightplanID() FlightplanID {
	return gen.flightplanID
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

// Flightoperation構成オブジェクトモック
type flightoperationComponentMock struct {
	id           string
	flightplanID string
	isCompleted  bool
	version      string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.id
}

func (f *flightoperationComponentMock) GetFlightplanID() string {
	return f.flightplanID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.isCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.version
}
