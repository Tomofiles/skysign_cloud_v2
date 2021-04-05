package app

import (
	"context"
	"flightoperation/pkg/flightoperation/adapters/postgresql"
	"flightoperation/pkg/flightoperation/adapters/uuid"
	"flightoperation/pkg/flightoperation/domain/event"
	"flightoperation/pkg/flightoperation/domain/txmanager"
	"flightoperation/pkg/flightoperation/service"
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
			ManageFlightoperation:  service.NewManageFlightoperationService(flightoperationGen, flightoperationRepo, txm, psm),
			OperateFlightoperation: service.NewOperateFlightoperationService(flightoperationGen, flightoperationRepo, txm, psm),
		},
	}
}
