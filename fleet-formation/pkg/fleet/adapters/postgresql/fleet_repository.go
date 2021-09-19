package postgresql

import (
	fl "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"gorm.io/gorm"
)

type fleetRepository struct {
	gen fl.Generator
}

// NewFleetRepository .
func NewFleetRepository(gen fl.Generator) fl.Repository {
	return &fleetRepository{
		gen: gen,
	}
}

// GetByID .
func (r *fleetRepository) GetByID(
	tx txmanager.Tx,
	id fl.ID,
) (*fl.Fleet, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	fleetRecord := Fleet{}
	var assignmentRecords []*Assignment
	var eventRecords []*Event

	if err := txGorm.Limit(1).Find(&fleetRecord, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if fleetRecord.ID == "" {
		return nil, fl.ErrNotFound
	}
	if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Find(&assignmentRecords).Error; err != nil {
		return nil, err
	}
	if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Find(&eventRecords).Error; err != nil {
		return nil, err
	}

	fleetRecord.Assignments = assignmentRecords
	fleetRecord.Events = eventRecords

	fleet := fl.AssembleFrom(r.gen, &fleetRecord)

	return fleet, nil
}

// Save .
func (r *fleetRepository) Save(
	tx txmanager.Tx,
	fleet *fl.Fleet,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	fleetRecord := Fleet{}
	assignmentRecords := []*Assignment{}
	eventRecords := []*Event{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&fleetRecord, "id = ?", string(fleet.GetID())).Error; err != nil {
		return err
	}

	if fleetRecord.ID == "" {
		isCreate = true
		fleetRecord.ID = string(fleet.GetID())
	}

	fl.TakeApart(
		fleet,
		func(id, version string, isCarbonCopy bool) {
			fleetRecord.ID = id
			fleetRecord.Version = version
			fleetRecord.IsCarbonCopy = isCarbonCopy
		},
		func(id, fleetID, vehicleID string) {
			assignmentRecords = append(
				assignmentRecords,
				&Assignment{
					ID:        id,
					FleetID:   fleetID,
					VehicleID: vehicleID,
				},
			)
		},
		func(id, fleetID, assignmentID, missionID string) {
			eventRecords = append(
				eventRecords,
				&Event{
					ID:           id,
					FleetID:      fleetID,
					AssignmentID: assignmentID,
					MissionID:    missionID,
				},
			)
		},
	)

	if isCreate {
		if err := txGorm.Create(&fleetRecord).Error; err != nil {
			return err
		}
		if len(assignmentRecords) != 0 {
			if err := txGorm.Create(&assignmentRecords).Error; err != nil {
				return err
			}
		}
		if len(eventRecords) != 0 {
			if err := txGorm.Create(&eventRecords).Error; err != nil {
				return err
			}
		}
	} else {
		if err := txGorm.Save(&fleetRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Delete(&Assignment{}).Error; err != nil {
			return err
		}
		if len(assignmentRecords) != 0 {
			if err := txGorm.Create(&assignmentRecords).Error; err != nil {
				return err
			}
		}
		if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Delete(&Event{}).Error; err != nil {
			return err
		}
		if len(eventRecords) != 0 {
			if err := txGorm.Create(&eventRecords).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Delete .
func (r *fleetRepository) Delete(
	tx txmanager.Tx,
	id fl.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	fleetRecord := Fleet{}
	assignmentRecord := Assignment{}
	eventRecord := Event{}

	if err := txGorm.Limit(1).Find(&fleetRecord, "id = ?", string(id)).Error; err != nil {
		return err
	}
	if fleetRecord.ID == "" {
		return fl.ErrNotFound
	}
	if err := txGorm.Delete(&fleetRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Delete(&assignmentRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("fleet_id = ?", fleetRecord.ID).Delete(&eventRecord).Error; err != nil {
		return err
	}

	return nil
}
