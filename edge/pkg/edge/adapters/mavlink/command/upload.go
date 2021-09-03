package mavlink

import (
	"context"
	"errors"
	"log"

	"edge/pkg/edge"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/common"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

var (
	ErrUploadCommand = errors.New("no upload command success")
)

// AdapterUpload .
func AdapterUpload(ctx context.Context, url string, missionModel *edge.Mission) error {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	return AdapterUploadInternal(ctx, glog.NewSupport(), mission, missionModel)
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
