package flightoperation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightoperationを一つ新しく作成し、初期状態を検証する。
func TestCreateNewFlightoperation(t *testing.T) {
	a := assert.New(t)

	var (
		NewID = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		version: DefaultVersion,
	}
	flightoperation := NewInstance(gen, NewID)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Empty(flightoperation.GetName())
	a.Empty(flightoperation.GetDescription())
	a.Equal(flightoperation.GetFleetID(), NewID)
	a.Equal(flightoperation.isCompleted, Operating)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), DefaultVersion)
	a.Equal(flightoperation.gen, gen)
}

// Flightoperationを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightoperationAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightoperationComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     string(DefaultFleetID),
		isCompleted: DefaultIsCompleted,
		version:     string(DefaultVersion),
	}
	gen := &generatorMock{}
	flightoperation := AssembleFrom(gen, comp)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Equal(flightoperation.GetName(), DefaultName)
	a.Equal(flightoperation.GetDescription(), DefaultDescription)
	a.Equal(flightoperation.GetFleetID(), DefaultFleetID)
	a.Equal(flightoperation.isCompleted, DefaultIsCompleted)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), DefaultVersion)
	a.Equal(flightoperation.gen, gen)
}

// Flightoperationを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightoperation(t *testing.T) {
	a := assert.New(t)

	flightoperation := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		isCompleted: DefaultIsCompleted,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
	}
	comp := &flightoperationComponentMock{}
	TakeApart(
		flightoperation,
		func(id, name, description, fleetID, version string, isCompleted bool) {
			comp.id = id
			comp.name = name
			comp.description = description
			comp.fleetID = fleetID
			comp.isCompleted = isCompleted
			comp.version = version
		},
	)

	expectComp := &flightoperationComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     string(DefaultFleetID),
		isCompleted: DefaultIsCompleted,
		version:     string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
