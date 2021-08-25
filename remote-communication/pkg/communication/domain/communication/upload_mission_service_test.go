package communication

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// CommunicationにUploadMissionを追加するドメインサービスをテストする。
// 指定されたIDのCommunicationにUploadMissionが追加されることを検証する。
func TestPushUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	testCommunication := NewInstance(gen, DefaultID)

	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	id, ret := PushUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultMissionID)

	expectCommunication := NewInstance(gen, DefaultID)
	expectCommunication.commands = append(
		expectCommunication.commands,
		&Command{
			id:    DefaultCommandID,
			cType: CommandTypeUPLOAD,
			time:  DefaultTime,
		},
	)
	expectCommunication.uploadMissions = append(
		expectCommunication.uploadMissions,
		&UploadMission{
			commandID: DefaultCommandID,
			missionID: DefaultMissionID,
		},
	)

	a.Equal(id, DefaultCommandID)
	a.Len(repo.saveCommunications, 1)
	a.Equal(repo.saveCommunications[0], expectCommunication)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// CommunicationにUploadMissionを追加するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// UploadMissionの追加が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPushUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	id, ret := PushUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultMissionID)

	a.Empty(id)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// CommunicationにUploadMissionを追加するドメインサービスをテストする。
// Communicationの保存時にリポジトリがエラーとなった場合、
// UploadMissionの追加が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenPushUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	testCommunication := NewInstance(gen, DefaultID)

	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	id, ret := PushUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultMissionID)

	a.Empty(id)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}

// CommunicationからUploadMissionを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationからUploadMissionを取得されることを検証する。
func TestPullUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	testCommunication := NewInstance(gen, DefaultID)
	testCommunication.commands = append(
		testCommunication.commands,
		&Command{
			id:    DefaultCommandID,
			cType: CommandTypeUPLOAD,
			time:  DefaultTime,
		},
	)
	testCommunication.uploadMissions = append(
		testCommunication.uploadMissions,
		&UploadMission{
			commandID: DefaultCommandID,
			missionID: DefaultMissionID,
		},
	)

	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	missionID, ret := PullUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	expectCommunication := NewInstance(gen, DefaultID)
	expectCommunication.commands = append(
		expectCommunication.commands,
		&Command{
			id:    DefaultCommandID,
			cType: CommandTypeUPLOAD,
			time:  DefaultTime,
		},
	)

	a.Equal(missionID, DefaultMissionID)
	a.Len(repo.saveCommunications, 1)
	a.Equal(repo.saveCommunications[0], expectCommunication)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// CommunicationからUploadMissionを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// UploadMissionの取得が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPullUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	missionID, ret := PullUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	a.Empty(missionID)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// CommunicationからUploadMissionを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationから指定されたUploadMissionのIDが取得できない場合、
// UploadMissionの取得が失敗し、エラーが返却されることを検証する。
func TestPullErrorWhenPullUploadMissionService(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommandID1 = DefaultCommandID + "-1"
		DefaultCommandID2 = DefaultCommandID + "-2"
	)

	ctx := context.Background()

	gen := &generatorMock{}
	testCommunication := NewInstance(gen, DefaultID)
	testCommunication.commands = append(
		testCommunication.commands,
		&Command{
			id:    DefaultCommandID1,
			cType: CommandTypeARM,
			time:  DefaultTime,
		},
	)
	testCommunication.uploadMissions = append(
		testCommunication.uploadMissions,
		&UploadMission{
			commandID: DefaultCommandID,
			missionID: DefaultMissionID,
		},
	)

	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	missionID, ret := PullUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultCommandID2)

	a.Empty(missionID)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotPullUploadMission)
}

// CommunicationからUploadMissionを取得するドメインサービスをテストする。
// Communicationの保存時にリポジトリがエラーとなった場合、
// UploadMissionの取得が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenPullUploadMissionService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	testCommunication := NewInstance(gen, DefaultID)
	testCommunication.commands = append(
		testCommunication.commands,
		&Command{
			id:    DefaultCommandID,
			cType: CommandTypeARM,
			time:  DefaultTime,
		},
	)
	testCommunication.uploadMissions = append(
		testCommunication.uploadMissions,
		&UploadMission{
			commandID: DefaultCommandID,
			missionID: DefaultMissionID,
		},
	)

	repo := &repositoryMockCommandService{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	missionID, ret := PullUploadMissionService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	a.Empty(missionID)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}
