package flightplan

// FleetIDGaveEvent .
type FleetIDGaveEvent struct {
	FleetID          FleetID
	NumberOfVehicles int
}

// GetFleetID .
func (e *FleetIDGaveEvent) GetFleetID() FleetID {
	return e.FleetID
}

// GetNumberOfVehicles .
func (e *FleetIDGaveEvent) GetNumberOfVehicles() int {
	return e.NumberOfVehicles
}

// FleetIDRemovedEvent .
type FleetIDRemovedEvent struct {
	FleetID FleetID
}

// GetFleetID .
func (e *FleetIDRemovedEvent) GetFleetID() FleetID {
	return e.FleetID
}

// FlightplanExecutedEvent .
type FlightplanExecutedEvent struct {
	ID          ID
	Name        string
	Description string
	FleetID     FleetID
}

// GetID .
func (e *FlightplanExecutedEvent) GetID() ID {
	return e.ID
}

// GetName .
func (e *FlightplanExecutedEvent) GetName() string {
	return e.Name
}

// GetDescription .
func (e *FlightplanExecutedEvent) GetDescription() string {
	return e.Description
}

// GetFleetID .
func (e *FlightplanExecutedEvent) GetFleetID() FleetID {
	return e.FleetID
}
