package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/common"
	mavsdk_rpc_action "edge/pkg/protos/action"
)

var (
	ErrLandCommand = errors.New("no land command success")
)

// AdapterLand .
func AdapterLand(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	return AdapterLandInternal(ctx, glog.NewSupport(), action)
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
