package mission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Missionを一つ新しく作成し、初期状態を検証する。
func TestCreateNewMission(t *testing.T) {
	a := assert.New(t)

	mission := NewInstance(DefaultID)

	a.Equal(mission.GetID(), DefaultID)
	a.Len(mission.waypoints, 0)
}

// Missionを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestMissionAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultID),
		waypoints: waypointComps,
	}
	mission := AssembleFrom(missionComp)

	expectWaypoints := []*Waypoint{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}

	a.Equal(mission.GetID(), DefaultID)
	a.Equal(mission.waypoints, expectWaypoints)
}

// Missionを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartMission(t *testing.T) {
	a := assert.New(t)

	waypoints := []*Waypoint{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	mission := NewInstance(DefaultID)
	mission.waypoints = waypoints

	var missionComp missionComponentMock
	var waypointComps []*waypointComponentMock

	TakeApart(
		mission,
		func(id string) {
			missionComp.id = id
		},
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			waypointComps = append(
				waypointComps,
				&waypointComponentMock{
					pointOrder:      pointOrder,
					latitudeDegree:  latitudeDegree,
					longitudeDegree: longitudeDegree,
					relativeHeightM: relativeHeightM,
					speedMS:         speedMS,
				},
			)
		},
	)

	missionComp.waypoints = waypointComps

	expectWaypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}
	expectMissionComp := &missionComponentMock{
		id:        string(DefaultID),
		waypoints: expectWaypointComps,
	}
	a.Equal(&missionComp, expectMissionComp)
}
