package flightoperation

// ID .
type ID string

// FlightplanID .
type FlightplanID string

// Flightoperation .
type Flightoperation struct {
	id           ID
	flightplanID FlightplanID
}

// GetID .
func (f *Flightoperation) GetID() ID {
	return f.id
}

// GetFlightplanID .
func (f *Flightoperation) GetFlightplanID() FlightplanID {
	return f.flightplanID
}

// Generator .
type Generator interface {
	NewID() ID
	NewFlightplanID() FlightplanID
}
