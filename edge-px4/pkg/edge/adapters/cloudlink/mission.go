package cloudlink

import (
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// GetUploadMission .
func GetUploadMission(
	cloud string,
	support common.Support,
	missionID string,
) (*model.Mission, error) {
	request := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	support.NotifyInfo("Send CLOUD data=%s", request)

	respBody, err := http.HttpClientDo(
		http.MethodGet,
		cloud+"/api/v1/uploadmissions/"+missionID,
		request,
	)
	if err != nil {
		return nil, fmt.Errorf("cloud mission http client error: %w", err)
	}

	var response skysign_proto.UploadMission
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("cloud mission response error: %w", err)
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	mission := &model.Mission{
		ID:        response.Id,
		Waypoints: []*model.Waypoints{},
	}
	for _, waypoint := range response.Waypoints {
		mission.Waypoints = append(mission.Waypoints, &model.Waypoints{
			Latitude:       waypoint.Latitude,
			Longitude:      waypoint.Longitude,
			RelativeHeight: waypoint.RelativeHeight,
			Speed:          waypoint.Speed,
		})
	}

	return mission, nil
}
