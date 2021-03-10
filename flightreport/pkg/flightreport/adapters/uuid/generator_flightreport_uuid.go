package uuid

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"

	"github.com/google/uuid"
)

// FlightreportUUID .
type FlightreportUUID struct{}

// NewFlightreportUUID .
func NewFlightreportUUID() *FlightreportUUID {
	return &FlightreportUUID{}
}

// NewID .
func (g *FlightreportUUID) NewID() frep.ID {
	uuid, _ := uuid.NewRandom()
	return frep.ID(uuid.String())
}

// NewFlightoperationID .
func (g *FlightreportUUID) NewFlightoperationID() frep.FlightoperationID {
	uuid, _ := uuid.NewRandom()
	return frep.FlightoperationID(uuid.String())
}
