package postgresql

import (
	c "remote-communication/pkg/communication/domain/communication"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultCommunicationID = c.ID("communication-id")
const DefaultCommunicationCommandID = c.CommandID("command-id")
const DefaultCommunicationMissionID = c.MissionID("mission-id")

var DefaultCommunicationTime = time.Now()

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
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
