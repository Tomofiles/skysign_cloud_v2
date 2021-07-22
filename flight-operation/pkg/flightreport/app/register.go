package app

import (
	"context"
	"flight-operation/pkg/common/domain/event"
	"flight-operation/pkg/common/domain/txmanager"
	"flight-operation/pkg/flightreport/adapters/postgresql"
	"flight-operation/pkg/flightreport/adapters/uuid"
	"flight-operation/pkg/flightreport/service"
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
	flightreportGen := uuid.NewFlightreportUUID()
	flightreportRepo := postgresql.NewFlightreportRepository(flightreportGen)
	return Application{
		Services: Services{
			ManageFlightreport: service.NewManageFlightreportService(flightreportGen, flightreportRepo, txm, psm),
		},
	}
}
