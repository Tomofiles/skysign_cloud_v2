package mission

import (
	"errors"
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
)

// CarbonCopyMission .
func CarbonCopyMission(
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

	mission := Copy(gen, newID, original)

	if err := repo.Save(tx, mission); err != nil {
		return err
	}

	return nil
}
