package flightplan

import (
	"context"
	"flight-operation/pkg/common/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightplanを実行するドメインサービスをテストする。
// 実行されたことを表すドメインイベントが発行されたことを
// 検証する。
func TestExecuteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         nil,
	}
	repo := &repositoryMockExecuteService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	pub := &publisherMock{}

	ret := ExecuteFlightplan(ctx, repo, pub, DefaultID)

	expectEvent := FlightplanExecutedEvent{
		ID:          DefaultID,
		Name:        DefaultName,
		Description: DefaultDescription,
		FleetID:     DefaultFleetID,
	}

	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Flightplanを実行するドメインサービスをテストする。
// 指定されたIDのFlightplanの取得がエラーとなった場合、
// 実行が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenExecuteFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	repo := &repositoryMockExecuteService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)

	pub := &publisherMock{}

	ret := ExecuteFlightplan(ctx, repo, pub, DefaultID)

	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Flightplan実行サービス用リポジトリモック
type repositoryMockExecuteService struct {
	mock.Mock
}

func (rm *repositoryMockExecuteService) GetAll(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockExecuteService) GetByID(tx txmanager.Tx, id ID) (*Flightplan, error) {
	ret := rm.Called(id)
	var v *Flightplan
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Flightplan)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockExecuteService) Save(tx txmanager.Tx, v *Flightplan) error {
	panic("implement me")
}
func (rm *repositoryMockExecuteService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
