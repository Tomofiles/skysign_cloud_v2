package mavlink

import (
	"context"
	"errors"

	"edge-px4/pkg/edge/domain/common"
	mavsdk_rpc_action "edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrTakeoffCommand = errors.New("no takeoff command success")
)

// AdapterTakeoff .
func AdapterTakeoff(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterTakeoffInternal(ctx, support, action)
}

// AdapterTakeoffInternal .
func AdapterTakeoffInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("takeoff command error: %v", err)
		}
	}()

	takeoffRequest := mavsdk_rpc_action.TakeoffRequest{}
	response, err := action.Takeoff(ctx, &takeoffRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrTakeoffCommand
		return
	}

	return nil
}
