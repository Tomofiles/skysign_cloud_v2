package action

import "errors"

const DefaultID = ID("action-id")
const DefaultCommunicationID = CommunicationID("communication-id")
const DefaultFleetID = FleetID("fleet-id")

var DefaultTelemetrySnapshot = TelemetrySnapshot{
	Latitude:         1.0,
	Longitude:        2.0,
	Altitude:         3.0,
	RelativeAltitude: 4.0,
	Speed:            5.0,
	Armed:            true,
	FlightMode:       "state",
	OrientationX:     6.0,
	OrientationY:     7.0,
	OrientationZ:     8.0,
	OrientationW:     9.0,
}
var DefaultTrajectoryPoint = trajectoryPoint{
	PointNumber:      1,
	Latitude:         1.0,
	Longitude:        2.0,
	Altitude:         3.0,
	RelativeAltitude: 4.0,
	Speed:            5.0,
	Armed:            true,
	FlightMode:       "state",
	OrientationX:     6.0,
	OrientationY:     7.0,
	OrientationZ:     8.0,
	OrientationW:     9.0,
}

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

type componentMock struct {
	ID               string
	CommunicationID  string
	FleetID          string
	IsCompleted      bool
	TrajectoryPoints []TrajectoryPointComponent
}

func (c *componentMock) GetID() string {
	return c.ID
}
func (c *componentMock) GetCommunicationID() string {
	return c.CommunicationID
}
func (c *componentMock) GetFleetID() string {
	return c.FleetID
}
func (c *componentMock) GetIsCompleted() bool {
	return c.IsCompleted
}
func (c *componentMock) GetTrajectory() []TrajectoryPointComponent {
	return c.TrajectoryPoints
}

type trajectoryPointComponentMock struct {
	trajectoryPoint trajectoryPoint
}

func (c *trajectoryPointComponentMock) GetPointNumber() int {
	return c.trajectoryPoint.PointNumber
}
func (c *trajectoryPointComponentMock) GetLatitude() float64 {
	return c.trajectoryPoint.Latitude
}
func (c *trajectoryPointComponentMock) GetLongitude() float64 {
	return c.trajectoryPoint.Longitude
}
func (c *trajectoryPointComponentMock) GetAltitude() float64 {
	return c.trajectoryPoint.Altitude
}
func (c *trajectoryPointComponentMock) GetRelativeAltitude() float64 {
	return c.trajectoryPoint.RelativeAltitude
}
func (c *trajectoryPointComponentMock) GetSpeed() float64 {
	return c.trajectoryPoint.Speed
}
func (c *trajectoryPointComponentMock) GetArmed() bool {
	return c.trajectoryPoint.Armed
}
func (c *trajectoryPointComponentMock) GetFlightMode() string {
	return c.trajectoryPoint.FlightMode
}
func (c *trajectoryPointComponentMock) GetOrientationX() float64 {
	return c.trajectoryPoint.OrientationX
}
func (c *trajectoryPointComponentMock) GetOrientationY() float64 {
	return c.trajectoryPoint.OrientationY
}
func (c *trajectoryPointComponentMock) GetOrientationZ() float64 {
	return c.trajectoryPoint.OrientationZ
}
func (c *trajectoryPointComponentMock) GetOrientationW() float64 {
	return c.trajectoryPoint.OrientationW
}
