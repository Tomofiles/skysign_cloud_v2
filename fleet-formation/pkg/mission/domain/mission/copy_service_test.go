package mission

import (
	"context"
	"fleet-formation/pkg/common/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Missionをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのMissionを、指定されたIDでコピーする。
// コピーが成功するとイベントを発行されることを検証する。
func TestCarbonCopyMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID       = DefaultID + "-new"
		NewUploadID = DefaultUploadID + "-new"
	)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	navigation.uploadID = DefaultUploadID
	navigation.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	navigation.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	navigation.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{
		uploadID: NewUploadID,
	}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	id, ret := CarbonCopyMission(ctx, gen, repo, pub, DefaultID, NewID)

	expectNav := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	expectNav.uploadID = NewUploadID
	expectNav.PushNextWaypoint(11.0, 21.0, 31.0, 41.0)
	expectNav.PushNextWaypoint(12.0, 22.0, 32.0, 42.0)
	expectNav.PushNextWaypoint(13.0, 23.0, 33.0, 43.0)

	expectMission := Mission{
		id:           NewID,
		name:         DefaultName,
		navigation:   expectNav,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}
	expectEvent := CopiedMissionCreatedEvent{
		ID:      NewID,
		Mission: &expectMission,
	}

	a.Equal(id, string(NewUploadID))
	a.Len(repo.saveMissions, 1)
	a.Equal(repo.saveMissions[0], &expectMission)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})

	a.Nil(ret)
}

// Missionをカーボンコピーするドメインサービスをテストする。
// コピー後のIDのMissionのがすでに存在する場合、コピーを行わず
// 正常終了することを検証する。
func TestCopySuccessWhenAlreadyExistsMissionWhenCarbonCopyMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)
	navigation.uploadID = DefaultUploadID

	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(&testMission, nil)

	pub := &publisherMock{}

	id, ret := CarbonCopyMission(ctx, gen, repo, pub, DefaultID, NewID)

	a.Equal(id, string(DefaultUploadID))
	a.Len(repo.saveMissions, 0)
	a.Len(pub.events, 0)
	a.Nil(ret)
}

// Missionをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのMissionの取得がエラーとなった場合、
// コピーが失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenCarbonCopyMissionService(t *testing.T) {
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

	id, ret := CarbonCopyMission(ctx, gen, repo, pub, DefaultID, NewID)

	a.Empty(id)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Missionをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのMissionの取得がエラーとなった場合、
// コピーが失敗し、エラーが返却されることを検証する。
func TestGetError2WhenCarbonCopyMissionService(t *testing.T) {
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

	id, ret := CarbonCopyMission(ctx, gen, repo, pub, DefaultID, NewID)

	a.Empty(id)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Missionをカーボンコピーするドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、、
// 保存が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenCarbonCopyMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewID = DefaultID + "-new"
	)

	navigation := NewNavigation(DefaultTakeoffPointGroundHeightWGS84EllipsoidM)

	testMission := Mission{
		id:           DefaultID,
		name:         DefaultName,
		navigation:   navigation,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMockCopyService{}
	repo.On("GetByID", NewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultID).Return(&testMission, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	id, ret := CarbonCopyMission(ctx, gen, repo, pub, DefaultID, NewID)

	a.Empty(id)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrSave)
}

// Mission削除サービス用リポジトリモック
type repositoryMockCopyService struct {
	mock.Mock

	saveMissions []*Mission
}

func (rm *repositoryMockCopyService) GetAll(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetAllOrigin(tx txmanager.Tx) ([]*Mission, error) {
	panic("implement me")
}
func (rm *repositoryMockCopyService) GetByID(tx txmanager.Tx, id ID) (*Mission, error) {
	ret := rm.Called(id)
	var v *Mission
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Mission)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockCopyService) Save(tx txmanager.Tx, v *Mission) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveMissions = append(rm.saveMissions, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCopyService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
