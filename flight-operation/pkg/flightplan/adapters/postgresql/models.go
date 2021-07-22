package postgresql

// Flightplan .
type Flightplan struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	FleetID     string
	Version     string
}

// GetID .
func (f *Flightplan) GetID() string {
	return f.ID
}

// GetName .
func (f *Flightplan) GetName() string {
	return f.Name
}

// GetDescription .
func (f *Flightplan) GetDescription() string {
	return f.Description
}

// GetFleetID .
func (f *Flightplan) GetFleetID() string {
	return f.FleetID
}

// GetVersion .
func (f *Flightplan) GetVersion() string {
	return f.Version
}
