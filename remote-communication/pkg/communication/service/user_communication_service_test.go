package service

import (
	c "remote-communication/pkg/communication/domain/communication"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPushCommandTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	testCommunication := c.NewInstance(gen, DefaultCommunicationID)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &pushCommandCommandMock{
		ID:   string(DefaultCommunicationID),
		Type: string(c.CommandTypeARM),
	}
	var resCall bool
	ret := service.PushCommand(
		command,
		func(commandID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPushCommandOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &pushCommandCommandMock{
		ID:   string(DefaultCommunicationID),
		Type: string(c.CommandTypeARM),
	}
	var resCommandID string
	ret := service.pushCommandOperation(
		nil,
		pub,
		command,
		func(commandID string) {
			resCommandID = commandID
		},
	)

	a.Nil(ret)
	a.Equal(resCommandID, string(DefaultCommunicationCommandID))
}

func TestPushUploadMissionTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	testCommunication := c.NewInstance(gen, DefaultCommunicationID)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &pushUploadMissionCommandMock{
		ID:        string(DefaultCommunicationID),
		MissionID: string(DefaultCommunicationMissionID),
	}
	var resCall bool
	ret := service.PushUploadMission(
		command,
		func(commandID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPushUploadMissionOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &pushUploadMissionCommandMock{
		ID:        string(DefaultCommunicationID),
		MissionID: string(DefaultCommunicationMissionID),
	}
	var resCommandID string
	ret := service.pushUploadMissionOperation(
		nil,
		pub,
		command,
		func(commandID string) {
			resCommandID = commandID
		},
	)

	a.Nil(ret)
	a.Equal(resCommandID, string(DefaultCommunicationCommandID))
}

func TestPullTelemetryTransaction(t *testing.T) {
	a := assert.New(t)

	snapshot := c.TelemetrySnapshot{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
		Speed:            5.0,
		Armed:            c.Armed,
		FlightMode:       "NONE",
		X:                6.0,
		Y:                7.0,
		Z:                8.0,
		W:                9.0,
	}

	gen := &generatorMock{}
	repo := &repositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	testCommunication := c.NewInstance(gen, DefaultCommunicationID)
	testCommunication.PushTelemetry(snapshot)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	var resCall bool
	ret := service.PullTelemetry(
		command,
		func(telemetry UserTelemetry) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 0)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPullTelemetryOperation(t *testing.T) {
	a := assert.New(t)

	snapshot := c.TelemetrySnapshot{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
		Speed:            5.0,
		Armed:            c.Armed,
		FlightMode:       "NONE",
		X:                6.0,
		Y:                7.0,
		Z:                8.0,
		W:                9.0,
	}

	gen := &generatorMock{}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)
	testCommunication.PushTelemetry(snapshot)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &communicationIDCommandMock{
		ID: string(DefaultCommunicationID),
	}
	var resTelemetry UserTelemetry
	ret := service.pullTelemetryOperation(
		nil,
		pub,
		command,
		func(telemetry UserTelemetry) {
			resTelemetry = telemetry
		},
	)

	a.Nil(ret)
	a.Equal(resTelemetry.GetLatitude(), 1.0)
	a.Equal(resTelemetry.GetLongitude(), 2.0)
	a.Equal(resTelemetry.GetAltitude(), 3.0)
	a.Equal(resTelemetry.GetRelativeAltitude(), 4.0)
	a.Equal(resTelemetry.GetSpeed(), 5.0)
	a.Equal(resTelemetry.GetArmed(), c.Armed)
	a.Equal(resTelemetry.GetFlightMode(), "NONE")
	a.Equal(resTelemetry.GetX(), 6.0)
	a.Equal(resTelemetry.GetY(), 7.0)
	a.Equal(resTelemetry.GetZ(), 8.0)
	a.Equal(resTelemetry.GetW(), 9.0)
}
