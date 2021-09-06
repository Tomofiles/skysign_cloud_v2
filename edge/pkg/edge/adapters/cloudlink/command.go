package cloudlink

import (
	"edge/pkg/edge"
	"edge/pkg/edge/adapters/http"
	"edge/pkg/edge/domain/common"
	"encoding/json"
)

// PullCommand .
func PullCommand(
	cloud string,
	support common.Support,
	vehicleID, commandID string,
) (*edge.Command, error) {
	support.NotifyInfo("Send CLOUD data=%s", "{}")

	respBody, err := http.HttpClientDo(
		support,
		http.MethodPost,
		cloud+"/api/v1/communications/"+vehicleID+"/commands/"+commandID,
		[]byte("{}"),
	)
	if err != nil {
		support.NotifyError("cloud command http client error: %v", err)
		return nil, err
	}

	var command edge.Command
	err = json.Unmarshal(respBody, &command)
	if err != nil {
		support.NotifyError("cloud command response error: %v", err)
		return nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	return &command, nil
}
