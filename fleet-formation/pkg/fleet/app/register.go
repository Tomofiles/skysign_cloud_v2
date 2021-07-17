package app

import (
	"context"
	"fleet-formation/pkg/fleet/adapters/postgresql"
	"fleet-formation/pkg/fleet/adapters/uuid"
	"fleet-formation/pkg/fleet/domain/event"
	"fleet-formation/pkg/fleet/domain/txmanager"
	"fleet-formation/pkg/fleet/service"
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
