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
				PointOrder:      w.GetPointOrder(),
				LatitudeDegree:  w.GetLatitudeDegree(),
				LongitudeDegree: w.GetLongitudeDegree(),
				RelativeHeightM: w.GetRelativeHeightM(),
				SpeedMS:         w.GetSpeedMS(),
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
	waypointComp func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64),
) {
	missionComp(
		string(mission.id),
	)
	for _, w := range mission.waypoints {
		waypointComp(
			w.PointOrder,
			w.LatitudeDegree,
			w.LongitudeDegree,
			w.RelativeHeightM,
			w.SpeedMS,
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
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetRelativeHeightM() float64
	GetSpeedMS() float64
}
