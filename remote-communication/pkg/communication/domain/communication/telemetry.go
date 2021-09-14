package communication

const (
	// Armed .
	Armed = true
	// Disarmed .
	Disarmed = false
)

// TelemetrySnapshot .
type TelemetrySnapshot struct {
	LatitudeDegree    float64
	LongitudeDegree   float64
	AltitudeM         float64
	RelativeAltitudeM float64
	SpeedMS           float64
	Armed             bool
	FlightMode        string
	X                 float64
	Y                 float64
	Z                 float64
	W                 float64
}

// Telemetry .
type Telemetry struct {
	latitudeDegree    float64
	longitudeDegree   float64
	altitudeM         float64
	relativeAltitudeM float64
	speedMS           float64
	armed             bool
	flightMode        string
	x                 float64
	y                 float64
	z                 float64
	w                 float64
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
		latitudeDegree:    snapshot.LatitudeDegree,
		longitudeDegree:   snapshot.LongitudeDegree,
		altitudeM:         snapshot.AltitudeM,
		relativeAltitudeM: snapshot.RelativeAltitudeM,
		speedMS:           snapshot.SpeedMS,
		armed:             snapshot.Armed,
		flightMode:        snapshot.FlightMode,
		x:                 snapshot.X,
		y:                 snapshot.Y,
		z:                 snapshot.Z,
		w:                 snapshot.W,
	}
}

// GetSnapshot .
func (t *Telemetry) GetSnapshot() TelemetrySnapshot {
	return TelemetrySnapshot{
		LatitudeDegree:    t.latitudeDegree,
		LongitudeDegree:   t.longitudeDegree,
		AltitudeM:         t.altitudeM,
		RelativeAltitudeM: t.relativeAltitudeM,
		SpeedMS:           t.speedMS,
		Armed:             t.armed,
		FlightMode:        t.flightMode,
		X:                 t.x,
		Y:                 t.y,
		Z:                 t.z,
		W:                 t.w,
	}
}

// IsDisarmed .
func (t *Telemetry) IsDisarmed() bool {
	return t.armed == Disarmed
}
