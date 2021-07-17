package postgresql

import (
	"errors"
	"fleet-formation/pkg/vehicle/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTxManagerDo(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectCommit()

	txm := NewGormTransactionManager(db)

	var retOpe txmanager.Tx
	ret := txm.Do(func(tx txmanager.Tx) error {
		retOpe = tx
		return nil
	})

	_, ok := retOpe.(*gorm.DB)
	a.True(ok)
	a.Nil(ret)
}

func TestTxManagerErrorAndRollbackWhenDo(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectRollback()

	txm := NewGormTransactionManager(db)

	errDo := errors.New("do error")

	var retOpe txmanager.Tx
	ret := txm.Do(func(tx txmanager.Tx) error {
		retOpe = tx
		return errDo
	})

	_, ok := retOpe.(*gorm.DB)
	a.True(ok)
	a.Equal(ret, errDo)
}

func TestTxManagerDoAndEndHook(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectCommit()

	txm := NewGormTransactionManager(db)

	var retOpe txmanager.Tx
	retEndHook := false
	ret := txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			retOpe = tx
			return nil
		},
		func() error {
			retEndHook = true
			return nil
		},
	)

	_, ok := retOpe.(*gorm.DB)
	a.True(ok)
	a.True(retEndHook)
	a.Nil(ret)
}

func TestTxManagerOpeErrorAndRollbackAndNoInvokeEndHookDoAndEndHook(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectRollback()

	txm := NewGormTransactionManager(db)

	errDo := errors.New("do error")

	var retOpe txmanager.Tx
	retEndHook := false
	ret := txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			retOpe = tx
			return errDo
		},
		func() error {
			retEndHook = true
			return nil
		},
	)

	_, ok := retOpe.(*gorm.DB)
	a.True(ok)
	a.False(retEndHook)
	a.Equal(ret, errDo)
}

func TestTxManagerEndHookErrorAndCommitAndInvokeEndHookDoAndEndHook(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectCommit()

	txm := NewGormTransactionManager(db)

	errDo := errors.New("do error")

	var retOpe txmanager.Tx
	retEndHook := false
	ret := txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			retOpe = tx
			return nil
		},
		func() error {
			retEndHook = true
			return errDo
		},
	)

	_, ok := retOpe.(*gorm.DB)
	a.True(ok)
	a.True(retEndHook)
	a.Equal(ret, errDo)
}
