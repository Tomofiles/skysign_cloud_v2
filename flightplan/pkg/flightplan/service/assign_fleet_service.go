package service

import (
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
		return s.changeNumberOfVehiclesOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *AssignFleetService) changeNumberOfVehiclesOperation(
	tx txmanager.Tx,
	requestDpo ChangeNumberOfVehiclesRequestDpo,
	responseDpo ChangeNumberOfVehiclesResponseDpo,
) error {
	oldFleet, err := s.repo.GetByFlightplanID(
		tx,
		flightplan.ID(requestDpo.GetFlightplanID()),
	)
	if err != nil {
		return err
	}
	if oldFleet == nil {
		return fleet.ErrNotFound
	}

	if ret := s.repo.DeleteByFlightplanID(
		tx,
		flightplan.ID(requestDpo.GetFlightplanID()),
	); ret != nil {
		return ret
	}

	newFleet := fleet.NewInstance(
		s.gen,
		flightplan.ID(requestDpo.GetFlightplanID()),
		requestDpo.GetNumberOfVehicles())
	for _, assignmentID := range newFleet.GetAllAssignmentID() {
		newFleet.AddNewEvent(assignmentID)
	}
	if ret := s.repo.Save(tx, newFleet); ret != nil {
		return ret
	}

	responseDpo(requestDpo.GetFlightplanID(), requestDpo.GetNumberOfVehicles())
	return nil
}

// GetAssignments .
func (s *AssignFleetService) GetAssignments(
	requestDpo GetAssignmentsRequestDpo,
	responseEachDpo GetAssignmentsResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getAssignmentsOperation(
			tx,
			requestDpo,
			responseEachDpo,
		)
	})
}

func (s *AssignFleetService) getAssignmentsOperation(
	tx txmanager.Tx,
	requestDpo GetAssignmentsRequestDpo,
	responseEachDpo GetAssignmentsResponseDpo,
) error {
	aFleet, err := s.repo.GetByFlightplanID(
		tx,
		flightplan.ID(requestDpo.GetFlightplanID()),
	)
	if err != nil {
		return err
	}
	if aFleet == nil {
		return fleet.ErrNotFound
	}

	var assignments []assignmentVehicle
	var events []eventMission
	aFleet.ProvideAssignmentsInterest(
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
}

// UpdateAssignment .
func (s *AssignFleetService) UpdateAssignment(
	requestDpo UpdateAssignmentRequestDpo,
	responseDpo UpdateAssignmentResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.updateAssignmentOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *AssignFleetService) updateAssignmentOperation(
	tx txmanager.Tx,
	requestDpo UpdateAssignmentRequestDpo,
	responseDpo UpdateAssignmentResponseDpo,
) error {
	aFleet, err := s.repo.GetByFlightplanID(
		tx,
		flightplan.ID(requestDpo.GetFlightplanID()),
	)
	if err != nil {
		return err
	}
	if aFleet == nil {
		return fleet.ErrNotFound
	}

	if requestDpo.GetVehicleID() != "" {
		if ret := aFleet.AssignVehicle(
			fleet.AssignmentID(requestDpo.GetAssignmentID()),
			fleet.VehicleID(requestDpo.GetVehicleID()),
		); ret != nil {
			return ret
		}
	} else {
		if ret := aFleet.CancelVehiclesAssignment(
			fleet.AssignmentID(requestDpo.GetAssignmentID()),
		); ret != nil {
			return ret
		}
	}
	if requestDpo.GetMissionID() != "" {
		if ret := aFleet.AssignMission(
			fleet.EventID(requestDpo.GetEventID()),
			fleet.MissionID(requestDpo.GetMissionID()),
		); ret != nil {
			return ret
		}
	} else {
		if ret := aFleet.CancelMission(
			fleet.EventID(requestDpo.GetEventID()),
		); ret != nil {
			return ret
		}
	}
	if ret := s.repo.Save(tx, aFleet); ret != nil {
		return ret
	}

	responseDpo(
		requestDpo.GetEventID(),
		requestDpo.GetAssignmentID(),
		requestDpo.GetVehicleID(),
		requestDpo.GetMissionID(),
	)
	return nil
}

// ChangeNumberOfVehiclesRequestDpo .
type ChangeNumberOfVehiclesRequestDpo interface {
	GetFlightplanID() string
	GetNumberOfVehicles() int32
}

// ChangeNumberOfVehiclesResponseDpo .
type ChangeNumberOfVehiclesResponseDpo = func(id string, numberOfVehicles int32)

// GetAssignmentsRequestDpo .
type GetAssignmentsRequestDpo interface {
	GetFlightplanID() string
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
	GetFlightplanID() string
	GetEventID() string
	GetAssignmentID() string
	GetVehicleID() string
	GetMissionID() string
}

// UpdateAssignmentResponseDpo .
type UpdateAssignmentResponseDpo = func(id, assignmentId, vehicleId, missionId string)
