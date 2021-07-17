package service

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/fleet"
	f "flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFleetService .
type ManageFleetService interface {
	CreateFleet(command CreateFleetCommand) error
	DeleteFleet(command DeleteFleetCommand) error
	CarbonCopyFleet(command CarbonCopyFleetCommand) error
}

// CreateFleetCommand .
type CreateFleetCommand interface {
	GetID() string
	GetNumberOfVehicles() int
}

// DeleteFleetCommand .
type DeleteFleetCommand interface {
	GetID() string
}

// CarbonCopyFleetCommand .
type CarbonCopyFleetCommand interface {
	GetOriginalID() string
	GetNewID() string
}

// NewManageFleetService .
func NewManageFleetService(
	gen f.Generator,
	repo f.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFleetService {
	return &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageFleetService struct {
	gen  f.Generator
	repo f.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageFleetService) CreateFleet(
	command CreateFleetCommand,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.createFleetOperation(
			tx,
			command,
		)
	})
}

func (s *manageFleetService) createFleetOperation(
	tx txmanager.Tx,
	command CreateFleetCommand,
) error {
	fleet := f.NewInstance(
		s.gen,
		f.ID(command.GetID()),
		command.GetNumberOfVehicles())
	for _, assignmentID := range fleet.GetAllAssignmentID() {
		fleet.AddNewEvent(assignmentID)
	}
	if ret := s.repo.Save(tx, fleet); ret != nil {
		return ret
	}

	return nil
}

func (s *manageFleetService) DeleteFleet(
	command DeleteFleetCommand,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.deleteFleetOperation(
			tx,
			command,
		)
	})
}

func (s *manageFleetService) deleteFleetOperation(
	tx txmanager.Tx,
	command DeleteFleetCommand,
) error {
	if ret := s.repo.Delete(
		tx,
		f.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

func (s *manageFleetService) CarbonCopyFleet(
	command CarbonCopyFleetCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.carbonCopyFleetOperation(
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

func (s *manageFleetService) carbonCopyFleetOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CarbonCopyFleetCommand,
) error {
	if ret := fleet.CarbonCopyFleet(
		tx,
		s.gen,
		s.repo,
		pub,
		f.ID(command.GetOriginalID()),
		f.ID(command.GetNewID()),
	); ret != nil {
		return ret
	}

	return nil
}
