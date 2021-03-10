package postgresql

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/flightreport/domain/txmanager"

	"gorm.io/gorm"
)

// FlightreportRepository .
type FlightreportRepository struct {
	gen frep.Generator
}

// NewFlightreportRepository .
func NewFlightreportRepository(gen frep.Generator) *FlightreportRepository {
	return &FlightreportRepository{
		gen: gen,
	}
}

// GetAll .
func (r *FlightreportRepository) GetAll(
	tx txmanager.Tx,
) ([]*frep.Flightreport, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Flightreport

	if err := txGorm.Find(&records).Error; err != nil {
		return nil, err
	}

	var flightreports []*frep.Flightreport
	for _, record := range records {
		flightreport := frep.AssembleFrom(r.gen, record)
		flightreports = append(flightreports, flightreport)
	}

	return flightreports, nil
}

// GetByID .
func (r *FlightreportRepository) GetByID(
	tx txmanager.Tx,
	id frep.ID,
) (*frep.Flightreport, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightreport{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}

	if record.ID == "" {
		return nil, frep.ErrNotFound
	}

	flightreport := frep.AssembleFrom(r.gen, &record)

	return flightreport, nil
}

// Save .
func (r *FlightreportRepository) Save(
	tx txmanager.Tx,
	flightreport *frep.Flightreport,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightreport{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(flightreport.GetID())).Error; err != nil {
		return err
	}

	if record.ID != "" {
		return nil
	}

	frep.TakeApart(
		flightreport,
		func(id, flightoperationID string) {
			record.ID = id
			record.FlightoperationID = flightoperationID
		},
	)

	if err := txGorm.Create(&record).Error; err != nil {
		return err
	}

	return nil
}
