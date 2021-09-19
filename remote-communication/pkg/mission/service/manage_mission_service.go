package service

import (
	m "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageMissionService .
type ManageMissionService interface {
	CreateMission(command CreateMissionCommand) error
}

// CreateMissionCommand .
type CreateMissionCommand interface {
	GetID() string
	GetWaypoints() []Waypoint
}

// Waypoint .
type Waypoint interface {
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
}

// NewManageMissionService .
func NewManageMissionService(
	repo m.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageMissionService {
	return &manageMissionService{
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageMissionService struct {
	repo m.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageMissionService) CreateMission(
	command CreateMissionCommand,
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
) error {
	mission := m.NewInstance(m.ID(command.GetID()))

	for _, w := range command.GetWaypoints() {
		mission.PushWaypoint(
			w.GetLatitudeDegree(),
			w.GetLongitudeDegree(),
			w.GetRelativeAltitudeM(),
			w.GetSpeedMS(),
		)
	}

	if ret := s.repo.Save(tx, mission); ret != nil {
		return ret
	}

	return nil
}
