package mission

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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

	return string(mission.GetID()), string(mission.GetNavigation().GetUploadID()), nil
}
