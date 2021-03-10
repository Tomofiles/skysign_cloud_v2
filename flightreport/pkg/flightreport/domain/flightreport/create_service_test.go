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
// 保存が成功すると、Flightreportが作成されたことを表すメインイベントと、
// Flightoperationがコピーされたことを表すドメインイベントを発行する。
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
	pub := &publisherMock{}

	ret := CreateNewFlightreport(ctx, gen, repo, pub, OriginalID)

	expectFlightreport := Flightreport{
		id:                DefaultID,
		flightoperationID: NewID,
	}
	expectEvent1 := CreatedEvent{ID: DefaultID, FlightoperationID: NewID}
	expectEvent2 := FlightoperationCopiedWhenCreatedEvent{OriginalID: OriginalID, NewID: NewID}

	a.Len(repo.saveFlightreports, 1)
	a.Equal(repo.saveFlightreports[0], &expectFlightreport)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent1, expectEvent2})

	a.Nil(ret)
}

// Flightreportを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
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
	pub := &publisherMock{}

	ret := CreateNewFlightreport(ctx, gen, repo, pub, OriginalID)

	a.Len(repo.saveFlightreports, 0)
	a.Len(pub.events, 0)
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
