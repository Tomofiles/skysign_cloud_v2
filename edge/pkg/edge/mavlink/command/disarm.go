package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// AdapterDisarm .
func AdapterDisarm(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	disarmRequest := mavsdk_rpc_action.DisarmRequest{}
	response, err := action.Disarm(ctx, &disarmRequest)
	if err != nil {
		log.Println("disarm command error:", err)
		return err
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		log.Println("disarm command error:", err)
		return errors.New("no disarm command success")
	}

	return nil
}
