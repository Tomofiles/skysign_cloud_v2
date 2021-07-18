package service

import (
	"flight-operation/pkg/common/domain/event"
	"flight-operation/pkg/common/domain/txmanager"
	f "flight-operation/pkg/flightplan/domain/flightplan"
)

// ChangeFlightplanService .
type ChangeFlightplanService interface {
	ChangeNumberOfVehicles(command ChangeNumberOfVehiclesCommand) error
}

// ChangeNumberOfVehiclesCommand .
type ChangeNumberOfVehiclesCommand interface {
	GetID() string
	GetNumberOfVehicles() int
}

// NewChangeFlightplanService .
func NewChangeFlightplanService(
	gen f.Generator,
	repo f.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ChangeFlightplanService {
	return &changeFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type changeFlightplanService struct {
	gen  f.Generator
	repo f.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *changeFlightplanService) ChangeNumberOfVehicles(
	command ChangeNumberOfVehiclesCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.changeNumberOfVehiclesOperation(
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

func (s *changeFlightplanService) changeNumberOfVehiclesOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command ChangeNumberOfVehiclesCommand,
) error {
	if ret := f.ChangeNumberOfVehicles(
		tx,
		s.gen,
		s.repo,
		pub,
		f.ID(command.GetID()),
		command.GetNumberOfVehicles(),
	); ret != nil {
		return ret
	}

	return nil
}
