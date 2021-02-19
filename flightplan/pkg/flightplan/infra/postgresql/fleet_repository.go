package postgresql

import (
	"errors"
	fl "flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/txmanager"

	"gorm.io/gorm"
)

// FleetRepository .
type FleetRepository struct {
	gen fl.Generator
}

// NewFleetRepository .
func NewFleetRepository(gen fl.Generator) *FleetRepository {
	return &FleetRepository{
		gen: gen,
	}
}

// GetByFlightplanID .
func (r *FleetRepository) GetByFlightplanID(
	tx txmanager.Tx,
	flightplanID flightplan.ID,
) (*fl.Fleet, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	fleetRecord := Fleet{}
	var assignmentRecords []Assignment
	var eventRecords []Event

	if err := txGorm.First(&fleetRecord, "flightplan_id = ?", string(flightplanID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
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
func (r *FleetRepository) Save(
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
	if err := txGorm.First(&fleetRecord, "flightplan_id = ?", string(fleet.GetFlightplanID())).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		isCreate = true
		fleetRecord.ID = string(fleet.GetID())
	}

	fl.TakeApart(
		fleet,
		func(id, flightplanID, version string) {
			fleetRecord.ID = id
			fleetRecord.FlightplanID = flightplanID
			fleetRecord.Version = version
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

// DeleteByFlightplanID .
func (r *FleetRepository) DeleteByFlightplanID(
	tx txmanager.Tx,
	flightplanID flightplan.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	fleetRecord := Fleet{}
	assignmentRecord := Assignment{}
	eventRecord := Event{}

	if err := txGorm.First(&fleetRecord, "flightplan_id = ?", string(flightplanID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
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
