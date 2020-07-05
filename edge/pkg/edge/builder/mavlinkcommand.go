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
				default:
					continue
				}
			}
		}
	}()

	return nil
}
