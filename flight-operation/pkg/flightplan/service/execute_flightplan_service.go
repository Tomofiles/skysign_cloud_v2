package service

import (
	f "flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ExecuteFlightplanService .
type ExecuteFlightplanService interface {
	ExecuteFlightplan(command ExecuteFlightplanCommand) error
}

// ExecuteFlightplanCommand .
type ExecuteFlightplanCommand interface {
	GetID() string
}

// NewExecuteFlightplanService .
func NewExecuteFlightplanService(
	gen f.Generator,
	repo f.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ExecuteFlightplanService {
	return &executeFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type executeFlightplanService struct {
	gen  f.Generator
	repo f.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *executeFlightplanService) ExecuteFlightplan(
	command ExecuteFlightplanCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.executeFlightplanOperation(
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

func (s *executeFlightplanService) executeFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command ExecuteFlightplanCommand,
) error {
	if ret := f.ExecuteFlightplan(
		tx,
		s.repo,
		pub,
		f.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}
