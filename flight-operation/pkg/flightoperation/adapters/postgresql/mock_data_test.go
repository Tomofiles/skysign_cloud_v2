package postgresql

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"
)

const DefaultID = fope.ID("flightoperation-id")
const DefaultName = "flightoperation-name"
const DefaultDescription = "flightoperation-description"
const DefaultFleetID = fope.FleetID("fleet-id")
const DefaultIsCompleted = fope.Completed
const DefaultVersion = fope.Version("version")

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

// GetName .
func (f *flightoperationComponentMock) GetName() string {
	return f.name
}

// GetDescription .
func (f *flightoperationComponentMock) GetDescription() string {
	return f.description
}

// GetFleetID .
func (f *flightoperationComponentMock) GetFleetID() string {
	return f.fleetID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.isCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.version
}
