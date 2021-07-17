package fleet

import (
	"fleet-formation/pkg/fleet/domain/event"
)

// NewInstance .
func NewInstance(gen Generator, id ID, numberOfVehicles int) *Fleet {
	var vehicleAssignments []*VehicleAssignment
	var vaIndex int
	for vaIndex < numberOfVehicles {
		vehicleAssignments = append(vehicleAssignments, &VehicleAssignment{
			assignmentID: gen.NewAssignmentID(),
		})
		vaIndex++
	}
	version := gen.NewVersion()
	return &Fleet{
		id:                 id,
		vehicleAssignments: vehicleAssignments,
		isCarbonCopy:       Original,
		version:            version,
		newVersion:         version,
		gen:                gen,
	}
}

// Copy .
func Copy(
	gen Generator,
	pub event.Publisher,
	id ID,
	original *Fleet,
) *Fleet {
	var vehicleAssignments []*VehicleAssignment
	var eventPlannings []*EventPlanning

	assignmentIDMap := map[AssignmentID]AssignmentID{}
	vehicleIDMap := map[VehicleID]VehicleID{"": ""}
	missionIDMap := map[MissionID]MissionID{"": ""}
	for _, va := range original.vehicleAssignments {
		assignmentIDMap[va.assignmentID] = gen.NewAssignmentID()
		if va.vehicleID != "" {
			vehicleIDMap[va.vehicleID] = gen.NewVehicleID()
		}
		vehicleAssignments = append(
			vehicleAssignments,
			&VehicleAssignment{
				assignmentID: assignmentIDMap[va.assignmentID],
				vehicleID:    vehicleIDMap[va.vehicleID],
			},
		)
	}
	for _, ep := range original.eventPlannings {
		if ep.missionID != "" {
			missionIDMap[ep.missionID] = gen.NewMissionID()
		}
		eventPlannings = append(
			eventPlannings,
			&EventPlanning{
				eventID:      gen.NewEventID(),
				assignmentID: assignmentIDMap[ep.assignmentID],
				missionID:    missionIDMap[ep.missionID],
			},
		)
	}

	if pub != nil {
		for k, v := range vehicleIDMap {
			if k == "" {
				continue
			}
			event := VehicleCopiedEvent{
				FleetID:    id,
				OriginalID: k,
				NewID:      v,
			}
			pub.Publish(event)
		}
		for k, v := range missionIDMap {
			if k == "" {
				continue
			}
			event := MissionCopiedEvent{
				FleetID:    id,
				OriginalID: k,
				NewID:      v,
			}
			pub.Publish(event)
		}
	}

	return &Fleet{
		id:                 id,
		vehicleAssignments: vehicleAssignments,
		eventPlannings:     eventPlannings,
		isCarbonCopy:       CarbonCopy,
		version:            original.version,
		newVersion:         original.newVersion,
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
		vehicleAssignments: vehicleAssignments,
		eventPlannings:     eventPlannings,
		isCarbonCopy:       comp.GetIsCarbonCopy(),
		version:            Version(comp.GetVersion()),
		newVersion:         Version(comp.GetVersion()),
		gen:                gen,
	}
}

// TakeApart .
func TakeApart(
	fleet *Fleet,
	fleetComp func(id, version string, isCarbonCopy bool),
	assignmentComp func(id, fleetID, vehicleID string),
	eventComp func(id, fleetID, assignmentID, missionID string),
) {
	fleetComp(
		string(fleet.id),
		string(fleet.version),
		fleet.isCarbonCopy,
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
	GetIsCarbonCopy() bool
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
