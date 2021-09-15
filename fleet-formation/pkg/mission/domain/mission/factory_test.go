package mission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Missionを一つ新しく作成し、初期状態を検証する。
func TestCreateNewMission(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	mission := NewInstance(gen)

	a.Equal(mission.GetID(), DefaultID)
	a.Empty(mission.GetName())
	a.Nil(mission.GetNavigation())
	a.Equal(mission.isCarbonCopy, Original)
	a.Equal(mission.GetVersion(), DefaultVersion)
	a.Equal(mission.GetNewVersion(), DefaultVersion)
	a.Equal(mission.gen, gen)
	a.Nil(mission.pub)
}

// Missionのカーボンコピーを作成し、初期状態を検証する。
func TestCopyMission(t *testing.T) {
	a := assert.New(t)

	var (
		CopiedID    = DefaultID + "-copied"
		NewUploadID = DefaultUploadID + "-new"
	)

	gen := &generatorMock{
		uploadID: NewUploadID,
	}
	navigation := NewNavigation(DefaultTakeoffPointGroundAltitudeM)
	navigation.uploadID = DefaultUploadID
	navigation.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	navigation.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	navigation.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)
	original := &Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		pub:          &publisherMock{},
	}
	mission := Copy(gen, CopiedID, original)

	expectNav := NewNavigation(DefaultTakeoffPointGroundAltitudeM)
	expectNav.uploadID = NewUploadID
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	expectNav.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	expectNav.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	a.Equal(mission.GetID(), CopiedID)
	a.Equal(mission.GetName(), DefaultName)
	a.Equal(mission.GetNavigation(), expectNav)
	a.Equal(mission.isCarbonCopy, CarbonCopy)
	a.Equal(mission.GetVersion(), DefaultVersion)
	a.Equal(mission.GetNewVersion(), DefaultVersion)
	a.Equal(mission.gen, gen)
	a.Nil(mission.pub)
}

// Missionを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestMissionAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	comp := &missionComponentMock{
		id:   string(DefaultID),
		name: DefaultName,
		navigation: navigationComponentMock{
			takeoffPointGroundAltitudeM: DefaultTakeoffPointGroundAltitudeM,
			waypoints: []waypointComponentMock{
				{
					pointOrder:        1,
					latitudeDegree:    11.0,
					longitudeDegree:   21.0,
					relativeAltitudeM: 31.0,
					speedMS:           41.0,
				},
				{
					pointOrder:        2,
					latitudeDegree:    12.0,
					longitudeDegree:   22.0,
					relativeAltitudeM: 32.0,
					speedMS:           42.0,
				},
				{
					pointOrder:        3,
					latitudeDegree:    13.0,
					longitudeDegree:   23.0,
					relativeAltitudeM: 33.0,
					speedMS:           43.0,
				},
			},
			uploadID: string(DefaultUploadID),
		},
		isCarbonCopy: CarbonCopy,
		version:      string(DefaultVersion),
	}
	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion},
	}
	mission := AssembleFrom(gen, comp)

	expectNav := NewNavigation(DefaultTakeoffPointGroundAltitudeM)
	expectNav.uploadID = DefaultUploadID
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	expectNav.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	expectNav.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	a.Equal(mission.GetID(), DefaultID)
	a.Equal(mission.GetName(), DefaultName)
	a.Equal(mission.GetNavigation(), expectNav)
	a.Equal(mission.isCarbonCopy, CarbonCopy)
	a.Equal(mission.GetVersion(), DefaultVersion)
	a.Equal(mission.GetNewVersion(), DefaultVersion)
	a.Equal(mission.gen, gen)
	a.Nil(mission.pub)
}

// Missionを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartMission(t *testing.T) {
	a := assert.New(t)

	navigation := NewNavigation(DefaultTakeoffPointGroundAltitudeM)
	navigation.uploadID = DefaultUploadID
	navigation.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	navigation.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	navigation.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)
	mission := &Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		pub:          &publisherMock{},
	}
	comp := &missionComponentMock{
		navigation: navigationComponentMock{},
	}
	TakeApart(
		mission,
		func(id, name, version string, isCarbonCopy bool) {
			comp.id = id
			comp.name = name
			comp.isCarbonCopy = isCarbonCopy
			comp.version = version
		},
		func(takeoffPointGroundAltitudeM float64, uploadID string) {
			comp.navigation.takeoffPointGroundAltitudeM = takeoffPointGroundAltitudeM
			comp.navigation.uploadID = uploadID
		},
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64) {
			comp.navigation.waypoints = append(
				comp.navigation.waypoints,
				waypointComponentMock{
					pointOrder:        pointOrder,
					latitudeDegree:    latitudeDegree,
					longitudeDegree:   longitudeDegree,
					relativeAltitudeM: relativeAltitudeM,
					speedMS:           speedMS,
				},
			)
		},
	)

	expectComp := &missionComponentMock{
		id:   string(DefaultID),
		name: DefaultName,
		navigation: navigationComponentMock{
			takeoffPointGroundAltitudeM: DefaultTakeoffPointGroundAltitudeM,
			waypoints: []waypointComponentMock{
				{
					pointOrder:        1,
					latitudeDegree:    11.0,
					longitudeDegree:   21.0,
					relativeAltitudeM: 31.0,
					speedMS:           41.0,
				},
				{
					pointOrder:        2,
					latitudeDegree:    12.0,
					longitudeDegree:   22.0,
					relativeAltitudeM: 32.0,
					speedMS:           42.0,
				},
				{
					pointOrder:        3,
					latitudeDegree:    13.0,
					longitudeDegree:   23.0,
					relativeAltitudeM: 33.0,
					speedMS:           43.0,
				},
			},
			uploadID: string(DefaultUploadID),
		},
		isCarbonCopy: CarbonCopy,
		version:      string(DefaultVersion),
	}
	a.Equal(comp, expectComp)
}
