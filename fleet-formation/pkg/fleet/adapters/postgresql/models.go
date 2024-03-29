package postgresql

import "fleet-formation/pkg/fleet/domain/fleet"

// Fleet .
type Fleet struct {
	ID           string `gorm:"primaryKey"`
	IsCarbonCopy bool
	Assignments  []*Assignment `gorm:"-"`
	Events       []*Event      `gorm:"-"`
	Version      string
}

// GetID .
func (f *Fleet) GetID() string {
	return f.ID
}

// GetIsCarbonCopy .
func (f *Fleet) GetIsCarbonCopy() bool {
	return f.IsCarbonCopy
}

// GetVersion .
func (f *Fleet) GetVersion() string {
	return f.Version
}

// GetAssignments .
func (f *Fleet) GetAssignments() []fleet.AssignmentComponent {
	var assignments []fleet.AssignmentComponent
	for _, a := range f.Assignments {
		assignments = append(assignments, a)
	}
	return assignments
}

// GetEvents .
func (f *Fleet) GetEvents() []fleet.EventComponent {
	var events []fleet.EventComponent
	for _, e := range f.Events {
		events = append(events, e)
	}
	return events
}

// Assignment .
type Assignment struct {
	ID        string `gorm:"primaryKey"`
	FleetID   string
	VehicleID string
}

// GetID .
func (a *Assignment) GetID() string {
	return a.ID
}

// GetFleetID .
func (a *Assignment) GetFleetID() string {
	return a.FleetID
}

// GetVehicleID .
func (a *Assignment) GetVehicleID() string {
	return a.VehicleID
}

// Event .
type Event struct {
	ID           string `gorm:"primaryKey"`
	FleetID      string
	AssignmentID string
	MissionID    string
}

// GetID .
func (e *Event) GetID() string {
	return e.ID
}

// GetFleetID .
func (e *Event) GetFleetID() string {
	return e.FleetID
}

// GetAssignmentID .
func (e *Event) GetAssignmentID() string {
	return e.AssignmentID
}

// GetMissionID .
func (e *Event) GetMissionID() string {
	return e.MissionID
}
