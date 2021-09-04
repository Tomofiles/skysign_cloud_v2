package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/domain/common"
	mavsdk_rpc_action "edge/pkg/protos/action"

	"google.golang.org/grpc"
)

var (
	ErrArmCommand = errors.New("no arm command success")
)

// AdapterArm .
func AdapterArm(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	action := mavsdk_rpc_action.NewActionServiceClient(gr)
	return AdapterArmInternal(ctx, support, action)
}

// AdapterArmInternal .
func AdapterArmInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("arm command error: %v", err)
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
