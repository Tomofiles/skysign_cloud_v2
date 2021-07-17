package mission

import (
	"fleet-formation/pkg/mission/domain/event"
	"fleet-formation/pkg/mission/domain/txmanager"
)

// UpdateMission .
func UpdateMission(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	name string,
	navigation *Navigation,
) (string, error) {
	mission, err := repo.GetByID(tx, id)
	if err != nil {
		return "", err
	}

	mission.SetPublisher(pub)

	if err := mission.NameMission(name); err != nil {
		return "", err
	}
	if err := mission.ReplaceNavigationWith(navigation); err != nil {
		return "", err
	}

	if ret := repo.Save(tx, mission); ret != nil {
		return "", ret
	}

	return string(mission.GetNavigation().GetUploadID()), nil
}
