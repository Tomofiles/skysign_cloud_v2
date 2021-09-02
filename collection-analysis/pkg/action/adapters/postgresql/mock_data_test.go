package postgresql

import (
	"collection-analysis/pkg/action/domain/action"
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
	PointNumber      int
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
	Speed            float64
	Armed            bool
	FlightMode       string
	OrientationX     float64
	OrientationY     float64
	OrientationZ     float64
	OrientationW     float64
}

func (c *trajectoryPointComponentMock) GetPointNumber() int {
	return c.PointNumber
}
func (c *trajectoryPointComponentMock) GetLatitude() float64 {
	return c.Latitude
}
func (c *trajectoryPointComponentMock) GetLongitude() float64 {
	return c.Longitude
}
func (c *trajectoryPointComponentMock) GetAltitude() float64 {
	return c.Altitude
}
func (c *trajectoryPointComponentMock) GetRelativeAltitude() float64 {
	return c.RelativeAltitude
}
func (c *trajectoryPointComponentMock) GetSpeed() float64 {
	return c.Speed
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
