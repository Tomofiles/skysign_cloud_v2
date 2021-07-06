package mission

import (
	"mission/pkg/mission/domain/event"
	"mission/pkg/mission/domain/txmanager"
)

// CreateNewMission .
func CreateNewMission(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	navigation *Navigation,
) (string, string, error) {
	mission := NewInstance(gen)

	mission.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	mission.NameMission(name)
	mission.ReplaceNavigationWith(navigation)

	if err := repo.Save(tx, mission); err != nil {
		return "", "", err
	}

	pub.Publish(CreatedEvent{
		ID:      mission.GetID(),
		Mission: mission,
	})
	return string(mission.GetID()), string(mission.GetNavigation().GetUploadID()), nil
}
