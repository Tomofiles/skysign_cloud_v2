package communication

import (
	"errors"
	"time"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
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

// Communication用リポジトリモック
type repositoryMock struct {
	mock.Mock

	saveCommunications []*Communication
}

func (rm *repositoryMock) GetByID(tx txmanager.Tx, id ID) (*Communication, error) {
	ret := rm.Called(id)
	var v *Communication
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Communication)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMock) Save(tx txmanager.Tx, v *Communication) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveCommunications = append(rm.saveCommunications, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMock) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
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
