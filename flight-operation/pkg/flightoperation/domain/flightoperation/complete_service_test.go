package flightoperation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Flightoperationを完了させるドメインサービスをテストする。
// 完了操作が成功すると、Flightoperationが完了されたことを表す
// ドメインイベントを発行する。
func TestCompleteFlightoperationService(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		isCompleted: Operating,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         gen,
	}

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := CompleteFlightoperation(nil, repo, pub, DefaultID)

	expectFope := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		isCompleted: Completed,
		version:     DefaultVersion,
		newVersion:  NewVersion,
		gen:         gen,
		pub:         pub,
	}
	expectEvent := FlightoperationCompletedEvent{
		ID:          DefaultID,
		Name:        DefaultName,
		Description: DefaultDescription,
		FleetID:     DefaultFleetID,
	}

	a.Nil(ret)
	a.Len(repo.saveFlightoperations, 1)
	a.Equal(repo.saveFlightoperations[0], expectFope)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
}

// Flightoperationを更新するドメインサービスをテストする。
// 指定されたIDのFlightoperationの取得がエラーとなった場合、
// 完了操作が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenCompleteFlightoperationService(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := CompleteFlightoperation(nil, repo, pub, DefaultID)

	a.Len(repo.saveFlightoperations, 0)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrGet)
}

// Flightoperationを完了させるドメインサービスをテストする。
// Flightoperationが完了済みの場合、完了操作が失敗し、
// エラーが返却されることを検証する。
// また、ドメインイベントは発行されないことを検証する。
func TestCannotChangeErrorWhenCompleteFlightoperationService(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		isCompleted: Completed,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         gen,
	}

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	ret := CompleteFlightoperation(nil, repo, pub, DefaultID)

	a.Len(repo.saveFlightoperations, 0)
	a.Len(pub.events, 0)
	a.Equal(ret, ErrCannotChange)
}

// Flightoperationを完了させるドメインサービスをテストする。
// 保存時にリポジトリがエラーとなった場合、
// 完了操作が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenCompleteFlightoperationService(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := &Flightoperation{
		id:          DefaultID,
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     DefaultFleetID,
		isCompleted: Operating,
		version:     DefaultVersion,
		newVersion:  DefaultVersion,
		gen:         gen,
	}

	repo := &flightoperationRepositoryMock{}
	repo.On("GetByID", DefaultID).Return(flightoperation, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)
	pub := &publisherMock{}

	ret := CompleteFlightoperation(nil, repo, pub, DefaultID)

	a.Len(repo.saveFlightoperations, 0)
	// a.Len(pub.events, 0) // 1件PublishされるがFlushされない想定
	a.Equal(ret, ErrSave)
}
