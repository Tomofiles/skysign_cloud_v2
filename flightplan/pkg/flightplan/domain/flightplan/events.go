package flightplan

// CreatedEvent .
type CreatedEvent struct {
	ID ID
}

// GetFlightplanID .
func (e *CreatedEvent) GetFlightplanID() string {
	return string(e.ID)
}

// CopiedEvent .
type CopiedEvent struct {
	OriginalID ID
	NewID      ID
}

// GetOriginalID .
func (e *CopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *CopiedEvent) GetNewID() string {
	return string(e.NewID)
}

// DeletedEvent .
type DeletedEvent struct {
	ID ID
}

// GetFlightplanID .
func (e *DeletedEvent) GetFlightplanID() string {
	return string(e.ID)
}
