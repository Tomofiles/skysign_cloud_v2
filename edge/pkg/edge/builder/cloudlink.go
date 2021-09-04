package builder

import (
	"context"
	"edge/pkg/edge"
	"edge/pkg/edge/cloudlink"
	"edge/pkg/edge/domain/telemetry"
	"log"
	"time"
)

// CommandStream .
type CommandStream struct {
	CommandStream <-chan *edge.Command
	MissionStream <-chan *edge.Mission
}

// Cloudlink .
func Cloudlink(ctx context.Context, cloud string, telemetry telemetry.Telemetry) *CommandStream {
	commandStream := make(chan *edge.Command)
	missionStream := make(chan *edge.Mission)

	go func() {
		defer close(commandStream)
		defer close(missionStream)
		t := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				log.Println("telemetry ticker stop.")
				t.Stop()
				return
			case <-t.C:
				id, commIDs, err := cloudlink.PushTelemetry(cloud, telemetry)
				if err == nil {
					for _, commID := range commIDs.CommandIds {
						command, err := cloudlink.PullCommand(cloud, id, commID)
						if err == nil {
							if command.Type == "UPLOAD" {
								mission, err := cloudlink.PullMission(cloud, id, commID)
								if err == nil {
									missionStream <- mission
								}
							} else {
								commandStream <- command
							}
						}
					}
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
