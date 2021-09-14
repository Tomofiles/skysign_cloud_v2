package mission

// ID .
type ID string

// Waypoint .
type Waypoint struct {
	PointOrder      int
	LatitudeDegree  float64
	LongitudeDegree float64
	RelativeHeightM float64
	SpeedMS         float64
}

// Mission .
type Mission struct {
	id        ID
	waypoints []*Waypoint
}

// GetID .
func (m *Mission) GetID() ID {
	return m.id
}

// PushWaypoint .
func (m *Mission) PushWaypoint(
	latitudeDegree float64,
	longitudeDegree float64,
	relativeHeightM float64,
	speedMS float64,
) int {
	order := len(m.waypoints) + 1
	m.waypoints = append(
		m.waypoints,
		&Waypoint{
			PointOrder:      order,
			LatitudeDegree:  latitudeDegree,
			LongitudeDegree: longitudeDegree,
			RelativeHeightM: relativeHeightM,
			SpeedMS:         speedMS,
		},
	)
	return order
}

// GetWaypoints .
func (m *Mission) GetWaypoints() []*Waypoint {
	return m.waypoints
}
