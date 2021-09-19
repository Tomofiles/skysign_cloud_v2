package mavlink

import (
	"context"
	"errors"
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_mission "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/mission"

	"google.golang.org/grpc"
)

var (
	ErrUploadCommand = errors.New("no upload command success")
)

// AdapterUpload .
func AdapterUpload(ctx context.Context, gr *grpc.ClientConn, missionModel *model.Mission) error {
	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)
	return AdapterUploadInternal(ctx, mission, missionModel)
}

// AdapterUploadInternal .
func AdapterUploadInternal(ctx context.Context, mission mavsdk_rpc_mission.MissionServiceClient, missionModel *model.Mission) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("upload command error: %w", err)
		}
	}()

	missionItems := make([]*mavsdk_rpc_mission.MissionItem, 0)

	for _, waypoint := range missionModel.Waypoints {
		missionItems = append(missionItems,
			&mavsdk_rpc_mission.MissionItem{
				LatitudeDeg:       waypoint.LatitudeDegree,
				LongitudeDeg:      waypoint.LongitudeDegree,
				RelativeAltitudeM: float32(waypoint.RelativeAltitudeM),
				SpeedMS:           float32(waypoint.SpeedMS),
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
