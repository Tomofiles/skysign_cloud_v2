package builder

import (
	"context"
	"edge/pkg/edge"
	"edge/pkg/edge/adapters/cloudlink"
	"edge/pkg/edge/domain/common"
	"edge/pkg/edge/domain/mavlink/telemetry"
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
			return cloudlink.PushTelemetry(cloud, support, telemetry)
		},
		PullCommand: func(vehicleID, commandID string) (*edge.Command, error) {
			return cloudlink.PullCommand(cloud, support, vehicleID, commandID)
		},
		PullUploadMission: func(vehicleID, commandID string) (*edge.UploadMission, error) {
			return cloudlink.PullUploadMission(cloud, support, vehicleID, commandID)
		},
		GetUploadMission: func(missionID string) (*edge.Mission, error) {
			return cloudlink.GetUploadMission(cloud, support, missionID)
		},
	}
}
