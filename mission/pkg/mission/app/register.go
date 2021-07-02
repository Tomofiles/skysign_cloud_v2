package app

import (
	"context"
	"mission/pkg/mission/adapters/postgresql"
	"mission/pkg/mission/adapters/uuid"
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
	"mission/pkg/mission/service"
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
