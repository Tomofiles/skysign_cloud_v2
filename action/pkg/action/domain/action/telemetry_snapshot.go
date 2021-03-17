package action

// TelemetrySnapshot .
type TelemetrySnapshot struct {
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

// IsArmed .
func (ts *TelemetrySnapshot) IsArmed() bool {
	return ts.Armed == true
}

// IsDisarmed .
func (ts *TelemetrySnapshot) IsDisarmed() bool {
	return ts.Armed == false
}
