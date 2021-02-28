package flightplan

// NewInstance .
func NewInstance(gen Generator) *Flightplan {
	id := gen.NewID()
	version := gen.NewVersion()
	return &Flightplan{
		id:           id,
		isCarbonCopy: Original,
		version:      version,
		newVersion:   version,
		gen:          gen,
	}
}

// Copy .
func Copy(gen Generator, id ID, original *Flightplan) *Flightplan {
	return &Flightplan{
		id:           id,
		name:         original.name,
		description:  original.description,
		isCarbonCopy: CarbonCopy,
		version:      original.version,
		newVersion:   original.newVersion,
		gen:          gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightplan {
	return &Flightplan{
		id:           ID(comp.GetID()),
		name:         comp.GetName(),
		description:  comp.GetDescription(),
		isCarbonCopy: comp.GetIsCarbonCopy(),
		version:      Version(comp.GetVersion()),
		newVersion:   Version(comp.GetVersion()),
		gen:          gen,
	}
}

// TakeApart .
func TakeApart(
	flightplan *Flightplan,
	component func(id, name, description, version string, isCarbonCopy bool),
) {
	component(
		string(flightplan.id),
		flightplan.name,
		flightplan.description,
		string(flightplan.version),
		flightplan.isCarbonCopy,
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetIsCarbonCopy() bool
	GetVersion() string
}
