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

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Mission {
	mission := &Mission{
		id:           ID(comp.GetID()),
		name:         comp.GetName(),
		isCarbonCopy: comp.GetIsCarbonCopy(),
		version:      Version(comp.GetVersion()),
		newVersion:   Version(comp.GetVersion()),
		gen:          gen,
	}

	navigation := NewNavigation(comp.GetTakeoffPointGroundHeightWGS84M())
	for _, waypointComp := range comp.GetWaypoints() {
		navigation.PushNextWaypoint(
			waypointComp.GetLatitudeDegree(),
			waypointComp.GetLongitudeDegree(),
			waypointComp.GetHeightWGS84M(),
			waypointComp.GetSpeedMS(),
		)
	}
	mission.navigation = navigation

	return mission
}

// TakeApart .
func TakeApart(
	mission *Mission,
	component func(id, name, version string, takeoffPointGroundHeightM float64, isCarbonCopy bool),
	waypointComponent func(order int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64),
) {
	component(
		string(mission.id),
		mission.name,
		string(mission.version),
		mission.navigation.GetTakeoffPointGroundHeightM(),
		mission.isCarbonCopy,
	)
	mission.navigation.ProvideWaypointsInterest(
		waypointComponent,
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetTakeoffPointGroundHeightWGS84M() float64
	GetIsCarbonCopy() bool
	GetVersion() string
	GetWaypoints() []WaypointComponent
}

// WaypointComponent .
type WaypointComponent interface {
	GetOrder() int
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetHeightWGS84M() float64
	GetSpeedMS() float64
}
