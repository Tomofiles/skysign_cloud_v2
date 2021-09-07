package cloudlink

import (
	"context"
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/domain/common"
)

// CommandStream .
type CommandStream struct {
	CommandStream <-chan *edge.Command
	MissionStream <-chan *edge.Mission
}

// CloudTicker .
func CloudTicker(
	ctx context.Context,
	support common.Support,
	ticker common.Ticker,
	pushTelemetry func() (string, *edge.CommandIDs, error),
	pullCommand func(vehicleID, commandID string) (*edge.Command, error),
	pullUploadMission func(vehicleID, commandID string) (*edge.UploadMission, error),
	getUploadMission func(missionID string) (*edge.Mission, error),
) *CommandStream {
	commandStream := make(chan *edge.Command)
	missionStream := make(chan *edge.Mission)

	go func() {
		defer close(commandStream)
		defer close(missionStream)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("cloud ticker done")
				ticker.Stop()
				return
			case <-ticker.Tick():
				if err := CloudTickerInternal(
					pushTelemetry,
					pullCommand,
					pullUploadMission,
					getUploadMission,
					commandStream,
					missionStream,
				); err != nil {
					support.NotifyError("cloud ticker error: %v", err)
					continue
				}
			}
		}
	}()

	stream := &CommandStream{
		CommandStream: commandStream,
		MissionStream: missionStream,
	}
	return stream
}

// CloudTickerInternal .
func CloudTickerInternal(
	pushTelemetry func() (string, *edge.CommandIDs, error),
	pullCommand func(vehicleID, commandID string) (*edge.Command, error),
	pullUploadMission func(vehicleID, commandID string) (*edge.UploadMission, error),
	getUploadMission func(missionID string) (*edge.Mission, error),
	commandStream chan<- *edge.Command,
	missionStream chan<- *edge.Mission,
) error {
	id, commandIDs, err := pushTelemetry()
	if err != nil {
		return err
	}
	for _, commandID := range commandIDs.CommandIds {
		command, err := pullCommand(id, commandID)
		if err != nil {
			return err
		}
		if command.Type == "UPLOAD" {
			upload, err := pullUploadMission(id, commandID)
			if err != nil {
				return err
			}
			mission, err := getUploadMission(upload.MissionID)
			if err != nil {
				return err
			}
			missionStream <- mission
		} else {
			commandStream <- command
		}
	}

	return nil
}
