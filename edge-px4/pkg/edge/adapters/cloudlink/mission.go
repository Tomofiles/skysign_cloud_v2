package cloudlink

import (
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/http"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/json"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/common"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// GetUploadMission .
func GetUploadMission(
	cloud string,
	support common.Support,
	missionID string,
) (*model.Mission, error) {
	method := http.MethodGet
	url := fmt.Sprintf("/api/v1/uploadmissions/%s", missionID)
	request := json.Marshal(&skysign_proto.GetUploadMissionRequest{})

	support.NotifyInfo("SEND   , Mission  , Method=%s , API=%s, Message=%s", method, url, request)

	respBody, err := http.HttpClientDo(
		method,
		cloud+url,
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

	support.NotifyInfo("RECEIVE, Mission  , data=%s", respBody)

	mission := &model.Mission{
		ID:        response.Id,
		Waypoints: []*model.Waypoints{},
	}
	for _, waypoint := range response.Waypoints {
		mission.Waypoints = append(mission.Waypoints, &model.Waypoints{
			LatitudeDegree:    waypoint.Latitude,
			LongitudeDegree:   waypoint.Longitude,
			RelativeAltitudeM: waypoint.RelativeAltitude,
			SpeedMS:           waypoint.Speed,
		})
	}

	return mission, nil
}
