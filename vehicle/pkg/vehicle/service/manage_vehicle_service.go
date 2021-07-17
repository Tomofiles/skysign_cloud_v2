package service

import (
	"vehicle/pkg/vehicle/domain/event"
	"vehicle/pkg/vehicle/domain/txmanager"
	v "vehicle/pkg/vehicle/domain/vehicle"
)

// ManageVehicleService .
type ManageVehicleService interface {
	GetVehicle(command GetVehicleCommand, retrievedModel RetrievedModel) error
	ListVehicles(retrievedModel RetrievedModel) error
	CreateVehicle(command CreateVehicleCommand, createdID CreatedID) error
	UpdateVehicle(command UpdateVehicleCommand) error
	DeleteVehicle(command DeleteVehicleCommand) error
	CarbonCopyVehicle(command CarbonCopyVehicleCommand) error
}

// CreateVehicleCommand .
type CreateVehicleCommand interface {
	GetVehicle() Vehicle
}

// UpdateVehicleCommand .
type UpdateVehicleCommand interface {
	GetID() string
	GetVehicle() Vehicle
}

// GetVehicleCommand .
type GetVehicleCommand interface {
	GetID() string
}

// DeleteVehicleCommand .
type DeleteVehicleCommand interface {
	GetID() string
}

// CarbonCopyVehicleCommand .
type CarbonCopyVehicleCommand interface {
	GetOriginalID() string
	GetNewID() string
	GetFleetID() string
}

// VehiclePresentationModel .
type VehiclePresentationModel interface {
	GetVehicle() Vehicle
}

// Vehicle .
type Vehicle interface {
	GetID() string
	GetName() string
	GetCommunicationID() string
}

// CreatedID .
type CreatedID = func(id string)

// RetrievedModel .
type RetrievedModel = func(model VehiclePresentationModel)

// NewManageVehicleService .
func NewManageVehicleService(
	gen v.Generator,
	repo v.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageVehicleService {
	return &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageVehicleService struct {
	gen  v.Generator
	repo v.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageVehicleService) GetVehicle(
	command GetVehicleCommand,
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getVehicleOperation(
			tx,
			command,
			retrievedModel,
		)
	})
}

func (s *manageVehicleService) getVehicleOperation(
	tx txmanager.Tx,
	command GetVehicleCommand,
	retrievedModel RetrievedModel,
) error {
	vehicle, err := s.repo.GetByID(tx, v.ID(command.GetID()))
	if err != nil {
		return err
	}

	retrievedModel(
		&vehicleModel{
			vehicle: vehicle,
		},
	)
	return nil
}

func (s *manageVehicleService) ListVehicles(
	retrievedModel RetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listVehiclesOperation(
			tx,
			retrievedModel,
		)
	})
}

func (s *manageVehicleService) listVehiclesOperation(
	tx txmanager.Tx,
	retrievedModel RetrievedModel,
) error {
	vehicles, err := s.repo.GetAllOrigin(tx)
	if err != nil {
		return err
	}

	for _, v := range vehicles {
		retrievedModel(
			&vehicleModel{
				vehicle: v,
			},
		)
	}
	return nil
}

func (s *manageVehicleService) CreateVehicle(
	command CreateVehicleCommand,
	createdID CreatedID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createVehicleOperation(
				tx,
				pub,
				command,
				createdID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageVehicleService) createVehicleOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CreateVehicleCommand,
	createdID CreatedID,
) error {
	id, ret := v.CreateNewVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		command.GetVehicle().GetName(),
		v.CommunicationID(command.GetVehicle().GetCommunicationID()))
	if ret != nil {
		return ret
	}

	createdID(id)
	return nil
}

func (s *manageVehicleService) UpdateVehicle(
	command UpdateVehicleCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.updateVehicleOperation(
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

func (s *manageVehicleService) updateVehicleOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command UpdateVehicleCommand,
) error {
	ret := v.UpdateVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		v.ID(command.GetID()),
		command.GetVehicle().GetName(),
		v.CommunicationID(command.GetVehicle().GetCommunicationID()))
	if ret != nil {
		return ret
	}

	return nil
}

func (s *manageVehicleService) DeleteVehicle(
	command DeleteVehicleCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteVehicleOperation(
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

func (s *manageVehicleService) deleteVehicleOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command DeleteVehicleCommand,
) error {
	if ret := v.DeleteVehicle(
		tx,
		s.repo,
		pub,
		v.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

func (s *manageVehicleService) CarbonCopyVehicle(
	command CarbonCopyVehicleCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.carbonCopyVehicleOperation(
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

func (s *manageVehicleService) carbonCopyVehicleOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CarbonCopyVehicleCommand,
) error {
	if ret := v.CarbonCopyVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		v.ID(command.GetOriginalID()),
		v.ID(command.GetNewID()),
		v.FleetID(command.GetFleetID()),
	); ret != nil {
		return ret
	}

	return nil
}

type vehicleModel struct {
	vehicle *v.Vehicle
}

func (f *vehicleModel) GetVehicle() Vehicle {
	return &vehicle{
		vehicle: f.vehicle,
	}
}

type vehicle struct {
	vehicle *v.Vehicle
}

func (f *vehicle) GetID() string {
	return string(f.vehicle.GetID())
}

func (f *vehicle) GetName() string {
	return f.vehicle.GetName()
}

func (f *vehicle) GetCommunicationID() string {
	return string(f.vehicle.GetCommunicationID())
}
