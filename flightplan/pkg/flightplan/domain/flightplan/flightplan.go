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

// NewInstance .
func NewInstance(generator Generator) *Flightplan {
	id := generator.NewID()
	version := generator.NewVersion()
	return &Flightplan{
		id:         id,
		version:    version,
		newVersion: version,
		gen:        generator,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightplan {
	return &Flightplan{
		id:          ID(comp.GetID()),
		name:        comp.GetName(),
		description: comp.GetDescription(),
		version:     Version(comp.GetVersion()),
		newVersion:  Version(comp.GetVersion()),
		gen:         gen,
	}
}

// TakeApart .
func TakeApart(
	flightplan *Flightplan,
	component func(id, name, description, version string),
) {
	component(
		string(flightplan.id),
		flightplan.name,
		flightplan.description,
		string(flightplan.version),
	)
}

// Generator .
type Generator interface {
	NewID() ID
	NewVersion() Version
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetVersion() string
}
