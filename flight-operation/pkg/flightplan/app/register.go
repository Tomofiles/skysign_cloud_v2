package app

import (
	"context"
	"flight-operation/pkg/flightplan/adapters/postgresql"
	"flight-operation/pkg/flightplan/adapters/uuid"
	"flight-operation/pkg/flightplan/service"

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
	flightplanGen := uuid.NewFlightplanUUID()
	flightplanRepo := postgresql.NewFlightplanRepository(flightplanGen)
	return Application{
		Services: Services{
			ManageFlightplan:  service.NewManageFlightplanService(flightplanGen, flightplanRepo, txm, psm),
			ChangeFlightplan:  service.NewChangeFlightplanService(flightplanGen, flightplanRepo, txm, psm),
			ExecuteFlightplan: service.NewExecuteFlightplanService(flightplanGen, flightplanRepo, txm, psm),
		},
	}
}
