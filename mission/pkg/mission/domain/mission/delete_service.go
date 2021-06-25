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

	// 取得したmissionに対しては特にやることなし

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	return nil
}
