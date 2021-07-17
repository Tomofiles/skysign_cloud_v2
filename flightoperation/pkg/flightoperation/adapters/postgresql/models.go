package postgresql

// Flightoperation .
type Flightoperation struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	FleetID     string
	IsCompleted bool
	Version     string
}

// GetID .
func (f *Flightoperation) GetID() string {
	return f.ID
}

// GetName .
func (f *Flightoperation) GetName() string {
	return f.Name
}

// GetDescription .
func (f *Flightoperation) GetDescription() string {
	return f.Description
}

// GetFleetID .
func (f *Flightoperation) GetFleetID() string {
	return f.FleetID
}

// GetIsCompleted .
func (f *Flightoperation) GetIsCompleted() bool {
	return f.IsCompleted
}

// GetVersion .
func (f *Flightoperation) GetVersion() string {
	return f.Version
}
