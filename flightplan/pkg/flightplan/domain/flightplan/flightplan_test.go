package flightplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewFlightplan(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1},
	}
	flightplan := NewInstance(gen)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), "")
	a.Equal(flightplan.GetDescription(), "")
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

func TestChangeFlightplansDescription(t *testing.T) {
	a := assert.New(t)

	gen := &testGenerator{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	flightplan.ChangeDescription(DefaultDescription)

	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
}
