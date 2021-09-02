package app

import (
	"context"
	"remote-communication/pkg/mission/adapters/postgresql"
	"remote-communication/pkg/mission/service"

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
	repo := postgresql.NewMissionRepository()
	return Application{
		Services: Services{
			ManageMission: service.NewManageMissionService(repo, txm, psm),
			EdgeMission:   service.NewEdgeMissionService(repo, txm, psm),
		},
	}
}
