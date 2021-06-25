package service

import (
	"mission/pkg/mission/domain/event"
	m "mission/pkg/mission/domain/mission"
	"mission/pkg/mission/domain/txmanager"
)

// ManageMissionService .
type ManageMissionService interface {
	// GetVehicle(requestDpo GetVehicleRequestDpo, responseDpo GetVehicleResponseDpo) error
	// ListVehicles(responseEachDpo ListVehiclesResponseDpo) error
	CreateMission(command CreateMissionCommand, createdID CreatedID) error
	UpdateMission(command UpdateMissionCommand) error
	DeleteMission(command DeleteMissionCommand) error
	// CarbonCopyVehicle(requestDpo CarbonCopyVehicleRequestDpo) error
}

// CreateMissionCommand .
type CreateMissionCommand interface {
	GetName() string
	GetNavigation() Navigation
}

// UpdateMissionCommand .
type UpdateMissionCommand interface {
	GetID() string
	GetName() string
	GetNavigation() Navigation
}

// Navigation .
type Navigation interface {
	GetTakeoffPointGroundHeight() float64
	GetWaypoints() []Waypoint
}

// Waypoint .
type Waypoint interface {
	GetLatitude() float64
	GetLongitude() float64
	GetRelativeHeight() float64
	GetSpeed() float64
}

// CreatedID .
type CreatedID = func(id string)

// // GetVehicleRequestDpo .
// type GetVehicleRequestDpo interface {
// 	GetID() string
// }

// // GetVehicleResponseDpo .
// type GetVehicleResponseDpo = func(id, name, communicationID string)

// // ListVehiclesResponseDpo .
// type ListVehiclesResponseDpo = func(id, name, communicationID string)

// DeleteMissionCommand .
type DeleteMissionCommand interface {
	GetID() string
}

// // CarbonCopyVehicleRequestDpo .
// type CarbonCopyVehicleRequestDpo interface {
// 	GetOriginalID() string
// 	GetNewID() string
// 	GetFlightplanID() string
// }

// NewManageMissionService .
func NewManageMissionService(
	gen m.Generator,
	repo m.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageMissionService {
	return &manageMissionService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageMissionService struct {
	gen  m.Generator
	repo m.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

// func (s *manageVehicleService) GetVehicle(
// 	requestDpo GetVehicleRequestDpo,
// 	responseDpo GetVehicleResponseDpo,
// ) error {
// 	return s.txm.Do(func(tx txmanager.Tx) error {
// 		return s.getVehicleOperation(
// 			tx,
// 			requestDpo,
// 			responseDpo,
// 		)
// 	})
// }

// func (s *manageVehicleService) getVehicleOperation(
// 	tx txmanager.Tx,
// 	requestDpo GetVehicleRequestDpo,
// 	responseDpo GetVehicleResponseDpo,
// ) error {
// 	vehicle, err := s.repo.GetByID(tx, v.ID(requestDpo.GetID()))
// 	if err != nil {
// 		return err
// 	}

// 	responseDpo(
// 		string(vehicle.GetID()),
// 		vehicle.GetName(),
// 		string(vehicle.GetCommunicationID()),
// 	)
// 	return nil
// }

// func (s *manageVehicleService) ListVehicles(
// 	responseEachDpo ListVehiclesResponseDpo,
// ) error {
// 	return s.txm.Do(func(tx txmanager.Tx) error {
// 		return s.listVehiclesOperation(
// 			tx,
// 			responseEachDpo,
// 		)
// 	})
// }

// func (s *manageVehicleService) listVehiclesOperation(
// 	tx txmanager.Tx,
// 	responseEachDpo ListVehiclesResponseDpo,
// ) error {
// 	vehicles, err := s.repo.GetAllOrigin(tx)
// 	if err != nil {
// 		return err
// 	}

// 	for _, v := range vehicles {
// 		responseEachDpo(
// 			string(v.GetID()),
// 			v.GetName(),
// 			string(v.GetCommunicationID()),
// 		)
// 	}
// 	return nil
// }

func (s *manageMissionService) CreateMission(
	command CreateMissionCommand,
	createdID CreatedID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createMissionOperation(
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

func (s *manageMissionService) createMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CreateMissionCommand,
	createdID CreatedID,
) error {
	navigation := m.NewNavigation(command.GetNavigation().GetTakeoffPointGroundHeight())
	for _, w := range command.GetNavigation().GetWaypoints() {
		navigation.PushNextWaypoint(
			w.GetLatitude(),
			w.GetLongitude(),
			w.GetRelativeHeight(),
			w.GetSpeed(),
		)
	}
	id, ret := m.CreateNewMission(
		tx,
		s.gen,
		s.repo,
		pub,
		command.GetName(),
		navigation,
	)
	if ret != nil {
		return ret
	}

	createdID(id)
	return nil
}

func (s *manageMissionService) UpdateMission(
	command UpdateMissionCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.updateMissionOperation(
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

func (s *manageMissionService) updateMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command UpdateMissionCommand,
) error {
	navigation := m.NewNavigation(command.GetNavigation().GetTakeoffPointGroundHeight())
	for _, w := range command.GetNavigation().GetWaypoints() {
		navigation.PushNextWaypoint(
			w.GetLatitude(),
			w.GetLongitude(),
			w.GetRelativeHeight(),
			w.GetSpeed(),
		)
	}
	ret := m.UpdateMission(
		tx,
		s.gen,
		s.repo,
		pub,
		m.ID(command.GetID()),
		command.GetName(),
		navigation,
	)
	if ret != nil {
		return ret
	}

	return nil
}

func (s *manageMissionService) DeleteMission(
	command DeleteMissionCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteMissionOperation(
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

func (s *manageMissionService) deleteMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command DeleteMissionCommand,
) error {
	if ret := m.DeleteMission(
		tx,
		s.repo,
		pub,
		m.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

// func (s *manageVehicleService) CarbonCopyVehicle(
// 	requestDpo CarbonCopyVehicleRequestDpo,
// ) error {
// 	pub, chClose, err := s.psm.GetPublisher()
// 	if err != nil {
// 		return err
// 	}
// 	defer chClose()

// 	return s.txm.DoAndEndHook(
// 		func(tx txmanager.Tx) error {
// 			return s.carbonCopyVehicleOperation(
// 				tx,
// 				pub,
// 				requestDpo,
// 			)
// 		},
// 		func() error {
// 			return pub.Flush()
// 		},
// 	)
// }

// func (s *manageVehicleService) carbonCopyVehicleOperation(
// 	tx txmanager.Tx,
// 	pub event.Publisher,
// 	requestDpo CarbonCopyVehicleRequestDpo,
// ) error {
// 	if ret := v.CarbonCopyVehicle(
// 		tx,
// 		s.gen,
// 		s.repo,
// 		pub,
// 		v.ID(requestDpo.GetOriginalID()),
// 		v.ID(requestDpo.GetNewID()),
// 		v.FlightplanID(requestDpo.GetFlightplanID()),
// 	); ret != nil {
// 		return ret
// 	}

// 	return nil
// }
