package postgresql

import (
	"flight-operation/pkg/common/domain/txmanager"
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"

	"gorm.io/gorm"
)

// FlightoperationRepository .
type FlightoperationRepository struct {
	gen fope.Generator
}

// NewFlightoperationRepository .
func NewFlightoperationRepository(gen fope.Generator) *FlightoperationRepository {
	return &FlightoperationRepository{
		gen: gen,
	}
}

// GetAll .
func (r *FlightoperationRepository) GetAll(
	tx txmanager.Tx,
) ([]*fope.Flightoperation, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Flightoperation

	if err := txGorm.Find(&records).Error; err != nil {
		return nil, err
	}

	var flightoperations []*fope.Flightoperation
	for _, record := range records {
		flightoperation := fope.AssembleFrom(r.gen, record)
		flightoperations = append(flightoperations, flightoperation)
	}

	return flightoperations, nil
}

// GetAllOperating .
func (r *FlightoperationRepository) GetAllOperating(
	tx txmanager.Tx,
) ([]*fope.Flightoperation, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Flightoperation

	if err := txGorm.Where("is_completed = false").Find(&records).Error; err != nil {
		return nil, err
	}

	var flightoperations []*fope.Flightoperation
	for _, record := range records {
		flightoperation := fope.AssembleFrom(r.gen, record)
		flightoperations = append(flightoperations, flightoperation)
	}

	return flightoperations, nil
}

// GetByID .
func (r *FlightoperationRepository) GetByID(
	tx txmanager.Tx,
	id fope.ID,
) (*fope.Flightoperation, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightoperation{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}

	if record.ID == "" {
		return nil, fope.ErrNotFound
	}

	flightoperation := fope.AssembleFrom(r.gen, &record)

	return flightoperation, nil
}

// Save .
func (r *FlightoperationRepository) Save(
	tx txmanager.Tx,
	flightoperation *fope.Flightoperation,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightoperation{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&record, "id = ?", string(flightoperation.GetID())).Error; err != nil {
		return err
	}

	if record.ID == "" {
		isCreate = true
		record.ID = string(flightoperation.GetID())
	}

	fope.TakeApart(
		flightoperation,
		func(id, name, description, fleetID, version string, isCompleted bool) {
			record.Name = name
			record.Description = description
			record.FleetID = fleetID
			record.IsCompleted = isCompleted
			record.Version = version
		},
	)

	if isCreate {
		if err := txGorm.Create(&record).Error; err != nil {
			return err
		}
	} else {
		if err := txGorm.Save(&record).Error; err != nil {
			return err
		}
	}
	return nil
}
