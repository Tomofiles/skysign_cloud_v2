package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_mission "edge-px4/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrPauseCommand = errors.New("no pause command success")
)

// AdapterPause .
func AdapterPause(ctx context.Context, gr *grpc.ClientConn) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterPauseInternal(ctx, mission)
}

// AdapterPauseInternal .
func AdapterPauseInternal(ctx context.Context, mission mavsdk_rpc_mission.MissionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("pause command error: %w", err)
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
