package cloudlink

import (
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	"fmt"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// PushTelemetry .
func PushTelemetry(
	cloud string,
	support common.Support,
	telemetry model.Telemetry,
) (string, *model.CommandIDs, error) {
	snapshot, err := telemetry.Get()
	if err != nil {
		return "", nil, fmt.Errorf("cloud telemetry request error: %w", err)
	}

	method := http.MethodPost
	url := fmt.Sprintf("/api/v1/communications/%s/telemetry", snapshot.ID)
	request := json.Marshal(&skysign_proto.PushTelemetryRequest{
		Id: snapshot.ID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         snapshot.State.Latitude,
			Longitude:        snapshot.State.Longitude,
			Altitude:         snapshot.State.Altitude,
			RelativeAltitude: snapshot.State.RelativeAltitude,
			Speed:            snapshot.State.Speed,
			Armed:            snapshot.State.Armed,
			FlightMode:       snapshot.State.FlightMode,
			OrientationX:     snapshot.State.OrientationX,
			OrientationY:     snapshot.State.OrientationY,
			OrientationZ:     snapshot.State.OrientationZ,
			OrientationW:     snapshot.State.OrientationW,
		},
	})

	support.NotifyInfo("SEND   , Telemetry, Method=%s, API=%s, Message=%s", method, url, request)

	respBody, err := http.HttpClientDo(
		method,
		cloud+url,
		request,
	)
	if err != nil {
		return "", nil, fmt.Errorf("cloud telemetry http client error: %w", err)
	}

	var response skysign_proto.PushTelemetryResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", nil, fmt.Errorf("cloud telemetry response error: %w", err)
	}

	support.NotifyInfo("RECEIVE, Telemetry, data=%s", respBody)

	commandIDs := &model.CommandIDs{
		CommandIds: []string{},
	}
	for _, commandID := range response.CommandIds {
		commandIDs.CommandIds = append(commandIDs.CommandIds, commandID)
	}

	return snapshot.ID, commandIDs, nil
}
