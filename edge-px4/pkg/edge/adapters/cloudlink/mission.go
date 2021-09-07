package cloudlink

import (
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/domain/common"
	"encoding/json"
)

// GetUploadMission .
func GetUploadMission(
	cloud string,
	support common.Support,
	missionID string,
) (*edge.Mission, error) {
	support.NotifyInfo("Send CLOUD data=%s", "{}")

	respBody, err := http.HttpClientDo(
		support,
		http.MethodGet,
		cloud+"/api/v1/uploadmissions/"+missionID,
		[]byte("{}"),
	)
	if err != nil {
		support.NotifyError("cloud mission http client error: %v", err)
		return nil, err
	}

	var mission edge.Mission
	err = json.Unmarshal(respBody, &mission)
	if err != nil {
		support.NotifyError("cloud mission response error: %v", err)
		return nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	return &mission, nil
}
