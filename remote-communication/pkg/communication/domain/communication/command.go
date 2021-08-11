package communication

import "time"

// CommandID .
type CommandID string

// Command .
type Command struct {
	id    CommandID
	cType string
	time  time.Time
}

// NewCommand .
func NewCommand(id CommandID, cType string, time time.Time) *Command {
	return nil
}
