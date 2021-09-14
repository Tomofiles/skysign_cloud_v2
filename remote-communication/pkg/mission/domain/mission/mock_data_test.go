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
	pointOrder        int
	latitudeDegree    float64
	longitudeDegree   float64
	relativeAltitudeM float64
	speedMS           float64
}

func (m *waypointComponentMock) GetPointOrder() int {
	return m.pointOrder
}

func (m *waypointComponentMock) GetLatitudeDegree() float64 {
	return m.latitudeDegree
}

func (m *waypointComponentMock) GetLongitudeDegree() float64 {
	return m.longitudeDegree
}

func (m *waypointComponentMock) GetRelativeAltitudeM() float64 {
	return m.relativeAltitudeM
}

func (m *waypointComponentMock) GetSpeedMS() float64 {
	return m.speedMS
}
