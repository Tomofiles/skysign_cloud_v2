package communication

import (
	"remote-communication/pkg/common/domain/event"
	"time"
)

// ID .
type ID string

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
}

// PullTelemetry .
func (c *Communication) PullTelemetry() TelemetrySnapshot {
	return TelemetrySnapshot{}
}

// GetCommandIDs .
func (c *Communication) GetCommandIDs() []CommandID {
	var commandIDs []CommandID
	for _, command := range c.commands {
		// sort?
		commandIDs = append(commandIDs, command.id)
	}
	return commandIDs
}

// PushCommand .
func (c *Communication) PushCommand(cType string) CommandID {
	if IsFollowArmCommandPushPolicy(cType, c) {
		c.pushCommandDo("ARM")
	}
	return c.pushCommandDo(cType)
}

// PushUploadMission .
func (c *Communication) PushUploadMission(missionID MissionID) CommandID {
	id := c.pushCommandDo("UPLOAD")
	c.uploadMissions = append(c.uploadMissions, NewUploadMission(id, missionID))
	return id
}

func (c *Communication) pushCommandDo(cType string) CommandID {
	id := c.gen.NewCommandID()
	time := c.gen.NewTime()
	c.commands = append(c.commands, NewCommand(id, cType, time))
	return id
}

// PullCommandById .
func (c *Communication) PullCommandByID(commandID CommandID) string {
	return ""
}

// PullUploadMissionByID .
func (c *Communication) PullUploadMissionByID(commandID CommandID) MissionID {
	return ""
}

// Generator .
type Generator interface {
	NewCommandID() CommandID
	NewTime() time.Time
}
