package cloudlink

import (
	"edge-px4/pkg/edge/adapters/http"
	"edge-px4/pkg/edge/adapters/json"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"

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
		support.NotifyInfo("cloud telemetry request error: %v", err)
		return "", nil, err
	}

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

	support.NotifyInfo("Send CLOUD data=%s", request)

	respBody, err := http.HttpClientDo(
		support,
		http.MethodPost,
		cloud+"/api/v1/communications/"+snapshot.ID+"/telemetry",
		request,
	)
	if err != nil {
		support.NotifyError("cloud telemetry http client error: %v", err)
		return "", nil, err
	}

	var response skysign_proto.PushTelemetryResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		support.NotifyError("cloud telemetry response error: %v", err)
		return "", nil, err
	}

	support.NotifyInfo("Receive CLOUD data=%s", respBody)

	commandIDs := &model.CommandIDs{
		CommandIds: []string{},
	}
	for _, commandID := range response.CommandIds {
		commandIDs.CommandIds = append(commandIDs.CommandIds, commandID)
	}

	return snapshot.ID, commandIDs, nil
}
