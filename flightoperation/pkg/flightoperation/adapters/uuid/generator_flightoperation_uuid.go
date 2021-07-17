package uuid

import (
	"flightoperation/pkg/flightoperation/domain/flightoperation"

	"github.com/google/uuid"
)

// FlightoperationUUID .
type FlightoperationUUID struct{}

// NewFlightoperationUUID .
func NewFlightoperationUUID() *FlightoperationUUID {
	return &FlightoperationUUID{}
}

// NewID .
func (g *FlightoperationUUID) NewID() flightoperation.ID {
	uuid, _ := uuid.NewRandom()
	return flightoperation.ID(uuid.String())
}

// NewFleetID .
func (g *FlightoperationUUID) NewFleetID() flightoperation.FleetID {
	uuid, _ := uuid.NewRandom()
	return flightoperation.FleetID(uuid.String())
}

// NewVersion .
func (g *FlightoperationUUID) NewVersion() flightoperation.Version {
	uuid, _ := uuid.NewRandom()
	return flightoperation.Version(uuid.String())
}
