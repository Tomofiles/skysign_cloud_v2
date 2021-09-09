package service

import (
	f "flight-operation/pkg/flightplan/domain/flightplan"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageFlightplanService .
type ManageFlightplanService interface {
	GetFlightplan(command GetFlightplanCommand, retrievedModel RetrievedModel) error
	ListFlightplans(retrievedModel RetrievedModel) error
	CreateFlightplan(command CreateFlightplanCommand, createdID CreatedID, fleetID FleetID) error
	UpdateFlightplan(command UpdateFlightplanCommand, fleetID FleetID) error
	DeleteFlightplan(command DeleteFlightplanCommand) error
}

// CreateFlightplanCommand .
type CreateFlightplanCommand interface {
	GetFlightplan() Flightplan
}

// UpdateFlightplanCommand .
type UpdateFlightplanCommand interface {
	GetID() string
	GetFlightplan() Flightplan
}

// GetFlightplanCommand .
type GetFlightplanCommand interface {
	GetID() string
}

// DeleteFlightplanCommand .
type DeleteFlightplanCommand interface {
	GetID() string
}

// FlightplanPresentationModel .
type FlightplanPresentationModel interface {
	GetFlightplan() Flightplan
}

// Flightplan .
type Flightplan interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
}

// CreatedID .
type CreatedID = func(id string)

// UploadID .
type FleetID = func(fleetID string)

// RetrievedModel .
type RetrievedModel = func(model FlightplanPresentationModel)

// NewManageFlightplanService .
func NewManageFlightplanService(
	gen f.Generator,
	repo f.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFlightplanService {
	return &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageFlightplanService struct {
	gen  f.Generator
	repo f.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageFlightplanService) GetFlightplan(
	command GetFlightplanCommand,
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightplanOperation(
			tx,
			command,
			retrievedModel,
		)
	})
}

func (s *manageFlightplanService) getFlightplanOperation(
	tx txmanager.Tx,
	command GetFlightplanCommand,
	retrievedModel RetrievedModel,
) error {
	flightplan, err := s.repo.GetByID(tx, f.ID(command.GetID()))
	if err != nil {
		return err
	}

	retrievedModel(
		&flightplanModel{
			flightplan: flightplan,
		},
	)
	return nil
}

func (s *manageFlightplanService) ListFlightplans(
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightplansOperation(
			tx,
			retrievedModel,
		)
	})
}

func (s *manageFlightplanService) listFlightplansOperation(
	tx txmanager.Tx,
	retrievedModel RetrievedModel,
) error {
	flightplans, err := s.repo.GetAll(tx)
	if err != nil {
		return err
	}

	for _, f := range flightplans {
		retrievedModel(
			&flightplanModel{
				flightplan: f,
			},
		)
	}
	return nil
}

func (s *manageFlightplanService) CreateFlightplan(
	command CreateFlightplanCommand,
	createdID CreatedID,
	fleetID FleetID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createFlightplanOperation(
				tx,
				pub,
				command,
				createdID,
				fleetID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageFlightplanService) createFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CreateFlightplanCommand,
	createdID CreatedID,
	fleetID FleetID,
) error {
	id, fid, ret := f.CreateNewFlightplan(
		tx,
		s.gen,
		s.repo,
		pub,
		command.GetFlightplan().GetName(),
		command.GetFlightplan().GetDescription())
	if ret != nil {
		return ret
	}

	createdID(string(id))
	fleetID(string(fid))
	return nil
}

func (s *manageFlightplanService) UpdateFlightplan(
	command UpdateFlightplanCommand,
	fleetID FleetID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.updateFlightplanOperation(
				tx,
				pub,
				command,
				fleetID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageFlightplanService) updateFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command UpdateFlightplanCommand,
	fleetID FleetID,
) error {
	ret := f.UpdateFlightplan(
		tx,
		s.gen,
		s.repo,
		pub,
		f.ID(command.GetID()),
		command.GetFlightplan().GetName(),
		command.GetFlightplan().GetDescription())
	if ret != nil {
		return ret
	}

	flightplan, _ := s.repo.GetByID(tx, f.ID(command.GetID()))
	fleetID(string(flightplan.GetFleetID()))
	return nil
}

func (s *manageFlightplanService) DeleteFlightplan(
	command DeleteFlightplanCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteFlightplanOperation(
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

func (s *manageFlightplanService) deleteFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command DeleteFlightplanCommand,
) error {
	if ret := f.DeleteFlightplan(
		tx,
		s.repo,
		pub,
		f.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

type flightplanModel struct {
	flightplan *f.Flightplan
}

func (f *flightplanModel) GetFlightplan() Flightplan {
	return &flightplan{
		flightplan: f.flightplan,
	}
}

type flightplan struct {
	flightplan *f.Flightplan
}

func (f *flightplan) GetID() string {
	return string(f.flightplan.GetID())
}

func (f *flightplan) GetName() string {
	return f.flightplan.GetName()
}

func (f *flightplan) GetDescription() string {
	return f.flightplan.GetDescription()
}

func (f *flightplan) GetFleetID() string {
	return string(f.flightplan.GetFleetID())
}
