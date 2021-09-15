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
			original.navigation.GetTakeoffPointGroundAltitudeM(),
		)
		original.navigation.ProvideWaypointsInterest(
			func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64) {
				navigation.PushNextWaypoint(
					latitudeDegree,
					longitudeDegree,
					relativeAltitudeM,
					speedMS,
				)
			},
		)
		navigation.uploadID = gen.NewUploadID()
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

	navigation := NewNavigation(comp.GetNavigation().GetTakeoffPointGroundAltitudeM())
	for _, waypointComp := range comp.GetNavigation().GetWaypoints() {
		navigation.PushNextWaypoint(
			waypointComp.GetLatitudeDegree(),
			waypointComp.GetLongitudeDegree(),
			waypointComp.GetRelativeAltitudeM(),
			waypointComp.GetSpeedMS(),
		)
	}
	navigation.uploadID = UploadID(comp.GetNavigation().GetUploadID())
	mission.navigation = navigation

	return mission
}

// TakeApart .
func TakeApart(
	mission *Mission,
	component func(id, name, version string, isCarbonCopy bool),
	navigationComponent func(takeoffPointGroundAltitudeM float64, uploadID string),
	waypointComponent func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64),
) {
	component(
		string(mission.id),
		mission.name,
		string(mission.version),
		mission.isCarbonCopy,
	)
	navigationComponent(
		mission.navigation.GetTakeoffPointGroundAltitudeM(),
		string(mission.navigation.uploadID),
	)
	mission.navigation.ProvideWaypointsInterest(
		waypointComponent,
	)
}

// Component .
type Component interface {
	GetID() string
	GetName() string
	GetNavigation() NavigationComponent
	GetIsCarbonCopy() bool
	GetVersion() string
}

// NavigationComponent .
type NavigationComponent interface {
	GetTakeoffPointGroundAltitudeM() float64
	GetWaypoints() []WaypointComponent
	GetUploadID() string
}

// WaypointComponent .
type WaypointComponent interface {
	GetPointOrder() int
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
}
