package mission

// GeodesicCoordinates .
type GeodesicCoordinates struct {
	latitudeDegree  float64
	longitudeDegree float64
}

func NewGeodesicCoordinatesFromDegree(
	latitudeDegree, longitudeDegree float64,
) GeodesicCoordinates {
	return GeodesicCoordinates{
		latitudeDegree:  latitudeDegree,
		longitudeDegree: longitudeDegree,
	}
}

// Height .
type Height struct {
	heightM float64
}

// NewHeightFromM .
func NewHeightFromM(heightM float64) Height {
	return Height{heightM: heightM}
}

// Speed .
type Speed struct {
	speedMS float64
}

// NewSpeedFromMS .
func NewSpeedFromMS(speedMS float64) Speed {
	return Speed{speedMS: speedMS}
}

// Waypoint .
type Waypoint struct {
	order          int
	coordinates    GeodesicCoordinates
	relativeHeight Height
	speed          Speed
}

func NewWaypoint(
	order int,
	latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64,
) Waypoint {
	return Waypoint{
		order,
		NewGeodesicCoordinatesFromDegree(latitudeDegree, longitudeDegree),
		NewHeightFromM(relativeHeightM),
		NewSpeedFromMS(speedMS),
	}
}

// Navigation .
type Navigation struct {
	currentOrder             int
	takeoffPointGroundHeight Height
	waypoints                []Waypoint
}

// NewNavigation .
func NewNavigation(takeoffPointGroundHeightM float64) *Navigation {
	return &Navigation{
		currentOrder:             1,
		takeoffPointGroundHeight: NewHeightFromM(takeoffPointGroundHeightM),
		waypoints:                []Waypoint{},
	}
}

// GetTakeoffPointGroundHeightM .
func (n *Navigation) GetTakeoffPointGroundHeightM() float64 {
	return n.takeoffPointGroundHeight.heightM
}

// PushNextWaypoint .
func (n *Navigation) PushNextWaypoint(
	latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64,
) {
	n.currentOrder = n.currentOrder + 1
	n.waypoints = append(
		n.waypoints,
		NewWaypoint(
			n.currentOrder,
			latitudeDegree,
			longitudeDegree,
			relativeHeightM,
			speedMS,
		),
	)
}

// ProvideWaypointsInterest .
func (n *Navigation) ProvideWaypointsInterest(
	waypoint func(latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64),
) {
	for _, w := range n.waypoints {
		waypoint(
			w.coordinates.latitudeDegree,
			w.coordinates.longitudeDegree,
			w.relativeHeight.heightM,
			w.speed.speedMS,
		)
	}
}
