package fleet

import (
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
	return nil
}

// CancelAssignment .
func (f *Fleet) CancelAssignment(assignmentID AssignmentID) error {
	return nil
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
