package postgresql

import (
	m "remote-communication/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"gorm.io/gorm"
)

// MissionRepository .
type MissionRepository struct {
}

// NewMissionRepository .
func NewMissionRepository() *MissionRepository {
	return &MissionRepository{}
}

// GetByID .
func (r *MissionRepository) GetByID(
	tx txmanager.Tx,
	id m.ID,
) (*m.Mission, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	missionRecord := Mission{}
	var waypointRecords []*Waypoint

	if err := txGorm.Limit(1).Find(&missionRecord, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if missionRecord.ID == "" {
		return nil, m.ErrNotFound
	}
	if err := txGorm.Where("mission_id = ?", missionRecord.ID).Find(&waypointRecords).Error; err != nil {
		return nil, err
	}

	missionRecord.Waypoints = waypointRecords

	mission := m.AssembleFrom(&missionRecord)

	return mission, nil
}

// Save .
func (r *MissionRepository) Save(
	tx txmanager.Tx,
	mission *m.Mission,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	missionRecord := Mission{}
	waypointRecords := []*Waypoint{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&missionRecord, "id = ?", string(mission.GetID())).Error; err != nil {
		return err
	}

	if missionRecord.ID == "" {
		isCreate = true
		missionRecord.ID = string(mission.GetID())
	}

	m.TakeApart(
		mission,
		func(id string) {},
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			waypointRecords = append(
				waypointRecords,
				&Waypoint{
					MissionID:       missionRecord.ID,
					PointOrder:      pointOrder,
					LatitudeDegree:  latitudeDegree,
					LongitudeDegree: longitudeDegree,
					RelativeHeightM: relativeHeightM,
					SpeedMS:         speedMS,
				},
			)
		},
	)

	if isCreate {
		if err := txGorm.Create(&missionRecord).Error; err != nil {
			return err
		}
		if len(waypointRecords) != 0 {
			if err := txGorm.Create(&waypointRecords).Error; err != nil {
				return err
			}
		}
	} else {
		if err := txGorm.Save(&missionRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Where("mission_id = ?", missionRecord.ID).Delete(&Waypoint{}).Error; err != nil {
			return err
		}
		if len(waypointRecords) != 0 {
			if err := txGorm.Create(&waypointRecords).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Delete .
func (r *MissionRepository) Delete(
	tx txmanager.Tx,
	id m.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	missionRecord := Mission{}
	waypointRecord := Waypoint{}

	if err := txGorm.Limit(1).Find(&missionRecord, "id = ?", string(id)).Error; err != nil {
		return err
	}
	if missionRecord.ID == "" {
		return m.ErrNotFound
	}
	if err := txGorm.Delete(&missionRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("mission_id = ?", missionRecord.ID).Delete(&waypointRecord).Error; err != nil {
		return err
	}

	return nil
}
