package flightplan

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// CarbonCopyFlightplan .
func CarbonCopyFlightplan(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	originalID ID,
	newID ID,
) error {
	original, err := repo.GetByID(tx, originalID)
	if err != nil {
		return err
	}

	flightplan := Copy(gen, newID, original)

	if err := repo.Save(tx, flightplan); err != nil {
		return err
	}

	pub.Publish(CopiedEvent{
		OriginalID: originalID,
		NewID:      newID,
	})
	return nil
}
