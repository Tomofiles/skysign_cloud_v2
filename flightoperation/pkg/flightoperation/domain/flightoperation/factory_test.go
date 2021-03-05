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
		id: DefaultID,
	}
	flightoperation := NewInstance(gen, NewID)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Equal(flightoperation.GetFlightplanID(), NewID)
}

// Flightoperationを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightoperationAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightoperationComponentMock{
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
	}
	gen := &generatorMock{}
	flightoperation := AssembleFrom(gen, comp)

	a.Equal(flightoperation.GetID(), DefaultID)
	a.Equal(flightoperation.GetFlightplanID(), DefaultFlightplanID)
}

// Flightoperationを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightoperation(t *testing.T) {
	a := assert.New(t)

	flightoperation := &Flightoperation{
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
	}
	comp := &flightoperationComponentMock{}
	TakeApart(
		flightoperation,
		func(id, flightplanID string) {
			comp.id = id
			comp.flightplanID = flightplanID
		},
	)

	expectComp := &flightoperationComponentMock{
		id:           string(DefaultID),
		flightplanID: string(DefaultFlightplanID),
	}
	a.Equal(comp, expectComp)
}
