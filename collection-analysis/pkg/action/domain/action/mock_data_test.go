package action

import "errors"

const DefaultID = ID("action-id")
const DefaultCommunicationID = CommunicationID("communication-id")
const DefaultFleetID = FleetID("fleet-id")

var DefaultTelemetrySnapshot = TelemetrySnapshot{
	LatitudeDegree:    1.0,
	LongitudeDegree:   2.0,
	AltitudeM:         3.0,
	RelativeAltitudeM: 4.0,
	SpeedMS:           5.0,
	Armed:             true,
	FlightMode:        "state",
	OrientationX:      6.0,
	OrientationY:      7.0,
	OrientationZ:      8.0,
	OrientationW:      9.0,
}
var DefaultTrajectoryPoint = trajectoryPoint{
	PointNumber:       1,
	LatitudeDegree:    1.0,
	LongitudeDegree:   2.0,
	AltitudeM:         3.0,
	RelativeAltitudeM: 4.0,
	SpeedMS:           5.0,
	Armed:             true,
	FlightMode:        "state",
	OrientationX:      6.0,
	OrientationY:      7.0,
	OrientationZ:      8.0,
	OrientationW:      9.0,
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
func (c *trajectoryPointComponentMock) GetLatitudeDegree() float64 {
	return c.trajectoryPoint.LatitudeDegree
}
func (c *trajectoryPointComponentMock) GetLongitudeDegree() float64 {
	return c.trajectoryPoint.LongitudeDegree
}
func (c *trajectoryPointComponentMock) GetAltitudeM() float64 {
	return c.trajectoryPoint.AltitudeM
}
func (c *trajectoryPointComponentMock) GetRelativeAltitudeM() float64 {
	return c.trajectoryPoint.RelativeAltitudeM
}
func (c *trajectoryPointComponentMock) GetSpeedMS() float64 {
	return c.trajectoryPoint.SpeedMS
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
