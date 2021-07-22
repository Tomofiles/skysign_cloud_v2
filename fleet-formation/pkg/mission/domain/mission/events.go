package mission

// CopiedMissionCreatedEvent .
type CopiedMissionCreatedEvent struct {
	ID      ID
	Mission *Mission
}

// GetID .
func (e *CopiedMissionCreatedEvent) GetID() ID {
	return e.ID
}

// GetMission .
func (e *CopiedMissionCreatedEvent) GetMission() *Mission {
	return e.Mission
}
