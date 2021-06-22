package mission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Missionの名前を変更する。
// バージョンが更新されていることを検証する。
func TestChangeMissionsName(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	mission := NewInstance(gen)

	err := mission.NameMission(DefaultName)

	a.Equal(mission.GetName(), DefaultName)
	a.Equal(mission.GetVersion(), DefaultVersion1)
	a.Equal(mission.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Missionに新たなNavigationを追加する。
// Navigation自体の詳細なテストは、MissionNavigationTestsにまとめている。
func TestAddNavigationToMission(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2},
	}
	mission := NewInstance(gen)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	err := mission.ReplaceNavigationWith(navigation)

	a.Equal(mission.GetNavigation(), navigation)
	a.Equal(mission.GetVersion(), DefaultVersion1)
	a.Equal(mission.GetNewVersion(), DefaultVersion2)
	a.Nil(err)
}

// Missionの名前を変更する。
// カーボンコピーされたMissionの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenChangeMissionsName(t *testing.T) {
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
	original := &Mission{
		id:           DefaultID,
		name:         DefaultName1,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	mission := Copy(gen, CopiedID, original)

	err := mission.NameMission(DefaultName2)

	a.Equal(mission.GetName(), DefaultName1)
	a.Equal(mission.GetVersion(), DefaultVersion1)
	a.Equal(mission.GetNewVersion(), DefaultVersion1)
	a.Equal(err, ErrCannotChange)
}

// Missionに対してNavigationを付与する。
// カーボンコピーされたMissionの変更がエラーとなることを検証する。
func TestCannotChangeErrorWhenChangeCarbonCopiedMissionsNavigation(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID        = DefaultID + "-copied"
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2},
	}
	original := &Mission{
		id:           DefaultID,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
	}
	mission := Copy(gen, CopiedID, original)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	err := mission.ReplaceNavigationWith(navigation)

	a.Nil(mission.GetNavigation())
	a.Equal(mission.GetVersion(), DefaultVersion1)
	a.Equal(mission.GetNewVersion(), DefaultVersion1)
	a.Equal(err, ErrCannotChange)
}
