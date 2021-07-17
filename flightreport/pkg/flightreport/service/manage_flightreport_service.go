package service

import (
	"flightreport/pkg/flightreport/domain/event"
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/flightreport/domain/txmanager"
)

// ManageFlightreportService .
type ManageFlightreportService interface {
	GetFlightreport(command GetFlightreportCommand, retrievedModel RetrievedModel) error
	ListFlightreports(retrievedModel RetrievedModel) error
	CreateFlightreport(command CreateFlightreportCommand) error
}

// CreateFlightreportCommand .
type CreateFlightreportCommand interface {
	GetFlightreport() Flightreport
}

// GetFlightreportCommand .
type GetFlightreportCommand interface {
	GetID() string
}

// FlightreportPresentationModel .
type FlightreportPresentationModel interface {
	GetFlightreport() Flightreport
}

// Flightreport .
type Flightreport interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
}

// RetrievedModel .
type RetrievedModel = func(model FlightreportPresentationModel)

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
	command GetFlightreportCommand,
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightreportOperation(
			tx,
			command,
			retrievedModel,
		)
	})
}

func (s *manageFlightreportService) getFlightreportOperation(
	tx txmanager.Tx,
	command GetFlightreportCommand,
	retrievedModel RetrievedModel,
) error {
	flightreport, err := s.repo.GetByID(tx, frep.ID(command.GetID()))
	if err != nil {
		return err
	}

	retrievedModel(
		&flightreportModel{
			flightreport: flightreport,
		},
	)
	return nil
}

func (s *manageFlightreportService) ListFlightreports(
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightreportsOperation(
			tx,
			retrievedModel,
		)
	})
}

func (s *manageFlightreportService) listFlightreportsOperation(
	tx txmanager.Tx,
	retrievedModel RetrievedModel,
) error {
	flightreports, err := s.repo.GetAll(tx)
	if err != nil {
		return err
	}

	for _, f := range flightreports {
		retrievedModel(
			&flightreportModel{
				flightreport: f,
			},
		)
	}
	return nil
}

func (s *manageFlightreportService) CreateFlightreport(
	command CreateFlightreportCommand,
) error {
	return s.txm.Do(
		func(tx txmanager.Tx) error {
			return s.createFlightreportOperation(
				tx,
				command,
			)
		},
	)
}

func (s *manageFlightreportService) createFlightreportOperation(
	tx txmanager.Tx,
	command CreateFlightreportCommand,
) error {
	if ret := frep.CreateNewFlightreport(
		tx,
		s.gen,
		s.repo,
		command.GetFlightreport().GetName(),
		command.GetFlightreport().GetDescription(),
		frep.FleetID(command.GetFlightreport().GetFleetID()),
	); ret != nil {
		return ret
	}

	return nil
}

type flightreportModel struct {
	flightreport *frep.Flightreport
}

func (f *flightreportModel) GetFlightreport() Flightreport {
	return &flightreport{
		flightreport: f.flightreport,
	}
}

type flightreport struct {
	flightreport *frep.Flightreport
}

func (f *flightreport) GetID() string {
	return string(f.flightreport.GetID())
}

func (f *flightreport) GetName() string {
	return f.flightreport.GetName()
}

func (f *flightreport) GetDescription() string {
	return f.flightreport.GetDescription()
}

func (f *flightreport) GetFleetID() string {
	return string(f.flightreport.GetFleetID())
}
