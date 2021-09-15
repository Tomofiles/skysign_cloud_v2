package action

// TelemetrySnapshot .
type TelemetrySnapshot struct {
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

// IsArmed .
func (ts *TelemetrySnapshot) IsArmed() bool {
	return ts.Armed == true
}

// IsDisarmed .
func (ts *TelemetrySnapshot) IsDisarmed() bool {
	return ts.Armed == false
}
