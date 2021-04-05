package postgresql

// Flightoperation .
type Flightoperation struct {
	ID           string `gorm:"primaryKey"`
	FlightplanID string
	IsCompleted  bool
	Version      string
}

// GetID .
func (f *Flightoperation) GetID() string {
	return f.ID
}

// GetFlightplanID .
func (f *Flightoperation) GetFlightplanID() string {
	return f.FlightplanID
}

// GetIsCompleted .
func (f *Flightoperation) GetIsCompleted() bool {
	return f.IsCompleted
}

// GetVersion .
func (f *Flightoperation) GetVersion() string {
	return f.Version
}
