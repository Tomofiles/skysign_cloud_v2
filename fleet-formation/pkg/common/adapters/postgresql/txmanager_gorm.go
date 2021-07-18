package postgresql

import (
	"fleet-formation/pkg/common/domain/txmanager"

	"gorm.io/gorm"
)

// GormTransactionManager .
type GormTransactionManager struct {
	db *gorm.DB
}

// NewGormTransactionManager .
func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{
		db: db,
	}
}

// Do .
func (txm *GormTransactionManager) Do(operation func(tx txmanager.Tx) error) error {
	return txm.db.Transaction(
		func(tx *gorm.DB) error {
			return operation(tx)
		},
	)
}

// DoAndEndHook .
func (txm *GormTransactionManager) DoAndEndHook(
	operation func(tx txmanager.Tx) error,
	endHook func() error,
) error {
	if err := txm.db.Transaction(
		func(tx *gorm.DB) error {
			return operation(tx)
		},
	); err != nil {
		return err
	}
	return endHook()
}
