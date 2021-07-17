package fleet

// VehicleCopiedEvent .
type VehicleCopiedEvent struct {
	FleetID    ID
	OriginalID VehicleID
	NewID      VehicleID
}

// GetFleetID .
func (e *VehicleCopiedEvent) GetFleetID() string {
	return string(e.FleetID)
}

// GetOriginalID .
func (e *VehicleCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *VehicleCopiedEvent) GetNewID() string {
	return string(e.NewID)
}

// MissionCopiedEvent .
type MissionCopiedEvent struct {
	FleetID    ID
	OriginalID MissionID
	NewID      MissionID
}

// GetFleetID .
func (e *MissionCopiedEvent) GetFleetID() string {
	return string(e.FleetID)
}

// GetOriginalID .
func (e *MissionCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *MissionCopiedEvent) GetNewID() string {
	return string(e.NewID)
}
