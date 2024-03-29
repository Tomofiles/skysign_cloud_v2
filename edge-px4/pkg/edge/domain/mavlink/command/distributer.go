package command

import (
	"context"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
)

// CommandStream .
type CommandStream struct {
	ArmStream     <-chan struct{}
	DisarmStream  <-chan struct{}
	StartStream   <-chan struct{}
	PauseStream   <-chan struct{}
	TakeoffStream <-chan struct{}
	LandStream    <-chan struct{}
	ReturnStream  <-chan struct{}
}

// CommandDistributer .
func CommandDistributer(
	ctx context.Context,
	support common.Support,
	commandStream <-chan *model.Command,
) *CommandStream {
	armStream := make(chan struct{})
	disarmStream := make(chan struct{})
	startStream := make(chan struct{})
	pauseStream := make(chan struct{})
	takeoffStream := make(chan struct{})
	landStream := make(chan struct{})
	returnStream := make(chan struct{})

	go func() {
		defer close(armStream)
		defer close(disarmStream)
		defer close(startStream)
		defer close(pauseStream)
		defer close(takeoffStream)
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
					takeoffStream <- struct{}{}
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
		TakeoffStream: takeoffStream,
		LandStream:    landStream,
		ReturnStream:  returnStream,
	}
	return stream
}

// MissionDistributer .
func MissionDistributer(
	ctx context.Context,
	support common.Support,
	missionStream <-chan *model.Mission,
) <-chan *model.Mission {
	stream := make(chan *model.Mission)

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
