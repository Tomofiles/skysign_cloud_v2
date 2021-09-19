package uuid

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/google/uuid"
)

// FleetUUID .
type FleetUUID struct{}

// NewFleetUUID .
func NewFleetUUID() *FleetUUID {
	return &FleetUUID{}
}

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

// NewVehicleID .
func (g *FleetUUID) NewVehicleID() fleet.VehicleID {
	uuid, _ := uuid.NewRandom()
	return fleet.VehicleID(uuid.String())
}

// NewMissionID .
func (g *FleetUUID) NewMissionID() fleet.MissionID {
	uuid, _ := uuid.NewRandom()
	return fleet.MissionID(uuid.String())
}

// NewVersion .
func (g *FleetUUID) NewVersion() fleet.Version {
	uuid, _ := uuid.NewRandom()
	return fleet.Version(uuid.String())
}
