package flightreport

// CreatedEvent .
type CreatedEvent struct {
	ID                ID
	FlightoperationID FlightoperationID
}

// GetID .
func (e *CreatedEvent) GetID() string {
	return string(e.ID)
}

// GetFlightoperationID .
func (e *CreatedEvent) GetFlightoperationID() string {
	return string(e.FlightoperationID)
}

// FlightoperationCopiedWhenCreatedEvent .
type FlightoperationCopiedWhenCreatedEvent struct {
	OriginalID FlightoperationID
	NewID      FlightoperationID
}

// GetOriginalID .
func (e *FlightoperationCopiedWhenCreatedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *FlightoperationCopiedWhenCreatedEvent) GetNewID() string {
	return string(e.NewID)
}
