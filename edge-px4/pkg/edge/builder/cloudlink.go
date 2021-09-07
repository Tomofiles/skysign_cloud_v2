package builder

import (
	"context"
	"edge-px4/pkg/edge"
	cloudlink_adapter "edge-px4/pkg/edge/adapters/cloudlink"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/telemetry"
)

// CloudlinkAdapter .
type CloudlinkAdapter struct {
	PushTelemetry     func() (string, *edge.CommandIDs, error)
	PullCommand       func(vehicleID, commandID string) (*edge.Command, error)
	PullUploadMission func(vehicleID, commandID string) (*edge.UploadMission, error)
	GetUploadMission  func(missionID string) (*edge.Mission, error)
}

// Cloudlink .
func Cloudlink(
	ctx context.Context,
	cloud string,
	support common.Support,
	telemetry telemetry.Telemetry,
) *CloudlinkAdapter {
	return &CloudlinkAdapter{
		PushTelemetry: func() (string, *edge.CommandIDs, error) {
			return cloudlink_adapter.PushTelemetry(cloud, support, telemetry)
		},
		PullCommand: func(vehicleID, commandID string) (*edge.Command, error) {
			return cloudlink_adapter.PullCommand(cloud, support, vehicleID, commandID)
		},
		PullUploadMission: func(vehicleID, commandID string) (*edge.UploadMission, error) {
			return cloudlink_adapter.PullUploadMission(cloud, support, vehicleID, commandID)
		},
		GetUploadMission: func(missionID string) (*edge.Mission, error) {
			return cloudlink_adapter.GetUploadMission(cloud, support, missionID)
		},
	}
}
