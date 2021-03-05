package postgresql

// Flightoperation .
type Flightoperation struct {
	ID           string `gorm:"primaryKey"`
	FlightplanID string
}

// GetID .
func (f *Flightoperation) GetID() string {
	return f.ID
}

// GetFlightplanID .
func (f *Flightoperation) GetFlightplanID() string {
	return f.FlightplanID
}
