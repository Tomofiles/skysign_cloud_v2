package generator

import (
	"flightplan/pkg/flightplan/domain/fleet"

	"github.com/google/uuid"
)

// FleetUUID .
type FleetUUID struct{}

// NewID .
func (g *FleetUUID) NewID() fleet.ID {
	uuid, _ := uuid.NewRandom()
	return fleet.ID(uuid.String())
}

// NewAssignmentID .
func (g *FleetUUID) NewAssignmentID() fleet.AssignmentID {
	uuid, _ := uuid.NewRandom()
	return fleet.AssignmentID(uuid.String())
}

// NewEventID .
func (g *FleetUUID) NewEventID() fleet.EventID {
	uuid, _ := uuid.NewRandom()
	return fleet.EventID(uuid.String())
}

// NewVersion .
func (g *FleetUUID) NewVersion() fleet.Version {
	uuid, _ := uuid.NewRandom()
	return fleet.Version(uuid.String())
}
