package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/common"
	mavsdk_rpc_action "edge/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrDisarmCommand = errors.New("no disarm command success")
)

// AdapterDisarm .
func AdapterDisarm(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterDisarmInternal(ctx, support, action)
}

// AdapterDisarmInternal .
func AdapterDisarmInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("disarm command error: %v", err)
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
