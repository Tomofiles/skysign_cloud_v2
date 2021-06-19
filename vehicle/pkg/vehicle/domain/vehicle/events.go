package vehicle

// FlightplanID .
type FlightplanID string

// CommunicationIDGaveEvent .
type CommunicationIDGaveEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIDGaveEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// CommunicationIDRemovedEvent .
type CommunicationIDRemovedEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIDRemovedEvent) GetCommunicationID() string {
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
