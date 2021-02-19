package postgresql

import (
	"flightplan/pkg/flightplan/domain/txmanager"

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
func (txm *GormTransactionManager) Do(f func(tx txmanager.Tx) error) error {
	return txm.db.Transaction(func(tx *gorm.DB) error {
		return f(tx)
	})
}
