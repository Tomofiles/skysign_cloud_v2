package vehicle

// FleetID .
type FleetID string

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
	FleetID         FleetID
}

// GetVehicleID .
func (e *CopiedVehicleCreatedEvent) GetVehicleID() string {
	return string(e.ID)
}

// GetCommunicationID .
func (e *CopiedVehicleCreatedEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// GetFleetID .
func (e *CopiedVehicleCreatedEvent) GetFleetID() string {
	return string(e.FleetID)
}
