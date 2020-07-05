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
func Cloudlink(ctx context.Context, cloud string, telemetry telemetry.Telemetry) <-chan *edge.Command {
	commandStream := make(chan *edge.Command)

	go func() {
		defer close(commandStream)
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
							commandStream <- command
						}
					}
				}
			}
		}
	}()

	return commandStream
}
