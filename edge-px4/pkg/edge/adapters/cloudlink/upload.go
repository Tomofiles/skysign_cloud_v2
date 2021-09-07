package cloudlink

import (
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// PullUploadMission .
func PullUploadMission(
	cloud string,
	support common.Support,
	vehicleID, commandID string,
) (*edge.UploadMission, error) {
	request := json.Marshal(&skysign_proto.PullUploadMissionRequest{})

	support.NotifyInfo("Send CLOUD data=%s", request)

	respBody, err := http.HttpClientDo(
		support,
		http.MethodPost,
		cloud+"/api/v1/communications/"+vehicleID+"/uploadmissions/"+commandID,
		request,
	)
	if err != nil {
		support.NotifyError("cloud upload http client error: %v", err)
		return nil, err
	}

	var response skysign_proto.PullUploadMissionResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		support.NotifyError("cloud upload response error: %v", err)
		return nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	uploadMission := &edge.UploadMission{
		ID:        response.CommandId,
		MissionID: response.MissionId,
	}

	return uploadMission, nil
}
