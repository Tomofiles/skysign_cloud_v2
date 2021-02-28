package service

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFleetService .
type ManageFleetService interface {
	CreateFleet(requestDpo CreateFleetRequestDpo) error
	DeleteFleet(requestDpo DeleteFleetRequestDpo) error
	CarbonCopyFleet(requestDpo CarbonCopyFleetRequestDpo) error
}

// CreateFleetRequestDpo .
type CreateFleetRequestDpo interface {
	GetFlightplanID() string
}

// DeleteFleetRequestDpo .
type DeleteFleetRequestDpo interface {
	GetFlightplanID() string
}

// CarbonCopyFleetRequestDpo .
type CarbonCopyFleetRequestDpo interface {
	GetOriginalID() string
	GetNewID() string
}

// NewManageFleetService .
func NewManageFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	txm txmanager.TransactionManager,
) ManageFleetService {
	return &manageFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}
}

type manageFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	txm  txmanager.TransactionManager
}

func (s *manageFleetService) CreateFleet(
	requestDpo CreateFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.createFleetOperation(
			tx,
			requestDpo,
		)
	})
}

func (s *manageFleetService) createFleetOperation(
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

func (s *manageFleetService) DeleteFleet(
	requestDpo DeleteFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.deleteFleetOperation(
			tx,
			requestDpo,
		)
	})
}

func (s *manageFleetService) deleteFleetOperation(
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

func (s *manageFleetService) CarbonCopyFleet(
	requestDpo CarbonCopyFleetRequestDpo,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.carbonCopyFleetOperation(
				tx,
				requestDpo,
			)
		},
	)
}

func (s *manageFleetService) carbonCopyFleetOperation(
	tx txmanager.Tx,
	requestDpo CarbonCopyFleetRequestDpo,
) error {
	if ret := fleet.CarbonCopyFleet(
		tx,
		s.gen,
		s.repo,
		flightplan.ID(requestDpo.GetOriginalID()),
		flightplan.ID(requestDpo.GetNewID()),
	); ret != nil {
		return ret
	}

	return nil
}
