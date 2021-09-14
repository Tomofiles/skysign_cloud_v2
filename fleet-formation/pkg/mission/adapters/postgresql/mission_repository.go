package postgresql

import (
	m "fleet-formation/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"gorm.io/gorm"
)

// MissionRepository .
type MissionRepository struct {
	gen m.Generator
}

// NewMissionRepository .
func NewMissionRepository(gen m.Generator) *MissionRepository {
	return &MissionRepository{
		gen: gen,
	}
}

// GetAll .
func (r *MissionRepository) GetAll(
	tx txmanager.Tx,
) ([]*m.Mission, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var missionRecords []*Mission

	if err := txGorm.Find(&missionRecords).Error; err != nil {
		return nil, err
	}

	var missions []*m.Mission
	for _, missionRecord := range missionRecords {
		navigationRecord := Navigation{}
		var waypointRecords []*Waypoint

		if err := txGorm.Limit(1).Find(&navigationRecord, "mission_id = ?", string(missionRecord.ID)).Error; err != nil {
			return nil, err
		}
		if err := txGorm.Order("point_order").Where("mission_id = ?", string(missionRecord.ID)).Find(&waypointRecords).Error; err != nil {
			return nil, err
		}

		missionRecord.Navigation = &navigationRecord
		missionRecord.Navigation.Waypoints = waypointRecords

		mission := m.AssembleFrom(r.gen, missionRecord)
		missions = append(missions, mission)
	}

	return missions, nil
}

// GetAllOrigin .
func (r *MissionRepository) GetAllOrigin(
	tx txmanager.Tx,
) ([]*m.Mission, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var missionRecords []*Mission

	if err := txGorm.Where("is_carbon_copy = false").Find(&missionRecords).Error; err != nil {
		return nil, err
	}

	var missions []*m.Mission
	for _, missionRecord := range missionRecords {
		navigationRecord := Navigation{}
		var waypointRecords []*Waypoint

		if err := txGorm.Limit(1).Find(&navigationRecord, "mission_id = ?", string(missionRecord.ID)).Error; err != nil {
			return nil, err
		}
		if err := txGorm.Order("point_order").Where("mission_id = ?", string(missionRecord.ID)).Find(&waypointRecords).Error; err != nil {
			return nil, err
		}

		missionRecord.Navigation = &navigationRecord
		missionRecord.Navigation.Waypoints = waypointRecords

		mission := m.AssembleFrom(r.gen, missionRecord)
		missions = append(missions, mission)
	}

	return missions, nil
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
	navigationRecord := Navigation{}
	var waypointRecords []*Waypoint

	if err := txGorm.Limit(1).Find(&missionRecord, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}

	if missionRecord.ID == "" {
		return nil, m.ErrNotFound
	}

	if err := txGorm.Limit(1).Find(&navigationRecord, "mission_id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if err := txGorm.Order("point_order").Where("mission_id = ?", string(id)).Find(&waypointRecords).Error; err != nil {
		return nil, err
	}

	missionRecord.Navigation = &navigationRecord
	missionRecord.Navigation.Waypoints = waypointRecords

	mission := m.AssembleFrom(r.gen, &missionRecord)

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
	navigationRecord := Navigation{}
	var waypointRecords []*Waypoint

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
		func(id, name, version string, isCarbonCopy bool) {
			missionRecord.Name = name
			missionRecord.IsCarbonCopy = isCarbonCopy
			missionRecord.Version = version
		},
		func(takeoffPointGroundHeightWGS84EllipsoidM float64, uploadID string) {
			navigationRecord.MissionID = string(mission.GetID())
			navigationRecord.TakeoffPointGroundAltitudeM = takeoffPointGroundHeightWGS84EllipsoidM
			navigationRecord.UploadID = uploadID
		},
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			waypointRecords = append(
				waypointRecords,
				&Waypoint{
					MissionID:         string(mission.GetID()),
					PointOrder:        pointOrder,
					LatitudeDegree:    latitudeDegree,
					LongitudeDegree:   longitudeDegree,
					RelativeAltitudeM: relativeHeightM,
					SpeedMS:           speedMS,
				},
			)
		},
	)

	if isCreate {
		if err := txGorm.Create(&missionRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Create(&navigationRecord).Error; err != nil {
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
		if err := txGorm.Where("mission_id = ?", navigationRecord.MissionID).Delete(&Navigation{}).Error; err != nil {
			return err
		}
		if err := txGorm.Create(&navigationRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Where("mission_id = ?", navigationRecord.MissionID).Delete(&Waypoint{}).Error; err != nil {
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
	navigationRecord := Navigation{}
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
	if err := txGorm.Where("mission_id = ?", string(id)).Delete(&navigationRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("mission_id = ?", string(id)).Delete(&waypointRecord).Error; err != nil {
		return err
	}

	return nil
}
