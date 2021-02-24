package postgresql

import (
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"

	"gorm.io/gorm"
)

// FlightplanRepository .
type FlightplanRepository struct {
	gen fpl.Generator
}

// NewFlightplanRepository .
func NewFlightplanRepository(gen fpl.Generator) *FlightplanRepository {
	return &FlightplanRepository{
		gen: gen,
	}
}

// GetAll .
func (r *FlightplanRepository) GetAll(
	tx txmanager.Tx,
) ([]*fpl.Flightplan, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Flightplan

	if err := txGorm.Find(&records).Error; err != nil {
		return nil, err
	}

	var flightplans []*fpl.Flightplan
	for _, record := range records {
		flightplan := fpl.AssembleFrom(r.gen, record)
		flightplans = append(flightplans, flightplan)
	}

	return flightplans, nil
}

// GetByID .
func (r *FlightplanRepository) GetByID(
	tx txmanager.Tx,
	id fpl.ID,
) (*fpl.Flightplan, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightplan{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}

	if record.ID == "" {
		return nil, fpl.ErrNotFound
	}

	flightplan := fpl.AssembleFrom(r.gen, &record)

	return flightplan, nil
}

// Save .
func (r *FlightplanRepository) Save(
	tx txmanager.Tx,
	flightplan *fpl.Flightplan,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightplan{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&record, "id = ?", string(flightplan.GetID())).Error; err != nil {
		return err
	}

	if record.ID == "" {
		isCreate = true
		record.ID = string(flightplan.GetID())
	}

	fpl.TakeApart(
		flightplan,
		func(id, name, description, version string) {
			record.Name = name
			record.Description = description
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

// Delete .
func (r *FlightplanRepository) Delete(
	tx txmanager.Tx,
	id fpl.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Flightplan{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return err
	}

	if record.ID == "" {
		return fpl.ErrNotFound
	}

	if err := txGorm.Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
