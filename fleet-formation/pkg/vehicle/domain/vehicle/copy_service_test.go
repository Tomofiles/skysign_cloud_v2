package vehicle

import (
	"context"
	"fleet-formation/pkg/vehicle/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Vehicleをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのVehicleを、指定されたIDでコピーする。
// コピーが成功すると、Vehicleのコピーが作成されたことを表す
// ドメインイベントを発行する。
func TestCarbonCopyVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := CarbonCopyVehicle(ctx, gen, repo, pub, DefaultID, NewID, DefaultFleetID)

	expectVehicle := Vehicle{
		id:              NewID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    CarbonCopy,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             gen,
	}
	expectEvent := CopiedVehicleCreatedEvent{
		ID:              NewID,
		CommunicationID: DefaultCommunicationID,
		FleetID:         DefaultFleetID,
	}

	a.Len(repo.saveVehicles, 1)
	a.Equal(repo.saveVehicles[0], &expectVehicle)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Vehicleをカーボンコピーするドメインサービスをテストする。
// コピー後のIDのVehicleのがすでに存在する場合、コピーを行わず
// 正常終了することを検証する。
func TestCopySuccessWhenAlreadyExistsVehicleWhenCarbonCopyVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(&testVehicle, nil)

	pub := &publisherMock{}

	ret := CarbonCopyVehicle(ctx, gen, repo, pub, DefaultID, NewID, DefaultFleetID)

	a.Len(repo.saveVehicles, 0)
	a.Len(pub.events, 0)
	a.Nil(ret)
}

// Vehicleをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのVehicleの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetErrorWhenCarbonCopyVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrGet)
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := CarbonCopyVehicle(ctx, gen, repo, pub, DefaultID, NewID, DefaultFleetID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Vehicleをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのVehicleの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetError2WhenCarbonCopyVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := CarbonCopyVehicle(ctx, gen, repo, pub, DefaultID, NewID, DefaultFleetID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Vehicleをカーボンコピーするドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、、
// 保存が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenCarbonCopyVehicleService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testVehicle := Vehicle{
		id:              DefaultID,
		name:            DefaultName,
		communicationID: DefaultCommunicationID,
		isCarbonCopy:    Original,
		version:         DefaultVersion,
		newVersion:      DefaultVersion,
		gen:             nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testVehicle, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	ret := CarbonCopyVehicle(ctx, gen, repo, pub, DefaultID, NewID, DefaultFleetID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrSave)
}

// Vehicle削除サービス用リポジトリモック
type repositoryMockCopyService struct {
	mock.Mock

	saveVehicles []*Vehicle
}

func (rm *repositoryMockCopyService) GetAll(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetAllOrigin(tx txmanager.Tx) ([]*Vehicle, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetByID(tx txmanager.Tx, id ID) (*Vehicle, error) {
	ret := rm.Called(id)
	var v *Vehicle
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Vehicle)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockCopyService) Save(tx txmanager.Tx, v *Vehicle) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveVehicles = append(rm.saveVehicles, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCopyService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
