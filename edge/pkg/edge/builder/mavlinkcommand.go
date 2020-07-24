package builder

import (
	"context"
	"edge/pkg/edge"
	mavlink "edge/pkg/edge/mavlink/command"
	"log"
)

// MavlinkCommand .
func MavlinkCommand(ctx context.Context, mavsdk string, commandStream <-chan *edge.Command) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("mavlink command done.")
				return
			case command, ok := <-commandStream:
				if !ok {
					log.Println("command none.")
					continue
				}
				switch command.Type {
				case "ARM":
					mavlink.AdapterArm(ctx, mavsdk)
				case "DISARM":
					mavlink.AdapterDisarm(ctx, mavsdk)
				case "UPLOAD":
					mavlink.AdapterUpload(ctx, mavsdk)
				case "START":
					mavlink.AdapterStart(ctx, mavsdk)
				case "PAUSE":
					mavlink.AdapterPause(ctx, mavsdk)
				case "TAKEOFF":
					mavlink.AdapterTakeOff(ctx, mavsdk)
				case "LAND":
					mavlink.AdapterLand(ctx, mavsdk)
				case "RETURN":
					mavlink.AdapterReturn(ctx, mavsdk)
				default:
					continue
				}
			}
		}
	}()

	return nil
}
