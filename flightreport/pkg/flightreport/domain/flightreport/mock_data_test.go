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
