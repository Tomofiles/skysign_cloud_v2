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
	ErrTakeoffCommand = errors.New("no takeoff command success")
)

// AdapterTakeOff .
func AdapterTakeOff(ctx context.Context, url string) error {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterTakeoffInternal(ctx, glog.NewSupport(), action)
}

// AdapterTakeoffInternal .
func AdapterTakeoffInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("takeoff command error: %v", err)
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
