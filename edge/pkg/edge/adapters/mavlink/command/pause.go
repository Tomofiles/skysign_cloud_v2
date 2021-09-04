package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/domain/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrPauseCommand = errors.New("no pause command success")
)

// AdapterPause .
func AdapterPause(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterPauseInternal(ctx, support, mission)
}

// AdapterPauseInternal .
func AdapterPauseInternal(ctx context.Context, support common.Support, mission mavsdk_rpc_mission.MissionServiceClient) (err error) {
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
