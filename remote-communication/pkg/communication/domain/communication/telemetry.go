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
	latitude         float64
	longitude        float64
	altitude         float64
	relativeAltitude float64
	speed            float64
	armed            bool
	flightMode       string
	x                float64
	y                float64
	z                float64
	w                float64
}

// NewTelemetry .
func NewTelemetry() *Telemetry {
	return &Telemetry{
		armed:      Disarmed,
		flightMode: "NONE",
	}
}

// NewTelemetryBySnapshot .
func NewTelemetryBySnapshot(snapshot TelemetrySnapshot) *Telemetry {
	return &Telemetry{
		latitude:         snapshot.Latitude,
		longitude:        snapshot.Longitude,
		altitude:         snapshot.Altitude,
		relativeAltitude: snapshot.RelativeAltitude,
		speed:            snapshot.Speed,
		armed:            snapshot.Armed,
		flightMode:       snapshot.FlightMode,
		x:                snapshot.X,
		y:                snapshot.Y,
		z:                snapshot.Z,
		w:                snapshot.W,
	}
}

// GetSnapshot .
func (t *Telemetry) GetSnapshot() TelemetrySnapshot {
	return TelemetrySnapshot{
		Latitude:         t.latitude,
		Longitude:        t.longitude,
		Altitude:         t.altitude,
		RelativeAltitude: t.relativeAltitude,
		Speed:            t.speed,
		Armed:            t.armed,
		FlightMode:       t.flightMode,
		X:                t.x,
		Y:                t.y,
		Z:                t.z,
		W:                t.w,
	}
}

// IsDisarmed .
func (t *Telemetry) IsDisarmed() bool {
	return t.armed == Disarmed
}
