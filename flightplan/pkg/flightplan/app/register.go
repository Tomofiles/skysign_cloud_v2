package app

import (
	"context"
	"flightplan/pkg/flightplan/generator"
	"flightplan/pkg/flightplan/infra"
	"flightplan/pkg/flightplan/infra/postgresql"
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
	flightplanGen := &generator.FlightplanUUID{}
	fleetGen := &generator.FleetUUID{}
	flightplanRepo := postgresql.NewFlightplanRepository(flightplanGen)
	fleetRepo := postgresql.NewFleetRepository(fleetGen)
	pub := &infra.PublisherDirect{}
	return Application{
		Pub: pub,
		Services: Services{
			ManageFlightplan: service.NewManageFlightplanService(flightplanGen, flightplanRepo, pub, txm),
			ManageFleet:      service.NewManageFleetService(fleetGen, fleetRepo, pub, txm),
			AssignFleet:      service.NewAssignFleetService(fleetGen, fleetRepo, pub, txm),
		},
	}
}
