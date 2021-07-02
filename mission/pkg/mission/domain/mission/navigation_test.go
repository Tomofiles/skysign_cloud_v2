package mission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Navigationを一つ新しく作成し、初期状態を検証する。
func TestCreateNewNavigation(t *testing.T) {
	a := assert.New(t)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	var resCall bool
	navigation.ProvideWaypointsInterest(
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			resCall = true
		},
	)

	a.Equal(navigation.currentOrder, 0)
	a.Equal(navigation.takeoffPointGroundHeightWGS84EllipsoidM, NewHeightFromM(DefaultTakeoffPointGroundHeightWGS84EllipsoidM))
	a.Len(navigation.waypoints, 0)
	a.False(resCall)
}

// Navigationを一つ新しく作成し、Waypointを1点追加する。
// Waypoint追加後の内部状態を検証する。
func TestCreateAndPushSingleWaypointNavigation(t *testing.T) {
	a := assert.New(t)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	navigation.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)

	var wpComp waypointComponentMock
	navigation.ProvideWaypointsInterest(
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			wpComp.pointOrder = pointOrder
			wpComp.latitudeDegree = latitudeDegree
			wpComp.longitudeDegree = longitudeDegree
			wpComp.relativeHeightM = relativeHeightM
			wpComp.speedMS = speedMS
		},
	)

	expectNav := []Waypoint{
		{
			pointOrder:     1,
			coordinates:    NewGeodesicCoordinatesFromDegree(11.0, 21.0),
			relativeHeight: NewHeightFromM(31.0),
			speed:          NewSpeedFromMS(41.0),
		},
	}

	expectWp := waypointComponentMock{
		pointOrder:      1,
		latitudeDegree:  11.0,
		longitudeDegree: 21.0,
		relativeHeightM: 31.0,
		speedMS:         41.0,
	}

	a.Equal(navigation.currentOrder, 1)
	a.Equal(navigation.takeoffPointGroundHeightWGS84EllipsoidM, NewHeightFromM(DefaultTakeoffPointGroundHeightWGS84EllipsoidM))
	a.Len(navigation.waypoints, 1)
	a.Equal(navigation.waypoints, expectNav)
	a.Equal(wpComp, expectWp)
}

// Navigationを一つ新しく作成し、Waypointを複数点追加する。
// Waypoint追加後の内部状態を検証する。
func TestCreateAndPushMultipleWaypointsNavigation(t *testing.T) {
	a := assert.New(t)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	navigation.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	navigation.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	navigation.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	var wpComps []waypointComponentMock
	navigation.ProvideWaypointsInterest(
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			wpComps = append(
				wpComps,
				waypointComponentMock{
					pointOrder:      pointOrder,
					latitudeDegree:  latitudeDegree,
					longitudeDegree: longitudeDegree,
					relativeHeightM: relativeHeightM,
					speedMS:         speedMS,
				},
			)
		},
	)

	expectNav := []Waypoint{
		{
			pointOrder:     1,
			coordinates:    NewGeodesicCoordinatesFromDegree(11.0, 21.0),
			relativeHeight: NewHeightFromM(31.0),
			speed:          NewSpeedFromMS(41.0),
		},
		{
			pointOrder:     2,
			coordinates:    NewGeodesicCoordinatesFromDegree(12.0, 22.0),
			relativeHeight: NewHeightFromM(32.0),
			speed:          NewSpeedFromMS(42.0),
		},
		{
			pointOrder:     3,
			coordinates:    NewGeodesicCoordinatesFromDegree(13.0, 23.0),
			relativeHeight: NewHeightFromM(33.0),
			speed:          NewSpeedFromMS(43.0),
		},
	}

	expectWps := []waypointComponentMock{
		{
			pointOrder:      1,
			latitudeDegree:  11.0,
			longitudeDegree: 21.0,
			relativeHeightM: 31.0,
			speedMS:         41.0,
		},
		{
			pointOrder:      2,
			latitudeDegree:  12.0,
			longitudeDegree: 22.0,
			relativeHeightM: 32.0,
			speedMS:         42.0,
		},
		{
			pointOrder:      3,
			latitudeDegree:  13.0,
			longitudeDegree: 23.0,
			relativeHeightM: 33.0,
			speedMS:         43.0,
		},
	}

	a.Equal(navigation.currentOrder, 3)
	a.Equal(navigation.takeoffPointGroundHeightWGS84EllipsoidM, NewHeightFromM(DefaultTakeoffPointGroundHeightWGS84EllipsoidM))
	a.Len(navigation.waypoints, 3)
	a.Equal(navigation.waypoints, expectNav)
	a.Equal(wpComps, expectWps)
}
