package flightoperation

// NewInstance .
func NewInstance(gen Generator, fleetID FleetID) *Flightoperation {
	version := gen.NewVersion()
	return &Flightoperation{
		id:          gen.NewID(),
		fleetID:     fleetID,
		isCompleted: Operating,
		version:     version,
		newVersion:  version,
		gen:         gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightoperation {
	return &Flightoperation{
		id:          ID(comp.GetID()),
		name:        comp.GetName(),
		description: comp.GetDescription(),
		fleetID:     FleetID(comp.GetFleetID()),
		isCompleted: comp.GetIsCompleted(),
		version:     Version(comp.GetVersion()),
		newVersion:  Version(comp.GetVersion()),
		gen:         gen,
	}
}

// TakeApart .
func TakeApart(
	flightoperation *Flightoperation,
	component func(id, name, description, fleetID, version string, isCompleted bool),
) {
	component(
		string(flightoperation.id),
		flightoperation.name,
		flightoperation.description,
		string(flightoperation.fleetID),
		string(flightoperation.version),
		flightoperation.isCompleted,
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
	GetIsCompleted() bool
	GetVersion() string
}
