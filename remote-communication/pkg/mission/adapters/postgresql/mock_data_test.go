package postgresql

import (
	"remote-communication/pkg/mission/domain/mission"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultMissionID = mission.ID("mission-id")

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

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
