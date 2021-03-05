package flightoperation

// NewInstance .
func NewInstance(gen Generator, flightplanID FlightplanID) *Flightoperation {
	return &Flightoperation{
		id:           gen.NewID(),
		flightplanID: flightplanID,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightoperation {
	return &Flightoperation{
		id:           ID(comp.GetID()),
		flightplanID: FlightplanID(comp.GetFlightplanID()),
	}
}

// TakeApart .
func TakeApart(
	flightoperation *Flightoperation,
	component func(id, flightplanID string),
) {
	component(
		string(flightoperation.id),
		string(flightoperation.flightplanID),
	)
}

// Component .
type Component interface {
	GetID() string
	GetFlightplanID() string
}
