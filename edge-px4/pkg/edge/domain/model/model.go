package model

// ConnectionState .
type ConnectionState struct {
	VehicleID string
}

// Position .
type Position struct {
	LatitudeDegree    float64
	LongitudeDegree   float64
	AltitudeM         float64
	RelativeAltitudeM float64
}

// Quaternion .
type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Velocity .
type Velocity struct {
	NorthMS float64
	EastMS  float64
	DownMS  float64
}

// Armed .
type Armed struct {
	Armed bool
}

// FlightMode .
type FlightMode struct {
	FlightMode string
}

// PushTelemetry .
type PushTelemetry struct {
	ID    string
	State *State
}

// State .
type State struct {
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

// CommandIDs .
type CommandIDs struct {
	CommandIds []string
}

// Command .
type Command struct {
	Type string
}

// UploadMission .
type UploadMission struct {
	ID        string
	MissionID string
}

// Mission .
type Mission struct {
	ID        string
	Waypoints []*Waypoints
}

// Waypoints .
type Waypoints struct {
	LatitudeDegree    float64
	LongitudeDegree   float64
	RelativeAltitudeM float64
	SpeedMS           float64
}
