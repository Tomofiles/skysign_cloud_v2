package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// AdapterLand .
func AdapterLand(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	landRequest := mavsdk_rpc_action.LandRequest{}
	response, err := action.Land(ctx, &landRequest)
	if err != nil {
		log.Println("land command error:", err)
		return err
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		log.Println("land command error:", err)
		return errors.New("no land command success")
	}

	return nil
}
