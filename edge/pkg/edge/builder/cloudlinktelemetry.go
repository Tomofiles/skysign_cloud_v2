package builder

import (
	"context"
	"edge/pkg/edge/cloudlink"
	"edge/pkg/edge/telemetry"
	"log"
	"time"
)

// CloudlinkTelemetry .
func CloudlinkTelemetry(ctx context.Context, cloud string, telemetry telemetry.Telemetry) {
	go func(done <-chan struct{}) {
		t := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-done:
				log.Println("telemetry ticker stop.")
				t.Stop()
				return
			case <-t.C:
				cloudlink.PushTelemetry(cloud, telemetry)
			}
		}
	}(ctx.Done())
}
