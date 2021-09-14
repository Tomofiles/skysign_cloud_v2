package service

import (
	m "remote-communication/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// EdgeMissionService .
type EdgeMissionService interface {
	PullMission(command PullMissionCommand, pulledMission PulledMission) error
}

// PullMissionCommand .
type PullMissionCommand interface {
	GetID() string
}

// PulledMission .
type PulledMission = func(id string, waypoints []Waypoint)

// NewEdgeMissionService .
func NewEdgeMissionService(
	repo m.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) EdgeMissionService {
	return &edgeMissionService{
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type edgeMissionService struct {
	repo m.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *edgeMissionService) PullMission(
	command PullMissionCommand,
	pulledMission PulledMission,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pullMissionOperation(
				tx,
				pub,
				command,
				pulledMission,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *edgeMissionService) pullMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PullMissionCommand,
	pulledMission PulledMission,
) error {
	mission, err := s.repo.GetByID(tx, m.ID(command.GetID()))
	if err != nil {
		return err
	}

	var waypoints []Waypoint
	for _, w := range mission.GetWaypoints() {
		waypoints = append(
			waypoints,
			&waypoint{
				PointOrder:      w.PointOrder,
				LatitudeDegree:  w.LatitudeDegree,
				LongitudeDegree: w.LongitudeDegree,
				RelativeHeightM: w.RelativeHeightM,
				SpeedMS:         w.SpeedMS,
			},
		)
	}

	pulledMission(
		string(mission.GetID()),
		waypoints,
	)
	return nil
}

type waypoint struct {
	PointOrder                                                int
	LatitudeDegree, LongitudeDegree, RelativeHeightM, SpeedMS float64
}

func (v *waypoint) GetPointOrder() int {
	return v.PointOrder
}

func (v *waypoint) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

func (v *waypoint) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

func (v *waypoint) GetRelativeHeightM() float64 {
	return v.RelativeHeightM
}

func (v *waypoint) GetSpeedMS() float64 {
	return v.SpeedMS
}
