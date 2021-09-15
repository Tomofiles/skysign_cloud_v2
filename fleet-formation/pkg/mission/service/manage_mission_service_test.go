package service

import (
	m "fleet-formation/pkg/mission/domain/mission"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMissionTransaction(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
	)

	repo := &missionRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultMissionID).Return(mission, nil)

	service := &manageMissionService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	var resCall bool
	ret := service.GetMission(
		command,
		func(model MissionPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetMissionOperation(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
	)

	repo := &missionRepositoryMock{}
	repo.On("GetByID", DefaultMissionID).Return(mission, nil)

	service := &manageMissionService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	var resModel MissionPresentationModel
	ret := service.getMissionOperation(
		nil,
		command,
		func(model MissionPresentationModel) {
			resModel = model
		},
	)

	a.Nil(ret)
	a.Equal(resModel.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(resModel.GetMission().GetName(), DefaultMissionName)
	a.Equal(resModel.GetMission().GetNavigation().GetTakeoffPointGroundAltitudeM(), DefaultMissionTakeoffPointGroundAltitudeM)
	a.Equal(resModel.GetMission().GetNavigation().GetUploadID(), string(DefaultMissionUploadID))
	a.Len(resModel.GetMission().GetNavigation().GetWaypoints(), 0)
}

func TestListMissionsTransaction(t *testing.T) {
	a := assert.New(t)

	missions := []*m.Mission{
		m.AssembleFrom(
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

	repo := &missionRepositoryMock{}
	repo.On("GetAllOrigin").Return(missions, nil)
	txm := &txManagerMock{}

	service := &manageMissionService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListMissions(
		func(model MissionPresentationModel) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestListMissionsOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionID1                          = string(DefaultMissionID) + "-1"
		DefaultMissionName1                        = DefaultMissionName + "-1"
		DefaultMissionTakeoffPointGroundAltitudeM1 = DefaultMissionTakeoffPointGroundAltitudeM + 1
		DefaultMissionVersion1                     = string(DefaultMissionVersion) + "-1"
		DefaultMissionUploadID1                    = string(DefaultMissionUploadID) + "-1"
		DefaultMissionID2                          = string(DefaultMissionID) + "-2"
		DefaultMissionName2                        = DefaultMissionName + "-2"
		DefaultMissionTakeoffPointGroundAltitudeM2 = DefaultMissionTakeoffPointGroundAltitudeM + 2
		DefaultMissionVersion2                     = string(DefaultMissionVersion) + "-2"
		DefaultMissionUploadID2                    = string(DefaultMissionUploadID) + "-2"
		DefaultMissionID3                          = string(DefaultMissionID) + "-3"
		DefaultMissionName3                        = DefaultMissionName + "-3"
		DefaultMissionTakeoffPointGroundAltitudeM3 = DefaultMissionTakeoffPointGroundAltitudeM + 3
		DefaultMissionVersion3                     = string(DefaultMissionVersion) + "-3"
		DefaultMissionUploadID3                    = string(DefaultMissionUploadID) + "-3"
	)

	missions := []*m.Mission{
		m.AssembleFrom(
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
		m.AssembleFrom(
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
		m.AssembleFrom(
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
	}

	repo := &missionRepositoryMock{}
	repo.On("GetAllOrigin").Return(missions, nil)

	service := &manageMissionService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resModels []MissionPresentationModel
	ret := service.listMissionsOperation(
		nil,
		func(model MissionPresentationModel) {
			resModels = append(resModels, model)
		},
	)

	a.Nil(ret)
	a.Len(resModels, 3)
}

func TestCreateMissionTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
		DefaultMissionVersion3 = DefaultMissionVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultMissionID,
		uploadID: DefaultMissionUploadID,
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2, DefaultMissionVersion3},
	}
	repo := &missionRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &missionCommandMock{
		Mission: missionMock{
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints:                   []waypointMock{},
			},
		},
	}
	var resCall1, resCall2 bool
	ret := service.CreateMission(
		command,
		func(id string) {
			resCall1 = true
		},
		func(uploadID string) {
			resCall2 = true
		},
	)

	a.Nil(ret)
	a.True(resCall1)
	a.True(resCall2)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateMissionOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
		DefaultMissionVersion3 = DefaultMissionVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultMissionID,
		uploadID: DefaultMissionUploadID,
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2, DefaultMissionVersion3},
	}
	repo := &missionRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionCommandMock{
		Mission: missionMock{
			Name: DefaultMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints:                   []waypointMock{},
			},
		},
	}
	var resID, resUploadID string
	ret := service.createMissionOperation(
		nil,
		pub,
		command,
		func(id string) {
			resID = id
		},
		func(uploadID string) {
			resUploadID = uploadID
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultMissionID))
	a.Equal(resUploadID, string(DefaultMissionUploadID))
	a.Len(pub.events, 0)
}

func TestUpdateMissionTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		AfterMissionName       = DefaultMissionName + "-after"
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultMissionID,
		uploadID: DefaultMissionUploadID,
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2},
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultMissionID).Return(mission, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: AfterMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints:                   []waypointMock{},
			},
		},
	}
	var resCall bool
	ret := service.UpdateMission(
		command,
		func(uploadID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestUpdateMissionOperation(t *testing.T) {
	a := assert.New(t)

	var (
		AfterMissionName       = DefaultMissionName + "-after"
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultMissionID,
		uploadID: DefaultMissionUploadID,
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2},
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	repo.On("GetByID", DefaultMissionID).Return(mission, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageMissionService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionCommandMock{
		Mission: missionMock{
			ID:   string(DefaultMissionID),
			Name: AfterMissionName,
			Navigation: navigationMock{
				TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
				Waypoints:                   []waypointMock{},
			},
		},
	}
	var resUploadID string
	ret := service.updateMissionOperation(
		nil,
		pub,
		command,
		func(uploadID string) {
			resUploadID = uploadID
		},
	)

	a.Nil(ret)
	a.Equal(resUploadID, string(DefaultMissionUploadID))
	a.Len(pub.events, 0)
}

func TestDeleteMissionTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		versions: []m.Version{DefaultMissionVersion},
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultMissionID).Return(mission, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	ret := service.DeleteMission(command)

	a.Nil(ret)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestDeleteMissionOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		versions: []m.Version{DefaultMissionVersion},
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	repo.On("GetByID", DefaultMissionID).Return(mission, nil)
	repo.On("Delete", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &missionIDCommandMock{
		ID: string(DefaultMissionID),
	}
	ret := service.deleteMissionOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
	a.Len(pub.events, 0)
}

func TestCarbonCopyMissionTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultMissionID + "-original"
		DefaultNewID      = DefaultMissionID + "-new"
	)

	gen := &generatorMock{
		uploadID: DefaultMissionUploadID,
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultNewID).Return(nil, m.ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(mission, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &carbonCopyCommandMock{
		OriginalID: string(DefaultOriginalID),
		NewID:      string(DefaultNewID),
	}
	var resCall bool
	ret := service.CarbonCopyMission(
		command,
		func(uploadID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCarbonCopyMissionOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultMissionID + "-original"
		DefaultNewID      = DefaultMissionID + "-new"
	)

	gen := &generatorMock{
		uploadID: DefaultMissionUploadID,
	}

	mission := m.AssembleFrom(
		gen,
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
	)

	repo := &missionRepositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, m.ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(mission, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &carbonCopyCommandMock{
		OriginalID: string(DefaultOriginalID),
		NewID:      string(DefaultNewID),
	}
	var resUploadID string
	ret := service.carbonCopyMissionOperation(
		nil,
		pub,
		command,
		func(uploadID string) {
			resUploadID = uploadID
		},
	)

	a.Nil(ret)
	a.Equal(resUploadID, string(DefaultMissionUploadID))
	a.Len(pub.events, 1)
}

func TestNoWaypointMissionModel(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundAltitudeM(), DefaultMissionTakeoffPointGroundAltitudeM)
	a.Equal(model.GetMission().GetNavigation().GetUploadID(), string(DefaultMissionUploadID))
	a.Len(model.GetMission().GetNavigation().GetWaypoints(), 0)
	a.Equal(model.GetMission().GetNavigation().GetWaypoints(), expectWps)
}

func TestSingleWaypointMissionModel(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
				UploadID: string(DefaultMissionUploadID),
			},
			Version: string(DefaultMissionVersion),
		},
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{
		&waypoint{
			latitudeDegree:    11.0,
			longitudeDegree:   21.0,
			relativeAltitudeM: 31.0,
			speedMS:           41.0,
		},
	}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundAltitudeM(), DefaultMissionTakeoffPointGroundAltitudeM)
	a.Equal(model.GetMission().GetNavigation().GetUploadID(), string(DefaultMissionUploadID))
	a.Len(model.GetMission().GetNavigation().GetWaypoints(), 1)
	a.Equal(model.GetMission().GetNavigation().GetWaypoints(), expectWps)
}

func TestMultipleWaypointsMissionModel(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
				UploadID: string(DefaultMissionUploadID),
			},
			Version: string(DefaultMissionVersion),
		},
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{
		&waypoint{
			latitudeDegree:    11.0,
			longitudeDegree:   21.0,
			relativeAltitudeM: 31.0,
			speedMS:           41.0,
		},
		&waypoint{
			latitudeDegree:    12.0,
			longitudeDegree:   22.0,
			relativeAltitudeM: 32.0,
			speedMS:           42.0,
		},
		&waypoint{
			latitudeDegree:    13.0,
			longitudeDegree:   23.0,
			relativeAltitudeM: 33.0,
			speedMS:           43.0,
		},
	}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundAltitudeM(), DefaultMissionTakeoffPointGroundAltitudeM)
	a.Equal(model.GetMission().GetNavigation().GetUploadID(), string(DefaultMissionUploadID))
	a.Len(model.GetMission().GetNavigation().GetWaypoints(), 3)
	a.Equal(model.GetMission().GetNavigation().GetWaypoints(), expectWps)
}
