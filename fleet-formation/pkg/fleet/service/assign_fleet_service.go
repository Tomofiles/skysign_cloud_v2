package service

import (
	"fleet-formation/pkg/common/domain/txmanager"
	f "fleet-formation/pkg/fleet/domain/fleet"
)

// AssignFleetService .
type AssignFleetService interface {
	GetAssignments(command GetAssignmentsCommand, model AssignmentRetrievedModel) error
	UpdateAssignment(command UpdateAssignmentCommand, model AssignmentRetrievedModel) error
}

// GetAssignmentsCommand .
type GetAssignmentsCommand interface {
	GetID() string
}

// UpdateAssignmentCommand .
type UpdateAssignmentCommand interface {
	GetID() string
	GetEventID() string
	GetAssignmentID() string
	GetVehicleID() string
	GetMissionID() string
}

// AssignmentPresentationModel .
type AssignmentPresentationModel interface {
	GetAssignment() Assignment
}

// Assignment .
type Assignment interface {
	GetID() string
	GetEventID() string
	GetAssignmentID() string
	GetVehicleID() string
	GetMissionID() string
}

// AssignmentRetrievedModel .
type AssignmentRetrievedModel = func(model AssignmentPresentationModel)

type assignmentVehicle struct {
	assignmentID string
	vehicleID    string
}
type eventMission struct {
	eventID      string
	assignmentID string
	missionID    string
}

// NewAssignFleetService .
func NewAssignFleetService(
	gen f.Generator,
	repo f.Repository,
	txm txmanager.TransactionManager,
) AssignFleetService {
	return &assignFleetService{
		gen:  gen,
		repo: repo,
		txm:  txm,
	}
}

type assignFleetService struct {
	gen  f.Generator
	repo f.Repository
	txm  txmanager.TransactionManager
}

func (s *assignFleetService) GetAssignments(
	command GetAssignmentsCommand,
	model AssignmentRetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getAssignmentsOperation(
			tx,
			command,
			model,
		)
	})
}

func (s *assignFleetService) getAssignmentsOperation(
	tx txmanager.Tx,
	command GetAssignmentsCommand,
	model AssignmentRetrievedModel,
) error {
	fleet, err := s.repo.GetByID(
		tx,
		f.ID(command.GetID()),
	)
	if err != nil {
		return err
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
		model(
			&assignmentModel{
				assignment: &assignment{
					ID:           string(fleet.GetID()),
					EventID:      eventID,
					AssignmentID: a.assignmentID,
					VehicleID:    a.vehicleID,
					MissionID:    missionID,
				},
			},
		)
	}
	return nil
}

func (s *assignFleetService) UpdateAssignment(
	command UpdateAssignmentCommand,
	model AssignmentRetrievedModel,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.updateAssignmentOperation(
			tx,
			command,
			model,
		)
	})
}

func (s *assignFleetService) updateAssignmentOperation(
	tx txmanager.Tx,
	command UpdateAssignmentCommand,
	model AssignmentRetrievedModel,
) error {
	fleet, err := s.repo.GetByID(
		tx,
		f.ID(command.GetID()),
	)
	if err != nil {
		return err
	}

	if command.GetVehicleID() != "" {
		if ret := fleet.AssignVehicle(
			f.AssignmentID(command.GetAssignmentID()),
			f.VehicleID(command.GetVehicleID()),
		); ret != nil {
			return ret
		}
	} else {
		if ret := fleet.CancelVehiclesAssignment(
			f.AssignmentID(command.GetAssignmentID()),
		); ret != nil {
			return ret
		}
	}
	if command.GetMissionID() != "" {
		if ret := fleet.AssignMission(
			f.EventID(command.GetEventID()),
			f.MissionID(command.GetMissionID()),
		); ret != nil {
			return ret
		}
	} else {
		if ret := fleet.CancelMission(
			f.EventID(command.GetEventID()),
		); ret != nil {
			return ret
		}
	}
	if ret := s.repo.Save(tx, fleet); ret != nil {
		return ret
	}

	model(
		&assignmentModel{
			assignment: &assignment{
				ID:           string(fleet.GetID()),
				EventID:      command.GetEventID(),
				AssignmentID: command.GetAssignmentID(),
				VehicleID:    command.GetVehicleID(),
				MissionID:    command.GetMissionID(),
			},
		},
	)
	return nil
}

type assignmentModel struct {
	assignment *assignment
}

func (f *assignmentModel) GetAssignment() Assignment {
	return f.assignment
}

type assignment struct {
	ID           string
	EventID      string
	AssignmentID string
	VehicleID    string
	MissionID    string
}

func (a *assignment) GetID() string {
	return a.ID
}

func (a *assignment) GetEventID() string {
	return a.EventID
}

func (a *assignment) GetAssignmentID() string {
	return a.AssignmentID
}

func (a *assignment) GetVehicleID() string {
	return a.VehicleID
}

func (a *assignment) GetMissionID() string {
	return a.MissionID
}
