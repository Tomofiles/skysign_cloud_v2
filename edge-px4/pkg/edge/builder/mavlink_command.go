package builder

import (
	"context"
	mavlink_command_adapter "edge-px4/pkg/edge/adapters/mavlink/command"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"

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
	support common.Support,
) *CommandAdapter {
	adapterArm := func() error {
		return mavlink_command_adapter.AdapterArm(ctx, gr, support)
	}
	adapterDisarm := func() error {
		return mavlink_command_adapter.AdapterDisarm(ctx, gr, support)
	}
	adapterStart := func() error {
		return mavlink_command_adapter.AdapterStart(ctx, gr, support)
	}
	adapterPause := func() error {
		return mavlink_command_adapter.AdapterPause(ctx, gr, support)
	}
	adapterTakeoff := func() error {
		return mavlink_command_adapter.AdapterTakeoff(ctx, gr, support)
	}
	adapterLand := func() error {
		return mavlink_command_adapter.AdapterLand(ctx, gr, support)
	}
	adapterReturn := func() error {
		return mavlink_command_adapter.AdapterReturn(ctx, gr, support)
	}
	adapterUpload := func(mission *model.Mission) error {
		return mavlink_command_adapter.AdapterUpload(ctx, gr, support, mission)
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
