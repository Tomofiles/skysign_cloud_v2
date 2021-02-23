package service

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFleetService .
type ManageFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	txm  txmanager.TransactionManager
}

// NewManageFleetService .
func NewManageFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	txm txmanager.TransactionManager,
) ManageFleetService {
	return ManageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}
}

// CreateFleet .
func (s *ManageFleetService) CreateFleet(
	requestDpo CreateFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.createFleetOperation(
			tx,
			requestDpo,
		)
	})
}

func (s *ManageFleetService) createFleetOperation(
	tx txmanager.Tx,
	requestDpo CreateFleetRequestDpo,
) error {
	fleet := fleet.NewInstance(
		s.gen,
		flightplan.ID(requestDpo.GetFlightplanID()),
		0)
	if ret := s.repo.Save(tx, fleet); ret != nil {
		return ret
	}

	return nil
}

// DeleteFleet .
func (s *ManageFleetService) DeleteFleet(
	requestDpo DeleteFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.deleteFleetOperation(
			tx,
			requestDpo,
		)
	})
}

func (s *ManageFleetService) deleteFleetOperation(
	tx txmanager.Tx,
	requestDpo DeleteFleetRequestDpo,
) error {
	if ret := s.repo.DeleteByFlightplanID(
		tx,
		flightplan.ID(requestDpo.GetFlightplanID()),
	); ret != nil {
		return ret
	}

	return nil
}

// CreateFleetRequestDpo .
type CreateFleetRequestDpo interface {
	GetFlightplanID() string
}

// DeleteFleetRequestDpo .
type DeleteFleetRequestDpo interface {
	GetFlightplanID() string
}
