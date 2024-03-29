package vehicle

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// UpdateVehicle .
func UpdateVehicle(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	name string,
	communicationID CommunicationID,
) error {
	vehicle, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	vehicle.SetPublisher(pub)
	if err := vehicle.NameVehicle(name); err != nil {
		return err
	}
	if err := vehicle.GiveCommunication(communicationID); err != nil {
		return err
	}

	if ret := repo.Save(tx, vehicle); ret != nil {
		return ret
	}

	return nil
}
