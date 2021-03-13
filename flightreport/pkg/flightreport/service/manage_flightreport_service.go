package service

import (
	"flightreport/pkg/flightreport/domain/event"
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/flightreport/domain/txmanager"
)

// ManageFlightreportService .
type ManageFlightreportService interface {
	GetFlightreport(requestDpo GetFlightreportRequestDpo, responseDpo GetFlightreportResponseDpo) error
	ListFlightreports(responseEachDpo ListFlightreportsResponseDpo) error
	CreateFlightreport(requestDpo CreateFlightreportRequestDpo) error
}

// CreateFlightreportRequestDpo .
type CreateFlightreportRequestDpo interface {
	GetFlightoperationID() string
}

// GetFlightreportRequestDpo .
type GetFlightreportRequestDpo interface {
	GetID() string
}

// GetFlightreportResponseDpo .
type GetFlightreportResponseDpo = func(id, flightoperationID string)

// ListFlightreportsResponseDpo .
type ListFlightreportsResponseDpo = func(id, flightoperationID string)

// NewManageFlightreportService .
func NewManageFlightreportService(
	gen frep.Generator,
	repo frep.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFlightreportService {
	return &manageFlightreportService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageFlightreportService struct {
	gen  frep.Generator
	repo frep.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageFlightreportService) GetFlightreport(
	requestDpo GetFlightreportRequestDpo,
	responseDpo GetFlightreportResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightreportOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *manageFlightreportService) getFlightreportOperation(
	tx txmanager.Tx,
	requestDpo GetFlightreportRequestDpo,
	responseDpo GetFlightreportResponseDpo,
) error {
	flightreport, err := s.repo.GetByID(tx, frep.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	responseDpo(
		string(flightreport.GetID()),
		string(flightreport.GetFlightoperationID()),
	)
	return nil
}

func (s *manageFlightreportService) ListFlightreports(
	responseEachDpo ListFlightreportsResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightreportsOperation(
			tx,
			responseEachDpo,
		)
	})
}

func (s *manageFlightreportService) listFlightreportsOperation(
	tx txmanager.Tx,
	responseEachDpo ListFlightreportsResponseDpo,
) error {
	flightreports, err := s.repo.GetAll(tx)
	if err != nil {
		return err
	}

	for _, f := range flightreports {
		responseEachDpo(
			string(f.GetID()),
			string(f.GetFlightoperationID()),
		)
	}
	return nil
}

func (s *manageFlightreportService) CreateFlightreport(
	requestDpo CreateFlightreportRequestDpo,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.createFlightreportOperation(
				tx,
				requestDpo,
			)
		},
	)
}

func (s *manageFlightreportService) createFlightreportOperation(
	tx txmanager.Tx,
	requestDpo CreateFlightreportRequestDpo,
) error {
	if ret := frep.CreateNewFlightreport(
		tx,
		s.gen,
		s.repo,
		frep.FlightoperationID(requestDpo.GetFlightoperationID()),
	); ret != nil {
		return ret
	}

	return nil
}
