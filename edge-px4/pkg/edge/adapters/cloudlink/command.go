package cloudlink

import (
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// PullCommand .
func PullCommand(
	cloud string,
	support common.Support,
	vehicleID, commandID string,
) (*model.Command, error) {
	method := http.MethodPost
	url := fmt.Sprintf("/api/v1/communications/%s/commands/%s", vehicleID, commandID)
	request := json.Marshal(&skysign_proto.PullCommandRequest{})

	support.NotifyInfo("SEND   , Command  , Method=%s, API=%s, Message=%s", method, url, request)

	respBody, err := http.HttpClientDo(
		method,
		cloud+url,
		request,
	)
	if err != nil {
		return nil, fmt.Errorf("cloud command http client error: %w", err)
	}

	var response skysign_proto.PullCommandResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("cloud command response error: %w", err)
	}

	support.NotifyInfo("RECEIVE, Command  , data=%s", respBody)

	command := &model.Command{
		Type: response.Type.String(),
	}

	return command, nil
}
