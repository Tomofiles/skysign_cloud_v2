package builder

import (
	"context"

	cloudlink_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/cloudlink"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/common"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
)

// CloudlinkAdapter .
type CloudlinkAdapter struct {
	PushTelemetry     func() (string, *model.CommandIDs, error)
	PullCommand       func(vehicleID, commandID string) (*model.Command, error)
	PullUploadMission func(vehicleID, commandID string) (*model.UploadMission, error)
	GetUploadMission  func(missionID string) (*model.Mission, error)
}

// Cloudlink .
func Cloudlink(
	ctx context.Context,
	cloud string,
	support common.Support,
	telemetry model.Telemetry,
) *CloudlinkAdapter {
	return &CloudlinkAdapter{
		PushTelemetry: func() (string, *model.CommandIDs, error) {
			return cloudlink_adapter.PushTelemetry(cloud, support, telemetry)
		},
		PullCommand: func(vehicleID, commandID string) (*model.Command, error) {
			return cloudlink_adapter.PullCommand(cloud, support, vehicleID, commandID)
		},
		PullUploadMission: func(vehicleID, commandID string) (*model.UploadMission, error) {
			return cloudlink_adapter.PullUploadMission(cloud, support, vehicleID, commandID)
		},
		GetUploadMission: func(missionID string) (*model.Mission, error) {
			return cloudlink_adapter.GetUploadMission(cloud, support, missionID)
		},
	}
}
