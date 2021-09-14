package postgresql

import (
	"remote-communication/pkg/mission/domain/mission"
)

// Mission .
type Mission struct {
	ID        string      `gorm:"primaryKey"`
	Waypoints []*Waypoint `gorm:"-"`
}

// GetID .
func (m *Mission) GetID() string {
	return m.ID
}

// GetWaypoints .
func (m *Mission) GetWaypoints() []mission.WaypointComponent {
	var waypoints []mission.WaypointComponent
	for _, w := range m.Waypoints {
		waypoints = append(waypoints, w)
	}
	return waypoints
}

// Waypoint .
type Waypoint struct {
	MissionID         string `gorm:"primaryKey"`
	PointOrder        int
	LatitudeDegree    float64
	LongitudeDegree   float64
	RelativeAltitudeM float64
	SpeedMS           float64
}

// GetMissionID .
func (w *Waypoint) GetMissionID() string {
	return w.MissionID
}

// GetPointOrder .
func (w *Waypoint) GetPointOrder() int {
	return w.PointOrder
}

// GetLatitudeDegree .
func (w *Waypoint) GetLatitudeDegree() float64 {
	return w.LatitudeDegree
}

// GetLongitudeDegree .
func (w *Waypoint) GetLongitudeDegree() float64 {
	return w.LongitudeDegree
}

// GetRelativeHeightM .
func (w *Waypoint) GetRelativeHeightM() float64 {
	return w.RelativeAltitudeM
}

// GetSpeedMS .
func (w *Waypoint) GetSpeedMS() float64 {
	return w.SpeedMS
}
