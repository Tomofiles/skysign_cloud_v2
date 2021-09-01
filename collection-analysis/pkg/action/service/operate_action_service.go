package service

import (
	"collection-analysis/pkg/action/domain/action"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// OperateActionService .
type OperateActionService interface {
	CompleteAction(command CompleteActionCommand) error
	PushTelemetry(command PushTelemetryCommand) error
}

// CompleteActionCommand .
type CompleteActionCommand interface {
	GetFleetID() string
}

// PushTelemetryCommand .
type PushTelemetryCommand interface {
	GetCommunicationID() string
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
	GetRelativeAltitude() float64
	GetSpeed() float64
	GetArmed() bool
	GetFlightMode() string
	GetOrientationX() float64
	GetOrientationY() float64
	GetOrientationZ() float64
	GetOrientationW() float64
}

// NewOperateActionService .
func NewOperateActionService(
	repo action.Repository,
	txm txmanager.TransactionManager,
) OperateActionService {
	return &operateActionService{
		repo: repo,
		txm:  txm,
	}
}

type operateActionService struct {
	repo action.Repository
	txm  txmanager.TransactionManager
}

func (s *operateActionService) CompleteAction(
	command CompleteActionCommand,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.completeActionOperation(
				tx,
				command,
			)
		},
	)
}

func (s *operateActionService) completeActionOperation(
	tx txmanager.Tx,
	command CompleteActionCommand,
) error {
	actions, err := s.repo.GetAllActiveByFleetID(tx, action.FleetID(command.GetFleetID()))
	if err != nil {
		return err
	}

	for _, action := range actions {
		if err := action.Complete(); err != nil {
			return err
		}
		if err := s.repo.Save(tx, action); err != nil {
			return err
		}
	}

	return nil
}

func (s *operateActionService) PushTelemetry(
	command PushTelemetryCommand,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.pushTelemetryOperation(
				tx,
				command,
			)
		},
	)
}

func (s *operateActionService) pushTelemetryOperation(
	tx txmanager.Tx,
	command PushTelemetryCommand,
) error {
	aAction, err := s.repo.GetActiveByCommunicationID(tx, action.CommunicationID(command.GetCommunicationID()))
	if err != nil {
		return err
	}

	snapshot := action.TelemetrySnapshot{
		Latitude:         command.GetLatitude(),
		Longitude:        command.GetLongitude(),
		Altitude:         command.GetAltitude(),
		RelativeAltitude: command.GetRelativeAltitude(),
		Speed:            command.GetSpeed(),
		Armed:            command.GetArmed(),
		FlightMode:       command.GetFlightMode(),
		OrientationX:     command.GetOrientationX(),
		OrientationY:     command.GetOrientationY(),
		OrientationZ:     command.GetOrientationZ(),
		OrientationW:     command.GetOrientationW(),
	}

	if err := aAction.PushTelemetry(snapshot); err != nil {
		return err
	}
	if err := s.repo.Save(tx, aAction); err != nil {
		return err
	}

	return nil
}
