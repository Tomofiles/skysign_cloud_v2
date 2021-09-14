package postgresql

import (
	"remote-communication/pkg/mission/domain/mission"
)

const DefaultMissionID = mission.ID("mission-id")

// Mission構成オブジェクトモック
type missionComponentMock struct {
	id        string
	waypoints []*waypointComponentMock
}

func (m *missionComponentMock) GetID() string {
	return m.id
}

func (m *missionComponentMock) GetWaypoints() []mission.WaypointComponent {
	var waypoints []mission.WaypointComponent
	for _, w := range m.waypoints {
		waypoints = append(waypoints, w)
	}
	return waypoints
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	pointOrder      int
	latitudeDegree  float64
	longitudeDegree float64
	relativeHeightM float64
	speedMS         float64
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

func (m *waypointComponentMock) GetRelativeHeightM() float64 {
	return m.relativeHeightM
}

func (m *waypointComponentMock) GetSpeedMS() float64 {
	return m.speedMS
}
