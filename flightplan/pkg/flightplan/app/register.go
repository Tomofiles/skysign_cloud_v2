package app

import (
	"context"
	"flightplan/pkg/flightplan/adapters/postgresql"
	"flightplan/pkg/flightplan/adapters/uuid"
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
	"flightplan/pkg/flightplan/service"
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
	flightplanGen := uuid.NewFlightplanUUID()
	fleetGen := uuid.NewFleetUUID()
	flightplanRepo := postgresql.NewFlightplanRepository(flightplanGen)
	fleetRepo := postgresql.NewFleetRepository(fleetGen)
	return Application{
		Services: Services{
			ManageFlightplan: service.NewManageFlightplanService(flightplanGen, flightplanRepo, txm, psm),
			ChangeFlightplan: service.NewChangeFlightplanService(flightplanGen, flightplanRepo, txm, psm),
			ManageFleet:      service.NewManageFleetService(fleetGen, fleetRepo, txm, psm),
			AssignFleet:      service.NewAssignFleetService(fleetGen, fleetRepo, txm),
		},
	}
}
