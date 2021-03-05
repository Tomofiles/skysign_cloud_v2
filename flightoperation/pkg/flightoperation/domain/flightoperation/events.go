package flightoperation

// CreatedEvent .
type CreatedEvent struct {
	ID           ID
	FlightplanID FlightplanID
}

// GetID .
func (e *CreatedEvent) GetID() string {
	return string(e.ID)
}

// GetFlightplanID .
func (e *CreatedEvent) GetFlightplanID() string {
	return string(e.FlightplanID)
}

// FlightplanCopiedWhenCreatedEvent .
type FlightplanCopiedWhenCreatedEvent struct {
	OriginalID FlightplanID
	NewID      FlightplanID
}

// GetOriginalID .
func (e *FlightplanCopiedWhenCreatedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *FlightplanCopiedWhenCreatedEvent) GetNewID() string {
	return string(e.NewID)
}
