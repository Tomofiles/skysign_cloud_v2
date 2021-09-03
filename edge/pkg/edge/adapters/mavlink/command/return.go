package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/common"
	mavsdk_rpc_action "edge/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrReturnCommand = errors.New("no rtl command success")
)

// AdapterReturn .
func AdapterReturn(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterReturnInternal(ctx, support, action)
}

// AdapterReturnInternal .
func AdapterReturnInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("rtl command error: %v", err)
		}
	}()

	rtlRequest := mavsdk_rpc_action.ReturnToLaunchRequest{}
	response, err := action.ReturnToLaunch(ctx, &rtlRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrReturnCommand
		return
	}

	return nil
}
