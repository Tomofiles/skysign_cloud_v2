package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// CarbonCopyFleet .
func CarbonCopyFleet(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	originalID ID,
	newID ID,
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

	fleet := Copy(gen, pub, newID, original)

	if err := repo.Save(tx, fleet); err != nil {
		return err
	}

	return nil
}
