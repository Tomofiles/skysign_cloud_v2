package app

import (
	"context"
	"fleet-formation/pkg/fleet/adapters/postgresql"
	"fleet-formation/pkg/fleet/adapters/uuid"
	"fleet-formation/pkg/fleet/service"

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
	gen := uuid.NewFleetUUID()
	repo := postgresql.NewFleetRepository(gen)
	return Application{
		Services: Services{
			ManageFleet: service.NewManageFleetService(gen, repo, txm, psm),
			AssignFleet: service.NewAssignFleetService(gen, repo, txm),
		},
	}
}
