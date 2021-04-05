package service

import (
	"action/pkg/action/domain/action"
	"action/pkg/action/domain/txmanager"
)

// OperateActionService .
type OperateActionService interface {
	CompleteAction(requestDpo CompleteActionRequestDpo) error
	PushTelemetry(requestDpo PushTelemetryRequestDpo) error
}

// CompleteActionRequestDpo .
type CompleteActionRequestDpo interface {
	GetFlightplanID() string
}

// PushTelemetryRequestDpo .
type PushTelemetryRequestDpo interface {
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
	requestDpo CompleteActionRequestDpo,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.completeActionOperation(
				tx,
				requestDpo,
			)
		},
	)
}

func (s *operateActionService) completeActionOperation(
	tx txmanager.Tx,
	requestDpo CompleteActionRequestDpo,
) error {
	actions, err := s.repo.GetAllActiveByFlightplanID(tx, action.FlightplanID(requestDpo.GetFlightplanID()))
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
	requestDpo PushTelemetryRequestDpo,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.pushTelemetryOperation(
				tx,
				requestDpo,
			)
		},
	)
}

func (s *operateActionService) pushTelemetryOperation(
	tx txmanager.Tx,
	requestDpo PushTelemetryRequestDpo,
) error {
	aAction, err := s.repo.GetActiveByCommunicationID(tx, action.CommunicationID(requestDpo.GetCommunicationID()))
	if err != nil {
		return err
	}

	snapshot := action.TelemetrySnapshot{
		Latitude:         requestDpo.GetLatitude(),
		Longitude:        requestDpo.GetLongitude(),
		Altitude:         requestDpo.GetAltitude(),
		RelativeAltitude: requestDpo.GetRelativeAltitude(),
		Speed:            requestDpo.GetSpeed(),
		Armed:            requestDpo.GetArmed(),
		FlightMode:       requestDpo.GetFlightMode(),
		OrientationX:     requestDpo.GetOrientationX(),
		OrientationY:     requestDpo.GetOrientationY(),
		OrientationZ:     requestDpo.GetOrientationZ(),
		OrientationW:     requestDpo.GetOrientationW(),
	}

	if err := aAction.PushTelemetry(snapshot); err != nil {
		return err
	}
	if err := s.repo.Save(tx, aAction); err != nil {
		return err
	}

	return nil
}
