package builder

import (
	"context"
	"edge/pkg/edge"
	mavlink "edge/pkg/edge/adapters/mavlink/command"
	"edge/pkg/edge/common"

	"google.golang.org/grpc"
)

// MavlinkCommand .
func MavlinkCommand(
	ctx context.Context,
	gr *grpc.ClientConn,
	support common.Support,
	commandStream <-chan *edge.Command,
	missionStream <-chan *edge.Mission,
) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("mavlink command done.")
				return
			case command, ok := <-commandStream:
				if !ok {
					support.NotifyInfo("command none.")
					continue
				}
				switch command.Type {
				case "ARM":
					mavlink.AdapterArm(ctx, gr, support)
				case "DISARM":
					mavlink.AdapterDisarm(ctx, gr, support)
				case "START":
					mavlink.AdapterStart(ctx, gr, support)
				case "PAUSE":
					mavlink.AdapterPause(ctx, gr, support)
				case "TAKEOFF":
					mavlink.AdapterTakeoff(ctx, gr, support)
				case "LAND":
					mavlink.AdapterLand(ctx, gr, support)
				case "RETURN":
					mavlink.AdapterReturn(ctx, gr, support)
				default:
					continue
				}
			case mission, ok := <-missionStream:
				if !ok {
					support.NotifyInfo("mission none.")
					continue
				}
				mavlink.AdapterUpload(ctx, gr, support, mission)
			}
		}
	}()

	return nil
}
