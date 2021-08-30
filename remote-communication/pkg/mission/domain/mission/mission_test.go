package mission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MissionにWaypointを1件追加する。
func TestPushSingleWaypoint(t *testing.T) {
	a := assert.New(t)

	mission := NewInstance(DefaultID)

	ret := mission.PushWaypoint(1.0, 2.0, 3.0, 4.0)

	expectMission := AssembleFrom(
		&missionComponentMock{
			id: string(DefaultID),
			waypoints: []*waypointComponentMock{
				{
					1, 1.0, 2.0, 3.0, 4.0,
				},
			},
		},
	)
	a.Equal(ret, 1)
	a.Equal(mission, expectMission)
	a.Equal(mission.GetWaypoints(), expectMission.GetWaypoints())
}

// MissionにWaypointを複数件追加する。
func TestPushMultipleWaypoints(t *testing.T) {
	a := assert.New(t)

	mission := NewInstance(DefaultID)

	ret1 := mission.PushWaypoint(11.0, 21.0, 31.0, 41.0)
	ret2 := mission.PushWaypoint(12.0, 22.0, 32.0, 42.0)
	ret3 := mission.PushWaypoint(13.0, 23.0, 33.0, 43.0)

	expectMission := AssembleFrom(
		&missionComponentMock{
			id: string(DefaultID),
			waypoints: []*waypointComponentMock{
				{
					1, 11.0, 21.0, 31.0, 41.0,
				},
				{
					2, 12.0, 22.0, 32.0, 42.0,
				},
				{
					3, 13.0, 23.0, 33.0, 43.0,
				},
			},
		},
	)
	a.Equal(ret1, 1)
	a.Equal(ret2, 2)
	a.Equal(ret3, 3)
	a.Equal(mission, expectMission)
	a.Equal(mission.GetWaypoints(), expectMission.GetWaypoints())
}
