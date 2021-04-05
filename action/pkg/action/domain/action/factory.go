package action

// NewInstance .
func NewInstance(
	id ID,
	communicationID CommunicationID,
	flightplanID FlightplanID,
) *Action {
	return &Action{
		id:              id,
		communicationID: communicationID,
		flightplanID:    flightplanID,
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
				PointNumber:      tp.GetPointNumber(),
				Latitude:         tp.GetLatitude(),
				Longitude:        tp.GetLongitude(),
				Altitude:         tp.GetAltitude(),
				RelativeAltitude: tp.GetRelativeAltitude(),
				Speed:            tp.GetSpeed(),
				Armed:            tp.GetArmed(),
				FlightMode:       tp.GetFlightMode(),
				OrientationX:     tp.GetOrientationX(),
				OrientationY:     tp.GetOrientationY(),
				OrientationZ:     tp.GetOrientationZ(),
				OrientationW:     tp.GetOrientationW(),
			},
		)
	}
	return &Action{
		id:              ID(comp.GetID()),
		communicationID: CommunicationID(comp.GetCommunicationID()),
		flightplanID:    FlightplanID(comp.GetFlightplanID()),
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
	GetFlightplanID() string
	GetIsCompleted() bool
}

// TrajectoryPointComponent .
type TrajectoryPointComponent interface {
	GetPointNumber() int
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
	GetRelativeAltitude() float64
	GetSpeed() float64
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
func (c *actionComponent) GetFlightplanID() string {
	return string(c.action.flightplanID)
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
func (c *trajectoryPointComponent) GetLatitude() float64 {
	return c.trajectoryPoint.Latitude
}
func (c *trajectoryPointComponent) GetLongitude() float64 {
	return c.trajectoryPoint.Longitude
}
func (c *trajectoryPointComponent) GetAltitude() float64 {
	return c.trajectoryPoint.Altitude
}
func (c *trajectoryPointComponent) GetRelativeAltitude() float64 {
	return c.trajectoryPoint.RelativeAltitude
}
func (c *trajectoryPointComponent) GetSpeed() float64 {
	return c.trajectoryPoint.Speed
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
