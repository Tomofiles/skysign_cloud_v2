package flightreport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightreportを一つ新しく作成し、初期状態を検証する。
func TestCreateNewFlightreport(t *testing.T) {
	a := assert.New(t)

	var (
		NewID = DefaultFlightoperationID + "-new"
	)

	gen := &generatorMock{
		id: DefaultID,
	}
	flightreport := NewInstance(gen, NewID)

	a.Equal(flightreport.GetID(), DefaultID)
	a.Equal(flightreport.GetFlightoperationID(), NewID)
}

// Flightreportを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightreportAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightreportComponentMock{
		id:                string(DefaultID),
		flightoperationID: string(DefaultFlightoperationID),
	}
	gen := &generatorMock{}
	flightreport := AssembleFrom(gen, comp)

	a.Equal(flightreport.GetID(), DefaultID)
	a.Equal(flightreport.GetFlightoperationID(), DefaultFlightoperationID)
}

// Flightreportを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightreport(t *testing.T) {
	a := assert.New(t)

	flightreport := &Flightreport{
		id:                DefaultID,
		flightoperationID: DefaultFlightoperationID,
	}
	comp := &flightreportComponentMock{}
	TakeApart(
		flightreport,
		func(id, flightoperationID string) {
			comp.id = id
			comp.flightoperationID = flightoperationID
		},
	)

	expectComp := &flightreportComponentMock{
		id:                string(DefaultID),
		flightoperationID: string(DefaultFlightoperationID),
	}
	a.Equal(comp, expectComp)
}
