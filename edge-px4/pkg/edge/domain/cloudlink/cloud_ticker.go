package cloudlink

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/common"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
)

// CommandStream .
type CommandStream struct {
	CommandStream <-chan *model.Command
	MissionStream <-chan *model.Mission
}

// CloudTicker .
func CloudTicker(
	ctx context.Context,
	support common.Support,
	ticker common.Ticker,
	pushTelemetry func() (string, *model.CommandIDs, error),
	pullCommand func(vehicleID, commandID string) (*model.Command, error),
	pullUploadMission func(vehicleID, commandID string) (*model.UploadMission, error),
	getUploadMission func(missionID string) (*model.Mission, error),
) *CommandStream {
	commandStream := make(chan *model.Command)
	missionStream := make(chan *model.Mission)

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
	pushTelemetry func() (string, *model.CommandIDs, error),
	pullCommand func(vehicleID, commandID string) (*model.Command, error),
	pullUploadMission func(vehicleID, commandID string) (*model.UploadMission, error),
	getUploadMission func(missionID string) (*model.Mission, error),
	commandStream chan<- *model.Command,
	missionStream chan<- *model.Mission,
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
