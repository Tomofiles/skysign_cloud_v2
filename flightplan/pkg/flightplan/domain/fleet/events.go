package fleet

// VehicleCopiedWhenCopiedEvent .
type VehicleCopiedWhenCopiedEvent struct {
	OriginalID VehicleID
	NewID      VehicleID
}

// GetOriginalID .
func (e *VehicleCopiedWhenCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *VehicleCopiedWhenCopiedEvent) GetNewID() string {
	return string(e.NewID)
}

// MissionCopiedWhenCopiedEvent .
type MissionCopiedWhenCopiedEvent struct {
	OriginalID MissionID
	NewID      MissionID
}

// GetOriginalID .
func (e *MissionCopiedWhenCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *MissionCopiedWhenCopiedEvent) GetNewID() string {
	return string(e.NewID)
}
