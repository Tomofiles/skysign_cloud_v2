package flightreport

import "errors"

const DefaultID = ID("flightreport-id")
const DefaultFlightoperationID = FlightoperationID("flightoperation-id")

var (
	ErrSave = errors.New("save error")
)

// Flightreport用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id                ID
	flightoperationID FlightoperationID
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewFlightoperationID() FlightoperationID {
	return gen.flightoperationID
}

// Flightreport用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

// Flightreport構成オブジェクトモック
type flightreportComponentMock struct {
	id                string
	flightoperationID string
}

func (f *flightreportComponentMock) GetID() string {
	return f.id
}

func (f *flightreportComponentMock) GetFlightoperationID() string {
	return f.flightoperationID
}
