package mission

// CreatedEvent .
type CreatedEvent struct {
	ID      ID
	Mission *Mission
}

// GetMissionID .
func (e *CreatedEvent) GetMissionID() string {
	return string(e.ID)
}

// GetMission .
func (e *CreatedEvent) GetMission() *Mission {
	return e.Mission
}

// DeletedEvent .
type DeletedEvent struct {
	ID       ID
	UploadID UploadID
}

// GetMissionID .
func (e *DeletedEvent) GetMissionID() string {
	return string(e.ID)
}

// GetUploadID .
func (e *DeletedEvent) GetUploadID() string {
	return string(e.UploadID)
}
