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
) (string, error) {
	old, err := repo.GetByID(tx, newID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return "", err
	} else if err == nil {
		return string(old.GetNavigation().GetUploadID()), nil
	}

	original, err := repo.GetByID(tx, originalID)
	if err != nil {
		return "", err
	}

	mission := Copy(gen, newID, original)

	if err := repo.Save(tx, mission); err != nil {
		return "", err
	}

	pub.Publish(CopiedMissionCreatedEvent{
		ID:      mission.GetID(),
		Mission: mission,
	})
	return string(mission.GetNavigation().GetUploadID()), nil
}
