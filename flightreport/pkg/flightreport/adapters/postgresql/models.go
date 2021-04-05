package postgresql

// Flightreport .
type Flightreport struct {
	ID                string `gorm:"primaryKey"`
	FlightoperationID string
}

// GetID .
func (f *Flightreport) GetID() string {
	return f.ID
}

// GetFlightoperationID .
func (f *Flightreport) GetFlightoperationID() string {
	return f.FlightoperationID
}
