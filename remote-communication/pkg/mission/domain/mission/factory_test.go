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
			PointOrder:       1,
			Latitude:         11.0,
			Longitude:        21.0,
			RelativeAltitude: 31.0,
			Speed:            41.0,
		},
		{
			PointOrder:       2,
			Latitude:         12.0,
			Longitude:        22.0,
			RelativeAltitude: 32.0,
			Speed:            42.0,
		},
		{
			PointOrder:       3,
			Latitude:         13.0,
			Longitude:        23.0,
			RelativeAltitude: 33.0,
			Speed:            43.0,
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
			PointOrder:       1,
			Latitude:         11.0,
			Longitude:        21.0,
			RelativeAltitude: 31.0,
			Speed:            41.0,
		},
		{
			PointOrder:       2,
			Latitude:         12.0,
			Longitude:        22.0,
			RelativeAltitude: 32.0,
			Speed:            42.0,
		},
		{
			PointOrder:       3,
			Latitude:         13.0,
			Longitude:        23.0,
			RelativeAltitude: 33.0,
			Speed:            43.0,
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
		func(pointOrder int, latitude, longitude, relativeAltitude, speed float64) {
			waypointComps = append(
				waypointComps,
				&waypointComponentMock{
					pointOrder:       pointOrder,
					latitude:         latitude,
					longitude:        longitude,
					relativeAltitude: relativeAltitude,
					speed:            speed,
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
