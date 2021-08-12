package communication

const (
	// Armed .
	Armed = true
	// Disarmed .
	Disarmed = false
)

// TelemetrySnapshot .
type TelemetrySnapshot struct {
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
	Speed            float64
	Armed            bool
	FlightMode       string
	X                float64
	Y                float64
	Z                float64
	W                float64
}

// Telemetry .
type Telemetry struct {
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
	Speed            float64
	Armed            bool
	FlightMode       string
	X                float64
	Y                float64
	Z                float64
	W                float64
}

// NewTelemetry .
func NewTelemetry() *Telemetry {
	return &Telemetry{
		Armed:      Disarmed,
		FlightMode: "NONE",
	}
}

// NewTelemetryBySnapshot .
func NewTelemetryBySnapshot(snapshot TelemetrySnapshot) *Telemetry {
	return &Telemetry{
		Latitude:         snapshot.Latitude,
		Longitude:        snapshot.Longitude,
		Altitude:         snapshot.Altitude,
		RelativeAltitude: snapshot.RelativeAltitude,
		Speed:            snapshot.Speed,
		Armed:            snapshot.Armed,
		FlightMode:       snapshot.FlightMode,
		X:                snapshot.X,
		Y:                snapshot.Y,
		Z:                snapshot.Z,
		W:                snapshot.W,
	}
}

// GetSnapshot .
func (t *Telemetry) GetSnapshot() TelemetrySnapshot {
	return TelemetrySnapshot{
		Latitude:         t.Latitude,
		Longitude:        t.Longitude,
		Altitude:         t.Altitude,
		RelativeAltitude: t.RelativeAltitude,
		Speed:            t.Speed,
		Armed:            t.Armed,
		FlightMode:       t.FlightMode,
		X:                t.X,
		Y:                t.Y,
		Z:                t.Z,
		W:                t.W,
	}
}

// IsDisarmed .
func (t *Telemetry) IsDisarmed() bool {
	return t.Armed == Disarmed
}
