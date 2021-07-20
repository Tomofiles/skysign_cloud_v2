package postgresql

import (
	"collection-analysis/pkg/action/domain/action"
	"collection-analysis/pkg/common/domain/txmanager"

	"gorm.io/gorm"
)

// ActionRepository .
type ActionRepository struct {
}

// NewActionRepository .
func NewActionRepository() *ActionRepository {
	return &ActionRepository{}
}

// GetByID .
func (r *ActionRepository) GetByID(
	tx txmanager.Tx,
	id action.ID,
) (*action.Action, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	actionRecord := Action{}
	var trajectoryPointRecords []*TrajectoryPoint

	if err := txGorm.Limit(1).Find(&actionRecord, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if actionRecord.ID == "" {
		return nil, action.ErrNotFound
	}
	if err := txGorm.Where("action_id = ?", string(id)).Order("point_number").Find(&trajectoryPointRecords).Error; err != nil {
		return nil, err
	}

	actionRecord.TrajectoryPoints = trajectoryPointRecords

	action := action.AssembleFrom(&actionRecord)

	return action, nil
}

// GetAllActiveByFleetID .
func (r *ActionRepository) GetAllActiveByFleetID(
	tx txmanager.Tx,
	fleetID action.FleetID,
) ([]*action.Action, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var actions []*action.Action

	var actionRecords []*Action

	if err := txGorm.Where("fleet_id = ? AND is_completed = false", string(fleetID)).Find(&actionRecords).Error; err != nil {
		return nil, err
	}

	for _, actionRecord := range actionRecords {
		var trajectoryPointRecords []*TrajectoryPoint

		if err := txGorm.Where("action_id = ?", string(actionRecord.ID)).Order("point_number").Find(&trajectoryPointRecords).Error; err != nil {
			return nil, err
		}

		actionRecord.TrajectoryPoints = trajectoryPointRecords

		action := action.AssembleFrom(actionRecord)

		actions = append(actions, action)
	}
	return actions, nil
}

// GetActiveByCommunicationID .
func (r *ActionRepository) GetActiveByCommunicationID(
	tx txmanager.Tx,
	communicationID action.CommunicationID,
) (*action.Action, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	actionRecord := Action{}
	var trajectoryPointRecords []*TrajectoryPoint

	if err := txGorm.Limit(1).Find(&actionRecord, "communication_id = ? AND is_completed = false", string(communicationID)).Error; err != nil {
		return nil, err
	}
	if actionRecord.ID == "" {
		return nil, action.ErrNotFound
	}
	if err := txGorm.Where("action_id = ?", string(actionRecord.ID)).Order("point_number").Find(&trajectoryPointRecords).Error; err != nil {
		return nil, err
	}

	actionRecord.TrajectoryPoints = trajectoryPointRecords

	action := action.AssembleFrom(&actionRecord)

	return action, nil
}

// Save .
func (r *ActionRepository) Save(
	tx txmanager.Tx,
	aAction *action.Action,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	actionRecord := Action{}
	trajectoryPointRecords := []*TrajectoryPoint{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&actionRecord, "id = ?", string(aAction.GetID())).Error; err != nil {
		return err
	}

	var trajectoryPointCount int64
	if err := txGorm.Model(&TrajectoryPoint{}).Where("action_id = ?", string(aAction.GetID())).Count(&trajectoryPointCount).Error; err != nil {
		return err
	}

	if actionRecord.ID == "" {
		isCreate = true
		actionRecord.ID = string(aAction.GetID())
	}

	action.TakeApart(
		aAction,
		func(comp action.ActionComponent) {
			actionRecord.CommunicationID = comp.GetCommunicationID()
			actionRecord.FleetID = comp.GetFleetID()
			actionRecord.IsCompleted = comp.GetIsCompleted()
		},
		func(comp action.TrajectoryPointComponent) {
			if trajectoryPointCount >= int64(comp.GetPointNumber()) {
				return
			}
			trajectoryPointRecords = append(
				trajectoryPointRecords,
				&TrajectoryPoint{
					ActionID:         actionRecord.ID,
					PointNumber:      comp.GetPointNumber(),
					Latitude:         comp.GetLatitude(),
					Longitude:        comp.GetLongitude(),
					Altitude:         comp.GetAltitude(),
					RelativeAltitude: comp.GetRelativeAltitude(),
					Speed:            comp.GetSpeed(),
					Armed:            comp.GetArmed(),
					FlightMode:       comp.GetFlightMode(),
					OrientationX:     comp.GetOrientationX(),
					OrientationY:     comp.GetOrientationY(),
					OrientationZ:     comp.GetOrientationZ(),
					OrientationW:     comp.GetOrientationW(),
				},
			)
		},
	)

	if isCreate {
		if err := txGorm.Create(&actionRecord).Error; err != nil {
			return err
		}
		if len(trajectoryPointRecords) != 0 {
			if err := txGorm.Create(&trajectoryPointRecords).Error; err != nil {
				return err
			}
		}
	} else {
		if err := txGorm.Save(&actionRecord).Error; err != nil {
			return err
		}
		if len(trajectoryPointRecords) != 0 {
			if err := txGorm.Create(&trajectoryPointRecords).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
