package communication

import (
	"errors"
	"remote-communication/pkg/common/domain/event"
	"sort"
	"time"
)

// ID .
type ID string

var (
	// ErrCannotPullCommand .
	ErrCannotPullCommand = errors.New("cannot pull command by id")
	// ErrCannotPullUploadMission .
	ErrCannotPullUploadMission = errors.New("cannot pull upload mission by id")
)

// Communication .
type Communication struct {
	id             ID
	telemetry      *Telemetry
	commands       []*Command
	uploadMissions []*UploadMission
	gen            Generator
	pub            event.Publisher
}

// SetPublisher .
func (c *Communication) SetPublisher(pub event.Publisher) {
	c.pub = pub
}

// GetID .
func (c *Communication) GetID() ID {
	return c.id
}

// PushTelemetry .
func (c *Communication) PushTelemetry(snapshot TelemetrySnapshot) {
	c.telemetry = NewTelemetryBySnapshot(snapshot)

	if c.pub != nil {
		c.pub.Publish(TelemetryUpdatedEvent{
			CommunicationID: c.id,
			Telemetry:       snapshot,
		})
	}
}

// PullTelemetry .
func (c *Communication) PullTelemetry() TelemetrySnapshot {
	return c.telemetry.GetSnapshot()
}

// GetCommandIDs .
func (c *Communication) GetCommandIDs() []CommandID {
	var commandIDs []CommandID
	commands := append([]*Command{}, c.commands...)
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].time.Before(commands[j].time)
	})
	for _, command := range commands {
		commandIDs = append(commandIDs, command.id)
	}
	return commandIDs
}

// PushCommand .
func (c *Communication) PushCommand(cType CommandType) CommandID {
	if IsFollowArmCommandPushPolicy(cType, c) {
		c.pushCommandDo(CommandTypeARM)
	}
	return c.pushCommandDo(cType)
}

// PushUploadMission .
func (c *Communication) PushUploadMission(missionID MissionID) CommandID {
	id := c.pushCommandDo(CommandTypeUPLOAD)
	c.uploadMissions = append(c.uploadMissions, NewUploadMission(id, missionID))
	return id
}

func (c *Communication) pushCommandDo(cType CommandType) CommandID {
	id := c.gen.NewCommandID()
	time := c.gen.NewTime()
	c.commands = append(c.commands, NewCommand(id, cType, time))
	return id
}

// PullCommandById .
func (c *Communication) PullCommandByID(commandID CommandID) (CommandType, error) {
	var command *Command
	var commands []*Command
	for _, cmd := range c.commands {
		if cmd.id == commandID {
			command = cmd
		} else {
			commands = append(commands, cmd)
		}
	}
	if command == nil {
		return "", ErrCannotPullCommand
	}
	c.commands = commands
	return command.cType, nil
}

// PullUploadMissionByID .
func (c *Communication) PullUploadMissionByID(commandID CommandID) (MissionID, error) {
	var uploadMission *UploadMission
	var uploadMissions []*UploadMission
	for _, um := range c.uploadMissions {
		if um.commandID == commandID {
			uploadMission = um
		} else {
			uploadMissions = append(uploadMissions, um)
		}
	}
	if uploadMission == nil {
		return "", ErrCannotPullUploadMission
	}
	c.uploadMissions = uploadMissions
	return uploadMission.missionID, nil
}

// Generator .
type Generator interface {
	NewCommandID() CommandID
	NewTime() time.Time
}
