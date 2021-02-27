package fleet

import "flightplan/pkg/flightplan/domain/flightplan"

// NewInstance .
func NewInstance(gen Generator, flightplanID flightplan.ID, numberOfVehicles int32) *Fleet {
	var vehicleAssignments []*VehicleAssignment
	var vaIndex int32
	for vaIndex < numberOfVehicles {
		vehicleAssignments = append(vehicleAssignments, &VehicleAssignment{
			assignmentID: gen.NewAssignmentID(),
		})
		vaIndex++
	}
	version := gen.NewVersion()
	return &Fleet{
		id:                 gen.NewID(),
		flightplanID:       flightplanID,
		vehicleAssignments: vehicleAssignments,
		version:            version,
		newVersion:         version,
		gen:                gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Fleet {
	var vehicleAssignments []*VehicleAssignment
	for _, a := range comp.GetAssignments() {
		vehicleAssignments = append(
			vehicleAssignments,
			&VehicleAssignment{
				assignmentID: AssignmentID(a.GetID()),
				vehicleID:    VehicleID(a.GetVehicleID()),
			},
		)
	}
	var eventPlannings []*EventPlanning
	for _, e := range comp.GetEvents() {
		eventPlannings = append(
			eventPlannings,
			&EventPlanning{
				eventID:      EventID(e.GetID()),
				assignmentID: AssignmentID(e.GetAssignmentID()),
				missionID:    MissionID(e.GetMissionID()),
			},
		)
	}
	return &Fleet{
		id:                 ID(comp.GetID()),
		flightplanID:       flightplan.ID(comp.GetFlightplanID()),
		vehicleAssignments: vehicleAssignments,
		eventPlannings:     eventPlannings,
		version:            Version(comp.GetVersion()),
		newVersion:         Version(comp.GetVersion()),
		gen:                gen,
	}
}

// TakeApart .
func TakeApart(
	fleet *Fleet,
	fleetComp func(id, flightplanID, version string),
	assignmentComp func(id, fleetID, vehicleID string),
	eventComp func(id, fleetID, assignmentID, missionID string),
) {
	fleetComp(
		string(fleet.id),
		string(fleet.flightplanID),
		string(fleet.version),
	)
	for _, a := range fleet.vehicleAssignments {
		assignmentComp(
			string(a.assignmentID),
			string(fleet.id),
			string(a.vehicleID),
		)
	}
	for _, e := range fleet.eventPlannings {
		eventComp(
			string(e.eventID),
			string(fleet.id),
			string(e.assignmentID),
			string(e.missionID),
		)
	}
}

// Component .
type Component interface {
	GetID() string
	GetFlightplanID() string
	GetVersion() string
	GetAssignments() []AssignmentComponent
	GetEvents() []EventComponent
}

// AssignmentComponent .
type AssignmentComponent interface {
	GetID() string
	GetFleetID() string
	GetVehicleID() string
}

// EventComponent .
type EventComponent interface {
	GetID() string
	GetFleetID() string
	GetAssignmentID() string
	GetMissionID() string
}
