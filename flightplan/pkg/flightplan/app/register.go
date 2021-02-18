package app

import (
	"context"
	"flightplan/pkg/flightplan/generator"
	"flightplan/pkg/flightplan/infra"
	"flightplan/pkg/flightplan/infra/postgresql"
	"flightplan/pkg/flightplan/service"
)

// NewApplication .
func NewApplication(ctx context.Context) Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) Application {
	db, err := postgresql.NewPostgresqlConnection()
	if err != nil {
		panic(err)
	}
	txm := postgresql.NewGormTransactionManager(db)
	flightplanGen := &generator.FlightplanUUID{}
	fleetGen := &generator.FleetUUID{}
	flightplanRepo := postgresql.NewFlightplanRepository(flightplanGen)
	fleetRepo := &infra.InmemoryFleetRepository{}
	pub := &infra.PublisherDirect{}
	return Application{
		Pub: pub,
		Services: Services{
			ManageFlightplan: service.NewManageFlightplanService(flightplanGen, flightplanRepo, pub, txm),
			ManageFleet:      service.NewManageFleetService(fleetGen, fleetRepo, pub),
			AssignFleet:      service.NewAssignFleetService(fleetGen, fleetRepo, pub),
		},
	}
}
