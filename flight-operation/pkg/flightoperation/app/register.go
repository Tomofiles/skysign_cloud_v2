package app

import (
	"context"
	"flight-operation/pkg/flightoperation/adapters/postgresql"
	"flight-operation/pkg/flightoperation/adapters/uuid"
	"flight-operation/pkg/flightoperation/service"

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
	flightoperationGen := uuid.NewFlightoperationUUID()
	flightoperationRepo := postgresql.NewFlightoperationRepository(flightoperationGen)
	return Application{
		Services: Services{
			ManageFlightoperation:  service.NewManageFlightoperationService(flightoperationGen, flightoperationRepo, txm, psm),
			OperateFlightoperation: service.NewOperateFlightoperationService(flightoperationGen, flightoperationRepo, txm, psm),
		},
	}
}
