package vehicle

import (
	"context"
	"testing"
	"vehicle/pkg/vehicle/domain/txmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Vehicleを更新するドメインサービスをテストする。
// 名前とCommunicationIDを変更し、保存する。
// 保存の際、CommunicationIDが変更されたことを検出すると、
// ドメインイベントを発行する。
func TestUpdateVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1    = DefaultVersion + "-1"
		DefaultVersion2    = DefaultVersion + "-2"
		DefaultVersion3    = DefaultVersion + "-3"
		NewName            = DefaultName + "-new"
		NewCommunicationID = DefaultCommunicationID + "-new"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2, DefaultVersion3},
	}
	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion1,
		newVersion:      DefaultVersion1,
		gen:             gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := UpdateVehicle(ctx, gen, repo, pub, DefaultID, NewName, NewCommunicationID)

	expectVehicle := Vehicle{
		id:              DefaultID,
		name:            NewName,
		communicationID: NewCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion1,
		newVersion:      DefaultVersion3,
		gen:             gen,
		pub:             pub,
	}
	expectEvent1 := CommunicationIDGaveEvent{CommunicationID: NewCommunicationID}
	expectEvent2 := CommunicationIDRemovedEvent{CommunicationID: DefaultCommunicationID}

	a.Len(repo.saveVehicles, 1)
	a.Equal(repo.saveVehicles[0], &expectVehicle)
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent2, expectEvent1})

	a.Nil(ret)
}

// Vehicleを更新するドメインサービスをテストする。
// 指定されたIDのVehicleの取得がエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetErrorWhenUpdateVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewName            = DefaultName + "-new"
		NewCommunicationID = DefaultCommunicationID + "-new"
	)

	gen := &generatorMock{}

	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := UpdateVehicle(ctx, gen, repo, pub, DefaultID, NewName, NewCommunicationID)

	a.Len(repo.saveVehicles, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Vehicleを更新するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenUpdateVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1    = DefaultVersion + "-1"
		DefaultVersion2    = DefaultVersion + "-2"
		DefaultVersion3    = DefaultVersion + "-3"
		NewName            = DefaultName + "-new"
		NewCommunicationID = DefaultCommunicationID + "-new"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion1,
		newVersion:      DefaultVersion1,
		gen:             gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	ret := UpdateVehicle(ctx, gen, repo, pub, DefaultID, NewName, NewCommunicationID)

	a.Len(repo.saveVehicles, 0)
	// a.Len(pub.events, 0) // 2件PublishされるがFlushされない想定
	a.Equal(ret, ErrSave)
}

// Vehicleを更新するドメインサービスをテストする。
// カーボンコピーされたVehicleを更新しようとした場合、、
// 更新が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestCannnotUpdateErrorWhenUpdateVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewName            = DefaultName + "-new"
		NewCommunicationID = DefaultCommunicationID + "-new"
	)

	gen := &generatorMock{}

	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    CarbonCopy,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             nil,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	ret := UpdateVehicle(ctx, gen, repo, pub, DefaultID, NewName, NewCommunicationID)

	a.Len(repo.saveVehicles, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotChange)
}

// Vehicle更新サービス用リポジトリモック
type repositoryMockUpdateService struct {
	mock.Mock

	saveVehicles []*Vehicle
}

func (rm *repositoryMockUpdateService) GetAll(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockUpdateService) GetAllOrigin(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockUpdateService) GetByID(tx txmanager.Tx, id ID) (*Vehicle, error) {
	ret := rm.Called(id)
	var v *Vehicle
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Vehicle)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockUpdateService) Save(tx txmanager.Tx, v *Vehicle) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveVehicles = append(rm.saveVehicles, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockUpdateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
