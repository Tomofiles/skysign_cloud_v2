package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

var (
	ErrStartCommand = errors.New("no start command success")
)

// AdapterStart .
func AdapterStart(ctx context.Context, url string) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	return AdapterStartInternal(ctx, common.NewSupport(), mission)
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
