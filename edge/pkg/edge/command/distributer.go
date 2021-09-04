package command

import (
	"context"
	"edge/pkg/edge"
	"edge/pkg/edge/common"
)

// CommandStream .
type CommandStream struct {
	ArmStream     <-chan struct{}
	DisarmStream  <-chan struct{}
	StartStream   <-chan struct{}
	PauseStream   <-chan struct{}
	TakeOffStream <-chan struct{}
	LandStream    <-chan struct{}
	ReturnStream  <-chan struct{}
}

// CommandDistributer .
func CommandDistributer(
	ctx context.Context,
	support common.Support,
	commandStream <-chan *edge.Command,
) *CommandStream {
	armStream := make(chan struct{})
	disarmStream := make(chan struct{})
	startStream := make(chan struct{})
	pauseStream := make(chan struct{})
	takeOffStream := make(chan struct{})
	landStream := make(chan struct{})
	returnStream := make(chan struct{})

	go func() {
		defer close(armStream)
		defer close(disarmStream)
		defer close(startStream)
		defer close(pauseStream)
		defer close(takeOffStream)
		defer close(landStream)
		defer close(returnStream)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("command distributer done")
				return
			case command, ok := <-commandStream:
				if !ok {
					support.NotifyInfo("command distributer close")
					return
				}
				switch command.Type {
				case "ARM":
					armStream <- struct{}{}
				case "DISARM":
					disarmStream <- struct{}{}
				case "START":
					startStream <- struct{}{}
				case "PAUSE":
					pauseStream <- struct{}{}
				case "TAKEOFF":
					takeOffStream <- struct{}{}
				case "LAND":
					landStream <- struct{}{}
				case "RETURN":
					returnStream <- struct{}{}
				default:
					support.NotifyError("command cannot distribute")
				}
			}
		}
	}()

	stream := &CommandStream{
		ArmStream:     armStream,
		DisarmStream:  disarmStream,
		StartStream:   startStream,
		PauseStream:   pauseStream,
		TakeOffStream: takeOffStream,
		LandStream:    landStream,
		ReturnStream:  returnStream,
	}
	return stream
}

// MissionDistributer .
func MissionDistributer(
	ctx context.Context,
	support common.Support,
	missionStream <-chan *edge.Mission,
) <-chan *edge.Mission {
	stream := make(chan *edge.Mission)

	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("mission distributer done")
				return
			case mission, ok := <-missionStream:
				if !ok {
					support.NotifyInfo("mission distributer close")
					return
				}
				stream <- mission
			}
		}
	}()

	return stream
}
