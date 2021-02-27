package builder

import (
	"context"
	"edge/pkg/edge"
	"edge/pkg/edge/cloudlink"
	"edge/pkg/edge/telemetry"
	"log"
	"time"
)

// Cloudlink .
func Cloudlink(ctx context.Context, cloud string, telemetry telemetry.Telemetry) (<-chan *edge.Command, <-chan *edge.Mission) {
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
					for _, commID := range commIDs.CommIds {
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

	return commandStream, missionStream
}
