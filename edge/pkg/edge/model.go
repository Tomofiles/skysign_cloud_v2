package edge

// ConnectionState struct
type ConnectionState struct {
	VehicleID string
}

// Position struct
type Position struct {
	Latitude         float64
	Longitude        float64
	AbsoluteAltitude float64
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
	ID               string  `json:"id"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	AbsoluteAltitude float64 `json:"absoluteAltitude"`
	RelativeAltitude float64 `json:"relativeAltitude"`
	SpeedNorth       float64 `json:"speedN"`
	SpeedEast        float64 `json:"speedE"`
	SpeedDown        float64 `json:"speedD"`
	Armed            bool    `json:"armed"`
	FlightMode       string  `json:"flightMode"`
	OrientationX     float64 `json:"orientationX"`
	OrientationY     float64 `json:"orientationY"`
	OrientationZ     float64 `json:"orientationZ"`
	OrientationW     float64 `json:"orientationW"`
}
