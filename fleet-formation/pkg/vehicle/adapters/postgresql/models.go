package postgresql

// Vehicle .
type Vehicle struct {
	ID              string `gorm:"primaryKey"`
	Name            string
	CommunicationID string
	IsCarbonCopy    bool
	Version         string
}

// GetID .
func (v *Vehicle) GetID() string {
	return v.ID
}

// GetName .
func (v *Vehicle) GetName() string {
	return v.Name
}

// GetCommunicationID .
func (v *Vehicle) GetCommunicationID() string {
	return v.CommunicationID
}

// GetIsCarbonCopy .
func (v *Vehicle) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

// GetVersion .
func (v *Vehicle) GetVersion() string {
	return v.Version
}
