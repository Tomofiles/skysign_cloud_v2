package fleet

import (
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ChangeNumberOfVehicles .
func ChangeNumberOfVehicles(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	id ID,
	numberOfVehicles int,
) error {
	if fleet, err := repo.GetByID(tx, id); err != nil {
		return err
	} else if fleet.isCarbonCopy {
		return ErrCannotChange
	}

	if ret := repo.Delete(tx, id); ret != nil {
		return ret
	}

	newFleet := NewInstance(
		gen,
		id,
		numberOfVehicles)
	for _, assignmentID := range newFleet.GetAllAssignmentID() {
		newFleet.AddNewEvent(assignmentID)
	}
	if ret := repo.Save(tx, newFleet); ret != nil {
		return ret
	}

	return nil
}
