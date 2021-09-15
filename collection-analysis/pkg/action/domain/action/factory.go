package action

// NewInstance .
func NewInstance(
	id ID,
	communicationID CommunicationID,
	fleetID FleetID,
) *Action {
	return &Action{
		id:              id,
		communicationID: communicationID,
		fleetID:         fleetID,
		isCompleted:     Active,
		trajectory:      newTrajectory(),
	}
}

// AssembleFrom .
func AssembleFrom(comp Component) *Action {
	var trajectoryPoints []trajectoryPoint
	for _, tp := range comp.GetTrajectory() {
		trajectoryPoints = append(
			trajectoryPoints,
			trajectoryPoint{
				PointNumber:       tp.GetPointNumber(),
				LatitudeDegree:    tp.GetLatitudeDegree(),
				LongitudeDegree:   tp.GetLongitudeDegree(),
				AltitudeM:         tp.GetAltitudeM(),
				RelativeAltitudeM: tp.GetRelativeAltitudeM(),
				SpeedMS:           tp.GetSpeedMS(),
				Armed:             tp.GetArmed(),
				FlightMode:        tp.GetFlightMode(),
				OrientationX:      tp.GetOrientationX(),
				OrientationY:      tp.GetOrientationY(),
				OrientationZ:      tp.GetOrientationZ(),
				OrientationW:      tp.GetOrientationW(),
			},
		)
	}
	return &Action{
		id:              ID(comp.GetID()),
		communicationID: CommunicationID(comp.GetCommunicationID()),
		fleetID:         FleetID(comp.GetFleetID()),
		isCompleted:     comp.GetIsCompleted(),
		trajectory: Trajectory{
			numberOfPoints:   len(trajectoryPoints),
			trajectoryPoints: trajectoryPoints,
		},
	}
}

// TakeApart .
func TakeApart(
	action *Action,
	actionComp func(comp ActionComponent),
	trajectoryPointComp func(comp TrajectoryPointComponent),
) {
	actionComp(
		&actionComponent{
			action: action,
		},
	)
	for _, tp := range action.trajectory.trajectoryPoints {
		trajectoryPointComp(
			&trajectoryPointComponent{
				trajectoryPoint: tp,
			},
		)
	}
}

// Component .
type Component interface {
	ActionComponent
	GetTrajectory() []TrajectoryPointComponent
}

// ActionComponent .
type ActionComponent interface {
	GetID() string
	GetCommunicationID() string
	GetFleetID() string
	GetIsCompleted() bool
}

// TrajectoryPointComponent .
type TrajectoryPointComponent interface {
	GetPointNumber() int
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetAltitudeM() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
	GetArmed() bool
	GetFlightMode() string
	GetOrientationX() float64
	GetOrientationY() float64
	GetOrientationZ() float64
	GetOrientationW() float64
}

func newTrajectory() Trajectory {
	var trajectoryPoints []trajectoryPoint
	return Trajectory{
		numberOfPoints:   0,
		trajectoryPoints: trajectoryPoints,
	}
}

type actionComponent struct {
	action *Action
}

func (c *actionComponent) GetID() string {
	return string(c.action.id)
}
func (c *actionComponent) GetCommunicationID() string {
	return string(c.action.communicationID)
}
func (c *actionComponent) GetFleetID() string {
	return string(c.action.fleetID)
}
func (c *actionComponent) GetIsCompleted() bool {
	return c.action.isCompleted
}

type trajectoryPointComponent struct {
	trajectoryPoint trajectoryPoint
}

func (c *trajectoryPointComponent) GetPointNumber() int {
	return c.trajectoryPoint.PointNumber
}
func (c *trajectoryPointComponent) GetLatitudeDegree() float64 {
	return c.trajectoryPoint.LatitudeDegree
}
func (c *trajectoryPointComponent) GetLongitudeDegree() float64 {
	return c.trajectoryPoint.LongitudeDegree
}
func (c *trajectoryPointComponent) GetAltitudeM() float64 {
	return c.trajectoryPoint.AltitudeM
}
func (c *trajectoryPointComponent) GetRelativeAltitudeM() float64 {
	return c.trajectoryPoint.RelativeAltitudeM
}
func (c *trajectoryPointComponent) GetSpeedMS() float64 {
	return c.trajectoryPoint.SpeedMS
}
func (c *trajectoryPointComponent) GetArmed() bool {
	return c.trajectoryPoint.Armed
}
func (c *trajectoryPointComponent) GetFlightMode() string {
	return c.trajectoryPoint.FlightMode
}
func (c *trajectoryPointComponent) GetOrientationX() float64 {
	return c.trajectoryPoint.OrientationX
}
func (c *trajectoryPointComponent) GetOrientationY() float64 {
	return c.trajectoryPoint.OrientationY
}
func (c *trajectoryPointComponent) GetOrientationZ() float64 {
	return c.trajectoryPoint.OrientationZ
}
func (c *trajectoryPointComponent) GetOrientationW() float64 {
	return c.trajectoryPoint.OrientationW
}
