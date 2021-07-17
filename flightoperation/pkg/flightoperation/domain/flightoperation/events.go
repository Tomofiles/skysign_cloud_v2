package flightoperation

// FleetCopiedEvent .
type FleetCopiedEvent struct {
	OriginalID FleetID
	NewID      FleetID
}

// GetOriginalID .
func (e *FleetCopiedEvent) GetOriginalID() string {
	return string(e.OriginalID)
}

// GetNewID .
func (e *FleetCopiedEvent) GetNewID() string {
	return string(e.NewID)
}

// FlightoperationCompletedEvent .
type FlightoperationCompletedEvent struct {
	ID          ID
	Name        string
	Description string
	FleetID     FleetID
}

// GetID .
func (e *FlightoperationCompletedEvent) GetID() string {
	return string(e.ID)
}

// GetName .
func (e *FlightoperationCompletedEvent) GetName() string {
	return e.Name
}

// GetDescription .
func (e *FlightoperationCompletedEvent) GetDescription() string {
	return e.Description
}

// GetFleetID .
func (e *FlightoperationCompletedEvent) GetFleetID() string {
	return string(e.FleetID)
}
