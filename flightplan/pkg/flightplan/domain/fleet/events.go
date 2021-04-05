package fleet

import "flightplan/pkg/flightplan/domain/flightplan"

// VehicleCopiedWhenFlightplanCopiedEvent .
type VehicleCopiedWhenFlightplanCopiedEvent struct {
	FlightplanID flightplan.ID
	OriginalID   VehicleID
	NewID        VehicleID
}

// GetFlightplanID .
func (e *VehicleCopiedWhenFlightplanCopiedEvent) GetFlightplanID() string {
	return string(e.FlightplanID)
}

// GetOriginalID .
func (e *VehicleCopiedWhenFlightplanCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *VehicleCopiedWhenFlightplanCopiedEvent) GetNewID() string {
	return string(e.NewID)
}

// MissionCopiedWhenFlightplanCopiedEvent .
type MissionCopiedWhenFlightplanCopiedEvent struct {
	FlightplanID flightplan.ID
	OriginalID   MissionID
	NewID        MissionID
}

// GetFlightplanID .
func (e *MissionCopiedWhenFlightplanCopiedEvent) GetFlightplanID() string {
	return string(e.FlightplanID)
}

// GetOriginalID .
func (e *MissionCopiedWhenFlightplanCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *MissionCopiedWhenFlightplanCopiedEvent) GetNewID() string {
	return string(e.NewID)
}
