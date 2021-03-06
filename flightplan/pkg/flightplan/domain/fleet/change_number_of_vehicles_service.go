package fleet

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ChangeNumberOfVehicles .
func ChangeNumberOfVehicles(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	id flightplan.ID,
	numberOfVehicles int32,
) error {
	if fleet, err := repo.GetByFlightplanID(tx, id); err != nil {
		return err
	} else if fleet.isCarbonCopy {
		return ErrCannotChange
	}

	if ret := repo.DeleteByFlightplanID(tx, id); ret != nil {
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
