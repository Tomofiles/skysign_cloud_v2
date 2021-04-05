package flightoperation

// NewInstance .
func NewInstance(gen Generator, flightplanID FlightplanID) *Flightoperation {
	version := gen.NewVersion()
	return &Flightoperation{
		id:           gen.NewID(),
		flightplanID: flightplanID,
		isCompleted:  Operating,
		version:      version,
		newVersion:   version,
		gen:          gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightoperation {
	return &Flightoperation{
		id:           ID(comp.GetID()),
		flightplanID: FlightplanID(comp.GetFlightplanID()),
		isCompleted:  comp.GetIsCompleted(),
		version:      Version(comp.GetVersion()),
		newVersion:   Version(comp.GetVersion()),
		gen:          gen,
	}
}

// TakeApart .
func TakeApart(
	flightoperation *Flightoperation,
	component func(id, flightplanID, version string, isCompleted bool),
) {
	component(
		string(flightoperation.id),
		string(flightoperation.flightplanID),
		string(flightoperation.version),
		flightoperation.isCompleted,
	)
}

// Component .
type Component interface {
	GetID() string
	GetFlightplanID() string
	GetIsCompleted() bool
	GetVersion() string
}
