package service

import (
	"errors"
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// AssignFleetService .
type AssignFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	txm  txmanager.TransactionManager
}

// NewAssignFleetService .
func NewAssignFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	txm txmanager.TransactionManager,
) AssignFleetService {
	return AssignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}
}

// ChangeNumberOfVehicles .
func (s *AssignFleetService) ChangeNumberOfVehicles(
	requestDpo ChangeNumberOfVehiclesRequestDpo,
	responseDpo ChangeNumberOfVehiclesResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		oldFleet, err := s.repo.GetByFlightplanID(
			tx,
			flightplan.ID(requestDpo.GetId()),
		)
		if err != nil {
			return err
		}
		if oldFleet == nil {
			return errors.New("fleet not found")
		}

		err = s.repo.DeleteByFlightplanID(
			tx,
			flightplan.ID(requestDpo.GetId()),
		)
		if err != nil {
			return err
		}

		newFleet := fleet.NewInstance(
			s.gen,
			flightplan.ID(requestDpo.GetId()),
			requestDpo.GetNumberOfVehicles())
		for _, assignmentID := range newFleet.GetAllAssignmentID() {
			newFleet.AddNewEvent(assignmentID)
		}
		ret := s.repo.Save(tx, newFleet)
		if ret != nil {
			return ret
		}

		responseDpo(requestDpo.GetId(), requestDpo.GetNumberOfVehicles())
		return nil
	})
}

// GetAssignments .
func (s *AssignFleetService) GetAssignments(
	requestDpo GetAssignmentsRequestDpo,
	responseEachDpo GetAssignmentsResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		fleet, err := s.repo.GetByFlightplanID(
			tx,
			flightplan.ID(requestDpo.GetId()),
		)
		if err != nil {
			return err
		}
		if fleet == nil {
			return errors.New("fleet not found")
		}

		var assignments []assignmentVehicle
		var events []eventMission
		fleet.ProvideAssignmentsInterest(
			func(assignmentID string, vehicleID string) {
				assignments = append(
					assignments,
					assignmentVehicle{
						assignmentID: assignmentID,
						vehicleID:    vehicleID,
					},
				)
			},
			func(eventID string, assignmentID string, missionID string) {
				events = append(
					events,
					eventMission{
						eventID:      eventID,
						assignmentID: assignmentID,
						missionID:    missionID,
					},
				)
			},
		)

		for _, a := range assignments {
			var eventID, missionID string
			for _, e := range events {
				if a.assignmentID == e.assignmentID {
					eventID = e.eventID
					missionID = e.missionID
				}
			}
			responseEachDpo(
				eventID,
				a.assignmentID,
				a.vehicleID,
				missionID)
		}
		return nil
	})
}

// UpdateAssignment .
func (s *AssignFleetService) UpdateAssignment(
	requestDpo UpdateAssignmentRequestDpo,
	responseDpo UpdateAssignmentResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		aFleet, err := s.repo.GetByFlightplanID(
			tx,
			flightplan.ID(requestDpo.GetFlightplanId()),
		)
		if err != nil {
			return err
		}
		if aFleet == nil {
			return errors.New("fleet not found")
		}

		if requestDpo.GetVehicleId() != "" {
			if ret := aFleet.AssignVehicle(
				fleet.AssignmentID(requestDpo.GetAssignmentId()),
				fleet.VehicleID(requestDpo.GetVehicleId()),
			); ret != nil {
				return ret
			}
		} else {
			if ret := aFleet.CancelVehiclesAssignment(
				fleet.AssignmentID(requestDpo.GetAssignmentId()),
			); ret != nil {
				return ret
			}
		}
		if requestDpo.GetMissionId() != "" {
			if ret := aFleet.AssignMission(
				fleet.EventID(requestDpo.GetId()),
				fleet.MissionID(requestDpo.GetMissionId()),
			); ret != nil {
				return ret
			}
		} else {
			if ret := aFleet.CancelMission(
				fleet.EventID(requestDpo.GetId()),
			); ret != nil {
				return ret
			}
		}
		if ret := s.repo.Save(tx, aFleet); ret != nil {
			return ret
		}

		responseDpo(
			requestDpo.GetId(),
			requestDpo.GetAssignmentId(),
			requestDpo.GetVehicleId(),
			requestDpo.GetMissionId(),
		)
		return nil
	})
}

// ChangeNumberOfVehiclesRequestDpo .
type ChangeNumberOfVehiclesRequestDpo interface {
	GetId() string
	GetNumberOfVehicles() int32
}

// ChangeNumberOfVehiclesResponseDpo .
type ChangeNumberOfVehiclesResponseDpo = func(id string, numberOfVehicles int32)

// GetAssignmentsRequestDpo .
type GetAssignmentsRequestDpo interface {
	GetId() string
}

// GetAssignmentsResponseDpo .
type GetAssignmentsResponseDpo = func(id, assignmentId, vehicleId, missionId string)

type assignmentVehicle struct {
	assignmentID string
	vehicleID    string
}
type eventMission struct {
	eventID      string
	assignmentID string
	missionID    string
}

// UpdateAssignmentRequestDpo .
type UpdateAssignmentRequestDpo interface {
	GetFlightplanId() string
	GetId() string
	GetAssignmentId() string
	GetVehicleId() string
	GetMissionId() string
}

// UpdateAssignmentResponseDpo .
type UpdateAssignmentResponseDpo = func(id, assignmentId, vehicleId, missionId string)
