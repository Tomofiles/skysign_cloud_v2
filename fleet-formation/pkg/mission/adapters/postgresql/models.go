package postgresql

import m "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/domain/mission"

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
	MissionID                   string `gorm:"primaryKey"`
	TakeoffPointGroundAltitudeM float64
	Waypoints                   []*Waypoint `gorm:"-"`
	UploadID                    string
}

// GetMissionID .
func (v *Navigation) GetMissionID() string {
	return v.MissionID
}

// GetTakeoffPointGroundAltitudeM .
func (v *Navigation) GetTakeoffPointGroundAltitudeM() float64 {
	return v.TakeoffPointGroundAltitudeM
}

// GetWaypoints .
func (v *Navigation) GetWaypoints() []m.WaypointComponent {
	waypoints := []m.WaypointComponent{}
	for _, w := range v.Waypoints {
		waypoints = append(waypoints, w)
	}
	return waypoints
}

// GetUploadID .
func (v *Navigation) GetUploadID() string {
	return v.UploadID
}

// Waypoint
type Waypoint struct {
	MissionID         string
	PointOrder        int
	LatitudeDegree    float64
	LongitudeDegree   float64
	RelativeAltitudeM float64
	SpeedMS           float64
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

// GetRelativeAltitudeM .
func (v *Waypoint) GetRelativeAltitudeM() float64 {
	return v.RelativeAltitudeM
}

// GetSpeedMS .
func (v *Waypoint) GetSpeedMS() float64 {
	return v.SpeedMS
}
