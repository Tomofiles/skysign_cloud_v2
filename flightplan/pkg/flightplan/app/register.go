package app

import (
	"context"
	"flightplan/pkg/flightplan/adapters/inmemory"
	"flightplan/pkg/flightplan/adapters/postgresql"
	"flightplan/pkg/flightplan/adapters/uuid"
	"flightplan/pkg/flightplan/service"
)

// NewApplication .
func NewApplication(
	ctx context.Context,
	txm *postgresql.GormTransactionManager,
) Application {
	return newApplication(ctx, txm)
}

func newApplication(
	ctx context.Context,
	txm *postgresql.GormTransactionManager,
) Application {
	flightplanGen := uuid.NewFlightplanUUID()
	fleetGen := uuid.NewFleetUUID()
	flightplanRepo := postgresql.NewFlightplanRepository(flightplanGen)
	fleetRepo := postgresql.NewFleetRepository(fleetGen)
	pub := &inmemory.PublisherDirect{}
	return Application{
		Pub: pub,
		Services: Services{
			ManageFlightplan: service.NewManageFlightplanService(flightplanGen, flightplanRepo, pub, txm),
			ManageFleet:      service.NewManageFleetService(fleetGen, fleetRepo, pub, txm),
			AssignFleet:      service.NewAssignFleetService(fleetGen, fleetRepo, pub, txm),
		},
	}
}
