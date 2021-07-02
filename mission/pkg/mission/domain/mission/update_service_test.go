package mission

import (
	"context"
	"mission/pkg/mission/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Missionを更新するドメインサービスをテストする。
// 名前とNavigationを変更し、保存する。
func TestUpdateMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		NewName         = DefaultName + "-new"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion2, DefaultVersion3},
	}
	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   nil,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
		gen:          gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	ret := UpdateMission(ctx, gen, repo, pub, DefaultID, NewName, navigation)

	expectMission := Mission{
		id:           DefaultID,
		name:         NewName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion3,
		gen:          gen,
		pub:          pub,
	}

	a.Len(repo.saveMissions, 1)
	a.Equal(repo.saveMissions[0], &expectMission)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// Missionを更新するドメインサービスをテストする。
// 指定されたIDのMissionの取得がエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenUpdateMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewName = DefaultName + "-new"
	)

	gen := &generatorMock{}

	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	ret := UpdateMission(ctx, gen, repo, pub, DefaultID, NewName, navigation)

	a.Len(repo.saveMissions, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Missionを更新するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenUpdateMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		NewName         = DefaultName + "-new"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3},
	}
	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   nil,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion1,
		gen:          gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	ret := UpdateMission(ctx, gen, repo, pub, DefaultID, NewName, navigation)

	a.Len(repo.saveMissions, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}

// Missionを更新するドメインサービスをテストする。
// カーボンコピーされたMissionを更新しようとした場合、、
// 更新が失敗し、エラーが返却されることを検証する。
func TestCannnotUpdateErrorWhenUpdateMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewName = DefaultName + "-new"
	)

	gen := &generatorMock{}

	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   nil,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	ret := UpdateMission(ctx, gen, repo, pub, DefaultID, NewName, navigation)

	a.Len(repo.saveMissions, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotChange)
}

// Mission更新サービス用リポジトリモック
type repositoryMockUpdateService struct {
	mock.Mock

	saveMissions []*Mission
}

func (rm *repositoryMockUpdateService) GetAll(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockUpdateService) GetAllOrigin(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockUpdateService) GetByID(tx txmanager.Tx, id ID) (*Mission, error) {
	ret := rm.Called(id)
	var v *Mission
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Mission)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockUpdateService) Save(tx txmanager.Tx, v *Mission) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveMissions = append(rm.saveMissions, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockUpdateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
