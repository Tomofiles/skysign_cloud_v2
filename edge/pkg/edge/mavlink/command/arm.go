package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

var (
	ErrArmCommand = errors.New("no arm command success")
)

// AdapterArm .
func AdapterArm(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterArmInternal(ctx, nil, action)
}

// AdapterArmInternal .
func AdapterArmInternal(ctx context.Context, support Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
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

// Support .
type Support interface {
	NotifyError(format string, args ...interface{})
}
