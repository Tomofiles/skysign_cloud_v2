package service

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/fleet"
	f "flightplan/pkg/flightplan/domain/flightplan"
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
	gen  fleet.Generator
	repo fleet.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
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
		f.ID(requestDpo.GetFlightplanID()),
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
		f.ID(requestDpo.GetFlightplanID()),
	); ret != nil {
		return ret
	}

	return nil
}

func (s *manageFleetService) CarbonCopyFleet(
	requestDpo CarbonCopyFleetRequestDpo,
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
				requestDpo,
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
	requestDpo CarbonCopyFleetRequestDpo,
) error {
	if ret := fleet.CarbonCopyFleet(
		tx,
		s.gen,
		s.repo,
		pub,
		f.ID(requestDpo.GetOriginalID()),
		f.ID(requestDpo.GetNewID()),
	); ret != nil {
		return ret
	}

	return nil
}
