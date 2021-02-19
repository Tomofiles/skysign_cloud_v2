package service

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFleetService .
type ManageFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	pub  event.Publisher
	txm  txmanager.TransactionManager
}

// NewManageFleetService .
func NewManageFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	pub event.Publisher,
	txm txmanager.TransactionManager,
) ManageFleetService {
	return ManageFleetService{
		gen:  gen,
		repo: repo,
		pub:  pub,
		txm:  txm,
	}
}

// CreateFleet .
func (s *ManageFleetService) CreateFleet(
	requestDpo CreateFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		fleet := fleet.NewInstance(
			s.gen,
			flightplan.ID(requestDpo.GetFlightplanID()),
			0)
		ret := s.repo.Save(tx, fleet)
		if ret != nil {
			return ret
		}

		return nil
	})
}

// DeleteFleet .
func (s *ManageFleetService) DeleteFleet(
	requestDpo DeleteFleetRequestDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		ret := s.repo.DeleteByFlightplanID(tx, flightplan.ID(requestDpo.GetFlightplanID()))
		if ret != nil {
			return ret
		}

		return nil
	})
}

// CreateFleetRequestDpo .
type CreateFleetRequestDpo interface {
	GetFlightplanID() string
}

// DeleteFleetRequestDpo .
type DeleteFleetRequestDpo interface {
	GetFlightplanID() string
}
