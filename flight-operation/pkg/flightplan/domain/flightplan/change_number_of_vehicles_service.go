package flightplan

import (
	"flight-operation/pkg/common/domain/event"
	"flight-operation/pkg/common/domain/txmanager"
)

// ChangeNumberOfVehicles .
func ChangeNumberOfVehicles(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	numberOfVehicles int,
) error {
	flightplan, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	flightplan.SetPublisher(pub)
	if err := flightplan.ChangeNumberOfVehicles(numberOfVehicles); err != nil {
		return err
	}

	if ret := repo.Save(tx, flightplan); ret != nil {
		return ret
	}

	return nil
}
