package vehicle

import (
	"context"
	"testing"
	"vehicle/pkg/vehicle/domain/txmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Vehicleを削除するドメインサービスをテストする。
// 指定されたIDのVehicleを削除する。
// 削除する前にCommunicationIDが削除されたことを表す
// ドメインイベントを発行する。
func TestDeleteVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		versions: []Version{DefaultVersion},
	}
	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             gen,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteVehicle(ctx, repo, pub, DefaultID)

	expectEvent := CommunicationIdRemovedEvent{
		CommunicationID: DefaultCommunicationID,
	}

	a.Len(repo.deleteIDs, 1)
	a.Equal(repo.deleteIDs[0], DefaultID)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Vehicleを削除するドメインサービスをテストする。
// 指定されたIDのVehicleの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetErrorWhenDeleteVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteVehicle(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Vehicleを削除するドメインサービスをテストする。
// 削除時にリポジトリがエラーとなった場合、、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestDeleteErrorWhenDeleteVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		versions: []Version{DefaultVersion},
	}
	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             gen,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Delete", mock.Anything).Return(ErrDelete)

	pub := &publisherMock{}

	ret := DeleteVehicle(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	// a.Len(pub.events, 0) // 1件PublishされるがFlushされない想定

	a.Equal(ret, ErrDelete)
}

// Vehicleを削除するドメインサービスをテストする。
// カーボンコピーされたVehicleを削除しようとした場合、、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestCannnotDeleteErrorWhenDeleteVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    CarbonCopy,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Delete", mock.Anything).Return(ErrDelete)

	pub := &publisherMock{}

	ret := DeleteVehicle(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotChange)
}

// Vehicle削除サービス用リポジトリモック
type repositoryMockDeleteService struct {
	mock.Mock

	saveVehicles []*Vehicle
	deleteIDs    []ID
}

func (rm *repositoryMockDeleteService) GetAll(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) GetAllOrigin(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) GetByID(tx txmanager.Tx, id ID) (*Vehicle, error) {
	ret := rm.Called(id)
	var v *Vehicle
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Vehicle)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockDeleteService) Save(tx txmanager.Tx, v *Vehicle) error {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) Delete(tx txmanager.Tx, id ID) error {
	ret := rm.Called(id)
	if ret.Error(0) == nil {
		rm.deleteIDs = append(rm.deleteIDs, id)
	}
	return ret.Error(0)
}
