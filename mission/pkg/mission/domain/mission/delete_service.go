package mission

import (
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
)

// DeleteMission .
func DeleteMission(
	tx txmanager.Tx,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	mission, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	mission.SetPublisher(pub)

	if mission.isCarbonCopy {
		return ErrCannotChange
	}

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	if mission.navigation != nil {
		pub.Publish(DeletedEvent{
			ID:       mission.GetID(),
			UploadID: mission.navigation.uploadID,
		})
	}
	return nil
}
