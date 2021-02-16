package flightplan

// CreatedEvent .
type CreatedEvent struct {
	id ID
}

// GetFlightplanID .
func (e *CreatedEvent) GetFlightplanID() string {
	return string(e.id)
}

// DeletedEvent .
type DeletedEvent struct {
	id ID
}

// GetFlightplanID .
func (e *DeletedEvent) GetFlightplanID() string {
	return string(e.id)
}
