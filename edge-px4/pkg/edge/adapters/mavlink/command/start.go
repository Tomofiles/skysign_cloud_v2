package mavlink

import (
	"context"
	"errors"
	"fmt"

	mavsdk_rpc_mission "edge-px4/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrStartCommand = errors.New("no start command success")
)

// AdapterStart .
func AdapterStart(ctx context.Context, gr *grpc.ClientConn) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterStartInternal(ctx, mission)
}

// AdapterStartInternal .
func AdapterStartInternal(ctx context.Context, mission mavsdk_rpc_mission.MissionServiceClient) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("start command error: %w", err)
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
