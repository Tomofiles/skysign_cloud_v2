package vehicle

// FlightplanID .
type FlightplanID string

// CommunicationIdGaveEvent .
type CommunicationIdGaveEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIdGaveEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// CommunicationIdRemovedEvent .
type CommunicationIdRemovedEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIdRemovedEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// CopiedVehicleCreatedEvent .
type CopiedVehicleCreatedEvent struct {
	ID              ID
	CommunicationID CommunicationID
	FlightplanID    FlightplanID
}

// GetVehicleID .
func (e *CopiedVehicleCreatedEvent) GetVehicleID() string {
	return string(e.ID)
}

// GetCommunicationID .
func (e *CopiedVehicleCreatedEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// GetFlightplanID .
func (e *CopiedVehicleCreatedEvent) GetFlightplanID() string {
	return string(e.FlightplanID)
}
