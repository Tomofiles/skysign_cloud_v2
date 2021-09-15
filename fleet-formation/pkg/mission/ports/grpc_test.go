package ports

import (
	"fleet-formation/pkg/mission/app"
	m "fleet-formation/pkg/mission/domain/mission"
	s "fleet-formation/pkg/mission/service"
	"testing"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

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
						TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
						Waypoints:                   []waypointComponentMock{},
						UploadID:                    string(DefaultMissionUploadID),
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
					TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                  []*proto.Waypoint{},
					UploadId:                   string(DefaultMissionUploadID),
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
		DefaultMissionID1                          = DefaultMissionID + "-1"
		DefaultMissionName1                        = DefaultMissionName + "-1"
		DefaultMissionTakeoffPointGroundAltitudeM1 = DefaultMissionTakeoffPointGroundAltitudeM + 1
		DefaultMissionUploadID1                    = DefaultMissionUploadID + "-1"
		DefaultMissionVersion1                     = DefaultMissionVersion + "-1"
		DefaultMissionID2                          = DefaultMissionID + "-2"
		DefaultMissionName2                        = DefaultMissionName + "-2"
		DefaultMissionTakeoffPointGroundAltitudeM2 = DefaultMissionTakeoffPointGroundAltitudeM + 2
		DefaultMissionUploadID2                    = DefaultMissionUploadID + "-2"
		DefaultMissionVersion2                     = DefaultMissionVersion + "-2"
		DefaultMissionID3                          = DefaultMissionID + "-3"
		DefaultMissionName3                        = DefaultMissionName + "-3"
		DefaultMissionTakeoffPointGroundAltitudeM3 = DefaultMissionTakeoffPointGroundAltitudeM + 3
		DefaultMissionUploadID3                    = DefaultMissionUploadID + "-3"
		DefaultMissionVersion3                     = DefaultMissionVersion + "-3"
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
						TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM1,
						Waypoints:                   []waypointComponentMock{},
						UploadID:                    string(DefaultMissionUploadID1),
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
						TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM2,
						Waypoints:                   []waypointComponentMock{},
						UploadID:                    string(DefaultMissionUploadID2),
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
						TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM3,
						Waypoints:                   []waypointComponentMock{},
						UploadID:                    string(DefaultMissionUploadID3),
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
					TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM1,
					Waypoints:                  []*proto.Waypoint{},
					UploadId:                   string(DefaultMissionUploadID1),
				},
			},
			{
				Id:   DefaultMissionID2,
				Name: DefaultMissionName2,
				Navigation: &proto.Navigation{
					TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM2,
					Waypoints:                  []*proto.Waypoint{},
					UploadId:                   string(DefaultMissionUploadID2),
				},
			},
			{
				Id:   DefaultMissionID3,
				Name: DefaultMissionName3,
				Navigation: &proto.Navigation{
					TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM3,
					Waypoints:                  []*proto.Waypoint{},
					UploadId:                   string(DefaultMissionUploadID3),
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
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                   []waypointComponentMock{},
					UploadID:                    string(DefaultMissionUploadID),
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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*proto.Waypoint{},
			UploadId:                   string(DefaultMissionUploadID),
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
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                   []waypointComponentMock{},
					UploadID:                    string(DefaultMissionUploadID),
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
		Name: DefaultMissionName,
		Navigation: &proto.Navigation{
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*proto.Waypoint{},
		},
	}
	response, err := grpc.CreateMission(
		nil,
		request,
	)

	expectResponse := &proto.Mission{
		Id:   DefaultMissionID,
		Name: DefaultMissionName,
		Navigation: &proto.Navigation{
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*proto.Waypoint{},
			UploadId:                   string(DefaultMissionUploadID),
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateMission(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                   []waypointComponentMock{},
					UploadID:                    string(DefaultMissionUploadID),
				},
				Version: string(DefaultMissionVersion),
			},
		),
	}
	service.On("UpdateMission", mock.Anything, mock.Anything).Return(missionModel, nil)

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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*proto.Waypoint{},
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
			TakeoffPointGroundAltitude: DefaultMissionTakeoffPointGroundAltitudeM,
			Waypoints:                  []*proto.Waypoint{},
			UploadId:                   string(DefaultMissionUploadID),
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
