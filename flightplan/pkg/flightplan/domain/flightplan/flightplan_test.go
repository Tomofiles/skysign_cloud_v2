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

	err := flightplan.NameFlightplan(DefaultName)

	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Flightplanの名前を変更する。
// カーボンコピーされたFlightplanの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenChangeFlightplansName(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultName1    = DefaultName + "-1"
		DefaultName2    = DefaultName + "-2"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2},
	}
	original := &Flightplan{
		id:           DefaultID,
		name:         DefaultName1,
		description:  DefaultDescription,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	flightplan := Copy(gen, CopiedID, original)

	err := flightplan.NameFlightplan(DefaultName2)

	a.Equal(flightplan.GetName(), DefaultName1)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
	a.Equal(err, ErrCannotChange)
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

	err := flightplan.ChangeDescription(DefaultDescription)

	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Flightplanの説明を変更する。
// カーボンコピーされたFlightplanの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenChangeFlightplansDescription(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID            = DefaultID + "-copied"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2},
	}
	original := &Flightplan{
		id:           DefaultID,
		name:         DefaultName,
		description:  DefaultDescription1,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	flightplan := Copy(gen, CopiedID, original)

	err := flightplan.ChangeDescription(DefaultDescription2)

	a.Equal(flightplan.GetDescription(), DefaultDescription1)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
	a.Equal(err, ErrCannotChange)
}
