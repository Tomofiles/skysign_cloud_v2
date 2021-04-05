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
	a.Equal(flightplan.isCarbonCopy, Original)
	a.Equal(flightplan.GetVersion(), DefaultVersion)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion)
	a.Equal(flightplan.gen, gen)
}

// Flightplanのカーボンコピーを作成し、初期状態を検証する。
func TestCopyFlightplan(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	id := DefaultID + "-copied"
	original := &Flightplan{
		id:           DefaultID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
	}
	flightplan := Copy(gen, id, original)

	a.Equal(flightplan.GetID(), id)
	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.isCarbonCopy, CarbonCopy)
	a.Equal(flightplan.GetVersion(), DefaultVersion)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion)
	a.Equal(flightplan.gen, gen)
}

// Flightplanを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightplanAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightplanComponentMock{
		id:           string(DefaultID),
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: CarbonCopy,
		version:      string(DefaultVersion),
	}
	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	flightplan := AssembleFrom(gen, comp)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.isCarbonCopy, CarbonCopy)
	a.Equal(flightplan.GetVersion(), DefaultVersion)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion)
	a.Equal(flightplan.gen, gen)
}

// Flightplanを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartFlightplan(t *testing.T) {
	a := assert.New(t)

	flightplan := &Flightplan{
		id:           DefaultID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
	}
	comp := &flightplanComponentMock{}
	TakeApart(
		flightplan,
		func(id, name, description, version string, isCarbonCopy bool) {
			comp.id = id
			comp.name = name
			comp.description = description
			comp.isCarbonCopy = isCarbonCopy
			comp.version = version
		},
	)

	expectComp := &flightplanComponentMock{
		id:           string(DefaultID),
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: CarbonCopy,
		version:      string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
