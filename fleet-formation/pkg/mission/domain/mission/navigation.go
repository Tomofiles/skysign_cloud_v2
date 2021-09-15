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

// Altitude .
type Altitude struct {
	altitudeM float64
}

// NewAltitudeFromM .
func NewAltitudeFromM(altitudeM float64) Altitude {
	return Altitude{altitudeM: altitudeM}
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
	pointOrder       int
	coordinates      GeodesicCoordinates
	relativeAltitude Altitude
	speed            Speed
}

func NewWaypoint(
	pointOrder int,
	latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64,
) Waypoint {
	return Waypoint{
		pointOrder,
		NewGeodesicCoordinatesFromDegree(latitudeDegree, longitudeDegree),
		NewAltitudeFromM(relativeAltitudeM),
		NewSpeedFromMS(speedMS),
	}
}

// UploadID .
type UploadID string

// Navigation .
type Navigation struct {
	currentOrder                int
	takeoffPointGroundAltitudeM Altitude
	waypoints                   []Waypoint
	uploadID                    UploadID
}

// NewNavigation .
func NewNavigation(takeoffPointGroundAltitudeM float64) *Navigation {
	return &Navigation{
		currentOrder:                0,
		takeoffPointGroundAltitudeM: NewAltitudeFromM(takeoffPointGroundAltitudeM),
		waypoints:                   []Waypoint{},
	}
}

// GetTakeoffPointGroundAltitudeM .
func (n *Navigation) GetTakeoffPointGroundAltitudeM() float64 {
	return n.takeoffPointGroundAltitudeM.altitudeM
}

// PushNextWaypoint .
func (n *Navigation) PushNextWaypoint(
	latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64,
) {
	n.currentOrder = n.currentOrder + 1
	n.waypoints = append(
		n.waypoints,
		NewWaypoint(
			n.currentOrder,
			latitudeDegree,
			longitudeDegree,
			relativeAltitudeM,
			speedMS,
		),
	)
}

// GetUploadID .
func (n *Navigation) GetUploadID() UploadID {
	return n.uploadID
}

// ProvideWaypointsInterest .
func (n *Navigation) ProvideWaypointsInterest(
	waypoint func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64),
) {
	for _, w := range n.waypoints {
		waypoint(
			w.pointOrder,
			w.coordinates.latitudeDegree,
			w.coordinates.longitudeDegree,
			w.relativeAltitude.altitudeM,
			w.speed.speedMS,
		)
	}
}
