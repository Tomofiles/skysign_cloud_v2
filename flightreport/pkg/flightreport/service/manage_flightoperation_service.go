package service

import (
	"flightreport/pkg/flightreport/domain/event"
	fope "flightreport/pkg/flightreport/domain/flightoperation"
	"flightreport/pkg/flightreport/domain/txmanager"
)

// ManageFlightoperationService .
type ManageFlightoperationService interface {
	GetFlightoperation(requestDpo GetFlightoperationRequestDpo, responseDpo GetFlightoperationResponseDpo) error
	ListFlightoperations(responseEachDpo ListFlightoperationsResponseDpo) error
	CreateFlightoperation(requestDpo CreateFlightoperationRequestDpo) error
}

// CreateFlightoperationRequestDpo .
type CreateFlightoperationRequestDpo interface {
	GetFlightplanID() string
}

// GetFlightoperationRequestDpo .
type GetFlightoperationRequestDpo interface {
	GetID() string
}

// GetFlightoperationResponseDpo .
type GetFlightoperationResponseDpo = func(id, flightplanID string)

// ListFlightoperationsResponseDpo .
type ListFlightoperationsResponseDpo = func(id, flightplanID string)

// NewManageFlightoperationService .
func NewManageFlightoperationService(
	gen fope.Generator,
	repo fope.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFlightoperationService {
	return &manageFlightoperationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageFlightoperationService struct {
	gen  fope.Generator
	repo fope.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageFlightoperationService) GetFlightoperation(
	requestDpo GetFlightoperationRequestDpo,
	responseDpo GetFlightoperationResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightoperationOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *manageFlightoperationService) getFlightoperationOperation(
	tx txmanager.Tx,
	requestDpo GetFlightoperationRequestDpo,
	responseDpo GetFlightoperationResponseDpo,
) error {
	flightoperation, err := s.repo.GetByID(tx, fope.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	responseDpo(
		string(flightoperation.GetID()),
		string(flightoperation.GetFlightplanID()),
	)
	return nil
}

func (s *manageFlightoperationService) ListFlightoperations(
	responseEachDpo ListFlightoperationsResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightoperationsOperation(
			tx,
			responseEachDpo,
		)
	})
}

func (s *manageFlightoperationService) listFlightoperationsOperation(
	tx txmanager.Tx,
	responseEachDpo ListFlightoperationsResponseDpo,
) error {
	flightoperations, err := s.repo.GetAll(tx)
	if err != nil {
		return err
	}

	for _, f := range flightoperations {
		responseEachDpo(
			string(f.GetID()),
			string(f.GetFlightplanID()),
		)
	}
	return nil
}

func (s *manageFlightoperationService) CreateFlightoperation(
	requestDpo CreateFlightoperationRequestDpo,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createFlightoperationOperation(
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

func (s *manageFlightoperationService) createFlightoperationOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo CreateFlightoperationRequestDpo,
) error {
	if ret := fope.CreateNewFlightoperation(
		tx,
		s.gen,
		s.repo,
		pub,
		fope.FlightplanID(requestDpo.GetFlightplanID()),
	); ret != nil {
		return ret
	}

	return nil
}
