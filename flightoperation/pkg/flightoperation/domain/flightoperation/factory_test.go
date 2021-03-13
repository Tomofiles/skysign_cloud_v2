package flightoperation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightoperationを一つ新しく作成し、初期状態を検証する。
func TestCreateNewFlightoperation(t *testing.T) {
	a := assert.New(t)

	var (
		NewID = DefaultFlightplanID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		version: DefaultVersion,
	}
	flightoperation := NewInstance(gen, NewID)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Equal(flightoperation.GetFlightplanID(), NewID)
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
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
		isCompleted:  DefaultIsCompleted,
		version:      string(DefaultVersion),
	}
	gen := &generatorMock{}
	flightoperation := AssembleFrom(gen, comp)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Equal(flightoperation.GetFlightplanID(), DefaultFlightplanID)
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
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
		isCompleted:  DefaultIsCompleted,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
	}
	comp := &flightoperationComponentMock{}
	TakeApart(
		flightoperation,
		func(id, flightplanID, version string, isCompleted bool) {
			comp.id = id
			comp.flightplanID = flightplanID
			comp.isCompleted = isCompleted
			comp.version = version
		},
	)

	expectComp := &flightoperationComponentMock{
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
		isCompleted:  DefaultIsCompleted,
		version:      string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
