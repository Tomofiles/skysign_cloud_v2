package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrReturnCommand = errors.New("no rtl command success")
)

// AdapterReturn .
func AdapterReturn(ctx context.Context, gr *grpc.ClientConn) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterReturnInternal(ctx, action)
}

// AdapterReturnInternal .
func AdapterReturnInternal(ctx context.Context, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rtl command error: %w", err)
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
