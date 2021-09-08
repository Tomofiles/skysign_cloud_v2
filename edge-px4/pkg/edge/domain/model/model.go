package model

// ConnectionState .
type ConnectionState struct {
	VehicleID string
}

// Position .
type Position struct {
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
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
	North float64
	East  float64
	Down  float64
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
	Latitude       float64
	Longitude      float64
	RelativeHeight float64
	Speed          float64
}
