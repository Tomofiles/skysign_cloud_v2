package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

// AdapterPause .
func AdapterPause(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	pauseRequest := mavsdk_rpc_mission.PauseMissionRequest{}
	response, err := mission.PauseMission(ctx, &pauseRequest)
	if err != nil {
		log.Println("pause command error:", err)
		return err
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		log.Println("pause command error:", err)
		return errors.New("no pause command success")
	}

	return nil
}
