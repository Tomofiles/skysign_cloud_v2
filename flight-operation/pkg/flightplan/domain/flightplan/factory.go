package flightplan

// NewInstance .
func NewInstance(gen Generator) *Flightplan {
	id := gen.NewID()
	version := gen.NewVersion()
	return &Flightplan{
		id:         id,
		fleetID:    BlankFleetID,
		version:    version,
		newVersion: version,
		gen:        gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightplan {
	return &Flightplan{
		id:          ID(comp.GetID()),
		name:        comp.GetName(),
		description: comp.GetDescription(),
		fleetID:     FleetID(comp.GetFleetID()),
		version:     Version(comp.GetVersion()),
		newVersion:  Version(comp.GetVersion()),
		gen:         gen,
	}
}

// TakeApart .
func TakeApart(
	flightplan *Flightplan,
	component func(id, name, description, fleetID, version string),
) {
	component(
		string(flightplan.id),
		flightplan.name,
		flightplan.description,
		string(flightplan.fleetID),
		string(flightplan.version),
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
	GetVersion() string
}
