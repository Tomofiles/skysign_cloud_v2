package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

var (
	ErrPauseCommand = errors.New("no pause command success")
)

// AdapterPause .
func AdapterPause(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	return AdapterPauseInternal(ctx, nil, mission)
}

// AdapterPauseInternal .
func AdapterPauseInternal(ctx context.Context, support Support, mission mavsdk_rpc_mission.MissionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("pause command error: %v", err)
		}
	}()

	pauseRequest := mavsdk_rpc_mission.PauseMissionRequest{}
	response, err := mission.PauseMission(ctx, &pauseRequest)
	if err != nil {
		return
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		err = ErrPauseCommand
		return
	}

	return nil
}
