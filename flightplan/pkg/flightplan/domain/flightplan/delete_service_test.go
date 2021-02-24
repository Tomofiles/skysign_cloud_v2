package flightplan

import (
	"context"
	"flightplan/pkg/flightplan/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightplanを削除するドメインサービスをテストする。
// 指定されたIDのFlightplanを削除する。
// 削除が成功すると、Flightplanが削除されたことを表す
// ドメインイベントを発行する。
func TestDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	expectEvent := DeletedEvent{ID: DefaultID}

	a.Len(repo.deleteIDs, 1)
	a.Equal(repo.deleteIDs[0], DefaultID)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Flightplanを削除するドメインサービスをテストする。
// 指定されたIDのFlightplanの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestGetErrorWhenDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Delete", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Flightplanを削除するドメインサービスをテストする。
// 削除時にリポジトリがエラーとなった場合、、
// 削除が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestDeleteErrorWhenDeleteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         nil,
	}
	repo := &repositoryMockDeleteService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Delete", mock.Anything).Return(ErrDelete)

	pub := &publisherMock{}

	ret := DeleteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(repo.deleteIDs, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrDelete)
}

// Flightplan削除サービス用リポジトリモック
type repositoryMockDeleteService struct {
	mock.Mock

	saveFlightplans []*Flightplan
	deleteIDs       []ID
}

func (rm *repositoryMockDeleteService) GetAll(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) GetByID(tx txmanager.Tx, id ID) (*Flightplan, error) {
	ret := rm.Called(id)
	var f *Flightplan
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*Flightplan)
	}
	return f, ret.Error(1)
}
func (rm *repositoryMockDeleteService) Save(tx txmanager.Tx, f *Flightplan) error {
	panic("implement me")
}
func (rm *repositoryMockDeleteService) Delete(tx txmanager.Tx, id ID) error {
	ret := rm.Called(id)
	if ret.Error(0) == nil {
		rm.deleteIDs = append(rm.deleteIDs, id)
	}
	return ret.Error(0)
}
