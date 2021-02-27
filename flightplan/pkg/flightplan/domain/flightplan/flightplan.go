package flightplan

// ID .
type ID string

// Version .
type Version string

// Flightplan .
type Flightplan struct {
	id          ID
	name        string
	description string
	version     Version
	newVersion  Version
	gen         Generator
}

// GetID .
func (f *Flightplan) GetID() ID {
	return f.id
}

// GetName .
func (f *Flightplan) GetName() string {
	return f.name
}

// GetDescription .
func (f *Flightplan) GetDescription() string {
	return f.description
}

// GetVersion .
func (f *Flightplan) GetVersion() Version {
	return f.version
}

// GetNewVersion .
func (f *Flightplan) GetNewVersion() Version {
	return f.newVersion
}

// NameFlightplan .
func (f *Flightplan) NameFlightplan(name string) {
	f.name = name
	f.newVersion = f.gen.NewVersion()
}

// ChangeDescription .
func (f *Flightplan) ChangeDescription(description string) {
	f.description = description
	f.newVersion = f.gen.NewVersion()
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
}
