package mission

import (
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
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

	var uploadID UploadID
	if mission.GetNavigation() != nil {
		uploadID = mission.GetNavigation().GetUploadID()
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

	if uploadID != "" {
		pub.Publish(DeletedEvent{
			ID:       mission.GetID(),
			UploadID: uploadID,
		})
	}
	pub.Publish(CreatedEvent{
		ID:      mission.GetID(),
		Mission: mission,
	})
	return string(mission.GetNavigation().GetUploadID()), nil
}
