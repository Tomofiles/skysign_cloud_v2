package flightreport

// ID .
type ID string

// FleetID .
type FleetID string

// Flightreport .
type Flightreport struct {
	id          ID
	name        string
	description string
	fleetID     FleetID
}

// GetID .
func (f *Flightreport) GetID() ID {
	return f.id
}

// GetName .
func (f *Flightreport) GetName() string {
	return f.name
}

// GetDescription .
func (f *Flightreport) GetDescription() string {
	return f.description
}

// GetFleetID .
func (f *Flightreport) GetFleetID() FleetID {
	return f.fleetID
}

// Generator .
type Generator interface {
	NewID() ID
}
