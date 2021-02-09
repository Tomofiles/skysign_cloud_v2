package flightplan

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

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetVersion() string
}