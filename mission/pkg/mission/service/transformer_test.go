package service

import (
	m "mission/pkg/mission/domain/mission"
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                []waypointMock{},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)

	a.Equal(navigation, expectNav)
}

func TestSingleNavigationTransformerFromCommand(t *testing.T) {
	a := assert.New(t)

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints: []waypointMock{
					{
						Latitude:       11.0,
						Longitude:      21.0,
						RelativeHeight: 31.0,
						Speed:          41.0,
					},
				},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints: []waypointMock{
					{
						Latitude:       11.0,
						Longitude:      21.0,
						RelativeHeight: 31.0,
						Speed:          41.0,
					},
					{
						Latitude:       12.0,
						Longitude:      22.0,
						RelativeHeight: 32.0,
						Speed:          42.0,
					},
					{
						Latitude:       13.0,
						Longitude:      23.0,
						RelativeHeight: 33.0,
						Speed:          43.0,
					},
				},
			},
		},
	}
	navigation := NavigationTransformerFromCommand(command.GetMission())

	expectNav := m.NewNavigation(DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	expectNav.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	expectNav.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	a.Equal(navigation, expectNav)
}
