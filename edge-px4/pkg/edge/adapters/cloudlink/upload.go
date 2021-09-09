package cloudlink

import (
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// PullUploadMission .
func PullUploadMission(
	cloud string,
	support common.Support,
	vehicleID, commandID string,
) (*model.UploadMission, error) {
	method := http.MethodPost
	url := fmt.Sprintf("/api/v1/communications/%s/uploadmissions/%s", vehicleID, commandID)
	request := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	support.NotifyInfo("SEND   , Upload   , Method=%s, API=%s, Message=%s", method, url, request)

	respBody, err := http.HttpClientDo(
		method,
		cloud+url,
		request,
	)
	if err != nil {
		return nil, fmt.Errorf("cloud upload http client error: %w", err)
	}

	var response skysign_proto.PullUploadMissionResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("cloud upload response error: %w", err)
	}

	support.NotifyInfo("RECEIVE, Upload   , data=%s", respBody)

	uploadMission := &model.UploadMission{
		ID:        response.CommandId,
		MissionID: response.MissionId,
	}

	return uploadMission, nil
}
