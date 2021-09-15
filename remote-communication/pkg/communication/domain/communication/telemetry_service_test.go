package communication

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// CommunicationのTelemetryを更新するドメインサービスをテストする。
// 指定されたIDのCommunicationのTelemetryが更新されることを検証する。
// Telemetryの更新が成功した場合、格納されているCommandのIDリストが返却されることを検証する。
// また、更新が成功した際、ドメインイベントが発行されること。
func TestPushTelemetryService(t *testing.T) {
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
		&Command{
			id:    DefaultCommandID2,
			cType: CommandTypeDISARM,
			time:  DefaultTime,
		},
	)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	snapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	ids, ret := PushTelemetryService(ctx, gen, repo, pub, DefaultID, snapshot)

	expectCommunication := NewInstance(gen, DefaultID)
	expectCommunication.commands = append(
		expectCommunication.commands,
		&Command{
			id:    DefaultCommandID1,
			cType: CommandTypeARM,
			time:  DefaultTime,
		},
		&Command{
			id:    DefaultCommandID2,
			cType: CommandTypeDISARM,
			time:  DefaultTime,
		},
	)
	expectCommunication.telemetry = &Telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}
	expectCommunication.SetPublisher(pub)

	expectEvent := TelemetryUpdatedEvent{
		CommunicationID: DefaultID,
		Telemetry:       snapshot,
	}

	a.Equal(ids, []CommandID{DefaultCommandID1, DefaultCommandID2})
	a.Len(repo.saveCommunications, 1)
	a.Equal(repo.saveCommunications[0], expectCommunication)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)

	a.Nil(ret)
}

// CommunicationのTelemetryを更新するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// Telemetryを更新が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPushTelemetryService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	snapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	ids, ret := PushTelemetryService(ctx, gen, repo, pub, DefaultID, snapshot)

	a.Empty(ids)
	a.Len(repo.saveCommunications, 0)
	a.Len(pub.events, 0)

	a.Equal(ret, ErrGet)
}

// CommunicationのTelemetryを更新するドメインサービスをテストする。
// Communicationの保存時にリポジトリがエラーとなった場合、
// Telemetryを更新が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenPushTelemetryService(t *testing.T) {
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

	snapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	ids, ret := PushTelemetryService(ctx, gen, repo, pub, DefaultID, snapshot)

	a.Empty(ids)
	a.Len(repo.saveCommunications, 0)
	// a.Len(pub.events, 0) // 1件PublishされるがFlushされない想定

	a.Equal(ret, ErrSave)
}

// CommunicationからTelemetryを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationからTelemetryを取得されることを検証する。
func TestPullTelemetryService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	testCommunication := NewInstance(gen, DefaultID)
	testCommunication.telemetry = &Telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(testCommunication, nil)

	pub := &publisherMock{}

	snapshot, ret := PullTelemetryService(ctx, gen, repo, pub, DefaultID)

	expectSnapshot := TelemetrySnapshot{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}

	a.Equal(snapshot, expectSnapshot)

	a.Nil(ret)
}

// CommunicationからTelemetryを取得するドメインサービスをテストする。
// 指定されたIDのCommunicationの取得がエラーとなった場合、
// Telemetryの取得が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenPullTelemetryService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	pub := &publisherMock{}

	snapshot, ret := PullTelemetryService(ctx, gen, repo, pub, DefaultID)

	a.Empty(snapshot)

	a.Equal(ret, ErrGet)
}
