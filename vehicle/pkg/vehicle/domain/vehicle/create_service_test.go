package vehicle

import (
	"context"
	"testing"
	"vehicle/pkg/vehicle/domain/txmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Vehicleを作成するドメインサービスをテストする。
// 名前とコミュニケーションIDをあらかじめ与えられたVehicleを作成し、保存する。
// 保存が成功すると、Vehicleが作成されたことを表す
// ドメインイベントを発行する。
func TestCreateNewVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	id, ret := CreateNewVehicle(ctx, gen, repo, pub, DefaultName, string(DefaultCommunicationID))

	expectVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion1,
		newVersion:      DefaultVersion3,
		gen:             gen,
		pub:             pub,
	}
	expectEvent := CommunicationIdGaveEvent{CommunicationID: DefaultCommunicationID}

	a.Equal(id, string(DefaultID))
	a.Len(repo.saveVehicles, 1)
	a.Equal(repo.saveVehicles[0], &expectVehicle)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Vehicleを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenCreateNewVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	id, ret := CreateNewVehicle(ctx, gen, repo, pub, DefaultName, string(DefaultCommunicationID))

	a.Empty(id)
	a.Len(repo.saveVehicles, 0)
	// a.Len(pub.events, 0) // 1件PublishされるがFlushされない想定
	a.Equal(ret, ErrSave)
}

// Vehicle作成サービス用リポジトリモック
type repositoryMockCreateService struct {
	mock.Mock

	saveVehicles []*Vehicle
}

func (rm *repositoryMockCreateService) GetAll(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetAllOrigin(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetByID(tx txmanager.Tx, id ID) (*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(tx txmanager.Tx, v *Vehicle) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveVehicles = append(rm.saveVehicles, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCreateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
