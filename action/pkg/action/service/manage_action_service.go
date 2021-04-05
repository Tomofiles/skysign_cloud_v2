package service

import (
	"action/pkg/action/domain/action"
	"action/pkg/action/domain/txmanager"
)

// ManageActionService .
type ManageActionService interface {
	CreateAction(requestDpo CreateActionRequestDpo) error
	GetTrajectory(requestDpo GetTrajectoryRequestDpo, responseEachDpo GetTrajectoryResponseDpo) error
}

// CreateActionRequestDpo .
type CreateActionRequestDpo interface {
	GetID() string
	GetCommunicationID() string
	GetFlightplanID() string
}

// GetTrajectoryRequestDpo .
type GetTrajectoryRequestDpo interface {
	GetID() string
}

// GetAssignmentsResponseDpo .
type GetTrajectoryResponseDpo = func(snapshot action.TelemetrySnapshot)

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
	requestDpo CreateActionRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.createActionOperation(
			tx,
			requestDpo,
		)
	})
}

func (s *manageActionService) createActionOperation(
	tx txmanager.Tx,
	requestDpo CreateActionRequestDpo,
) error {
	aAction := action.NewInstance(
		action.ID(requestDpo.GetID()),
		action.CommunicationID(requestDpo.GetCommunicationID()),
		action.FlightplanID(requestDpo.GetFlightplanID()),
	)
	if ret := s.repo.Save(tx, aAction); ret != nil {
		return ret
	}

	return nil
}

func (s *manageActionService) GetTrajectory(
	requestDpo GetTrajectoryRequestDpo,
	responseEachDpo GetTrajectoryResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getTrajectoryOperation(
			tx,
			requestDpo,
			responseEachDpo,
		)
	})
}

func (s *manageActionService) getTrajectoryOperation(
	tx txmanager.Tx,
	requestDpo GetTrajectoryRequestDpo,
	responseEachDpo GetTrajectoryResponseDpo,
) error {
	aAction, err := s.repo.GetByID(
		tx,
		action.ID(requestDpo.GetID()),
	)
	if err != nil {
		return err
	}

	aAction.ProvideTrajectoryInterest(
		responseEachDpo,
	)

	return nil
}
