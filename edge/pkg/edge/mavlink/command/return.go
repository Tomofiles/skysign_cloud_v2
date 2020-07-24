package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// AdapterReturn .
func AdapterReturn(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	action := mavsdk_rpc_action.NewActionServiceClient(gr)

	rtlRequest := mavsdk_rpc_action.ReturnToLaunchRequest{}
	response, err := action.ReturnToLaunch(ctx, &rtlRequest)
	if err != nil {
		log.Println("rtl command error:", err)
		return err
	}
	result := response.GetActionResult().GetResult()
	if result != mavsdk_rpc_action.ActionResult_SUCCESS {
		log.Println("rtl command error:", err)
		return errors.New("no rtl command success")
	}

	return nil
}
