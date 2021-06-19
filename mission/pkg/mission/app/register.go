package app

import (
	"context"
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
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
	// vehicleGen := uuid.NewVehicleUUID()
	// vehicleRepo := postgresql.NewVehicleRepository(vehicleGen)
	return Application{
		Services: Services{
			// ManageVehicle: service.NewManageVehicleService(vehicleGen, vehicleRepo, txm, psm),
		},
	}
}
