package flightreport

// NewInstance .
func NewInstance(gen Generator, flightoperationID FlightoperationID) *Flightreport {
	return &Flightreport{
		id:                gen.NewID(),
		flightoperationID: flightoperationID,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Flightreport {
	return &Flightreport{
		id:                ID(comp.GetID()),
		flightoperationID: FlightoperationID(comp.GetFlightoperationID()),
	}
}

// TakeApart .
func TakeApart(
	flightreport *Flightreport,
	component func(id, flightoperationID string),
) {
	component(
		string(flightreport.id),
		string(flightreport.flightoperationID),
	)
}

// Component .
type Component interface {
	GetID() string
	GetFlightoperationID() string
}
