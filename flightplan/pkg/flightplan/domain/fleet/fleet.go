package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
)

// ID .
type ID string

// AssignmentID .
type AssignmentID string

// VehicleID .
type VehicleID string

// EventID .
type EventID string

// MissionID .
type MissionID string

// Fleet .
type Fleet struct {
	id                 ID
	flightplanID       flightplan.ID
	vehicleAssignments []*VehicleAssignment
	eventPlannings     []*EventPlanning
	generator          Generator
}

// VehicleAssignment .
type VehicleAssignment struct {
	assignmentID AssignmentID
	vehicleID    VehicleID
}

// EventPlanning .
type EventPlanning struct {
	eventID      EventID
	assignmentID AssignmentID
	missionID    MissionID
}

// GetID .
func (f *Fleet) GetID() ID {
	return f.id
}

// GetFlightplanID .
func (f *Fleet) GetFlightplanID() flightplan.ID {
	return f.flightplanID
}

// GetNumberOfVehicles .
func (f *Fleet) GetNumberOfVehicles() int {
	return len(f.vehicleAssignments)
}

// GetVehicleAssignments .
func (f *Fleet) GetVehicleAssignments() []*VehicleAssignment {
	return f.vehicleAssignments
}

// GetEventPlannings .
func (f *Fleet) GetEventPlannings() []*EventPlanning {
	return f.eventPlannings
}

// AssignVehicle .
func (f *Fleet) AssignVehicle(assignmentID AssignmentID, vehicleID VehicleID) error {
	contains := false
	for _, va := range f.vehicleAssignments {
		if va.assignmentID != assignmentID && va.vehicleID == vehicleID {
			contains = true
		}
	}
	if contains {
		return errors.New("this vehicle has already assigned")
	}

	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			va.vehicleID = vehicleID
			return nil
		}
	}
	return errors.New("assignment not found")
}

// CancelVehiclesAssignment .
func (f *Fleet) CancelVehiclesAssignment(assignmentID AssignmentID) error {
	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			va.vehicleID = ""
			return nil
		}
	}
	return errors.New("assignment not found")
}

// AddNewEvent .
func (f *Fleet) AddNewEvent(assignmentID AssignmentID) (EventID, error) {
	notContains := true
	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			notContains = false
		}
	}
	if notContains {
		return "", errors.New("this id not assigned")
	}

	eventID := f.generator.NewEventID()
	f.eventPlannings = append(
		f.eventPlannings,
		&EventPlanning{
			eventID:      eventID,
			assignmentID: assignmentID,
		},
	)
	return eventID, nil
}

// RemoveEvent .
func (f *Fleet) RemoveEvent(eventID EventID) error {
	var eventPlannings []*EventPlanning
	for _, ep := range f.eventPlannings {
		if ep.eventID != eventID {
			eventPlannings = append(eventPlannings, ep)
		}
	}
	if len(eventPlannings) != len(f.eventPlannings) {
		f.eventPlannings = eventPlannings
		return nil
	}
	return errors.New("event not found")
}

// AssignMission .
func (f *Fleet) AssignMission(eventID EventID, missionID MissionID) error {
	contains := false
	for _, ep := range f.eventPlannings {
		if ep.eventID != eventID && ep.missionID == missionID {
			contains = true
		}
	}
	if contains {
		return errors.New("this mission has already assigned")
	}

	for _, ep := range f.eventPlannings {
		if ep.eventID == eventID {
			ep.missionID = missionID
			return nil
		}
	}
	return errors.New("event not found")
}

// CancelMission .
func (f *Fleet) CancelMission(eventID EventID) error {
	for _, ep := range f.eventPlannings {
		if ep.eventID == eventID {
			ep.missionID = ""
			return nil
		}
	}
	return errors.New("event not found")
}

// NewInstance .
func NewInstance(generator Generator, flightplanID flightplan.ID, numberOfVehicles int) *Fleet {
	var vehicleAssignments []*VehicleAssignment
	vaIndex := 0
	for vaIndex < numberOfVehicles {
		vehicleAssignments = append(vehicleAssignments, &VehicleAssignment{
			assignmentID: generator.NewAssignmentID(),
		})
		vaIndex++
	}
	return &Fleet{
		id:                 generator.NewID(),
		flightplanID:       flightplanID,
		vehicleAssignments: vehicleAssignments,
		generator:          generator,
	}
}

// Generator .
type Generator interface {
	NewID() ID
	NewAssignmentID() AssignmentID
	NewEventID() EventID
}
