package ports

import (
	"mission/pkg/mission/app"
	m "mission/pkg/mission/domain/mission"
	s "mission/pkg/mission/service"
	proto "mission/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleMissionsListMissions(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

	missionModels := []s.MissionPresentationModel{
		&missionModelMock{
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
		},
	}
	service.On("ListMissions", mock.Anything).Return(missionModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.Empty{}
	response, err := grpc.ListMissions(
		nil,
		request,
	)

	expectResponse := &proto.ListMissionsResponses{
		Missions: []*proto.Mission{
			{
				Id:   DefaultMissionID,
				Name: DefaultMissionName,
				Navigation: &proto.Navigation{
					TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
					Waypoints:                []*proto.Waypoint{},
				},
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleMissionsListMissions(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionID1                                      = DefaultMissionID + "-1"
		DefaultMissionName1                                    = DefaultMissionName + "-1"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM1 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 1
		DefaultMissionVersion1                                 = DefaultMissionVersion + "-1"
		DefaultMissionID2                                      = DefaultMissionID + "-2"
		DefaultMissionName2                                    = DefaultMissionName + "-2"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM2 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 2
		DefaultMissionVersion2                                 = DefaultMissionVersion + "-2"
		DefaultMissionID3                                      = DefaultMissionID + "-3"
		DefaultMissionName3                                    = DefaultMissionName + "-3"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM3 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 3
		DefaultMissionVersion3                                 = DefaultMissionVersion + "-3"
	)

	service := manageMissionServiceMock{}

	missionModels := []s.MissionPresentationModel{
		&missionModelMock{
			mission: m.AssembleFrom(
				nil,
				&missionComponentMock{
					ID:   string(DefaultMissionID1),
					Name: DefaultMissionName1,
					Navigation: navigationComponentMock{
						TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM1,
						Waypoints:                               []waypointComponentMock{},
					},
					Version: string(DefaultMissionVersion1),
				},
			),
		},
		&missionModelMock{
			mission: m.AssembleFrom(
				nil,
				&missionComponentMock{
					ID:   string(DefaultMissionID2),
					Name: DefaultMissionName2,
					Navigation: navigationComponentMock{
						TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM2,
						Waypoints:                               []waypointComponentMock{},
					},
					Version: string(DefaultMissionVersion2),
				},
			),
		},
		&missionModelMock{
			mission: m.AssembleFrom(
				nil,
				&missionComponentMock{
					ID:   string(DefaultMissionID3),
					Name: DefaultMissionName3,
					Navigation: navigationComponentMock{
						TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM3,
						Waypoints:                               []waypointComponentMock{},
					},
					Version: string(DefaultMissionVersion3),
				},
			),
		},
	}
	service.On("ListMissions", mock.Anything).Return(missionModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.Empty{}
	response, err := grpc.ListMissions(
		nil,
		request,
	)

	expectResponse := &proto.ListMissionsResponses{
		Missions: []*proto.Mission{
			{
				Id:   DefaultMissionID1,
				Name: DefaultMissionName1,
				Navigation: &proto.Navigation{
					TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM1,
					Waypoints:                []*proto.Waypoint{},
				},
			},
			{
				Id:   DefaultMissionID2,
				Name: DefaultMissionName2,
				Navigation: &proto.Navigation{
					TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM2,
					Waypoints:                []*proto.Waypoint{},
				},
			},
			{
				Id:   DefaultMissionID3,
				Name: DefaultMissionName3,
				Navigation: &proto.Navigation{
					TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM3,
					Waypoints:                []*proto.Waypoint{},
				},
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneMissionsListMissions(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

	missionModels := []s.MissionPresentationModel{}
	service.On("ListMissions", mock.Anything).Return(missionModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.Empty{}
	response, err := grpc.ListMissions(
		nil,
		request,
	)

	expectResponse := &proto.ListMissionsResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetMission(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

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
	service.On("GetMission", mock.Anything, mock.Anything).Return(missionModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.GetMissionRequest{
		Id: DefaultMissionID,
	}
	response, err := grpc.GetMission(
		nil,
		request,
	)

	expectResponse := &proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionName,
		Navigation: &proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*proto.Waypoint{},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateMission(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

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
	service.On("CreateMission", mock.Anything, mock.Anything).Return(missionModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.Mission{
		Name: DefaultMissionID,
		Navigation: &proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*proto.Waypoint{},
		},
	}
	response, err := grpc.CreateMission(
		nil,
		request,
	)

	expectResponse := &proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionID,
		Navigation: &proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*proto.Waypoint{},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateMission(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

	service.On("UpdateMission", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionID,
		Navigation: &proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*proto.Waypoint{},
		},
	}
	response, err := grpc.UpdateMission(
		nil,
		request,
	)

	expectResponse := &proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionID,
		Navigation: &proto.Navigation{
			TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
			Waypoints:                []*proto.Waypoint{},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestDeleteMission(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

	service.On("DeleteMission", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.DeleteMissionRequest{
		Id: DefaultMissionID,
	}
	response, err := grpc.DeleteMission(
		nil,
		request,
	)

	expectResponse := &proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
