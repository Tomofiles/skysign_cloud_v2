package postgresql

import (
	"vehicle/pkg/vehicle/domain/txmanager"
	v "vehicle/pkg/vehicle/domain/vehicle"

	"gorm.io/gorm"
)

// VehicleRepository .
type VehicleRepository struct {
	gen v.Generator
}

// NewVehicleRepository .
func NewVehicleRepository(gen v.Generator) *VehicleRepository {
	return &VehicleRepository{
		gen: gen,
	}
}

// GetAll .
func (r *VehicleRepository) GetAll(
	tx txmanager.Tx,
) ([]*v.Vehicle, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Vehicle

	if err := txGorm.Find(&records).Error; err != nil {
		return nil, err
	}

	var vehicles []*v.Vehicle
	for _, record := range records {
		vehicle := v.AssembleFrom(r.gen, record)
		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

// GetAllOrigin .
func (r *VehicleRepository) GetAllOrigin(
	tx txmanager.Tx,
) ([]*v.Vehicle, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	var records []*Vehicle

	if err := txGorm.Where("is_carbon_copy = false").Find(&records).Error; err != nil {
		return nil, err
	}

	var vehicles []*v.Vehicle
	for _, record := range records {
		vehicle := v.AssembleFrom(r.gen, record)
		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

// GetByID .
func (r *VehicleRepository) GetByID(
	tx txmanager.Tx,
	id v.ID,
) (*v.Vehicle, error) {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Vehicle{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return nil, err
	}

	if record.ID == "" {
		return nil, v.ErrNotFound
	}

	vehicle := v.AssembleFrom(r.gen, &record)

	return vehicle, nil
}

// Save .
func (r *VehicleRepository) Save(
	tx txmanager.Tx,
	vehicle *v.Vehicle,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Vehicle{}

	isCreate := false
	if err := txGorm.Limit(1).Find(&record, "id = ?", string(vehicle.GetID())).Error; err != nil {
		return err
	}

	if record.ID == "" {
		isCreate = true
		record.ID = string(vehicle.GetID())
	}

	v.TakeApart(
		vehicle,
		func(id, name, communicationID, version string, isCarbonCopy bool) {
			record.Name = name
			record.CommunicationID = communicationID
			record.Version = version
			record.IsCarbonCopy = isCarbonCopy
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
func (r *VehicleRepository) Delete(
	tx txmanager.Tx,
	id v.ID,
) error {
	txGorm, ok := tx.(*gorm.DB)
	if !ok {
		panic("developer error")
	}

	record := Vehicle{}

	if err := txGorm.Limit(1).Find(&record, "id = ?", string(id)).Error; err != nil {
		return err
	}

	if record.ID == "" {
		return v.ErrNotFound
	}

	if err := txGorm.Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
