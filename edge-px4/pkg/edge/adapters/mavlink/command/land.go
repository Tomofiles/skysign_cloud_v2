package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrLandCommand = errors.New("no land command success")
)

// AdapterLand .
func AdapterLand(ctx context.Context, gr *grpc.ClientConn) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterLandInternal(ctx, action)
}

// AdapterLandInternal .
func AdapterLandInternal(ctx context.Context, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("land command error: %w", err)
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
