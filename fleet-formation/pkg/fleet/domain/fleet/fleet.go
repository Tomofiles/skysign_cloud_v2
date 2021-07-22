package fleet

import (
	"errors"
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

// Version .
type Version string

var (
	// ErrVehicleHasAlreadyAssigned .
	ErrVehicleHasAlreadyAssigned = errors.New("this vehicle has already assigned")
	// ErrAssignmentNotFound .
	ErrAssignmentNotFound = errors.New("assignment not found")
	// ErrMissionHasAlreadyAssigned .
	ErrMissionHasAlreadyAssigned = errors.New("this mission has already assigned")
	// ErrEventNotFound .
	ErrEventNotFound = errors.New("event not found")
)

const (
	// Original .
	Original = false
	// CarbonCopy .
	CarbonCopy = true
)

var (
	// ErrCannotChange .
	ErrCannotChange = errors.New("cannnot change carbon copied fleet")
)

// Fleet .
type Fleet struct {
	id                 ID
	vehicleAssignments []*VehicleAssignment
	eventPlannings     []*EventPlanning
	isCarbonCopy       bool
	version            Version
	newVersion         Version
	gen                Generator
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

// GetAllAssignmentID .
func (f *Fleet) GetAllAssignmentID() []AssignmentID {
	var assignmentIDs []AssignmentID
	for _, va := range f.vehicleAssignments {
		assignmentIDs = append(assignmentIDs, va.assignmentID)
	}
	return assignmentIDs
}

// GetVersion .
func (f *Fleet) GetVersion() Version {
	return f.version
}

// GetNewVersion .
func (f *Fleet) GetNewVersion() Version {
	return f.newVersion
}

// AssignVehicle .
func (f *Fleet) AssignVehicle(assignmentID AssignmentID, vehicleID VehicleID) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}

	contains := false
	for _, va := range f.vehicleAssignments {
		if va.assignmentID != assignmentID && va.vehicleID == vehicleID {
			contains = true
		}
	}
	if contains {
		return ErrVehicleHasAlreadyAssigned
	}

	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			va.vehicleID = vehicleID
			f.newVersion = f.gen.NewVersion()
			return nil
		}
	}
	return ErrAssignmentNotFound
}

// CancelVehiclesAssignment .
func (f *Fleet) CancelVehiclesAssignment(assignmentID AssignmentID) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}

	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			va.vehicleID = ""
			f.newVersion = f.gen.NewVersion()
			return nil
		}
	}
	return ErrAssignmentNotFound
}

// AddNewEvent .
func (f *Fleet) AddNewEvent(assignmentID AssignmentID) (EventID, error) {
	if f.isCarbonCopy {
		return "", ErrCannotChange
	}

	notContains := true
	for _, va := range f.vehicleAssignments {
		if va.assignmentID == assignmentID {
			notContains = false
		}
	}
	if notContains {
		return "", ErrAssignmentNotFound
	}

	eventID := f.gen.NewEventID()
	f.eventPlannings = append(
		f.eventPlannings,
		&EventPlanning{
			eventID:      eventID,
			assignmentID: assignmentID,
		},
	)
	f.newVersion = f.gen.NewVersion()
	return eventID, nil
}

// RemoveEvent .
func (f *Fleet) RemoveEvent(eventID EventID) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}

	var eventPlannings []*EventPlanning
	for _, ep := range f.eventPlannings {
		if ep.eventID != eventID {
			eventPlannings = append(eventPlannings, ep)
		}
	}
	if len(eventPlannings) != len(f.eventPlannings) {
		f.eventPlannings = eventPlannings
		f.newVersion = f.gen.NewVersion()
		return nil
	}
	return ErrEventNotFound
}

// AssignMission .
func (f *Fleet) AssignMission(eventID EventID, missionID MissionID) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}

	contains := false
	for _, ep := range f.eventPlannings {
		if ep.eventID != eventID && ep.missionID == missionID {
			contains = true
		}
	}
	if contains {
		return ErrMissionHasAlreadyAssigned
	}

	for _, ep := range f.eventPlannings {
		if ep.eventID == eventID {
			ep.missionID = missionID
			f.newVersion = f.gen.NewVersion()
			return nil
		}
	}
	return ErrEventNotFound
}

// CancelMission .
func (f *Fleet) CancelMission(eventID EventID) error {
	if f.isCarbonCopy {
		return ErrCannotChange
	}

	for _, ep := range f.eventPlannings {
		if ep.eventID == eventID {
			ep.missionID = ""
			f.newVersion = f.gen.NewVersion()
			return nil
		}
	}
	return ErrEventNotFound
}

// ProvideAssignmentsInterest .
func (f *Fleet) ProvideAssignmentsInterest(
	assignment func(assignmentID, vehicleID string),
	event func(eventID, assignmentID, missionID string),
) {
	for _, va := range f.vehicleAssignments {
		assignment(string(va.assignmentID), string(va.vehicleID))
		for _, ep := range f.eventPlannings {
			if ep.assignmentID == va.assignmentID {
				event(string(ep.eventID), string(ep.assignmentID), string(ep.missionID))
			}
		}
	}
}

// Generator .
type Generator interface {
	NewAssignmentID() AssignmentID
	NewEventID() EventID
	NewVehicleID() VehicleID
	NewMissionID() MissionID
	NewVersion() Version
}
