package communication

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// CommunicationのTelemetryを更新する。
// すべてのTelemetryのフィールドが更新されることを検証する。
// イベントパブリッシャを設定していないため、イベント発行がスキップされること。
func TestNoEventsWhenPushTelemetryToCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	snapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	communication.PushTelemetry(snapshot)

	expectTelemetry := &Telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	a.Equal(communication.telemetry, expectTelemetry)
}

// CommunicationのTelemetryを更新する。
// すべてのTelemetryのフィールドが更新されることを検証する。
// Telemetryが更新されたことを表すイベントが発行されることを検証する。
func TestPushTelemetryToCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	pub := &publisherMock{}
	communication := NewInstance(gen, DefaultID)
	communication.SetPublisher(pub)

	snapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	communication.PushTelemetry(snapshot)

	expectTelemetry := &Telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}
	expectEvent := TelemetryUpdatedEvent{
		CommunicationID: DefaultID,
		Telemetry:       snapshot,
	}

	a.Equal(communication.telemetry, expectTelemetry)
	a.Equal(pub.events, []interface{}{expectEvent})
}

// CommunicationからTelemetryのSnapshotを取得する。
func TestPullTelemetryFromCommunication(t *testing.T) {
	a := assert.New(t)

	telemetry := &Telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)
	communication.telemetry = telemetry

	snapshot := communication.PullTelemetry()

	expectSnapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}

	a.Equal(snapshot, expectSnapshot)
}

// CommunicationにCommandを追加する。
// Commandが追加され、IDとTimeが付与されていることを検証する。
// Armコマンド追加条件に合致し、Commandが2件であることを検証する。
func TestArmCommandPushWhenPushCommandToCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1 = DefaultCommandID + "-1"
		DefaultCommandID2 = DefaultCommandID + "-2"
		DefaultTime1      = DefaultTime.Add(1 * time.Minute)
		DefaultTime2      = DefaultTime.Add(2 * time.Minute)
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2},
		times:      []time.Time{DefaultTime1, DefaultTime2},
	}
	communication := NewInstance(gen, DefaultID)

	id := communication.PushCommand(CommandTypeTAKEOFF)

	expectCommand1 := &Command{
		id:    DefaultCommandID1,
		cType: CommandTypeARM,
		time:  DefaultTime1,
	}
	expectCommand2 := &Command{
		id:    DefaultCommandID2,
		cType: CommandTypeTAKEOFF,
		time:  DefaultTime2,
	}

	a.Equal(id, DefaultCommandID2)
	a.Equal(communication.commands, []*Command{expectCommand1, expectCommand2})
}

// CommunicationにCommandを追加する。
// Commandが追加され、IDとTimeが付与されていることを検証する。
func TestPushCommandToCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	communication := NewInstance(gen, DefaultID)

	id := communication.PushCommand(CommandTypeARM)

	expectCommand := &Command{
		id:    DefaultCommandID,
		cType: CommandTypeARM,
		time:  DefaultTime,
	}

	a.Equal(id, DefaultCommandID)
	a.Equal(communication.commands, []*Command{expectCommand})
}

// CommunicationにUploadMissionを追加する。
// Commandが追加され、IDとTimeが付与されていることを検証する。
// また、UploadMissionが追加され、CommandIDとMissionIDが付与されていることを検証する。
func TestPushUploadMissionToCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	communication := NewInstance(gen, DefaultID)

	id := communication.PushUploadMission(DefaultMissionID)

	expectCommand := &Command{
		id:    DefaultCommandID,
		cType: CommandTypeUPLOAD,
		time:  DefaultTime,
	}
	expectUploadMission := &UploadMission{
		commandID: DefaultCommandID,
		missionID: DefaultMissionID,
	}

	a.Equal(id, DefaultCommandID)
	a.Equal(communication.commands, []*Command{expectCommand})
	a.Equal(communication.uploadMissions, []*UploadMission{expectUploadMission})
}

// CommunicationからCommandを取得する。
// CommandIDに合致するCommandが返却され、CommunicationからCommandが削除されることを検証する。
func TestPullCommandByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1 = DefaultCommandID + "-1"
		DefaultCommandID2 = DefaultCommandID + "-2"
		DefaultCommandID3 = DefaultCommandID + "-3"
		DefaultTime1      = DefaultTime.Add(1 * time.Minute)
		DefaultTime2      = DefaultTime.Add(2 * time.Minute)
		DefaultTime3      = DefaultTime.Add(3 * time.Minute)
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3},
		times:      []time.Time{DefaultTime1, DefaultTime2, DefaultTime3},
	}
	communication := NewInstance(gen, DefaultID)
	communication.PushCommand(CommandTypeARM)
	id := communication.PushCommand(CommandTypeDISARM)
	communication.PushCommand(CommandTypeLAND)

	cType, err := communication.PullCommandByID(id)

	a.Equal(cType, CommandTypeDISARM)
	a.Equal(communication.GetCommandIDs(), []CommandID{DefaultCommandID1, DefaultCommandID3})
	a.Nil(err)
}

// CommunicationからCommandを取得する。
// CommandIDに合致するCommandが存在しない場合、エラーが返却されることを検証する。
func TestNotFoundErrorWhenPullCommandByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1        = DefaultCommandID + "-1"
		DefaultCommandID2        = DefaultCommandID + "-2"
		DefaultCommandID3        = DefaultCommandID + "-3"
		DefaultCommandIDNotFound = DefaultCommandID + "-NF"
		DefaultTime1             = DefaultTime.Add(1 * time.Minute)
		DefaultTime2             = DefaultTime.Add(2 * time.Minute)
		DefaultTime3             = DefaultTime.Add(3 * time.Minute)
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3},
		times:      []time.Time{DefaultTime1, DefaultTime2, DefaultTime3},
	}
	communication := NewInstance(gen, DefaultID)
	communication.PushCommand(CommandTypeARM)
	communication.PushCommand(CommandTypeDISARM)
	communication.PushCommand(CommandTypeLAND)

	cType, err := communication.PullCommandByID(DefaultCommandIDNotFound)

	a.Equal(cType, CommandType(""))
	a.Equal(communication.GetCommandIDs(), []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3})
	a.Equal(err, ErrCannotPullCommand)
}

// CommunicationからCommandを取得する。
// Commandが0件の場合、エラーが返却されることを検証する。
func TestNoneCommandErrorWhenPullCommandByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandIDNotFound = DefaultCommandID + "-NF"
	)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	cType, err := communication.PullCommandByID(DefaultCommandIDNotFound)

	a.Equal(cType, CommandType(""))
	a.Equal(err, ErrCannotPullCommand)
}

// CommunicationからCommandIDリストを取得する。
// CommandIDはCommandをEdgeから古い順でCloudに取得しに来るため、
// CommandのTimeの昇順でソートされていることを検証する。
func TestGetCommandIDsFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1 = DefaultCommandID + "-1"
		DefaultCommandID2 = DefaultCommandID + "-2"
		DefaultCommandID3 = DefaultCommandID + "-3"
		DefaultTime1      = DefaultTime.Add(1 * time.Minute)
		DefaultTime2      = DefaultTime.Add(2 * time.Minute)
		DefaultTime3      = DefaultTime.Add(3 * time.Minute)
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3},
		times:      []time.Time{DefaultTime3, DefaultTime1, DefaultTime2},
	}
	communication := NewInstance(gen, DefaultID)
	communication.PushCommand(CommandTypeARM)
	communication.PushCommand(CommandTypeDISARM)
	communication.PushCommand(CommandTypeLAND)

	commandIDs := communication.GetCommandIDs()

	a.Equal(commandIDs, []CommandID{DefaultCommandID2, DefaultCommandID3, DefaultCommandID1})
}

// CommunicationからCommandIDリストを取得する。
// Commandが0件の場合でも、空振りしてエラーが発生しないことを検証する。
func TestNoneCommandsWhenGetCommandIDsFromCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	commandIDs := communication.GetCommandIDs()

	a.Len(commandIDs, 0)
}

// CommunicationからUploadMissionを取得する。
// CommandIDに合致するUploadMissionが返却され、CommunicationからUploadMissionが削除されることを検証する。
func TestPullUploadMissionByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1 = DefaultCommandID + "-1"
		DefaultCommandID2 = DefaultCommandID + "-2"
		DefaultCommandID3 = DefaultCommandID + "-3"
		DefaultTime1      = DefaultTime.Add(1 * time.Minute)
		DefaultTime2      = DefaultTime.Add(2 * time.Minute)
		DefaultTime3      = DefaultTime.Add(3 * time.Minute)
		DefaultMissionID1 = DefaultMissionID + "-1"
		DefaultMissionID2 = DefaultMissionID + "-2"
		DefaultMissionID3 = DefaultMissionID + "-3"
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3},
		times:      []time.Time{DefaultTime1, DefaultTime2, DefaultTime3},
	}
	communication := NewInstance(gen, DefaultID)
	communication.PushUploadMission(DefaultMissionID1)
	id := communication.PushUploadMission(DefaultMissionID2)
	communication.PushUploadMission(DefaultMissionID3)

	missionID, err := communication.PullUploadMissionByID(id)

	expectUploadMissions := []*UploadMission{
		{
			commandID: DefaultCommandID1,
			missionID: DefaultMissionID1,
		},
		{
			commandID: DefaultCommandID3,
			missionID: DefaultMissionID3,
		},
	}

	a.Equal(missionID, DefaultMissionID2)
	a.Equal(communication.uploadMissions, expectUploadMissions)
	a.Nil(err)
}

// CommunicationからUploadMissionを取得する。
// CommandIDに合致するUploadMissionが存在しない場合、エラーが返却されることを検証する。
func TestNotFoundErrorWhenPullUploadMissionByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1        = DefaultCommandID + "-1"
		DefaultCommandID2        = DefaultCommandID + "-2"
		DefaultCommandID3        = DefaultCommandID + "-3"
		DefaultCommandIDNotFound = DefaultCommandID + "-NF"
		DefaultTime1             = DefaultTime.Add(1 * time.Minute)
		DefaultTime2             = DefaultTime.Add(2 * time.Minute)
		DefaultTime3             = DefaultTime.Add(3 * time.Minute)
		DefaultMissionID1        = DefaultMissionID + "-1"
		DefaultMissionID2        = DefaultMissionID + "-2"
		DefaultMissionID3        = DefaultMissionID + "-3"
	)

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID1, DefaultCommandID2, DefaultCommandID3},
		times:      []time.Time{DefaultTime1, DefaultTime2, DefaultTime3},
	}
	communication := NewInstance(gen, DefaultID)
	communication.PushUploadMission(DefaultMissionID1)
	communication.PushUploadMission(DefaultMissionID2)
	communication.PushUploadMission(DefaultMissionID3)

	missionID, err := communication.PullUploadMissionByID(DefaultCommandIDNotFound)

	expectUploadMissions := []*UploadMission{
		{
			commandID: DefaultCommandID1,
			missionID: DefaultMissionID1,
		},
		{
			commandID: DefaultCommandID2,
			missionID: DefaultMissionID2,
		},
		{
			commandID: DefaultCommandID3,
			missionID: DefaultMissionID3,
		},
	}

	a.Equal(missionID, MissionID(""))
	a.Equal(communication.uploadMissions, expectUploadMissions)
	a.Equal(err, ErrCannotPullUploadMission)
}

// CommunicationからUploadMissionを取得する。
// UploadMissionが0件の場合でも、空振りしてエラーが発生しないことを検証する。
func TestNoneUploadMissionWhenPullUploadMissionByIDFromCommunication(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandIDNotFound = DefaultCommandID + "-NF"
	)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	missionID, err := communication.PullUploadMissionByID(DefaultCommandIDNotFound)

	a.Equal(missionID, MissionID(""))
	a.Equal(err, ErrCannotPullUploadMission)
}
