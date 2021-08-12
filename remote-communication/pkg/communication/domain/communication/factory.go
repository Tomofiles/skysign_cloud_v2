package communication

import "time"

// NewInstance .
func NewInstance(gen Generator, id ID) *Communication {
	telemetry := NewTelemetry()
	var commands []*Command
	var uploadMissions []*UploadMission
	return &Communication{
		id:             id,
		telemetry:      telemetry,
		commands:       commands,
		uploadMissions: uploadMissions,
		gen:            gen,
	}
}

// AssembleFrom .
func AssembleFrom(gen Generator, comp Component) *Communication {
	telemetry := &Telemetry{
		Latitude:         comp.GetTelemetry().GetLatitude(),
		Longitude:        comp.GetTelemetry().GetLongitude(),
		Altitude:         comp.GetTelemetry().GetAltitude(),
		RelativeAltitude: comp.GetTelemetry().GetRelativeAltitude(),
		Speed:            comp.GetTelemetry().GetSpeed(),
		Armed:            comp.GetTelemetry().GetArmed(),
		FlightMode:       comp.GetTelemetry().GetFlightMode(),
		X:                comp.GetTelemetry().GetX(),
		Y:                comp.GetTelemetry().GetY(),
		Z:                comp.GetTelemetry().GetZ(),
		W:                comp.GetTelemetry().GetW(),
	}
	var commands []*Command
	for _, c := range comp.GetCommands() {
		commands = append(
			commands,
			&Command{
				id:    CommandID(c.GetID()),
				cType: CommandType(c.GetType()),
				time:  c.GetTime(),
			},
		)
	}
	var uploadMissions []*UploadMission
	for _, um := range comp.GetUploadMissions() {
		uploadMissions = append(
			uploadMissions,
			&UploadMission{
				commandID: CommandID(um.GetCommandID()),
				missionID: MissionID(um.GetMissionID()),
			},
		)
	}
	return &Communication{
		id:             ID(comp.GetID()),
		telemetry:      telemetry,
		commands:       commands,
		uploadMissions: uploadMissions,
		gen:            gen,
	}
}

// TakeApart .
func TakeApart(
	communication *Communication,
	communicationComp func(id string),
	telemetryComp func(latitude, longitude, altitude, relativeAltitude, speed, x, y, z, w float64, armed bool, flightMode string),
	commandComp func(id, cType string, time time.Time),
	uploadMissionComp func(commandID, missionID string),
) {
	communicationComp(
		string(communication.id),
	)
	telemetryComp(
		communication.telemetry.Latitude,
		communication.telemetry.Longitude,
		communication.telemetry.Altitude,
		communication.telemetry.RelativeAltitude,
		communication.telemetry.Speed,
		communication.telemetry.X,
		communication.telemetry.Y,
		communication.telemetry.Z,
		communication.telemetry.W,
		communication.telemetry.Armed,
		communication.telemetry.FlightMode,
	)
	for _, c := range communication.commands {
		commandComp(
			string(c.id),
			string(c.cType),
			c.time,
		)
	}
	for _, um := range communication.uploadMissions {
		uploadMissionComp(
			string(um.commandID),
			string(um.missionID),
		)
	}
}

// Component .
type Component interface {
	GetID() string
	GetTelemetry() TelemetryComponent
	GetCommands() []CommandComponent
	GetUploadMissions() []UploadMissionComponent
}

// TelemetryComponent .
type TelemetryComponent interface {
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
	GetRelativeAltitude() float64
	GetSpeed() float64
	GetArmed() bool
	GetFlightMode() string
	GetX() float64
	GetY() float64
	GetZ() float64
	GetW() float64
}

// CommandComponent .
type CommandComponent interface {
	GetID() string
	GetType() string
	GetTime() time.Time
}

// UploadMissionComponent .
type UploadMissionComponent interface {
	GetCommandID() string
	GetMissionID() string
}
