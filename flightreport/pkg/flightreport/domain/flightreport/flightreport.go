package flightreport

// ID .
type ID string

// FlightoperationID .
type FlightoperationID string

// Flightreport .
type Flightreport struct {
	id                ID
	flightoperationID FlightoperationID
}

// GetID .
func (f *Flightreport) GetID() ID {
	return f.id
}

// GetFlightoperationID .
func (f *Flightreport) GetFlightoperationID() FlightoperationID {
	return f.flightoperationID
}

// Generator .
type Generator interface {
	NewID() ID
}
