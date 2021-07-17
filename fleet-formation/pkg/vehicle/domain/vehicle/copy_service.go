package vehicle

import (
	"errors"
	"fleet-formation/pkg/vehicle/domain/event"
	"fleet-formation/pkg/vehicle/domain/txmanager"
)

// CarbonCopyVehicle .
func CarbonCopyVehicle(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	originalID ID,
	newID ID,
	fleetID FleetID,
) error {
	_, err := repo.GetByID(tx, newID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	} else if err == nil {
		return nil
	}

	original, err := repo.GetByID(tx, originalID)
	if err != nil {
		return err
	}

	vehicle := Copy(gen, newID, original)

	if err := repo.Save(tx, vehicle); err != nil {
		return err
	}

	pub.Publish(CopiedVehicleCreatedEvent{
		ID:              vehicle.GetID(),
		CommunicationID: vehicle.GetCommunicationID(),
		FleetID:         fleetID,
	})
	return nil
}
