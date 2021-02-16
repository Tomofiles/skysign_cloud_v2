package generator

import (
	"flightplan/pkg/flightplan/domain/flightplan"

	"github.com/google/uuid"
)

// FlightplanUUID .
type FlightplanUUID struct{}

// NewID .
func (g *FlightplanUUID) NewID() flightplan.ID {
	uuid, _ := uuid.NewRandom()
	return flightplan.ID(uuid.String())
}

// NewVersion .
func (g *FlightplanUUID) NewVersion() flightplan.Version {
	uuid, _ := uuid.NewRandom()
	return flightplan.Version(uuid.String())
}
