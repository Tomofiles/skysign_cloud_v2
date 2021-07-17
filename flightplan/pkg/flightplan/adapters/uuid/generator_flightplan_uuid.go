package uuid

import (
	"flightplan/pkg/flightplan/domain/flightplan"

	"github.com/google/uuid"
)

// FlightplanUUID .
type FlightplanUUID struct{}

// NewFlightplanUUID .
func NewFlightplanUUID() *FlightplanUUID {
	return &FlightplanUUID{}
}

// NewID .
func (g *FlightplanUUID) NewID() flightplan.ID {
	uuid, _ := uuid.NewRandom()
	return flightplan.ID(uuid.String())
}

// NewFleetID .
func (g *FlightplanUUID) NewFleetID() flightplan.FleetID {
	uuid, _ := uuid.NewRandom()
	return flightplan.FleetID(uuid.String())
}

// NewVersion .
func (g *FlightplanUUID) NewVersion() flightplan.Version {
	uuid, _ := uuid.NewRandom()
	return flightplan.Version(uuid.String())
}
