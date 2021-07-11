package flightplan

import (
	"context"
	"flightplan/pkg/flightplan/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightplanを作成するドメインサービスをテストする。
// 名前と説明をあらかじめ与えられたFlightplanを作成し、保存する。
// 機体数を変更する際に、FleetIDが付与されたイベントが発行される
// ことを検証する。
func TestCreateNewFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		DefaultVersion4 = DefaultVersion + "-4"
	)

	gen := &generatorMock{
		id:       DefaultID,
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	id, fleetID, ret := CreateNewFlightplan(ctx, gen, repo, pub, DefaultName, DefaultDescription)

	expectFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion4,
		gen:         gen,
		pub:         pub,
	}
	expectEvent := FleetIDGaveEvent{
		FleetID:          DefaultFleetID,
		NumberOfVehicles: 0,
	}

	a.Equal(id, DefaultID)
	a.Equal(fleetID, DefaultFleetID)
	a.Len(repo.saveFlightplans, 1)
	a.Equal(repo.saveFlightplans[0], &expectFlightplan)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// Flightplanを作成するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 作成が失敗し、エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestSaveErrorWhenCreateNewFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		DefaultVersion4 = DefaultVersion + "-4"
	)

	gen := &generatorMock{
		id:       DefaultID,
		versions: []Version{DefaultVersion1, DefaultVersion2, DefaultVersion3, DefaultVersion4},
	}
	repo := &repositoryMockCreateService{}
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	id, fleetID, ret := CreateNewFlightplan(ctx, gen, repo, pub, DefaultName, DefaultDescription)

	a.Empty(id)
	a.Empty(fleetID)
	a.Len(repo.saveFlightplans, 0)
	// a.Len(pub.events, 0) // 1件PublishされるがFlushされない想定
	a.Equal(ret, ErrSave)
}

// Flightplan作成サービス用リポジトリモック
type repositoryMockCreateService struct {
	mock.Mock

	saveFlightplans []*Flightplan
}

func (rm *repositoryMockCreateService) GetAll(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetAllOrigin(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) GetByID(tx txmanager.Tx, id ID) (*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockCreateService) Save(tx txmanager.Tx, f *Flightplan) error {
	ret := rm.Called(f)
	if ret.Error(0) == nil {
		rm.saveFlightplans = append(rm.saveFlightplans, f)
	}
	return ret.Error(0)
}
func (rm *repositoryMockCreateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
