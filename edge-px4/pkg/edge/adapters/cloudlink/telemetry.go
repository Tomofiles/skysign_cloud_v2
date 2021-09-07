package cloudlink

import (
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/telemetry"
	"encoding/json"
)

// PushTelemetry .
func PushTelemetry(
	cloud string,
	support common.Support,
	telemetry telemetry.Telemetry,
) (string, *edge.CommandIDs, error) {
	telem, err := telemetry.Get()
	if err != nil {
		support.NotifyInfo("cloud telemetry request error: %v", err)
		return "", nil, err
	}

	jsonData, _ := json.Marshal(telem)
	support.NotifyInfo("Send CLOUD data=%s", jsonData)

	respBody, err := http.HttpClientDo(
		support,
		http.MethodPost,
		cloud+"/api/v1/communications/"+telem.ID+"/telemetry",
		jsonData,
	)
	if err != nil {
		support.NotifyError("cloud telemetry http client error: %v", err)
		return "", nil, err
	}

	var commandIDs edge.CommandIDs
	err = json.Unmarshal(respBody, &commandIDs)
	if err != nil {
		support.NotifyError("cloud telemetry response error: %v", err)
		return "", nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	return telem.ID, &commandIDs, nil
}
