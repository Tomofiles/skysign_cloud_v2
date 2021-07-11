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
