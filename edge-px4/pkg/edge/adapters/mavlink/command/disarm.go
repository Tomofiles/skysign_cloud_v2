package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_action "edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrDisarmCommand = errors.New("no disarm command success")
)

// AdapterDisarm .
func AdapterDisarm(ctx context.Context, gr *grpc.ClientConn) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterDisarmInternal(ctx, action)
}

// AdapterDisarmInternal .
func AdapterDisarmInternal(ctx context.Context, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("disarm command error: %w", err)
		}
	}()

	disarmRequest := mavsdk_rpc_action.DisarmRequest{}
	response, err := action.Disarm(ctx, &disarmRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrDisarmCommand
		return
	}

	return
}
