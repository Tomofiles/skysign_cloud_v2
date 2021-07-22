package mavlink

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

// AdapterUpload .
func AdapterUpload(ctx context.Context, url string, missionModel *edge.Mission) error {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return err
	}

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

	mission := mavsdk_rpc_mission.NewMissionServiceClient(gr)

	uploadRequest := mavsdk_rpc_mission.UploadMissionRequest{
		MissionItems: missionItems,
	}
	response, err := mission.UploadMission(ctx, &uploadRequest)
	if err != nil {
		log.Println("upload command error:", err)
		return err
	}
	result := response.GetMissionResult().GetResult()
	if result != mavsdk_rpc_mission.MissionResult_SUCCESS {
		log.Println("upload command error:", err)
		return errors.New("no upload command success")
	}

	return nil
}
