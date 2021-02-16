package app

import (
	"context"
	"flightplan/pkg/flightplan/generator"
	"flightplan/pkg/flightplan/infra"
	"flightplan/pkg/flightplan/service"
)

// NewApplication .
func NewApplication(ctx context.Context) Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) Application {
	flightplanGen := &generator.FlightplanUUID{}
	fleetGen := &generator.FleetUUID{}
	flightplanRepo := &infra.InmemoryFlightplanRepository{}
	fleetRepo := &infra.InmemoryFleetRepository{}
	pub := &infra.PublisherDirect{}
	return Application{
		Pub: pub,
		Services: Services{
			ManageFlightplan: service.NewManageFlightplanService(flightplanGen, flightplanRepo, pub),
			ManageFleet:      service.NewManageFleetService(fleetGen, fleetRepo, pub),
		},
	}
}
