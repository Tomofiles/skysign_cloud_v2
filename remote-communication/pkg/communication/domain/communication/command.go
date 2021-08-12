package communication

import "time"

// CommandID .
type CommandID string

// CommandType .
type CommandType string

const (
	// CommandTypeARM .
	CommandTypeARM = CommandType("ARM")
	// CommandTypeDISARM .
	CommandTypeDISARM = CommandType("DISARM")
	// CommandTypeUPLOAD .
	CommandTypeUPLOAD = CommandType("UPLOAD")
	// CommandTypeSTART .
	CommandTypeSTART = CommandType("START")
	// CommandTypePAUSE .
	CommandTypePAUSE = CommandType("PAUSE")
	// CommandTypeTAKEOFF .
	CommandTypeTAKEOFF = CommandType("TAKEOFF")
	// CommandTypeLAND .
	CommandTypeLAND = CommandType("LAND")
	// CommandTypeRETURN .
	CommandTypeRETURN = CommandType("RETURN")
	// CommandTypeNONE .
	CommandTypeNONE = CommandType("NONE")
)

// Command .
type Command struct {
	id    CommandID
	cType CommandType
	time  time.Time
}

// NewCommand .
func NewCommand(id CommandID, cType CommandType, time time.Time) *Command {
	return &Command{
		id:    id,
		cType: cType,
		time:  time,
	}
}
