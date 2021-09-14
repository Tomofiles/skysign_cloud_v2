package mission

// NewInstance .
func NewInstance(id ID) *Mission {
	return &Mission{
		id:        id,
		waypoints: []*Waypoint{},
	}
}

// AssembleFrom .
func AssembleFrom(comp Component) *Mission {
	var waypoints []*Waypoint
	for _, w := range comp.GetWaypoints() {
		waypoints = append(
			waypoints,
			&Waypoint{
				PointOrder:       w.GetPointOrder(),
				Latitude:         w.GetLatitude(),
				Longitude:        w.GetLongitude(),
				RelativeAltitude: w.GetRelativeAltitude(),
				Speed:            w.GetSpeed(),
			},
		)
	}
	return &Mission{
		id:        ID(comp.GetID()),
		waypoints: waypoints,
	}
}

// TakeApart .
func TakeApart(
	mission *Mission,
	missionComp func(id string),
	waypointComp func(pointOrder int, latitude, longitude, relativeAltitude, speed float64),
) {
	missionComp(
		string(mission.id),
	)
	for _, w := range mission.waypoints {
		waypointComp(
			w.PointOrder,
			w.Latitude,
			w.Longitude,
			w.RelativeAltitude,
			w.Speed,
		)
	}
}

// Component .
type Component interface {
	GetID() string
	GetWaypoints() []WaypointComponent
}

// WaypointComponent .
type WaypointComponent interface {
	GetPointOrder() int
	GetLatitude() float64
	GetLongitude() float64
	GetRelativeAltitude() float64
	GetSpeed() float64
}
