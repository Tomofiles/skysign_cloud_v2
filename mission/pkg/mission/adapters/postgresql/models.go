package postgresql

import m "mission/pkg/mission/domain/mission"

// Mission .
type Mission struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Navigation   *Navigation `gorm:"-"`
	IsCarbonCopy bool
	Version      string
}

// GetID .
func (v *Mission) GetID() string {
	return v.ID
}

// GetName .
func (v *Mission) GetName() string {
	return v.Name
}

// GetIsCarbonCopy .
func (v *Mission) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

// GetVersion .
func (v *Mission) GetVersion() string {
	return v.Version
}

// GetNavigation .
func (v *Mission) GetNavigation() m.NavigationComponent {
	return v.Navigation
}

// Navigation .
type Navigation struct {
	MissionID                               string `gorm:"primaryKey"`
	TakeoffPointGroundHeightWGS84EllipsoidM float64
	Waypoints                               []*Waypoint `gorm:"-"`
}

// GetMissionID .
func (v *Navigation) GetMissionID() string {
	return v.MissionID
}

// GetTakeoffPointGroundHeightWGS84EllipsoidM .
func (v *Navigation) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return v.TakeoffPointGroundHeightWGS84EllipsoidM
}

// GetWaypoints .
func (v *Navigation) GetWaypoints() []m.WaypointComponent {
	waypoints := []m.WaypointComponent{}
	for _, w := range v.Waypoints {
		waypoints = append(waypoints, w)
	}
	return waypoints
}

// Waypoint
type Waypoint struct {
	MissionID       string
	PointOrder      int
	LatitudeDegree  float64
	LongitudeDegree float64
	RelativeHeightM float64
	SpeedMS         float64
}

// GetMissionID .
func (v *Waypoint) GetMissionID() string {
	return v.MissionID
}

// GetPointOrder .
func (v *Waypoint) GetPointOrder() int {
	return v.PointOrder
}

// GetLatitude .
func (v *Waypoint) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

// GetLongitude .
func (v *Waypoint) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

// GetRelativeHeightM .
func (v *Waypoint) GetRelativeHeightM() float64 {
	return v.RelativeHeightM
}

// GetSpeedMS .
func (v *Waypoint) GetSpeedMS() float64 {
	return v.SpeedMS
}
