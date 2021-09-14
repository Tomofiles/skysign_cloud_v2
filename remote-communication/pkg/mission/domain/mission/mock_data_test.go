package mission

const DefaultID = ID("mission-id")

// Mission構成オブジェクトモック
type missionComponentMock struct {
	id        string
	waypoints []*waypointComponentMock
}

func (m *missionComponentMock) GetID() string {
	return m.id
}

func (m *missionComponentMock) GetWaypoints() []WaypointComponent {
	var waypoints []WaypointComponent
	for _, w := range m.waypoints {
		waypoints = append(waypoints, w)
	}
	return waypoints
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	pointOrder       int
	latitude         float64
	longitude        float64
	relativeAltitude float64
	speed            float64
}

func (m *waypointComponentMock) GetPointOrder() int {
	return m.pointOrder
}

func (m *waypointComponentMock) GetLatitude() float64 {
	return m.latitude
}

func (m *waypointComponentMock) GetLongitude() float64 {
	return m.longitude
}

func (m *waypointComponentMock) GetRelativeAltitude() float64 {
	return m.relativeAltitude
}

func (m *waypointComponentMock) GetSpeed() float64 {
	return m.speed
}
