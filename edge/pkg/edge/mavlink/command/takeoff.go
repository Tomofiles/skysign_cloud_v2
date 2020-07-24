package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// AdapterTakeOff .
func AdapterTakeOff(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	takeoffRequest := mavsdk_rpc_action.TakeoffRequest{}
	response, err := action.Takeoff(ctx, &takeoffRequest)
	if err != nil {
		log.Println("takeoff command error:", err)
		return err
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		log.Println("takeoff command error:", err)
		return errors.New("no takeoff command success")
	}

	return nil
}
