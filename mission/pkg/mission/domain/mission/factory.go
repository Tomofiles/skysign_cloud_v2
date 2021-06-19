package mission

// NewInstance .
func NewInstance(gen Generator) *Mission {
	id := gen.NewID()
	version := gen.NewVersion()
	return &Mission{
		id:           id,
		isCarbonCopy: Original,
		version:      version,
		newVersion:   version,
		gen:          gen,
	}
}

// Copy .
func Copy(gen Generator, id ID, original *Mission) *Mission {
	mission := &Mission{
		id:           id,
		name:         original.name,
		isCarbonCopy: CarbonCopy,
		version:      original.version,
		newVersion:   original.newVersion,
		gen:          gen,
	}

	if original.navigation != nil {
		navigation := NewNavigation(
			original.navigation.takeoffPointGroundHeight.heightM,
		)
		original.navigation.ProvideWaypointsInterest(
			func(latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
				navigation.PushNextWaypoint(
					latitudeDegree,
					longitudeDegree,
					relativeHeightM,
					speedMS,
				)
			},
		)
		mission.navigation = navigation
	}

	return mission
}

// // AssembleFrom .
// func AssembleFrom(gen Generator, comp Component) *Vehicle {
// 	return &Vehicle{
// 		id:              ID(comp.GetID()),
// 		name:            comp.GetName(),
// 		communicationID: CommunicationID(comp.GetCommunicationID()),
// 		isCarbonCopy:    comp.GetIsCarbonCopy(),
// 		version:         Version(comp.GetVersion()),
// 		newVersion:      Version(comp.GetVersion()),
// 		gen:             gen,
// 	}
// }

// // TakeApart .
// func TakeApart(
// 	vehicle *Vehicle,
// 	component func(id, name, communicationID, version string, isCarbonCopy bool),
// ) {
// 	component(
// 		string(vehicle.id),
// 		vehicle.name,
// 		string(vehicle.communicationID),
// 		string(vehicle.version),
// 		vehicle.isCarbonCopy,
// 	)
// }

// // Component .
// type Component interface {
// 	GetID() string
// 	GetName() string
// 	GetCommunicationID() string
// 	GetIsCarbonCopy() bool
// 	GetVersion() string
// }
