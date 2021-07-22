package flightplan

import (
	"context"
	"flight-operation/pkg/common/domain/txmanager"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightplanを更新するドメインサービスをテストする。
// 名前と説明を変更し、保存する。
func TestUpdateFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		NewName         = DefaultName + "-new"
		NewDescription  = DefaultDescription + "-new"
	)

	gen := &generatorMock{
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion2, DefaultVersion3},
	}
	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := UpdateFlightplan(ctx, gen, repo, pub, DefaultID, NewName, NewDescription)

	expectFlightplan := Flightplan{
		id:          DefaultID,
		name:        NewName,
		description: NewDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion3,
		gen:         gen,
	}

	a.Len(repo.saveFlightplans, 1)
	a.Equal(repo.saveFlightplans[0], &expectFlightplan)
	a.Len(pub.events, 0)

	a.Nil(ret)
}

// Flightplanを更新するドメインサービスをテストする。
// 指定されたIDのFlightplanの取得がエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenUpdateFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		NewName        = DefaultName + "-new"
		NewDescription = DefaultDescription + "-new"
	)

	gen := &generatorMock{}

	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	ret := UpdateFlightplan(ctx, gen, repo, pub, DefaultID, NewName, NewDescription)

	a.Len(repo.saveFlightplans, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// Flightplanを更新するドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 更新が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenUpdateFlightplanService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultVersion1 = DefaultVersion + "-1"
		DefaultVersion2 = DefaultVersion + "-2"
		DefaultVersion3 = DefaultVersion + "-3"
		NewName         = DefaultName + "-new"
		NewDescription  = DefaultDescription + "-new"
	)

	gen := &generatorMock{
		fleetID:  DefaultFleetID,
		versions: []Version{DefaultVersion2, DefaultVersion3},
	}
	testFlightplan := Flightplan{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		version:     DefaultVersion1,
		newVersion:  DefaultVersion1,
		gen:         gen,
	}
	repo := &repositoryMockUpdateService{}
	repo.On("GetByID", DefaultID).Return(&testFlightplan, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	ret := UpdateFlightplan(ctx, gen, repo, pub, DefaultID, NewName, NewDescription)

	a.Len(repo.saveFlightplans, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrSave)
}

// Flightplan更新サービス用リポジトリモック
type repositoryMockUpdateService struct {
	mock.Mock

	saveFlightplans []*Flightplan
}

func (rm *repositoryMockUpdateService) GetAll(tx txmanager.Tx) ([]*Flightplan, error) {
	panic("implement me")
}
func (rm *repositoryMockUpdateService) GetByID(tx txmanager.Tx, id ID) (*Flightplan, error) {
	ret := rm.Called(id)
	var v *Flightplan
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*Flightplan)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMockUpdateService) Save(tx txmanager.Tx, v *Flightplan) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveFlightplans = append(rm.saveFlightplans, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMockUpdateService) Delete(tx txmanager.Tx, id ID) error {
	panic("implement me")
}
