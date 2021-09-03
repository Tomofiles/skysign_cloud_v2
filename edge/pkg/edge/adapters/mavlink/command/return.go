package mavlink

import (
	"context"
	"errors"
	"log"

	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/common"
	mavsdk_rpc_action "edge/pkg/protos/action"
)

var (
	ErrReturnCommand = errors.New("no rtl command success")
)

// AdapterReturn .
func AdapterReturn(ctx context.Context, url string) error {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterReturnInternal(ctx, glog.NewSupport(), action)
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
