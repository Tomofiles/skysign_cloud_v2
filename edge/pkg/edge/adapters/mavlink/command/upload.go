package mavlink

import (
	"context"
	"errors"

	"edge/pkg/edge"
	"edge/pkg/edge/domain/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrUploadCommand = errors.New("no upload command success")
)

// AdapterUpload .
func AdapterUpload(ctx context.Context, gr *grpc.ClientConn, support common.Support, missionModel *edge.Mission) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterUploadInternal(ctx, support, mission, missionModel)
}

// AdapterUploadInternal .
func AdapterUploadInternal(ctx context.Context, support common.Support, mission mavsdk_rpc_mission.MissionServiceClient, missionModel *edge.Mission) (err error) {
	defer func() {
		if err != nil {
			support.NotifyError("upload command error: %v", err)
		}
	}()

	missionItems := make([]*mavsdk_rpc_mission.MissionItem, 0)

	for _, waypoint := range missionModel.Waypoints {
		missionItems = append(missionItems,
			&mavsdk_rpc_mission.MissionItem{
				LatitudeDeg:       waypoint.Latitude,
				LongitudeDeg:      waypoint.Longitude,
				RelativeAltitudeM: float32(waypoint.RelativeHeight),
				SpeedMS:           float32(waypoint.Speed),
			},
		)
	}

	uploadRequest := mavsdk_rpc_mission.UploadMissionRequest{
		MissionItems: missionItems,
	}
	response, err := mission.UploadMission(ctx, &uploadRequest)
	if err != nil {
		return
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		err = ErrUploadCommand
		return
	}

	return nil
}
