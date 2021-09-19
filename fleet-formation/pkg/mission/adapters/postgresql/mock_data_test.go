package postgresql

import (
	m "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/domain/mission"
)

const DefaultMissionID = m.ID("mission-id")
const DefaultMissionVersion = m.Version("version")
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundAltitudeM float64 = 10
const DefaultMissionUploadID = m.UploadID("upload-id")

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
	TakeoffPointGroundAltitudeM float64
	Waypoints                   []waypointComponentMock
	UploadID                    string
}

func (v *navigationComponentMock) GetTakeoffPointGroundAltitudeM() float64 {
	return v.TakeoffPointGroundAltitudeM
}

func (v *navigationComponentMock) GetWaypoints() []m.WaypointComponent {
	var waypoints []m.WaypointComponent
	for _, w := range v.Waypoints {
		waypoints = append(
			waypoints,
			&waypointComponentMock{
				PointOrder:        w.PointOrder,
				LatitudeDegree:    w.LatitudeDegree,
				LongitudeDegree:   w.LongitudeDegree,
				RelativeAltitudeM: w.RelativeAltitudeM,
				SpeedMS:           w.SpeedMS,
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
	PointOrder                                                  int
	LatitudeDegree, LongitudeDegree, RelativeAltitudeM, SpeedMS float64
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

func (v *waypointComponentMock) GetRelativeAltitudeM() float64 {
	return v.RelativeAltitudeM
}

func (v *waypointComponentMock) GetSpeedMS() float64 {
	return v.SpeedMS
}
