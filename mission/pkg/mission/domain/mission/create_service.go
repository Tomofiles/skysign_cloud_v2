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
) (string, error) {
	mission := NewInstance(gen)

	mission.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	mission.NameMission(name)
	mission.ReplaceNavigationWith(navigation)

	if err := repo.Save(tx, mission); err != nil {
		return "", err
	}

	return string(mission.GetID()), nil
}
