package edge

import "encoding/json"

// // Telemetry struct
// type Telemetry struct {
// 	VehicleID      string       `json:"vehicleID"`
// 	Position       *Position    `json:"position,omitempty"`
// 	Orientation    *Orientation `json:"orientation,omitempty"`
// 	Armed          bool         `json:"armed"`
// 	FlightMode     string       `json:"flightMode,omitempty"`
// 	VideoStreaming bool         `json:"videoStreaming"`
// }

// Position struct
type Position struct {
	CartographicDegrees []float64 `json:"cartographicDegrees"`
}

// Orientation struct
type Orientation struct {
	UnitQuaternion []float64 `json:"unitQuaternion"`
}

// Command struct
type Command struct {
	VehicleID string          `json:"vehicleID"`
	MessageID string          `json:"messageID"`
	Payload   json.RawMessage `json:"payload"`
}

// Mission struct
type Mission struct {
	MissionItems []*MissionItem `json:"missionItems"`
}

// MissionItem struct
type MissionItem struct {
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
	Alt   float32 `json:"alt"`
	Speed float32 `json:"speed"`
}

// SignalingMessage .
type SignalingMessage struct {
	Type      string `json:"type"`
	VehicleID string `json:"vehicleID"`
}

// Telemetry .
type Telemetry struct {
	ID           string  `json:"id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Altitude     float64 `json:"altitude"`
	Speed        float64 `json:"speed"`
	Armed        bool    `json:"armed"`
	FlightMode   string  `json:"flightMode"`
	OrientationX float64 `json:"orientationX"`
	OrientationY float64 `json:"orientationY"`
	OrientationZ float64 `json:"orientationZ"`
	OrientationW float64 `json:"orientationW"`
}
