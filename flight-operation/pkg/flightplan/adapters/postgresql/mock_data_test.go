package postgresql

import (
	"flight-operation/pkg/flightplan/domain/flightplan"
)

const DefaultFlightplanID = flightplan.ID("flightplan-id")
const DefaultFlightplanVersion = flightplan.Version("version")
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFlightplanFleetID = flightplan.FleetID("fleet-id")

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	id          string
	name        string
	description string
	fleetID     string
	version     string
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

func (f *flightplanComponentMock) GetFleetID() string {
	return f.fleetID
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.version
}
