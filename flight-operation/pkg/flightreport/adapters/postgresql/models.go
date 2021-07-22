package postgresql

// Flightreport .
type Flightreport struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	FleetID     string
}

// GetID .
func (f *Flightreport) GetID() string {
	return f.ID
}

// GetName .
func (f *Flightreport) GetName() string {
	return f.Name
}

// GetDescription .
func (f *Flightreport) GetDescription() string {
	return f.Description
}

// GetFleetID .
func (f *Flightreport) GetFleetID() string {
	return f.FleetID
}
