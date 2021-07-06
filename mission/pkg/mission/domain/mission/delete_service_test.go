package mission

import (
	"context"
	"mission/pkg/mission/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Missionを削除するドメインサービスをテストする。
// 指定されたIDのMissionを削除する。
// Navigationがもともと存在しない場合、
// Deletedイベントが発行されないことを検証する。
func TestNoNavigationWhenDeleteMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		versions: []Version{DefaultVersion},
	}
	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   nil,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteMission(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 1)
	a.Equal(repo.deleteIDs[0], DefaultID)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// Missionを削除するドメインサービスをテストする。
// 指定されたIDのMissionを削除する。
// Deletedイベントが発行されることを検証する。
func TestDeleteMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		versions: []Version{DefaultVersion},
	}
	testNav := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	testNav.uploadID = DefaultUploadID
	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   testNav,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteMission(ctx, repo, pub, DefaultID)

	expectEvent := DeletedEvent{
		ID:       DefaultID,
		UploadID: DefaultUploadID,
	}

	a.Len(repo.deleteIDs, 1)
	a.Equal(repo.deleteIDs[0], DefaultID)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})

	a.Nil(ret)
}

// Missionを削除するドメインサービスをテストする。
// 指定されたIDのMissionの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenDeleteMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteMission(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Missionを削除するドメインサービスをテストする。
// 削除時にリポジトリがエラーとなった場合、、
// 削除が失敗し、エラーが返却されることを検証する。
func TestDeleteErrorWhenDeleteMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		versions: []Version{DefaultVersion},
	}
	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM),
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Delete", mock.Anything).Return(ErrDelete)

	pub := &publisherMock{}

	ret := DeleteMission(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrDelete)
}

// Missionを削除するドメインサービスをテストする。
// カーボンコピーされたMissionを削除しようとした場合、、
// 削除が失敗し、エラーが返却されることを検証する。
func TestCannnotDeleteErrorWhenDeleteMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM),
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Delete", mock.Anything).Return(ErrDelete)

	pub := &publisherMock{}

	ret := DeleteMission(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotChange)
}

// Mission削除サービス用リポジトリモック
type repositoryMockDeleteService struct {
	mock.Mock

	saveMissions []*Mission
	deleteIDs    []ID
}

func (rm *repositoryMockDeleteService) GetAll(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) GetAllOrigin(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) GetByID(tx txmanager.Tx, id ID) (*Mission, error) {
	ret := rm.Called(id)
	var v *Mission
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Mission)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockDeleteService) Save(tx txmanager.Tx, v *Mission) error {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) Delete(tx txmanager.Tx, id ID) error {
	ret := rm.Called(id)
	if ret.Error(0) == nil {
		rm.deleteIDs = append(rm.deleteIDs, id)
	}
	return ret.Error(0)
}
