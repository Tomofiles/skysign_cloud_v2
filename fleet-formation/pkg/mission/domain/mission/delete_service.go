package mission

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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

	return nil
}
