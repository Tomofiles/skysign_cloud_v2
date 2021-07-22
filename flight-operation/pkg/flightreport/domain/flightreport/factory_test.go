package flightreport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightreportを一つ新しく作成し、初期状態を検証する。
func TestCreateNewFlightreport(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id: DefaultID,
	}
	flightreport := NewInstance(gen, DefaultName, DefaultDescription, DefaultFleetID)

	a.Equal(flightreport.GetID(), DefaultID)
	a.Equal(flightreport.GetName(), DefaultName)
	a.Equal(flightreport.GetDescription(), DefaultDescription)
	a.Equal(flightreport.GetFleetID(), DefaultFleetID)
}

// Flightreportを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightreportAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightreportComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     string(DefaultFleetID),
	}
	gen := &generatorMock{}
	flightreport := AssembleFrom(gen, comp)

	a.Equal(flightreport.GetID(), DefaultID)
	a.Equal(flightreport.GetName(), DefaultName)
	a.Equal(flightreport.GetDescription(), DefaultDescription)
	a.Equal(flightreport.GetFleetID(), DefaultFleetID)
}

// Flightreportを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightreport(t *testing.T) {
	a := assert.New(t)

	flightreport := &Flightreport{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
	}
	comp := &flightreportComponentMock{}
	TakeApart(
		flightreport,
		func(id, name, description, fleetID string) {
			comp.id = id
			comp.name = name
			comp.description = description
			comp.fleetID = fleetID
		},
	)

	expectComp := &flightreportComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     string(DefaultFleetID),
	}
	a.Equal(comp, expectComp)
}
