package grpc

import (
	"testing"

	m "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
)

func TestNoWaypointMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionID = NewMissionID()
	)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                   []waypointComponentMock{},
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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*skysign_proto.Waypoint{},
		},
	}

	a.Equal(proto, expectProto)
}

func TestSingleWaypointMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionID = NewMissionID()
	)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints: []waypointComponentMock{
						{
							PointOrder:        1,
							LatitudeDegree:    11.0,
							LongitudeDegree:   21.0,
							RelativeAltitudeM: 31.0,
							SpeedMS:           41.0,
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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:         11.0,
					Longitude:        21.0,
					RelativeAltitude: 31.0,
					Speed:            41.0,
				},
			},
		},
	}

	a.Equal(proto, expectProto)
}

func TestMultipleWaypointsMissionProtoTransformerFromModel(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionID = NewMissionID()
	)

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints: []waypointComponentMock{
						{
							PointOrder:        1,
							LatitudeDegree:    11.0,
							LongitudeDegree:   21.0,
							RelativeAltitudeM: 31.0,
							SpeedMS:           41.0,
						},
						{
							PointOrder:        2,
							LatitudeDegree:    12.0,
							LongitudeDegree:   22.0,
							RelativeAltitudeM: 32.0,
							SpeedMS:           42.0,
						},
						{
							PointOrder:        3,
							LatitudeDegree:    13.0,
							LongitudeDegree:   23.0,
							RelativeAltitudeM: 33.0,
							SpeedMS:           43.0,
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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints: []*skysign_proto.Waypoint{
				{
					Latitude:         11.0,
					Longitude:        21.0,
					RelativeAltitude: 31.0,
					Speed:            41.0,
				},
				{
					Latitude:         12.0,
					Longitude:        22.0,
					RelativeAltitude: 32.0,
					Speed:            42.0,
				},
				{
					Latitude:         13.0,
					Longitude:        23.0,
					RelativeAltitude: 33.0,
					Speed:            43.0,
				},
			},
		},
	}

	a.Equal(proto, expectProto)
}
