package app

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/adapters/postgresql"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/adapters/uuid"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// NewApplication .
func NewApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) Application {
	return newApplication(ctx, txm, psm)
}

func newApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) Application {
	missionGen := uuid.NewMissionUUID()
	missionRepo := postgresql.NewMissionRepository(missionGen)
	return Application{
		Services: Services{
			ManageMission: service.NewManageMissionService(missionGen, missionRepo, txm, psm),
		},
	}
}
