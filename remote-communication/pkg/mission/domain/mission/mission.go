package mission

// ID .
type ID string

// Waypoint .
type Waypoint struct {
	PointOrder       int
	Latitude         float64
	Longitude        float64
	RelativeAltitude float64
	Speed            float64
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
	latitude float64,
	longitude float64,
	relativeAltitude float64,
	speed float64,
) int {
	order := len(m.waypoints) + 1
	m.waypoints = append(
		m.waypoints,
		&Waypoint{
			PointOrder:       order,
			Latitude:         latitude,
			Longitude:        longitude,
			RelativeAltitude: relativeAltitude,
			Speed:            speed,
		},
	)
	return order
}

// GetWaypoints .
func (m *Mission) GetWaypoints() []*Waypoint {
	return m.waypoints
}
