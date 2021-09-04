package builder

import (
	"context"
	"edge/pkg/edge"
	mavlink "edge/pkg/edge/adapters/mavlink/command"
	"edge/pkg/edge/common"

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
	AdapterUpload  func(*edge.Mission) error
}

// MavlinkCommand .
func MavlinkCommand(
	ctx context.Context,
	gr *grpc.ClientConn,
	support common.Support,
) *CommandAdapter {
	adapterArm := func() error {
		return mavlink.AdapterArm(ctx, gr, support)
	}
	adapterDisarm := func() error {
		return mavlink.AdapterDisarm(ctx, gr, support)
	}
	adapterStart := func() error {
		return mavlink.AdapterStart(ctx, gr, support)
	}
	adapterPause := func() error {
		return mavlink.AdapterPause(ctx, gr, support)
	}
	adapterTakeoff := func() error {
		return mavlink.AdapterTakeoff(ctx, gr, support)
	}
	adapterLand := func() error {
		return mavlink.AdapterLand(ctx, gr, support)
	}
	adapterReturn := func() error {
		return mavlink.AdapterReturn(ctx, gr, support)
	}
	adapterUpload := func(mission *edge.Mission) error {
		return mavlink.AdapterUpload(ctx, gr, support, mission)
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
