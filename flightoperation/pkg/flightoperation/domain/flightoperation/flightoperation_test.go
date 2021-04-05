package flightoperation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FlightoperationをComplete状態に更新する。
// Complete状態に更新されたことを検証する。
func TestCompleteFlightoperation(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	flightoperation := &Flightoperation{
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
		isCompleted:  Operating,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}

	err := flightoperation.Complete()

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), NewVersion)
	a.Nil(err)
}

// FlightoperationをComplete状態に更新する。
// イベントパブリッシャーが設定されている場合、
// Completeされたことを表すドメインイベントが発行
// されていることを検証する。
func TestPublishCompletedEventWhenCompleteFlightoperation(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	pub := &publisherMock{}

	flightoperation := &Flightoperation{
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
		isCompleted:  Operating,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}

	flightoperation.SetPublisher(pub)

	err := flightoperation.Complete()

	expectEvent := CompletedEvent{
		ID:           DefaultID,
		FlightplanID: DefaultFlightplanID,
	}

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), NewVersion)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
	a.Nil(err)
}

// FlightoperationをComplete状態に更新する。
// すでにComplete状態であれば、エラーが発生することを検証する。
func TestNonePublishWhenCompleteFlightoperation(t *testing.T) {
	a := assert.New(t)

	const (
		NewVersion = DefaultVersion + "-new"
	)

	gen := &generatorMock{
		version: NewVersion,
	}

	pub := &publisherMock{}

	flightoperation := &Flightoperation{
		id:           DefaultID,
		flightplanID: DefaultFlightplanID,
		isCompleted:  Completed,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}

	flightoperation.SetPublisher(pub)

	err := flightoperation.Complete()

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), DefaultVersion)
	a.Len(pub.events, 0)
	a.Equal(err, ErrCannotChange)
}
