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

// NewFlightplanID .
func (g *FlightoperationUUID) NewFlightplanID() flightoperation.FlightplanID {
	uuid, _ := uuid.NewRandom()
	return flightoperation.FlightplanID(uuid.String())
}
