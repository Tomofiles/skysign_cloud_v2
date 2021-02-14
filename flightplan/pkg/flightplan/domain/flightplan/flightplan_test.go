package flightplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const DefaultID = ID("flightplan-id")
const DefaultVersion1 = Version("version-1")
const DefaultVersion2 = Version("version-2")
const DefaultName = "flightplan-name"

type testGenerator struct {
	Generator
	id           ID
	versions     []Version
	versionCount int
}

func (gen *testGenerator) NewID() ID {
	return gen.id
}
func (gen *testGenerator) NewVersion() Version {
	version := gen.versions[gen.versionCount]
	gen.versionCount++
	return version
}

func TestCreateNewFlightplan(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1},
	}
	flightplan := NewInstance(gen)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), "")
	a.Equal(flightplan.GetNumberOfVehicles(), 0)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
	a.Equal(flightplan.generator, gen)
}

func TestChangeFlightplansName(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	flightplan.NameFlightplan(DefaultName)

	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
}
