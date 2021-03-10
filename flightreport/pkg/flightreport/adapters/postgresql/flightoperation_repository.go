package postgresql

import (
	fope "flightreport/pkg/flightreport/domain/flightoperation"
	"flightreport/pkg/flightreport/domain/txmanager"

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

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(flightoperation.GetID())).Error; err != nil {
		return err
	}

	if record.ID != "" {
		return nil
	}

	fope.TakeApart(
		flightoperation,
		func(id, flightplanID string) {
			record.ID = id
			record.FlightplanID = flightplanID
		},
	)

	if err := txGorm.Create(&record).Error; err != nil {
		return err
	}

	return nil
}
