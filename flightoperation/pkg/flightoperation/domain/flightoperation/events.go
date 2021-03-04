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

// GetOriginalFlightplanID .
func (e *FlightplanCopiedWhenCreatedEvent) GetOriginalFlightplanID() string {
	return string(e.OriginalID)
}

// GetNewFlightplanID .
func (e *FlightplanCopiedWhenCreatedEvent) GetNewFlightplanID() string {
	return string(e.NewID)
}
