package app

import (
	"context"
	"fleet-formation/pkg/vehicle/adapters/postgresql"
	"fleet-formation/pkg/vehicle/adapters/uuid"
	"fleet-formation/pkg/vehicle/service"

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
	vehicleGen := uuid.NewVehicleUUID()
	vehicleRepo := postgresql.NewVehicleRepository(vehicleGen)
	return Application{
		Services: Services{
			ManageVehicle: service.NewManageVehicleService(vehicleGen, vehicleRepo, txm, psm),
		},
	}
}
