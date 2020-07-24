package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

// AdapterStart .
func AdapterStart(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	startRequest := mavsdk_rpc_mission.StartMissionRequest{}
	response, err := mission.StartMission(ctx, &startRequest)
	if err != nil {
		log.Println("start command error:", err)
		return err
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		log.Println("start command error:", err)
		return errors.New("no start command success")
	}

	return nil
}
