package mission

import (
	"context"
	"mission/pkg/mission/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Missionを作成するドメインサービスをテストする。
// 名前とNavigationをあらかじめ与えられたMissionを作成し、保存する。
func TestCreateNewMissionService(t *testing.T) {
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

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	id, ret := CreateNewMission(ctx, gen, repo, pub, DefaultName, navigation)

	expectMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion1,
		newVersion:   DefaultVersion3,
		gen:          gen,
		pub:          pub,
	}

	a.Equal(id, string(DefaultID))
	a.Len(repo.saveMissions, 1)
	a.Equal(repo.saveMissions[0], &expectMission)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// Missionを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenCreateNewMissionService(t *testing.T) {
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

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	id, ret := CreateNewMission(ctx, gen, repo, pub, DefaultName, navigation)

	a.Empty(id)
	a.Len(repo.saveMissions, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}

// Mission作成サービス用リポジトリモック
type repositoryMockCreateService struct {
	mock.Mock

	saveMissions []*Mission
}

func (rm *repositoryMockCreateService) GetAll(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetAllOrigin(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetByID(tx txmanager.Tx, id ID) (*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(tx txmanager.Tx, v *Mission) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveMissions = append(rm.saveMissions, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCreateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
