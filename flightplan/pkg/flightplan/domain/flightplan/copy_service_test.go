package flightplan

import (
	"context"
	"flightplan/pkg/flightplan/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightplanをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFlightplanを、指定されたIDでコピーする。
// コピーが成功すると、Flightplanのコピーが作成されたことを表す
// ドメインイベントを発行する。
func TestCarbonCopyFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testFlightplan := Flightplan{
		id:           DefaultID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := CarbonCopyFlightplan(ctx, gen, repo, pub, DefaultID, NewID)

	expectFlightplan := Flightplan{
		id:           NewID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}
	expectEvent := CopiedEvent{
		OriginalID: DefaultID,
		NewID:      NewID,
	}

	a.Len(repo.saveFlightplans, 1)
	a.Equal(repo.saveFlightplans[0], &expectFlightplan)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Flightplanをカーボンコピーするドメインサービスをテストする。
// コピー後のIDのFlightplanのがすでに存在する場合、コピーを行わず
// 正常終了することを検証する。
func TestCopySuccessWhenAlreadyExistsFlightplanWhenCarbonCopyFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testFlightplan := Flightplan{
		id:           NewID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(&testFlightplan, nil)

	pub := &publisherMock{}

	ret := CarbonCopyFlightplan(ctx, gen, repo, pub, DefaultID, NewID)

	a.Len(repo.saveFlightplans, 0)
	a.Len(pub.events, 0)
	a.Nil(ret)
}

// Flightplanをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFlightplanの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetErrorWhenCarbonCopyFlightplanService(t *testing.T) {
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

	ret := CarbonCopyFlightplan(ctx, gen, repo, pub, DefaultID, NewID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Flightplanをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFlightplanの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetError2WhenCarbonCopyFlightplanService(t *testing.T) {
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

	ret := CarbonCopyFlightplan(ctx, gen, repo, pub, DefaultID, NewID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Flightplanをカーボンコピーするドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、、
// 保存が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenCarbonCopyFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	testFlightplan := Flightplan{
		id:           DefaultID,
		name:         DefaultName,
		description:  DefaultDescription,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	ret := CarbonCopyFlightplan(ctx, gen, repo, pub, DefaultID, NewID)

	a.Len(pub.events, 0)
	a.Equal(ret, ErrSave)
}

// Flightplan削除サービス用リポジトリモック
type repositoryMockCopyService struct {
	mock.Mock

	saveFlightplans []*Flightplan
}

func (rm *repositoryMockCopyService) GetAll(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetAllOrigin(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetByID(tx txmanager.Tx, id ID) (*Flightplan, error) {
	ret := rm.Called(id)
	var f *Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*Flightplan)
	}
	return f, ret.Error(1)
}
func (rm *repositoryMockCopyService) Save(tx txmanager.Tx, f *Flightplan) error {
	ret := rm.Called(f)
	if ret.Error(0) == nil {
		rm.saveFlightplans = append(rm.saveFlightplans, f)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCopyService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
