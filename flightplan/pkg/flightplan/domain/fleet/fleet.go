package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
	"log"
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

// Fleet .
type Fleet struct {
	id                 ID
	flightplanID       flightplan.ID
	vehicleAssignments []*VehicleAssignment
	eventPlannings     []*EventPlanning
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

// GetFlightplanID .
func (f *Fleet) GetFlightplanID() flightplan.ID {
	return f.flightplanID
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
			f.newVersion = f.gen.NewVersion()
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
			f.newVersion = f.gen.NewVersion()
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
			f.newVersion = f.gen.NewVersion()
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
			f.newVersion = f.gen.NewVersion()
			return nil
		}
	}
	return errors.New("event not found")
}

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
		log.Println(a)
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

// Generator .
type Generator interface {
	NewID() ID
	NewAssignmentID() AssignmentID
	NewEventID() EventID
	NewVersion() Version
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
