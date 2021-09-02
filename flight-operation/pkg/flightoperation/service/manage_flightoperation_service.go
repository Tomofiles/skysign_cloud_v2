package service

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageFlightoperationService .
type ManageFlightoperationService interface {
	GetFlightoperation(command GetFlightoperationCommand, model RetrievedModel) error
	ListFlightoperations(model RetrievedModel) error
	CreateFlightoperation(command CreateFlightoperationCommand) error
}

// CreateFlightoperationRequestDpo .
type CreateFlightoperationCommand interface {
	GetFlightoperation() Flightoperation
}

// GetFlightoperationRequestDpo .
type GetFlightoperationCommand interface {
	GetID() string
}

// FlightoperationPresentationModel .
type FlightoperationPresentationModel interface {
	GetFlightoperation() Flightoperation
}

// Flightoperation .
type Flightoperation interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
}

// RetrievedModel .
type RetrievedModel = func(model FlightoperationPresentationModel)

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
	command GetFlightoperationCommand,
	model RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightoperationOperation(
			tx,
			command,
			model,
		)
	})
}

func (s *manageFlightoperationService) getFlightoperationOperation(
	tx txmanager.Tx,
	command GetFlightoperationCommand,
	model RetrievedModel,
) error {
	flightoperation, err := s.repo.GetByID(tx, fope.ID(command.GetID()))
	if err != nil {
		return err
	}

	model(
		&flightoperationModel{
			flightoperation: flightoperation,
		},
	)
	return nil
}

func (s *manageFlightoperationService) ListFlightoperations(
	model RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightoperationsOperation(
			tx,
			model,
		)
	})
}

func (s *manageFlightoperationService) listFlightoperationsOperation(
	tx txmanager.Tx,
	model RetrievedModel,
) error {
	flightoperations, err := s.repo.GetAllOperating(tx)
	if err != nil {
		return err
	}

	for _, f := range flightoperations {
		model(
			&flightoperationModel{
				flightoperation: f,
			},
		)
	}
	return nil
}

func (s *manageFlightoperationService) CreateFlightoperation(
	command CreateFlightoperationCommand,
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
				command,
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
	command CreateFlightoperationCommand,
) error {
	if ret := fope.CreateNewFlightoperation(
		tx,
		s.gen,
		s.repo,
		pub,
		command.GetFlightoperation().GetName(),
		command.GetFlightoperation().GetDescription(),
		fope.FleetID(command.GetFlightoperation().GetFleetID()),
	); ret != nil {
		return ret
	}

	return nil
}

type flightoperationModel struct {
	flightoperation *fope.Flightoperation
}

func (f *flightoperationModel) GetFlightoperation() Flightoperation {
	return &flightoperation{
		flightoperation: f.flightoperation,
	}
}

type flightoperation struct {
	flightoperation *fope.Flightoperation
}

func (f *flightoperation) GetID() string {
	return string(f.flightoperation.GetID())
}

func (f *flightoperation) GetName() string {
	return f.flightoperation.GetName()
}

func (f *flightoperation) GetDescription() string {
	return f.flightoperation.GetDescription()
}

func (f *flightoperation) GetFleetID() string {
	return string(f.flightoperation.GetFleetID())
}
