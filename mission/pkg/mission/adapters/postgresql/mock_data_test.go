package postgresql

import (
	m "mission/pkg/mission/domain/mission"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultMissionID = m.ID("mission-id")
const DefaultMissionVersion = m.Version("version")
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM float64 = 10
const DefaultMissionUploadID = m.UploadID("upload-id")

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
	ID           string
	Name         string
	Navigation   navigationComponentMock
	IsCarbonCopy bool
	Version      string
}

func (v *missionComponentMock) GetID() string {
	return v.ID
}

func (v *missionComponentMock) GetName() string {
	return v.Name
}

func (v *missionComponentMock) GetNavigation() m.NavigationComponent {
	return &v.Navigation
}

func (v *missionComponentMock) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

func (v *missionComponentMock) GetVersion() string {
	return v.Version
}

// Navigation構成オブジェクトモック
type navigationComponentMock struct {
	TakeoffPointGroundHeightWGS84EllipsoidM float64
	Waypoints                               []waypointComponentMock
	UploadID                                string
}

func (v *navigationComponentMock) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return v.TakeoffPointGroundHeightWGS84EllipsoidM
}

func (v *navigationComponentMock) GetWaypoints() []m.WaypointComponent {
	var waypoints []m.WaypointComponent
	for _, w := range v.Waypoints {
		waypoints = append(
			waypoints,
			&waypointComponentMock{
				PointOrder:      w.PointOrder,
				LatitudeDegree:  w.LatitudeDegree,
				LongitudeDegree: w.LongitudeDegree,
				RelativeHeightM: w.RelativeHeightM,
				SpeedMS:         w.SpeedMS,
			},
		)
	}
	return waypoints
}

// GetUploadID .
func (v *navigationComponentMock) GetUploadID() string {
	return v.UploadID
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	PointOrder                                                int
	LatitudeDegree, LongitudeDegree, RelativeHeightM, SpeedMS float64
}

func (v *waypointComponentMock) GetPointOrder() int {
	return v.PointOrder
}

func (v *waypointComponentMock) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

func (v *waypointComponentMock) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

func (v *waypointComponentMock) GetRelativeHeightM() float64 {
	return v.RelativeHeightM
}

func (v *waypointComponentMock) GetSpeedMS() float64 {
	return v.SpeedMS
}
