package ports

import (
	m "mission/pkg/mission/domain/mission"
	"mission/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoWaypointMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
					Waypoints:                               []waypointComponentMock{},
				},
				Version: string(DefaultMissionVersion),
			},
		),
	}
	proto := MissionProtoTransformerFromModel(missionModel)

	expectProto := &skysign_proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionName,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*skysign_proto.Waypoint{},
		},
	}

	a.Equal(proto, expectProto)
}

func TestSingleWaypointMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
					Waypoints: []waypointComponentMock{
						{
							PointOrder:      1,
							LatitudeDegree:  11.0,
							LongitudeDegree: 21.0,
							RelativeHeightM: 31.0,
							SpeedMS:         41.0,
						},
					},
				},
				Version: string(DefaultMissionVersion),
			},
		),
	}
	proto := MissionProtoTransformerFromModel(missionModel)

	expectProto := &skysign_proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionName,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:       11.0,
					Longitude:      21.0,
					RelativeHeight: 31.0,
					Speed:          41.0,
				},
			},
		},
	}

	a.Equal(proto, expectProto)
}

func TestMultipleWaypointsMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
					Waypoints: []waypointComponentMock{
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
					},
				},
				Version: string(DefaultMissionVersion),
			},
		),
	}
	proto := MissionProtoTransformerFromModel(missionModel)

	expectProto := &skysign_proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionName,
		Navigation: &skysign_proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints: []*skysign_proto.Waypoint{
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
	}

	a.Equal(proto, expectProto)
}
