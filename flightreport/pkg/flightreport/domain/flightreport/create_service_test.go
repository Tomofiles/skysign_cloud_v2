package flightreport

import (
	"context"
	"flightreport/pkg/flightreport/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightreportを作成するドメインサービスをテストする。
// あらかじめFlightoperationIDを与えられたFlightreportを作成し、保存する。
func TestCreateNewFlightreportService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		OriginalID = DefaultFlightoperationID + "-original"
		NewID      = DefaultFlightoperationID + "-new"
	)

	gen := &generatorMock{
		id:                DefaultID,
		flightoperationID: NewID,
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(nil)

	ret := CreateNewFlightreport(ctx, gen, repo, OriginalID)

	expectFlightreport := Flightreport{
		id:                DefaultID,
		flightoperationID: NewID,
	}

	a.Len(repo.saveFlightreports, 1)
	a.Equal(repo.saveFlightreports[0], &expectFlightreport)

	a.Nil(ret)
}

// Flightreportを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenCreateNewFlightreportService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		OriginalID = DefaultFlightoperationID + "-original"
		NewID      = DefaultFlightoperationID + "-new"
	)

	gen := &generatorMock{
		id:                DefaultID,
		flightoperationID: NewID,
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(ErrSave)

	ret := CreateNewFlightreport(ctx, gen, repo, OriginalID)

	a.Len(repo.saveFlightreports, 0)
	a.Equal(ret, ErrSave)
}

// Flightreport作成サービス用リポジトリモック
type repositoryMockCreateService struct {
	mock.Mock

	saveFlightreports []*Flightreport
}

func (rm *repositoryMockCreateService) GetAll(tx txmanager.Tx) ([]*Flightreport, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetByID(tx txmanager.Tx, id ID) (*Flightreport, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(tx txmanager.Tx, f *Flightreport) error {
	ret := rm.Called(f)
	if ret.Error(0) == nil {
		rm.saveFlightreports = append(rm.saveFlightreports, f)
	}
	return ret.Error(0)
}
