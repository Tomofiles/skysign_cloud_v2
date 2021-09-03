package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge/common"
	mavsdk_rpc_action "edge/pkg/protos/action"
)

var (
	ErrDisarmCommand = errors.New("no disarm command success")
)

// AdapterDisarm .
func AdapterDisarm(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterDisarmInternal(ctx, common.NewSupport(), action)
}

// AdapterDisarmInternal .
func AdapterDisarmInternal(ctx context.Context, support common.Support, action mavsdk_rpc_action.ActionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("disarm command error: %v", err)
		}
	}()

	disarmRequest := mavsdk_rpc_action.DisarmRequest{}
	response, err := action.Disarm(ctx, &disarmRequest)
	if err != nil {
		return
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		err = ErrDisarmCommand
		return
	}

	return
}
