package flightoperation

// NewInstance .
func NewInstance(gen Generator, flightplanID FlightplanID) *Flightoperation {
	return &Flightoperation{
		id:           gen.NewID(),
		flightplanID: flightplanID,
	}
}

// // AssembleFrom .
// func AssembleFrom(gen Generator, comp Component) *Flightplan {
// 	return &Flightplan{
// 		id:          ID(comp.GetID()),
// 		name:        comp.GetName(),
// 		description: comp.GetDescription(),
// 		version:     Version(comp.GetVersion()),
// 		newVersion:  Version(comp.GetVersion()),
// 		gen:         gen,
// 	}
// }

// // TakeApart .
// func TakeApart(
// 	flightplan *Flightplan,
// 	component func(id, name, description, version string),
// ) {
// 	component(
// 		string(flightplan.id),
// 		flightplan.name,
// 		flightplan.description,
// 		string(flightplan.version),
// 	)
// }

// // Component .
// type Component interface {
// 	GetID() string
// 	GetName() string
// 	GetDescription() string
// 	GetVersion() string
// }
