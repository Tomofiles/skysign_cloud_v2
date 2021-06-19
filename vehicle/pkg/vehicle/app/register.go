package app

import (
	"context"
	"vehicle/pkg/vehicle/adapters/postgresql"
	"vehicle/pkg/vehicle/adapters/uuid"
	"vehicle/pkg/vehicle/domain/event"
	"vehicle/pkg/vehicle/domain/txmanager"
	"vehicle/pkg/vehicle/service"
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
	vehicleGen := uuid.NewVehicleUUID()
	vehicleRepo := postgresql.NewVehicleRepository(vehicleGen)
	return Application{
		Services: Services{
			ManageVehicle: service.NewManageVehicleService(vehicleGen, vehicleRepo, txm, psm),
		},
	}
}
