package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrStartCommand = errors.New("no start command success")
)

// AdapterStart .
func AdapterStart(ctx context.Context, gr *grpc.ClientConn, support common.Support) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterStartInternal(ctx, support, mission)
}

// AdapterStartInternal .
func AdapterStartInternal(ctx context.Context, support common.Support, mission mavsdk_rpc_mission.MissionServiceClient) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("start command error: %v", err)
		}
	}()

	startRequest := mavsdk_rpc_mission.StartMissionRequest{}
	response, err := mission.StartMission(ctx, &startRequest)
	if err != nil {
		return
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		err = ErrStartCommand
		return
	}

	return nil
}
