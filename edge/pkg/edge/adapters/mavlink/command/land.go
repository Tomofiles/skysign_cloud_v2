package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/domain/common"
	mavsdk_rpc_action "edge/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrLandCommand = errors.New("no land command success")
)

// AdapterLand .
func AdapterLand(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterLandInternal(ctx, support, action)
}

// AdapterLandInternal .
func AdapterLandInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("land command error: %v", err)
		}
	}()

	landRequest := mavsdk_rpc_action.LandRequest{}
	response, err := action.Land(ctx, &landRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrLandCommand
		return
	}

	return nil
}
