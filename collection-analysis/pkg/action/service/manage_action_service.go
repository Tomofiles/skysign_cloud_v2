package service

import (
	"collection-analysis/pkg/action/domain/action"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageActionService .
type ManageActionService interface {
	CreateAction(command CreateActionCommand) error
	GetTrajectory(command GetTrajectoryCommand, telemetry TelemetrySnapshot) error
}

// CreateActionCommand .
type CreateActionCommand interface {
	GetID() string
	GetCommunicationID() string
	GetFleetID() string
}

// GetTrajectoryCommand .
type GetTrajectoryCommand interface {
	GetID() string
}

// TelemetrySnapshot .
type TelemetrySnapshot = func(snapshot action.TelemetrySnapshot)

// NewManageActionService .
func NewManageActionService(
	repo action.Repository,
	txm txmanager.TransactionManager,
) ManageActionService {
	return &manageActionService{
		repo: repo,
		txm:  txm,
	}
}

type manageActionService struct {
	repo action.Repository
	txm  txmanager.TransactionManager
}

func (s *manageActionService) CreateAction(
	command CreateActionCommand,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.createActionOperation(
			tx,
			command,
		)
	})
}

func (s *manageActionService) createActionOperation(
	tx txmanager.Tx,
	command CreateActionCommand,
) error {
	aAction := action.NewInstance(
		action.ID(command.GetID()),
		action.CommunicationID(command.GetCommunicationID()),
		action.FleetID(command.GetFleetID()),
	)
	if ret := s.repo.Save(tx, aAction); ret != nil {
		return ret
	}

	return nil
}

func (s *manageActionService) GetTrajectory(
	command GetTrajectoryCommand,
	telemetry TelemetrySnapshot,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getTrajectoryOperation(
			tx,
			command,
			telemetry,
		)
	})
}

func (s *manageActionService) getTrajectoryOperation(
	tx txmanager.Tx,
	command GetTrajectoryCommand,
	telemetry TelemetrySnapshot,
) error {
	aAction, err := s.repo.GetByID(
		tx,
		action.ID(command.GetID()),
	)
	if err != nil {
		return err
	}

	aAction.ProvideTrajectoryInterest(
		telemetry,
	)

	return nil
}
