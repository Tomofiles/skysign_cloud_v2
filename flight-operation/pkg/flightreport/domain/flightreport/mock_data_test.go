package flightreport

import "errors"

const DefaultID = ID("flightreport-id")
const DefaultName = "flightreport-name"
const DefaultDescription = "flightreport-description"
const DefaultFleetID = FleetID("fleet-id")

var (
	ErrSave = errors.New("save error")
)

// Flightreport用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id ID
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}

// Flightreport構成オブジェクトモック
type flightreportComponentMock struct {
	id          string
	name        string
	description string
	fleetID     string
}

func (f *flightreportComponentMock) GetID() string {
	return f.id
}

func (f *flightreportComponentMock) GetName() string {
	return f.name
}

func (f *flightreportComponentMock) GetDescription() string {
	return f.description
}

func (f *flightreportComponentMock) GetFleetID() string {
	return f.fleetID
}
