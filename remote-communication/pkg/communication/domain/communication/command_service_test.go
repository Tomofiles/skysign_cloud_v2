package communication

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// CommunicationにCommandを追加するドメインサービスをテストする。
// 指定されたIDのCommunicationにCommandが追加されることを検証する。
func TestPushCommandService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	testCommunication := NewInstance(gen, DefaultID)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	id, ret := PushCommandService(ctx, gen, repo, pub, DefaultID, CommandTypeARM)

	expectCommunication := NewInstance(gen, DefaultID)
	expectCommunication.commands = append(
		expectCommunication.commands,
		&Command{
			id:    DefaultCommandID,
			cType: CommandTypeARM,
			time:  DefaultTime,
		},
	)

	a.Equal(id, DefaultCommandID)
	a.Len(repo.saveCommunications, 1)
	a.Equal(repo.saveCommunications[0], expectCommunication)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// CommunicationにCommandを追加するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// Commandの追加が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPushCommandService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	id, ret := PushCommandService(ctx, gen, repo, pub, DefaultID, CommandTypeARM)

	a.Empty(id)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// CommunicationにCommandを追加するドメインサービスをテストする。
// Communicationの保存時にリポジトリがエラーとなった場合、
// Commandの追加が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenPushCommandService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{
		commandIDs: []CommandID{DefaultCommandID},
		times:      []time.Time{DefaultTime},
	}
	testCommunication := NewInstance(gen, DefaultID)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	id, ret := PushCommandService(ctx, gen, repo, pub, DefaultID, CommandTypeARM)

	a.Empty(id)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}

// CommunicationからCommandを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationからCommandを取得されることを検証する。
func TestPullCommandService(t *testing.T) {
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

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	cType, ret := PullCommandService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	expectCommunication := NewInstance(gen, DefaultID)

	a.Equal(cType, CommandTypeARM)
	a.Len(repo.saveCommunications, 1)
	a.Equal(repo.saveCommunications[0], expectCommunication)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// CommunicationからCommandを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// Commandの取得が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPullCommandService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	cType, ret := PullCommandService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	a.Empty(cType)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// CommunicationからCommandを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationから指定されたCommandのIDが取得できない場合、
// Commandの取得が失敗し、エラーが返却されることを検証する。
func TestPullErrorWhenPullCommandService(t *testing.T) {
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

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	cType, ret := PullCommandService(ctx, gen, repo, pub, DefaultID, DefaultCommandID2)

	a.Empty(cType)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrCannotPullCommand)
}

// CommunicationからCommandを取得するドメインサービスをテストする。
// Communicationの保存時にリポジトリがエラーとなった場合、
// Commandの取得が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenPullCommandService(t *testing.T) {
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

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	pub := &publisherMock{}

	cType, ret := PullCommandService(ctx, gen, repo, pub, DefaultID, DefaultCommandID)

	a.Empty(cType)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}
