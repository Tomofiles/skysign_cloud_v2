package service

import (
	m "fleet-formation/pkg/mission/domain/mission"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoNavigationTransformerFromCommand(t *testing.T) {
	a := assert.New(t)

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints:                   []waypointMock{},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundAltitudeM)

	a.Equal(navigation, expectNav)
}

func TestSingleNavigationTransformerFromCommand(t *testing.T) {
	a := assert.New(t)

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints: []waypointMock{
					{
						LatitudeDegree:    11.0,
						LongitudeDegree:   21.0,
						RelativeAltitudeM: 31.0,
						SpeedMS:           41.0,
					},
				},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundAltitudeM)
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)

	a.Equal(navigation, expectNav)
}

func TestMultipleNavigationTransformerFromCommand(t *testing.T) {
	a := assert.New(t)

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints: []waypointMock{
					{
						LatitudeDegree:    11.0,
						LongitudeDegree:   21.0,
						RelativeAltitudeM: 31.0,
						SpeedMS:           41.0,
					},
					{
						LatitudeDegree:    12.0,
						LongitudeDegree:   22.0,
						RelativeAltitudeM: 32.0,
						SpeedMS:           42.0,
					},
					{
						LatitudeDegree:    13.0,
						LongitudeDegree:   23.0,
						RelativeAltitudeM: 33.0,
						SpeedMS:           43.0,
					},
				},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundAltitudeM)
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	expectNav.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	expectNav.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	a.Equal(navigation, expectNav)
}
