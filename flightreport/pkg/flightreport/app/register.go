package app

import (
	"context"
	"flightreport/pkg/flightreport/adapters/postgresql"
	"flightreport/pkg/flightreport/adapters/uuid"
	"flightreport/pkg/flightreport/domain/event"
	"flightreport/pkg/flightreport/domain/txmanager"
	"flightreport/pkg/flightreport/service"
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
	flightoperationGen := uuid.NewFlightoperationUUID()
	flightoperationRepo := postgresql.NewFlightoperationRepository(flightoperationGen)
	return Application{
		Services: Services{
			ManageFlightoperation: service.NewManageFlightoperationService(flightoperationGen, flightoperationRepo, txm, psm),
		},
	}
}
