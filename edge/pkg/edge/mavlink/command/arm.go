package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// AdapterArm .
func AdapterArm(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	armRequest := mavsdk_rpc_action.ArmRequest{}
	response, err := action.Arm(ctx, &armRequest)
	if err != nil {
		log.Println("arm command error:", err)
		return err
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		log.Println("arm command error:", err)
		return errors.New("no arm command success")
	}

	return nil
}
