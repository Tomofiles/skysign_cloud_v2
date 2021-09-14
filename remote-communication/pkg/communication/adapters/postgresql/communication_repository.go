package postgresql

import (
	c "remote-communication/pkg/communication/domain/communication"
	"time"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CommunicationRepository .
type CommunicationRepository struct {
	gen c.Generator
}

// NewCommunicationRepository .
func NewCommunicationRepository(gen c.Generator) *CommunicationRepository {
	return &CommunicationRepository{
		gen: gen,
	}
}

// GetByID .
func (r *CommunicationRepository) GetByID(
	tx txmanager.Tx,
	id c.ID,
) (*c.Communication, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	communicationRecord := Communication{}
	telemetryRecord := Telemetry{}
	var commandRecords []*Command
	var uploadMissionRecords []*UploadMission

	if err := txGorm.Clauses(clause.Locking{Strength: "UPDATE"}).Limit(1).Find(&communicationRecord, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if communicationRecord.ID == "" {
		return nil, c.ErrNotFound
	}
	if err := txGorm.Limit(1).Find(&telemetryRecord, "communication_id = ?", string(id)).Error; err != nil {
		return nil, err
	}
	if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Find(&commandRecords).Error; err != nil {
		return nil, err
	}
	if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Find(&uploadMissionRecords).Error; err != nil {
		return nil, err
	}

	communicationRecord.Telemetry = &telemetryRecord
	communicationRecord.Commands = commandRecords
	communicationRecord.UploadMissions = uploadMissionRecords

	communication := c.AssembleFrom(r.gen, &communicationRecord)

	return communication, nil
}

// Save .
func (r *CommunicationRepository) Save(
	tx txmanager.Tx,
	communication *c.Communication,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	communicationRecord := Communication{}
	telemetryRecord := Telemetry{}
	commandRecords := []*Command{}
	uploadMissionRecords := []*UploadMission{}

	isCreate := false
	if err := txGorm.Clauses(clause.Locking{Strength: "UPDATE"}).Limit(1).Find(&communicationRecord, "id = ?", string(communication.GetID())).Error; err != nil {
		return err
	}

	if communicationRecord.ID == "" {
		isCreate = true
		communicationRecord.ID = string(communication.GetID())
	}

	c.TakeApart(
		communication,
		func(id string) {},
		func(latitude, longitude, altitude, relativeAltitude, speed, x, y, z, w float64, armed bool, flightMode string) {
			telemetryRecord.CommunicationID = communicationRecord.ID
			telemetryRecord.LatitudeDegree = latitude
			telemetryRecord.LongitudeDegree = longitude
			telemetryRecord.AltitudeM = altitude
			telemetryRecord.RelativeAltitudeM = relativeAltitude
			telemetryRecord.SpeedMS = speed
			telemetryRecord.Armed = armed
			telemetryRecord.FlightMode = flightMode
			telemetryRecord.OrientationX = x
			telemetryRecord.OrientationY = y
			telemetryRecord.OrientationZ = z
			telemetryRecord.OrientationW = w
		},
		func(id, cType string, time time.Time) {
			commandRecords = append(
				commandRecords,
				&Command{
					ID:              id,
					Type:            cType,
					CommunicationID: communicationRecord.ID,
					Time:            time,
				},
			)
		},
		func(commandID, missionID string) {
			uploadMissionRecords = append(
				uploadMissionRecords,
				&UploadMission{
					ID:              commandID,
					CommunicationID: communicationRecord.ID,
					MissionID:       missionID,
				},
			)
		},
	)

	if isCreate {
		if err := txGorm.Create(&communicationRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Create(&telemetryRecord).Error; err != nil {
			return err
		}
		if len(commandRecords) != 0 {
			if err := txGorm.Create(&commandRecords).Error; err != nil {
				return err
			}
		}
		if len(uploadMissionRecords) != 0 {
			if err := txGorm.Create(&uploadMissionRecords).Error; err != nil {
				return err
			}
		}
	} else {
		if err := txGorm.Save(&communicationRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Save(&telemetryRecord).Error; err != nil {
			return err
		}
		if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Delete(&Command{}).Error; err != nil {
			return err
		}
		if len(commandRecords) != 0 {
			if err := txGorm.Create(&commandRecords).Error; err != nil {
				return err
			}
		}
		if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Delete(&UploadMission{}).Error; err != nil {
			return err
		}
		if len(uploadMissionRecords) != 0 {
			if err := txGorm.Create(&uploadMissionRecords).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Delete .
func (r *CommunicationRepository) Delete(
	tx txmanager.Tx,
	id c.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	communicationRecord := Communication{}
	telemetryRecord := Telemetry{}
	commandRecord := Command{}
	uploadMissionRecord := UploadMission{}

	if err := txGorm.Limit(1).Find(&communicationRecord, "id = ?", string(id)).Error; err != nil {
		return err
	}
	if communicationRecord.ID == "" {
		return c.ErrNotFound
	}
	if err := txGorm.Delete(&communicationRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Delete(&telemetryRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Delete(&commandRecord).Error; err != nil {
		return err
	}
	if err := txGorm.Where("communication_id = ?", communicationRecord.ID).Delete(&uploadMissionRecord).Error; err != nil {
		return err
	}

	return nil
}
