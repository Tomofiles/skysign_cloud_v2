package postgresql

import "flightplan/pkg/flightplan/domain/fleet"

// Flightplan .
type Flightplan struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	FleetID     string
	Version     string
}

// GetID .
func (f *Flightplan) GetID() string {
	return f.ID
}

// GetName .
func (f *Flightplan) GetName() string {
	return f.Name
}

// GetDescription .
func (f *Flightplan) GetDescription() string {
	return f.Description
}

// GetFleetID .
func (f *Flightplan) GetFleetID() string {
	return f.FleetID
}

// GetVersion .
func (f *Flightplan) GetVersion() string {
	return f.Version
}

// Fleet .
type Fleet struct {
	ID           string `gorm:"primaryKey"`
	FlightplanID string
	IsCarbonCopy bool
	Assignments  []*Assignment `gorm:"-"`
	Events       []*Event      `gorm:"-"`
	Version      string
}

// GetID .
func (f *Fleet) GetID() string {
	return f.ID
}

// GetFlightplanID .
func (f *Fleet) GetFlightplanID() string {
	return f.FlightplanID
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
