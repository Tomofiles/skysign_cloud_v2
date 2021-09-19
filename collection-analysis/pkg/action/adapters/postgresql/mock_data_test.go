package postgresql

import (
	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/domain/action"
)

const DefaultActionID = action.ID("action-id")
const DefaultActionCommunicationID = action.CommunicationID("communication-id")
const DefaultActionFleetID = action.FleetID("fleet-id")

type actionComponentMock struct {
	ID               string
	CommunicationID  string
	FleetID          string
	IsCompleted      bool
	TrajectoryPoints []action.TrajectoryPointComponent
}

func (c *actionComponentMock) GetID() string {
	return c.ID
}
func (c *actionComponentMock) GetCommunicationID() string {
	return c.CommunicationID
}
func (c *actionComponentMock) GetFleetID() string {
	return c.FleetID
}
func (c *actionComponentMock) GetIsCompleted() bool {
	return c.IsCompleted
}
func (c *actionComponentMock) GetTrajectory() []action.TrajectoryPointComponent {
	return c.TrajectoryPoints
}

type trajectoryPointComponentMock struct {
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

func (c *trajectoryPointComponentMock) GetPointNumber() int {
	return c.PointNumber
}
func (c *trajectoryPointComponentMock) GetLatitudeDegree() float64 {
	return c.LatitudeDegree
}
func (c *trajectoryPointComponentMock) GetLongitudeDegree() float64 {
	return c.LongitudeDegree
}
func (c *trajectoryPointComponentMock) GetAltitudeM() float64 {
	return c.AltitudeM
}
func (c *trajectoryPointComponentMock) GetRelativeAltitudeM() float64 {
	return c.RelativeAltitudeM
}
func (c *trajectoryPointComponentMock) GetSpeedMS() float64 {
	return c.SpeedMS
}
func (c *trajectoryPointComponentMock) GetArmed() bool {
	return c.Armed
}
func (c *trajectoryPointComponentMock) GetFlightMode() string {
	return c.FlightMode
}
func (c *trajectoryPointComponentMock) GetOrientationX() float64 {
	return c.OrientationX
}
func (c *trajectoryPointComponentMock) GetOrientationY() float64 {
	return c.OrientationY
}
func (c *trajectoryPointComponentMock) GetOrientationZ() float64 {
	return c.OrientationZ
}
func (c *trajectoryPointComponentMock) GetOrientationW() float64 {
	return c.OrientationW
}
