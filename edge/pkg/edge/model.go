package edge

// ConnectionState struct
type ConnectionState struct {
	VehicleID string
}

// Position struct
type Position struct {
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
}

// Quaternion struct
type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Velocity struct
type Velocity struct {
	North float64
	East  float64
	Down  float64
}

// Armed struct
type Armed struct {
	Armed bool
}

// FlightMode struct
type FlightMode struct {
	FlightMode string
}

// Telemetry struct
type Telemetry struct {
	ID    string `json:"id"`
	State *State `json:"telemetry"`
}

// State struct
type State struct {
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Altitude         float64 `json:"altitude"`
	RelativeAltitude float64 `json:"relative_altitude"`
	Speed            float64 `json:"speed"`
	Armed            bool    `json:"armed"`
	FlightMode       string  `json:"flight_mode"`
	OrientationX     float64 `json:"orientation_x"`
	OrientationY     float64 `json:"orientation_y"`
	OrientationZ     float64 `json:"orientation_z"`
	OrientationW     float64 `json:"orientation_w"`
}

// CommandIDs .
type CommandIDs struct {
	CommandIds []string `json:"command_ids"`
}

// Command .
type Command struct {
	Type string `json:"type"`
}

// UploadMission .
type UploadMission struct {
	ID        string `json:"id"`
	MissionID string `json:"mission_id"`
}

// Mission .
type Mission struct {
	ID        string       `json:"id"`
	Waypoints []*Waypoints `json:"waypoints"`
}

// Waypoints .
type Waypoints struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	RelativeHeight float64 `json:"relative_height"`
	Speed          float64 `json:"speed"`
}
