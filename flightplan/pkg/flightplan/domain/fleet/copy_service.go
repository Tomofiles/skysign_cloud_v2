package fleet

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// CarbonCopyFleet .
func CarbonCopyFleet(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	originalID flightplan.ID,
	newID flightplan.ID,
) error {
	original, err := repo.GetByFlightplanID(tx, originalID)
	if err != nil {
		return err
	}

	fleet := Copy(gen, newID, original)

	if err := repo.Save(tx, fleet); err != nil {
		return err
	}

	return nil
}
