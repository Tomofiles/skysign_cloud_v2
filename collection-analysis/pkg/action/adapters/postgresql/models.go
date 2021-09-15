package postgresql

import "collection-analysis/pkg/action/domain/action"

// Action .
type Action struct {
	ID               string `gorm:"primaryKey"`
	CommunicationID  string
	FleetID          string
	IsCompleted      bool
	TrajectoryPoints []*TrajectoryPoint `gorm:"-"`
}

// GetID .
func (a *Action) GetID() string {
	return a.ID
}

// GetCommunicationID .
func (a *Action) GetCommunicationID() string {
	return a.CommunicationID
}

// GetFleetID .
func (a *Action) GetFleetID() string {
	return a.FleetID
}

// GetIsCompleted .
func (a *Action) GetIsCompleted() bool {
	return a.IsCompleted
}

// GetTrajectory .
func (a *Action) GetTrajectory() []action.TrajectoryPointComponent {
	var trajectoryPoints []action.TrajectoryPointComponent
	for _, tp := range a.TrajectoryPoints {
		trajectoryPoints = append(trajectoryPoints, tp)
	}
	return trajectoryPoints
}

// TrajectoryPoint .
type TrajectoryPoint struct {
	ActionID          string `gorm:"primaryKey"`
	PointNumber       int
	LatitudeDegree    float64
	LongitudeDegree   float64
	AltitudeM         float64
	RelativeAltitudeM float64
	SpeedMS           float64
	Armed             bool
	FlightMode        string
	OrientationX      float64
	OrientationY      float64
	OrientationZ      float64
	OrientationW      float64
}

func (tp *TrajectoryPoint) GetActionID() string {
	return tp.ActionID
}
func (tp *TrajectoryPoint) GetPointNumber() int {
	return tp.PointNumber
}
func (tp *TrajectoryPoint) GetLatitudeDegree() float64 {
	return tp.LatitudeDegree
}
func (tp *TrajectoryPoint) GetLongitudeDegree() float64 {
	return tp.LongitudeDegree
}
func (tp *TrajectoryPoint) GetAltitudeM() float64 {
	return tp.AltitudeM
}
func (tp *TrajectoryPoint) GetRelativeAltitudeM() float64 {
	return tp.RelativeAltitudeM
}
func (tp *TrajectoryPoint) GetSpeedMS() float64 {
	return tp.SpeedMS
}
func (tp *TrajectoryPoint) GetArmed() bool {
	return tp.Armed
}
func (tp *TrajectoryPoint) GetFlightMode() string {
	return tp.FlightMode
}
func (tp *TrajectoryPoint) GetOrientationX() float64 {
	return tp.OrientationX
}
func (tp *TrajectoryPoint) GetOrientationY() float64 {
	return tp.OrientationY
}
func (tp *TrajectoryPoint) GetOrientationZ() float64 {
	return tp.OrientationZ
}
func (tp *TrajectoryPoint) GetOrientationW() float64 {
	return tp.OrientationW
}
