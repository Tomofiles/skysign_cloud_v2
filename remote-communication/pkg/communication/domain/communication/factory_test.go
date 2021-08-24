package communication

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Communicationを一つ新しく作成し、初期状態を検証する。
func TestCreateNewCommunication(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	expectTelemetry := NewTelemetry()
	var expectCommands []*Command
	var expectUploadMissions []*UploadMission

	a.Equal(communication.GetID(), DefaultID)
	a.Equal(communication.telemetry, expectTelemetry)
	a.Equal(communication.commands, expectCommands)
	a.Equal(communication.uploadMissions, expectUploadMissions)
	a.Equal(communication.gen, gen)
}

// Communicationを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestCommunicationAssembleFromComponent(t *testing.T) {
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

	gen := &generatorMock{}

	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	commandComps := []*commandComponentMock{
		{
			id:    string(DefaultCommandID1),
			cType: string(CommandTypeARM),
			time:  DefaultTime1,
		},
		{
			id:    string(DefaultCommandID2),
			cType: string(CommandTypeDISARM),
			time:  DefaultTime2,
		},
		{
			id:    string(DefaultCommandID3),
			cType: string(CommandTypeLAND),
			time:  DefaultTime3,
		},
	}
	uploadMissionComps := []*uploadMissionComponentMock{
		{
			commandID: string(DefaultCommandID1),
			missionID: string(DefaultMissionID1),
		},
		{
			commandID: string(DefaultCommandID2),
			missionID: string(DefaultMissionID2),
		},
		{
			commandID: string(DefaultCommandID3),
			missionID: string(DefaultMissionID3),
		},
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := AssembleFrom(gen, &communicationComp)

	expectTelemetry := &Telemetry{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	expectCommands := []*Command{
		{
			id:    DefaultCommandID1,
			cType: CommandTypeARM,
			time:  DefaultTime1,
		},
		{
			id:    DefaultCommandID2,
			cType: CommandTypeDISARM,
			time:  DefaultTime2,
		},
		{
			id:    DefaultCommandID3,
			cType: CommandTypeLAND,
			time:  DefaultTime3,
		},
	}
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

	a.Equal(communication.GetID(), DefaultID)
	a.Equal(communication.telemetry, expectTelemetry)
	a.Equal(communication.commands, expectCommands)
	a.Equal(communication.uploadMissions, expectUploadMissions)
	a.Equal(communication.gen, gen)
}

// Communicationを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartCommunication(t *testing.T) {
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

	telemetry := &Telemetry{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	commands := []*Command{
		{
			id:    DefaultCommandID1,
			cType: CommandTypeARM,
			time:  DefaultTime1,
		},
		{
			id:    DefaultCommandID2,
			cType: CommandTypeDISARM,
			time:  DefaultTime2,
		},
		{
			id:    DefaultCommandID3,
			cType: CommandTypeLAND,
			time:  DefaultTime3,
		},
	}
	uploadMissions := []*UploadMission{
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

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)
	communication.telemetry = telemetry
	communication.commands = commands
	communication.uploadMissions = uploadMissions

	var communicationComp communicationComponentMock
	var commandComps []*commandComponentMock
	var uploadMissionComps []*uploadMissionComponentMock

	TakeApart(
		communication,
		func(id string) {
			communicationComp.id = id
		},
		func(latitude, longitude, altitude, relativeAltitude, speed, x, y, z, w float64, armed bool, flightMode string) {
			communicationComp.telemetry = &telemetryComponentMock{
				latitude:         latitude,
				longitude:        longitude,
				altitude:         altitude,
				relativeAltitude: relativeAltitude,
				speed:            speed,
				armed:            armed,
				flightMode:       flightMode,
				x:                x,
				y:                y,
				z:                z,
				w:                w,
			}
		},
		func(id, cType string, time time.Time) {
			commandComps = append(
				commandComps,
				&commandComponentMock{
					id:    id,
					cType: cType,
					time:  time,
				},
			)
		},
		func(commandID, missionID string) {
			uploadMissionComps = append(
				uploadMissionComps,
				&uploadMissionComponentMock{
					commandID: commandID,
					missionID: missionID,
				},
			)
		},
	)

	communicationComp.commands = commandComps
	communicationComp.uploadMissions = uploadMissionComps

	expectTelemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	expectCommandComps := []*commandComponentMock{
		{
			id:    string(DefaultCommandID1),
			cType: string(CommandTypeARM),
			time:  DefaultTime1,
		},
		{
			id:    string(DefaultCommandID2),
			cType: string(CommandTypeDISARM),
			time:  DefaultTime2,
		},
		{
			id:    string(DefaultCommandID3),
			cType: string(CommandTypeLAND),
			time:  DefaultTime3,
		},
	}
	expectUploadMissionComps := []*uploadMissionComponentMock{
		{
			commandID: string(DefaultCommandID1),
			missionID: string(DefaultMissionID1),
		},
		{
			commandID: string(DefaultCommandID2),
			missionID: string(DefaultMissionID2),
		},
		{
			commandID: string(DefaultCommandID3),
			missionID: string(DefaultMissionID3),
		},
	}
	expectCommunicationComp := communicationComponentMock{
		id:             string(DefaultID),
		telemetry:      &expectTelemetryComp,
		commands:       expectCommandComps,
		uploadMissions: expectUploadMissionComps,
	}
	a.Equal(communicationComp, expectCommunicationComp)
}
