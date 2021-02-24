package flightplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightplanを一つ新しく作成し、初期状態を検証する。
func TestCreateNewFlightplan(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	flightplan := NewInstance(gen)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), "")
	a.Equal(flightplan.GetDescription(), "")
	a.Equal(flightplan.GetVersion(), DefaultVersion)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion)
	a.Equal(flightplan.gen, gen)
}

// Flightplanを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightplanAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightplanComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		version:     string(DefaultVersion),
	}
	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	flightplan := AssembleFrom(gen, comp)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion)
	a.Equal(flightplan.gen, gen)
}

// Flightplanを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightplan(t *testing.T) {
	a := assert.New(t)

	flightplan := &Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
	}
	comp := &flightplanComponentMock{}
	TakeApart(
		flightplan,
		func(id, name, description, version string) {
			comp.id = id
			comp.name = name
			comp.description = description
			comp.version = version
		},
	)

	expectComp := &flightplanComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		version:     string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
