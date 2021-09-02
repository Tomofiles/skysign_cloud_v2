package postgresql

import (
	frep "flight-operation/pkg/flightreport/domain/flightreport"
)

const DefaultID = frep.ID("flightreport-id")
const DefaultName = "flightreport-name"
const DefaultDescription = "flightreport-description"
const DefaultFleetID = frep.FleetID("fleet-id")

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
