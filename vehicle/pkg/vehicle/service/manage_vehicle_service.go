package service

import (
	"vehicle/pkg/vehicle/domain/event"
	"vehicle/pkg/vehicle/domain/txmanager"
	v "vehicle/pkg/vehicle/domain/vehicle"
)

// ManageVehicleService .
type ManageVehicleService interface {
	GetVehicle(requestDpo GetVehicleRequestDpo, responseDpo GetVehicleResponseDpo) error
	ListVehicles(responseEachDpo ListVehiclesResponseDpo) error
	CreateVehicle(requestDpo CreateVehicleRequestDpo, responseDpo CreateVehicleResponseDpo) error
	UpdateVehicle(requestDpo UpdateVehicleRequestDpo, responseDpo UpdateVehicleResponseDpo) error
	DeleteVehicle(requestDpo DeleteVehicleRequestDpo) error
	CarbonCopyVehicle(requestDpo CarbonCopyVehicleRequestDpo) error
}

// CreateVehicleRequestDpo .
type CreateVehicleRequestDpo interface {
	GetName() string
	GetCommunicationID() string
}

// CreateVehicleResponseDpo .
type CreateVehicleResponseDpo = func(id, name, communicationID string)

// UpdateVehicleRequestDpo .
type UpdateVehicleRequestDpo interface {
	GetID() string
	GetName() string
	GetCommunicationID() string
}

// UpdateVehicleResponseDpo .
type UpdateVehicleResponseDpo = func(id, name, communicationID string)

// GetVehicleRequestDpo .
type GetVehicleRequestDpo interface {
	GetID() string
}

// GetVehicleResponseDpo .
type GetVehicleResponseDpo = func(id, name, communicationID string)

// ListVehiclesResponseDpo .
type ListVehiclesResponseDpo = func(id, name, communicationID string)

// DeleteVehicleRequestDpo .
type DeleteVehicleRequestDpo interface {
	GetID() string
}

// CarbonCopyVehicleRequestDpo .
type CarbonCopyVehicleRequestDpo interface {
	GetOriginalID() string
	GetNewID() string
	GetFlightplanID() string
}

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
	requestDpo GetVehicleRequestDpo,
	responseDpo GetVehicleResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getVehicleOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *manageVehicleService) getVehicleOperation(
	tx txmanager.Tx,
	requestDpo GetVehicleRequestDpo,
	responseDpo GetVehicleResponseDpo,
) error {
	vehicle, err := s.repo.GetByID(tx, v.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	responseDpo(
		string(vehicle.GetID()),
		vehicle.GetName(),
		string(vehicle.GetCommunicationID()),
	)
	return nil
}

func (s *manageVehicleService) ListVehicles(
	responseEachDpo ListVehiclesResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listVehiclesOperation(
			tx,
			responseEachDpo,
		)
	})
}

func (s *manageVehicleService) listVehiclesOperation(
	tx txmanager.Tx,
	responseEachDpo ListVehiclesResponseDpo,
) error {
	vehicles, err := s.repo.GetAllOrigin(tx)
	if err != nil {
		return err
	}

	for _, v := range vehicles {
		responseEachDpo(
			string(v.GetID()),
			v.GetName(),
			string(v.GetCommunicationID()),
		)
	}
	return nil
}

func (s *manageVehicleService) CreateVehicle(
	requestDpo CreateVehicleRequestDpo,
	responseDpo CreateVehicleResponseDpo,
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
				requestDpo,
				responseDpo,
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
	requestDpo CreateVehicleRequestDpo,
	responseDpo CreateVehicleResponseDpo,
) error {
	id, ret := v.CreateNewVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		requestDpo.GetName(),
		v.CommunicationID(requestDpo.GetCommunicationID()))
	if ret != nil {
		return ret
	}

	responseDpo(id, requestDpo.GetName(), requestDpo.GetCommunicationID())
	return nil
}

func (s *manageVehicleService) UpdateVehicle(
	requestDpo UpdateVehicleRequestDpo,
	responseDpo UpdateVehicleResponseDpo,
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
				requestDpo,
				responseDpo,
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
	requestDpo UpdateVehicleRequestDpo,
	responseDpo UpdateVehicleResponseDpo,
) error {
	ret := v.UpdateVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		v.ID(requestDpo.GetID()),
		requestDpo.GetName(),
		v.CommunicationID(requestDpo.GetCommunicationID()))
	if ret != nil {
		return ret
	}

	responseDpo(requestDpo.GetID(), requestDpo.GetName(), requestDpo.GetCommunicationID())
	return nil
}

func (s *manageVehicleService) DeleteVehicle(
	requestDpo DeleteVehicleRequestDpo,
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
				requestDpo,
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
	requestDpo DeleteVehicleRequestDpo,
) error {
	if ret := v.DeleteVehicle(
		tx,
		s.repo,
		pub,
		v.ID(requestDpo.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

func (s *manageVehicleService) CarbonCopyVehicle(
	requestDpo CarbonCopyVehicleRequestDpo,
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
				requestDpo,
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
	requestDpo CarbonCopyVehicleRequestDpo,
) error {
	if ret := v.CarbonCopyVehicle(
		tx,
		s.gen,
		s.repo,
		pub,
		v.ID(requestDpo.GetOriginalID()),
		v.ID(requestDpo.GetNewID()),
		v.FlightplanID(requestDpo.GetFlightplanID()),
	); ret != nil {
		return ret
	}

	return nil
}
