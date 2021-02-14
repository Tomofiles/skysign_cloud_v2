package flightplan

import "fmt"

// AssignmentID .
type AssignmentID string

// VehicleID .
type VehicleID string

// Fleet .
type Fleet struct {
	flightplanID       ID
	numberOfVehicles   int
	vehicleAssignments []*VehicleAssignment
	generator          Generator
}

// VehicleAssignment .
type VehicleAssignment struct {
	assignmentID AssignmentID
	name         string
	vehicleID    VehicleID
}

// AddAssignment .
func (f *Fleet) AddAssignment() error {
	f.numberOfVehicles++
	f.vehicleAssignments = append(
		f.vehicleAssignments,
		&VehicleAssignment{
			assignmentID: f.generator.NewAssignmentID(),
			name:         "Vehicle -- " + fmt.Sprint(f.numberOfVehicles),
		},
	)
	return nil
}

// RemoveAssignment .
func (f *Fleet) RemoveAssignment() error {
	return nil
}

// AssignVehicle .
func (f *Fleet) AssignVehicle(assignmentID AssignmentID, vehicleID VehicleID) error {
	return nil
}

// CancelAssignment .
func (f *Fleet) CancelAssignment(assignmentID AssignmentID) error {
	return nil
}
