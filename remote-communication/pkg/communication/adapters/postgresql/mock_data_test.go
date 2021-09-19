package postgresql

import (
	c "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"

	"time"
)

const DefaultCommunicationID = c.ID("communication-id")
const DefaultCommunicationCommandID = c.CommandID("command-id")
const DefaultCommunicationMissionID = c.MissionID("mission-id")

var DefaultCommunicationTime = time.Now()

// Communication構成オブジェクトモック
type communicationComponentMock struct {
	id             string
	telemetry      *telemetryComponentMock
	commands       []*commandComponentMock
	uploadMissions []*uploadMissionComponentMock
}

func (m *communicationComponentMock) GetID() string {
	return m.id
}

func (m *communicationComponentMock) GetTelemetry() c.TelemetryComponent {
	return m.telemetry
}

func (m *communicationComponentMock) GetCommands() []c.CommandComponent {
	var commands []c.CommandComponent
	for _, cmd := range m.commands {
		commands = append(commands, cmd)
	}
	return commands
}

func (m *communicationComponentMock) GetUploadMissions() []c.UploadMissionComponent {
	var uploadMissions []c.UploadMissionComponent
	for _, um := range m.uploadMissions {
		uploadMissions = append(uploadMissions, um)
	}
	return uploadMissions
}

// Telemetry構成オブジェクトモック
type telemetryComponentMock struct {
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

func (m *telemetryComponentMock) GetLatitudeDegree() float64 {
	return m.latitudeDegree
}

func (m *telemetryComponentMock) GetLongitudeDegree() float64 {
	return m.longitudeDegree
}

func (m *telemetryComponentMock) GetAltitudeM() float64 {
	return m.altitudeM
}

func (m *telemetryComponentMock) GetRelativeAltitudeM() float64 {
	return m.relativeAltitudeM
}

func (m *telemetryComponentMock) GetSpeedMS() float64 {
	return m.speedMS
}

func (m *telemetryComponentMock) GetArmed() bool {
	return m.armed
}

func (m *telemetryComponentMock) GetFlightMode() string {
	return m.flightMode
}

func (m *telemetryComponentMock) GetX() float64 {
	return m.x
}

func (m *telemetryComponentMock) GetY() float64 {
	return m.y
}

func (m *telemetryComponentMock) GetZ() float64 {
	return m.z
}

func (m *telemetryComponentMock) GetW() float64 {
	return m.w
}

// Command構成オブジェクトモック
type commandComponentMock struct {
	id    string
	cType string
	time  time.Time
}

func (m *commandComponentMock) GetID() string {
	return m.id
}

func (m *commandComponentMock) GetType() string {
	return m.cType
}

func (m *commandComponentMock) GetTime() time.Time {
	return m.time
}

// UploadMission構成オブジェクトモック
type uploadMissionComponentMock struct {
	commandID string
	missionID string
}

func (m *uploadMissionComponentMock) GetCommandID() string {
	return m.commandID
}

func (m *uploadMissionComponentMock) GetMissionID() string {
	return m.missionID
}
