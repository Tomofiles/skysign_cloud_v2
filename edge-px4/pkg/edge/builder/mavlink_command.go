package builder

import (
	"context"

	mavlink_command_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/mavlink/command"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"

	"google.golang.org/grpc"
)

// CommandAdapter .
type CommandAdapter struct {
	AdapterArm     func() error
	AdapterDisarm  func() error
	AdapterStart   func() error
	AdapterPause   func() error
	AdapterTakeoff func() error
	AdapterLand    func() error
	AdapterReturn  func() error
	AdapterUpload  func(*model.Mission) error
}

// MavlinkCommand .
func MavlinkCommand(
	ctx context.Context,
	gr *grpc.ClientConn,
) *CommandAdapter {
	adapterArm := func() error {
		return mavlink_command_adapter.AdapterArm(ctx, gr)
	}
	adapterDisarm := func() error {
		return mavlink_command_adapter.AdapterDisarm(ctx, gr)
	}
	adapterStart := func() error {
		return mavlink_command_adapter.AdapterStart(ctx, gr)
	}
	adapterPause := func() error {
		return mavlink_command_adapter.AdapterPause(ctx, gr)
	}
	adapterTakeoff := func() error {
		return mavlink_command_adapter.AdapterTakeoff(ctx, gr)
	}
	adapterLand := func() error {
		return mavlink_command_adapter.AdapterLand(ctx, gr)
	}
	adapterReturn := func() error {
		return mavlink_command_adapter.AdapterReturn(ctx, gr)
	}
	adapterUpload := func(mission *model.Mission) error {
		return mavlink_command_adapter.AdapterUpload(ctx, gr, mission)
	}

	commandAdapter := &CommandAdapter{
		AdapterArm:     adapterArm,
		AdapterDisarm:  adapterDisarm,
		AdapterStart:   adapterStart,
		AdapterPause:   adapterPause,
		AdapterTakeoff: adapterTakeoff,
		AdapterLand:    adapterLand,
		AdapterReturn:  adapterReturn,
		AdapterUpload:  adapterUpload,
	}
	return commandAdapter
}
