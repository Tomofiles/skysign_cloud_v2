package cloudlink

import (
	"edge/pkg/edge"
	"edge/pkg/edge/adapters/http"
	"edge/pkg/edge/domain/common"
	"encoding/json"
)

// PullUploadMission .
func PullUploadMission(
	cloud string,
	support common.Support,
	vehicleID, commandID string,
) (*edge.UploadMission, error) {
	support.NotifyInfo("Send CLOUD data=%s", "{}")

	respBody, err := http.HttpClientDo(
		support,
		http.MethodPost,
		cloud+"/api/v1/communications/"+vehicleID+"/uploadmissions/"+commandID,
		[]byte("{}"),
	)
	if err != nil {
		support.NotifyError("cloud upload http client error: %v", err)
		return nil, err
	}

	var uploadMission edge.UploadMission
	err = json.Unmarshal(respBody, &uploadMission)
	if err != nil {
		support.NotifyError("cloud upload response error: %v", err)
		return nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	return &uploadMission, nil
}
