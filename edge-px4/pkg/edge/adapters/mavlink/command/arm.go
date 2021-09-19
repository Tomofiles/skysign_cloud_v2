package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrArmCommand = errors.New("no arm command success")
)

// AdapterArm .
func AdapterArm(ctx context.Context, gr *grpc.ClientConn) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterArmInternal(ctx, action)
}

// AdapterArmInternal .
func AdapterArmInternal(ctx context.Context, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("arm command error: %w", err)
		}
	}()

	armRequest := mavsdk_rpc_action.ArmRequest{}
	response, err := action.Arm(ctx, &armRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrArmCommand
		return
	}

	return
}
