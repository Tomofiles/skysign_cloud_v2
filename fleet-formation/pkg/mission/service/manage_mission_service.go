package service

import (
	m "fleet-formation/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageMissionService .
type ManageMissionService interface {
	GetMission(command GetMissionCommand, retrievedModel RetrievedModel) error
	ListMissions(retrievedModel RetrievedModel) error
	CreateMission(command CreateMissionCommand, createdID CreatedID, uploadID UploadID) error
	UpdateMission(command UpdateMissionCommand, uploadID UploadID) error
	DeleteMission(command DeleteMissionCommand) error
	CarbonCopyMission(command CarbonCopyMissionCommand, uploadID UploadID) error
}

// CreateMissionCommand .
type CreateMissionCommand interface {
	GetMission() Mission
}

// UpdateMissionCommand .
type UpdateMissionCommand interface {
	GetID() string
	GetMission() Mission
}

// GetMissionCommand .
type GetMissionCommand interface {
	GetID() string
}

// DeleteMissionCommand .
type DeleteMissionCommand interface {
	GetID() string
}

// CarbonCopyMissionCommand .
type CarbonCopyMissionCommand interface {
	GetOriginalID() string
	GetNewID() string
}

// MissionPresentationModel .
type MissionPresentationModel interface {
	GetMission() Mission
}

// Mission .
type Mission interface {
	GetID() string
	GetName() string
	GetNavigation() Navigation
}

// Navigation .
type Navigation interface {
	GetTakeoffPointGroundAltitudeM() float64
	GetWaypoints() []Waypoint
	GetUploadID() string
}

// Waypoint .
type Waypoint interface {
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
}

// CreatedID .
type CreatedID = func(id string)

// UploadID .
type UploadID = func(uploadID string)

// RetrievedModel .
type RetrievedModel = func(model MissionPresentationModel)

// NewManageMissionService .
func NewManageMissionService(
	gen m.Generator,
	repo m.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageMissionService {
	return &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageMissionService struct {
	gen  m.Generator
	repo m.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageMissionService) GetMission(
	command GetMissionCommand,
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getMissionOperation(
			tx,
			command,
			retrievedModel,
		)
	})
}

func (s *manageMissionService) getMissionOperation(
	tx txmanager.Tx,
	command GetMissionCommand,
	retrievedModel RetrievedModel,
) error {
	mission, err := s.repo.GetByID(tx, m.ID(command.GetID()))
	if err != nil {
		return err
	}

	retrievedModel(
		&missionModel{
			mission: mission,
		},
	)
	return nil
}

func (s *manageMissionService) ListMissions(
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listMissionsOperation(
			tx,
			retrievedModel,
		)
	})
}

func (s *manageMissionService) listMissionsOperation(
	tx txmanager.Tx,
	retrievedModel RetrievedModel,
) error {
	missions, err := s.repo.GetAllOrigin(tx)
	if err != nil {
		return err
	}

	for _, mission := range missions {
		retrievedModel(
			&missionModel{
				mission: mission,
			},
		)
	}
	return nil
}

func (s *manageMissionService) CreateMission(
	command CreateMissionCommand,
	createdID CreatedID,
	uploadID UploadID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createMissionOperation(
				tx,
				pub,
				command,
				createdID,
				uploadID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageMissionService) createMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CreateMissionCommand,
	createdID CreatedID,
	uploadID UploadID,
) error {
	navigation := NavigationTransformerFromCommand(command.GetMission())
	id, uID, ret := m.CreateNewMission(
		tx,
		s.gen,
		s.repo,
		pub,
		command.GetMission().GetName(),
		navigation,
	)
	if ret != nil {
		return ret
	}

	createdID(id)
	uploadID(uID)
	return nil
}

func (s *manageMissionService) UpdateMission(
	command UpdateMissionCommand,
	uploadID UploadID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.updateMissionOperation(
				tx,
				pub,
				command,
				uploadID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageMissionService) updateMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command UpdateMissionCommand,
	uploadID UploadID,
) error {
	navigation := NavigationTransformerFromCommand(command.GetMission())
	uID, ret := m.UpdateMission(
		tx,
		s.gen,
		s.repo,
		pub,
		m.ID(command.GetID()),
		command.GetMission().GetName(),
		navigation,
	)
	if ret != nil {
		return ret
	}

	uploadID(uID)
	return nil
}

func (s *manageMissionService) DeleteMission(
	command DeleteMissionCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteMissionOperation(
				tx,
				pub,
				command,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageMissionService) deleteMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command DeleteMissionCommand,
) error {
	if ret := m.DeleteMission(
		tx,
		s.repo,
		pub,
		m.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

func (s *manageMissionService) CarbonCopyMission(
	command CarbonCopyMissionCommand,
	uploadID UploadID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.carbonCopyMissionOperation(
				tx,
				pub,
				command,
				uploadID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageMissionService) carbonCopyMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CarbonCopyMissionCommand,
	uploadID UploadID,
) error {
	id, ret := m.CarbonCopyMission(
		tx,
		s.gen,
		s.repo,
		pub,
		m.ID(command.GetOriginalID()),
		m.ID(command.GetNewID()),
	)
	if ret != nil {
		return ret
	}

	uploadID(id)
	return nil
}

type missionModel struct {
	mission *m.Mission
}

func (f *missionModel) GetMission() Mission {
	return &mission{
		mission: f.mission,
	}
}

type mission struct {
	mission *m.Mission
}

func (f *mission) GetID() string {
	return string(f.mission.GetID())
}

func (f *mission) GetName() string {
	return f.mission.GetName()
}

func (f *mission) GetNavigation() Navigation {
	waypoints := []waypoint{}
	f.mission.GetNavigation().ProvideWaypointsInterest(
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64) {
			waypoints = append(
				waypoints,
				waypoint{
					latitudeDegree:    latitudeDegree,
					longitudeDegree:   longitudeDegree,
					relativeAltitudeM: relativeAltitudeM,
					speedMS:           speedMS,
				},
			)
		},
	)
	navigation := &navigation{
		takeoffPointGroundAltitudeM: f.mission.GetNavigation().GetTakeoffPointGroundAltitudeM(),
		waypoints:                   waypoints,
		uploadID:                    string(f.mission.GetNavigation().GetUploadID()),
	}
	return navigation
}

type navigation struct {
	takeoffPointGroundAltitudeM float64
	waypoints                   []waypoint
	uploadID                    string
}

func (f *navigation) GetTakeoffPointGroundAltitudeM() float64 {
	return f.takeoffPointGroundAltitudeM
}

func (f *navigation) GetWaypoints() []Waypoint {
	waypoints := []Waypoint{}
	for _, w := range f.waypoints {
		waypoints = append(
			waypoints,
			&waypoint{
				latitudeDegree:    w.latitudeDegree,
				longitudeDegree:   w.longitudeDegree,
				relativeAltitudeM: w.relativeAltitudeM,
				speedMS:           w.speedMS,
			},
		)
	}
	return waypoints
}

func (f *navigation) GetUploadID() string {
	return string(f.uploadID)
}

type waypoint struct {
	latitudeDegree    float64
	longitudeDegree   float64
	relativeAltitudeM float64
	speedMS           float64
}

func (f *waypoint) GetLatitudeDegree() float64 {
	return f.latitudeDegree
}

func (f *waypoint) GetLongitudeDegree() float64 {
	return f.longitudeDegree
}

func (f *waypoint) GetRelativeAltitudeM() float64 {
	return f.relativeAltitudeM
}

func (f *waypoint) GetSpeedMS() float64 {
	return f.speedMS
}
