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
	pointOrder     int
	coordinates    GeodesicCoordinates
	relativeHeight Height
	speed          Speed
}

func NewWaypoint(
	pointOrder int,
	latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64,
) Waypoint {
	return Waypoint{
		pointOrder,
		NewGeodesicCoordinatesFromDegree(latitudeDegree, longitudeDegree),
		NewHeightFromM(relativeHeightM),
		NewSpeedFromMS(speedMS),
	}
}

// UploadID .
type UploadID string

// Navigation .
type Navigation struct {
	currentOrder                            int
	takeoffPointGroundHeightWGS84EllipsoidM Height
	waypoints                               []Waypoint
	uploadID                                UploadID
}

// NewNavigation .
func NewNavigation(takeoffPointGroundHeightWGS84EllipsoidM float64) *Navigation {
	return &Navigation{
		currentOrder:                            0,
		takeoffPointGroundHeightWGS84EllipsoidM: NewHeightFromM(takeoffPointGroundHeightWGS84EllipsoidM),
		waypoints:                               []Waypoint{},
	}
}

// GetTakeoffPointGroundHeightWGS84EllipsoidM .
func (n *Navigation) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return n.takeoffPointGroundHeightWGS84EllipsoidM.heightM
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

// GetUploadID .
func (n *Navigation) GetUploadID() UploadID {
	return n.uploadID
}

// ProvideWaypointsInterest .
func (n *Navigation) ProvideWaypointsInterest(
	waypoint func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64),
) {
	for _, w := range n.waypoints {
		waypoint(
			w.pointOrder,
			w.coordinates.latitudeDegree,
			w.coordinates.longitudeDegree,
			w.relativeHeight.heightM,
			w.speed.speedMS,
		)
	}
}
