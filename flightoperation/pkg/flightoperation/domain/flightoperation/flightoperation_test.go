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

	flightoperation.Complete()

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), NewVersion)
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

	flightoperation.Complete()

	expectEvent := &CompletedEvent{
		ID: DefaultID,
	}

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), NewVersion)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}

// FlightoperationをComplete状態に更新する。
// イベントパブリッシャーが設定されている場合、
// すでにComplete状態であれば、イベントが発行されない
// ことを検証する。
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

	flightoperation.Complete()

	a.Equal(flightoperation.isCompleted, Completed)
	a.Equal(flightoperation.GetVersion(), DefaultVersion)
	a.Equal(flightoperation.GetNewVersion(), NewVersion)
	a.Len(pub.events, 0)
}
