package flightplan

// CreatedEvent .
type CreatedEvent struct {
	ID ID
}

// GetFlightplanID .
func (e *CreatedEvent) GetFlightplanID() string {
	return string(e.ID)
}

// DeletedEvent .
type DeletedEvent struct {
	ID ID
}

// GetFlightplanID .
func (e *DeletedEvent) GetFlightplanID() string {
	return string(e.ID)
}
