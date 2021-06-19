package vehicle

import (
	"vehicle/pkg/vehicle/domain/event"
	"vehicle/pkg/vehicle/domain/txmanager"
)

// DeleteVehicle .
func DeleteVehicle(
	tx txmanager.Tx,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	vehicle, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	vehicle.SetPublisher(pub)
	if err := vehicle.RemoveCommunication(); err != nil {
		return err
	}

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	return nil
}
