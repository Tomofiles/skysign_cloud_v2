package flightoperation

import (
	"context"
	"flight-operation/pkg/common/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightoperationを作成するドメインサービスをテストする。
// あらかじめFleetIDを与えられたFlightoperationを作成し、保存する。
// 保存が成功すると、Fleetがコピーされたことを表すドメインイベントを発行する。
func TestCreateNewFlightoperationService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		OriginalID = DefaultFleetID + "-original"
		NewID      = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		fleetID: NewID,
		version: DefaultVersion,
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := CreateNewFlightoperation(ctx, gen, repo, pub, DefaultName, DefaultDescription, OriginalID)

	expectFlightoperation := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     NewID,
		isCompleted: Operating,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         gen,
		pub:         pub,
	}
	expectEvent := FleetCopiedEvent{OriginalID: OriginalID, NewID: NewID}

	a.Len(repo.saveFlightoperations, 1)
	a.Equal(repo.saveFlightoperations[0], expectFlightoperation)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})

	a.Nil(ret)
}

// Flightoperationを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenCreateNewFlightoperationService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		OriginalID = DefaultFleetID + "-original"
		NewID      = DefaultFleetID + "-new"
	)

	gen := &generatorMock{
		id:      DefaultID,
		fleetID: NewID,
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	ret := CreateNewFlightoperation(ctx, gen, repo, pub, DefaultName, DefaultDescription, OriginalID)

	a.Len(repo.saveFlightoperations, 0)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrSave)
}

// Flightplan作成サービス用リポジトリモック
type repositoryMockCreateService struct {
	mock.Mock

	saveFlightoperations []*Flightoperation
}

func (rm *repositoryMockCreateService) GetAll(tx txmanager.Tx) ([]*Flightoperation, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetAllOperating(tx txmanager.Tx) ([]*Flightoperation, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetByID(tx txmanager.Tx, id ID) (*Flightoperation, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(tx txmanager.Tx, f *Flightoperation) error {
	ret := rm.Called(f)
	if ret.Error(0) == nil {
		rm.saveFlightoperations = append(rm.saveFlightoperations, f)
	}
	return ret.Error(0)
}
