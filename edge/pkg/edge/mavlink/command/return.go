package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

var (
	ErrReturnCommand = errors.New("no rtl command success")
)

// AdapterReturn .
func AdapterReturn(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterReturnInternal(ctx, nil, action)
}

// AdapterReturnInternal .
func AdapterReturnInternal(ctx context.Context, support Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
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
