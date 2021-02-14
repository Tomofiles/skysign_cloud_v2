package flightplan

// ID .
type ID string

// Version .
type Version string

// Flightplan .
type Flightplan struct {
	id               ID
	name             string
	numberOfVehicles int
	version          Version
	newVersion       Version
	generator        Generator
}

// GetID .
func (f *Flightplan) GetID() ID {
	return f.id
}

// GetName .
func (f *Flightplan) GetName() string {
	return f.name
}

// GetNumberOfVehicles .
func (f *Flightplan) GetNumberOfVehicles() int {
	return f.numberOfVehicles
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
	f.newVersion = f.generator.NewVersion()
}

// NewInstance .
func NewInstance(generator Generator) *Flightplan {
	id := generator.NewID()
	version := generator.NewVersion()
	return &Flightplan{
		id:         id,
		version:    version,
		newVersion: version,
		generator:  generator,
	}
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
	NewAssignmentID() AssignmentID
}
