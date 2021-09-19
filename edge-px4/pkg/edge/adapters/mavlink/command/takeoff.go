package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrTakeoffCommand = errors.New("no takeoff command success")
)

// AdapterTakeoff .
func AdapterTakeoff(ctx context.Context, gr *grpc.ClientConn) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterTakeoffInternal(ctx, action)
}

// AdapterTakeoffInternal .
func AdapterTakeoffInternal(ctx context.Context, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("takeoff command error: %w", err)
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
