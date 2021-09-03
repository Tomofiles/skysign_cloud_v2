package mavlink

import (
	"context"
	"errors"
	"log"

	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

var (
	ErrPauseCommand = errors.New("no pause command success")
)

// AdapterPause .
func AdapterPause(ctx context.Context, url string) error {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	return AdapterPauseInternal(ctx, glog.NewSupport(), mission)
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
