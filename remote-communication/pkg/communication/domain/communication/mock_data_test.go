package communication

import (
	"errors"
	"time"
)

const DefaultID = ID("communication-id")
const DefaultCommandID = CommandID("command-id")
const DefaultMissionID = MissionID("mission-id")

var DefaultTime = time.Now()

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

// Communication用汎用ジェネレータモック
type generatorMock struct {
	Generator
	commandIDs     []CommandID
	commandIDIndex int
	times          []time.Time
	timeIndex      int
}

func (gen *generatorMock) NewCommandID() CommandID {
	commandID := gen.commandIDs[gen.commandIDIndex]
	gen.commandIDIndex++
	return commandID
}
func (gen *generatorMock) NewTime() time.Time {
	time := gen.times[gen.timeIndex]
	gen.timeIndex++
	return time
}

// Communication用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

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

func (m *communicationComponentMock) GetTelemetry() TelemetryComponent {
	return m.telemetry
}

func (m *communicationComponentMock) GetCommands() []CommandComponent {
	var commands []CommandComponent
	for _, cmd := range m.commands {
		commands = append(commands, cmd)
	}
	return commands
}

func (m *communicationComponentMock) GetUploadMissions() []UploadMissionComponent {
	var uploadMissions []UploadMissionComponent
	for _, um := range m.uploadMissions {
		uploadMissions = append(uploadMissions, um)
	}
	return uploadMissions
}

// Telemetry構成オブジェクトモック
type telemetryComponentMock struct {
	latitude         float64
	longitude        float64
	altitude         float64
	relativeAltitude float64
	speed            float64
	armed            bool
	flightMode       string
	x                float64
	y                float64
	z                float64
	w                float64
}

func (m *telemetryComponentMock) GetLatitude() float64 {
	return m.latitude
}

func (m *telemetryComponentMock) GetLongitude() float64 {
	return m.longitude
}

func (m *telemetryComponentMock) GetAltitude() float64 {
	return m.altitude
}

func (m *telemetryComponentMock) GetRelativeAltitude() float64 {
	return m.relativeAltitude
}

func (m *telemetryComponentMock) GetSpeed() float64 {
	return m.speed
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
