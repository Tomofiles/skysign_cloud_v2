package action

// Trajectory .
type Trajectory struct {
	numberOfPoints   int
	trajectoryPoints []trajectoryPoint
}

// Extenstion .
func (t *Trajectory) Extension(snapshot TelemetrySnapshot) Trajectory {
	newNumberOfPoints := t.numberOfPoints + 1
	newTrajectoryPoints := append(
		t.trajectoryPoints,
		trajectoryPoint{
			PointNumber:      newNumberOfPoints,
			Latitude:         snapshot.Latitude,
			Longitude:        snapshot.Longitude,
			Altitude:         snapshot.Altitude,
			RelativeAltitude: snapshot.RelativeAltitude,
			Speed:            snapshot.Speed,
			Armed:            snapshot.Armed,
			FlightMode:       snapshot.FlightMode,
			OrientationX:     snapshot.OrientationX,
			OrientationY:     snapshot.OrientationY,
			OrientationZ:     snapshot.OrientationZ,
			OrientationW:     snapshot.OrientationW,
		},
	)
	return Trajectory{
		numberOfPoints:   newNumberOfPoints,
		trajectoryPoints: newTrajectoryPoints,
	}
}

// Extenstion .
func (t *Trajectory) TakeASnapshots() []TelemetrySnapshot {
	var telemetrySnapshots []TelemetrySnapshot
	for _, tp := range t.trajectoryPoints {
		telemetrySnapshots = append(
			telemetrySnapshots,
			TelemetrySnapshot{
				Latitude:         tp.Latitude,
				Longitude:        tp.Longitude,
				Altitude:         tp.Altitude,
				RelativeAltitude: tp.RelativeAltitude,
				Speed:            tp.Speed,
				Armed:            tp.Armed,
				FlightMode:       tp.FlightMode,
				OrientationX:     tp.OrientationX,
				OrientationY:     tp.OrientationY,
				OrientationZ:     tp.OrientationZ,
				OrientationW:     tp.OrientationW,
			},
		)
	}
	return telemetrySnapshots
}

type trajectoryPoint struct {
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
