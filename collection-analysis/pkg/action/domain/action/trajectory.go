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
			PointNumber:       newNumberOfPoints,
			LatitudeDegree:    snapshot.LatitudeDegree,
			LongitudeDegree:   snapshot.LongitudeDegree,
			AltitudeM:         snapshot.AltitudeM,
			RelativeAltitudeM: snapshot.RelativeAltitudeM,
			SpeedMS:           snapshot.SpeedMS,
			Armed:             snapshot.Armed,
			FlightMode:        snapshot.FlightMode,
			OrientationX:      snapshot.OrientationX,
			OrientationY:      snapshot.OrientationY,
			OrientationZ:      snapshot.OrientationZ,
			OrientationW:      snapshot.OrientationW,
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
				LatitudeDegree:    tp.LatitudeDegree,
				LongitudeDegree:   tp.LongitudeDegree,
				AltitudeM:         tp.AltitudeM,
				RelativeAltitudeM: tp.RelativeAltitudeM,
				SpeedMS:           tp.SpeedMS,
				Armed:             tp.Armed,
				FlightMode:        tp.FlightMode,
				OrientationX:      tp.OrientationX,
				OrientationY:      tp.OrientationY,
				OrientationZ:      tp.OrientationZ,
				OrientationW:      tp.OrientationW,
			},
		)
	}
	return telemetrySnapshots
}

type trajectoryPoint struct {
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
