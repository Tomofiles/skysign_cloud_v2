package postgresql

import (
	"remote-communication/pkg/communication/domain/communication"
	"time"
)

// Communication .
type Communication struct {
	ID             string           `gorm:"primaryKey"`
	Telemetry      *Telemetry       `gorm:"-"`
	Commands       []*Command       `gorm:"-"`
	UploadMissions []*UploadMission `gorm:"-"`
}

// GetID .
func (c *Communication) GetID() string {
	return c.ID
}

// GetID .
func (c *Communication) GetTelemetry() communication.TelemetryComponent {
	return c.Telemetry
}

// GetCommands .
func (c *Communication) GetCommands() []communication.CommandComponent {
	var commands []communication.CommandComponent
	for _, command := range c.Commands {
		commands = append(commands, command)
	}
	return commands
}

// GetEvents .
func (c *Communication) GetUploadMissions() []communication.UploadMissionComponent {
	var uploadMissions []communication.UploadMissionComponent
	for _, um := range c.UploadMissions {
		uploadMissions = append(uploadMissions, um)
	}
	return uploadMissions
}

// Telemetry .
type Telemetry struct {
	CommunicationID   string `gorm:"primaryKey"`
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

// GetCommunicationID .
func (t *Telemetry) GetCommunicationID() string {
	return t.CommunicationID
}

// GetLatitude .
func (t *Telemetry) GetLatitude() float64 {
	return t.LatitudeDegree
}

// GetLongitude .
func (t *Telemetry) GetLongitude() float64 {
	return t.LongitudeDegree
}

// GetAltitude .
func (t *Telemetry) GetAltitude() float64 {
	return t.AltitudeM
}

// GetRelativeAltitude .
func (t *Telemetry) GetRelativeAltitude() float64 {
	return t.RelativeAltitudeM
}

// GetSpeed .
func (t *Telemetry) GetSpeed() float64 {
	return t.SpeedMS
}

// GetArmed .
func (t *Telemetry) GetArmed() bool {
	return t.Armed
}

// GetFlightMode .
func (t *Telemetry) GetFlightMode() string {
	return t.FlightMode
}

// GetX .
func (t *Telemetry) GetX() float64 {
	return t.OrientationX
}

// GetY .
func (t *Telemetry) GetY() float64 {
	return t.OrientationY
}

// GetZ .
func (t *Telemetry) GetZ() float64 {
	return t.OrientationZ
}

// GetW .
func (t *Telemetry) GetW() float64 {
	return t.OrientationW
}

// Command .
type Command struct {
	ID              string `gorm:"primaryKey"`
	CommunicationID string
	Type            string
	Time            time.Time
}

// GetID .
func (c *Command) GetID() string {
	return c.ID
}

// GetCommunicationID .
func (c *Command) GetCommunicationID() string {
	return c.CommunicationID
}

// GetType .
func (c *Command) GetType() string {
	return c.Type
}

// GetTime .
func (c *Command) GetTime() time.Time {
	return c.Time
}

// UploadMission .
type UploadMission struct {
	ID              string `gorm:"primaryKey"`
	CommunicationID string
	MissionID       string
}

// GetCommandID .
func (um *UploadMission) GetCommandID() string {
	return um.ID
}

// GetCommunicationID .
func (um *UploadMission) GetCommunicationID() string {
	return um.CommunicationID
}

// GetMissionID .
func (um *UploadMission) GetMissionID() string {
	return um.MissionID
}
