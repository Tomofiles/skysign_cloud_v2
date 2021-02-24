package flightplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Flightplanの名前を変更する。
// バージョンが更新されていることを検証する。
func TestChangeFlightplansName(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	flightplan.NameFlightplan(DefaultName)

	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
}

// Flightplanの説明を変更する。
// バージョンが更新されていることを検証する。
func TestChangeFlightplansDescription(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	flightplan := NewInstance(gen)

	flightplan.ChangeDescription(DefaultDescription)

	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
}
