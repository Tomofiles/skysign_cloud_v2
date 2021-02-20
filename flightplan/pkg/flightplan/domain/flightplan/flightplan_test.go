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
		versions: []Version{DefaultVersion1},
	}
	flightplan := NewInstance(gen)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), "")
	a.Equal(flightplan.GetDescription(), "")
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
	a.Equal(flightplan.gen, gen)
}

// Flightplanの名前を変更する。
// バージョンが更新されていることを検証する。
func TestChangeFlightplansName(t *testing.T) {
	a := assert.New(t)

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

// Flightplanを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestFlightplanAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &flightplanComponentMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		version:     string(DefaultVersion1),
	}
	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1},
	}
	flightplan := AssembleFrom(gen, comp)

	a.Equal(flightplan.GetID(), DefaultID)
	a.Equal(flightplan.GetName(), DefaultName)
	a.Equal(flightplan.GetDescription(), DefaultDescription)
	a.Equal(flightplan.GetVersion(), DefaultVersion1)
	a.Equal(flightplan.GetNewVersion(), DefaultVersion1)
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
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
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
		version:     string(DefaultVersion1),
	}
	a.Equal(comp, expectComp)
}

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	id          string
	name        string
	description string
	version     string
}

func (f *flightplanComponentMock) GetID() string {
	return f.id
}

func (f *flightplanComponentMock) GetName() string {
	return f.name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.description
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.version
}
