package flightreport

// NewInstance .
func NewInstance(gen Generator, name, description string, fleetID FleetID) *Flightreport {
	return &Flightreport{
		id:          gen.NewID(),
		name:        name,
		description: description,
		fleetID:     fleetID,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightreport {
	return &Flightreport{
		id:          ID(comp.GetID()),
		name:        comp.GetName(),
		description: comp.GetDescription(),
		fleetID:     FleetID(comp.GetFleetID()),
	}
}

// TakeApart .
func TakeApart(
	flightreport *Flightreport,
	component func(id, name, description, fleetID string),
) {
	component(
		string(flightreport.id),
		flightreport.name,
		flightreport.description,
		string(flightreport.fleetID),
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetFleetID() string
}
