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
		latitudeDegree:    comp.GetTelemetry().GetLatitudeDegree(),
		longitudeDegree:   comp.GetTelemetry().GetLongitudeDegree(),
		altitudeM:         comp.GetTelemetry().GetAltitudeM(),
		relativeAltitudeM: comp.GetTelemetry().GetRelativeAltitudeM(),
		speedMS:           comp.GetTelemetry().GetSpeedMS(),
		armed:             comp.GetTelemetry().GetArmed(),
		flightMode:        comp.GetTelemetry().GetFlightMode(),
		x:                 comp.GetTelemetry().GetX(),
		y:                 comp.GetTelemetry().GetY(),
		z:                 comp.GetTelemetry().GetZ(),
		w:                 comp.GetTelemetry().GetW(),
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
	telemetryComp func(latitudeDegree, longitudeDegree, altitudeM, relativeAltitudeM, speedMS, x, y, z, w float64, armed bool, flightMode string),
	commandComp func(id, cType string, time time.Time),
	uploadMissionComp func(commandID, missionID string),
) {
	communicationComp(
		string(communication.id),
	)
	telemetryComp(
		communication.telemetry.latitudeDegree,
		communication.telemetry.longitudeDegree,
		communication.telemetry.altitudeM,
		communication.telemetry.relativeAltitudeM,
		communication.telemetry.speedMS,
		communication.telemetry.x,
		communication.telemetry.y,
		communication.telemetry.z,
		communication.telemetry.w,
		communication.telemetry.armed,
		communication.telemetry.flightMode,
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
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetAltitudeM() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
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
