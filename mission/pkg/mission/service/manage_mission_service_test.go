package service

import (
	m "mission/pkg/mission/domain/mission"
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
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
	a.Equal(resModel.GetMission().GetNavigation().GetTakeoffPointGroundHeight(), DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
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
					TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
					Waypoints:                               []waypointComponentMock{},
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
		DefaultMissionID1                                      = string(DefaultMissionID) + "-1"
		DefaultMissionName1                                    = DefaultMissionName + "-1"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM1 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 1
		DefaultMissionVersion1                                 = string(DefaultMissionVersion) + "-1"
		DefaultMissionID2                                      = string(DefaultMissionID) + "-2"
		DefaultMissionName2                                    = DefaultMissionName + "-2"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM2 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 2
		DefaultMissionVersion2                                 = string(DefaultMissionVersion) + "-2"
		DefaultMissionID3                                      = string(DefaultMissionID) + "-3"
		DefaultMissionName3                                    = DefaultMissionName + "-3"
		DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM3 = DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM + 3
		DefaultMissionVersion3                                 = string(DefaultMissionVersion) + "-3"
	)

	missions := []*m.Mission{
		m.AssembleFrom(
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
		m.AssembleFrom(
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
		m.AssembleFrom(
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                []waypointMock{},
			},
		},
	}
	var resCall bool
	ret := service.CreateMission(
		command,
		func(id string) {
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

func TestCreateMissionOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
		DefaultMissionVersion3 = DefaultMissionVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultMissionID,
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                []waypointMock{},
			},
		},
	}
	var resID string
	ret := service.createMissionOperation(
		nil,
		pub,
		command,
		func(id string) {
			resID = id
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultMissionID))
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
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2},
	}

	mission := m.AssembleFrom(
		gen,
		&missionComponentMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationComponentMock{
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                []waypointMock{},
			},
		},
	}
	ret := service.UpdateMission(command)

	a.Nil(ret)
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
		versions: []m.Version{DefaultMissionVersion1, DefaultMissionVersion2},
	}

	mission := m.AssembleFrom(
		gen,
		&missionComponentMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationComponentMock{
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
				TakeoffPointGroundHeight: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                []waypointMock{},
			},
		},
	}
	ret := service.updateMissionOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
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
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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

	gen := &generatorMock{}

	mission := m.AssembleFrom(
		gen,
		&missionComponentMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationComponentMock{
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
	ret := service.CarbonCopyMission(command)

	a.Nil(ret)
	a.Len(pub.events, 0)
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

	gen := &generatorMock{}

	mission := m.AssembleFrom(
		gen,
		&missionComponentMock{
			ID:   string(DefaultMissionID),
			Name: DefaultMissionName,
			Navigation: navigationComponentMock{
				TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
				Waypoints:                               []waypointComponentMock{},
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
	ret := service.carbonCopyMissionOperation(
		nil,
		pub,
		command,
	)

	a.Nil(ret)
	a.Len(pub.events, 0)
}

func TestNoWaypointMissionModel(t *testing.T) {
	a := assert.New(t)

	mission := m.AssembleFrom(
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
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundHeight(), DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
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
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{
		&waypoint{
			latitude:       11.0,
			longitude:      21.0,
			relativeHeight: 31.0,
			speed:          41.0,
		},
	}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundHeight(), DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
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
	)

	model := missionModel{
		mission: mission,
	}

	expectWps := []Waypoint{
		&waypoint{
			latitude:       11.0,
			longitude:      21.0,
			relativeHeight: 31.0,
			speed:          41.0,
		},
		&waypoint{
			latitude:       12.0,
			longitude:      22.0,
			relativeHeight: 32.0,
			speed:          42.0,
		},
		&waypoint{
			latitude:       13.0,
			longitude:      23.0,
			relativeHeight: 33.0,
			speed:          43.0,
		},
	}

	a.Equal(model.GetMission().GetID(), string(DefaultMissionID))
	a.Equal(model.GetMission().GetName(), DefaultMissionName)
	a.Equal(model.GetMission().GetNavigation().GetTakeoffPointGroundHeight(), DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM)
	a.Len(model.GetMission().GetNavigation().GetWaypoints(), 3)
	a.Equal(model.GetMission().GetNavigation().GetWaypoints(), expectWps)
}
