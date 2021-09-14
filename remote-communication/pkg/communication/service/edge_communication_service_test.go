package service

import (
	c "remote-communication/pkg/communication/domain/communication"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPullCommandTransaction(t *testing.T) {
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
	testCommunication.PushCommand(c.CommandTypeARM)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &pullCommandMock{
		ID:        string(DefaultCommunicationID),
		CommandID: string(DefaultCommunicationCommandID),
	}
	var resCall bool
	ret := service.PullCommand(
		command,
		func(cType string) {
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

func TestPullCommandOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)
	testCommunication.PushCommand(c.CommandTypeARM)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &pullCommandMock{
		ID:        string(DefaultCommunicationID),
		CommandID: string(DefaultCommunicationCommandID),
	}
	var resType string
	ret := service.pullCommandOperation(
		nil,
		pub,
		command,
		func(cType string) {
			resType = cType
		},
	)

	a.Nil(ret)
	a.Equal(resType, string(c.CommandTypeARM))
}

func TestPullUploadMissionTransaction(t *testing.T) {
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
	testCommunication.PushUploadMission(DefaultCommunicationMissionID)

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	command := &pullCommandMock{
		ID:        string(DefaultCommunicationID),
		CommandID: string(DefaultCommunicationCommandID),
	}
	var resCall bool
	ret := service.PullUploadMission(
		command,
		func(missionID string) {
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

func TestPullUploadMissionOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID},
		times:      []time.Time{DefaultCommunicationTime},
	}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)
	testCommunication.PushUploadMission(DefaultCommunicationMissionID)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	command := &pullCommandMock{
		ID:        string(DefaultCommunicationID),
		CommandID: string(DefaultCommunicationCommandID),
	}
	var resMissionID string
	ret := service.pullUploadMissionOperation(
		nil,
		pub,
		command,
		func(missionID string) {
			resMissionID = missionID
		},
	)

	a.Nil(ret)
	a.Equal(resMissionID, string(DefaultCommunicationMissionID))
}

func TestPushTelemetryTransaction(t *testing.T) {
	a := assert.New(t)

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

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	telemetry := &telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             c.Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	command := &pushTelemetryCommandMock{
		ID:        string(DefaultCommunicationID),
		Telemetry: telemetry,
	}
	var resCall bool
	ret := service.PushTelemetry(
		command,
		func(commandIDs []string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPushTelemetryOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommunicationCommandID1 = DefaultCommunicationCommandID + "-1"
		DefaultCommunicationCommandID2 = DefaultCommunicationCommandID + "-2"
		DefaultCommunicationTime1      = DefaultCommunicationTime.Add(1 * time.Minute)
		DefaultCommunicationTime2      = DefaultCommunicationTime.Add(2 * time.Minute)
	)

	gen := &generatorMock{
		commandIDs: []c.CommandID{DefaultCommunicationCommandID1, DefaultCommunicationCommandID2},
		times:      []time.Time{DefaultCommunicationTime1, DefaultCommunicationTime2},
	}
	testCommunication := c.NewInstance(gen, DefaultCommunicationID)
	testCommunication.PushCommand(c.CommandTypeARM)
	testCommunication.PushCommand(c.CommandTypeDISARM)

	repo := &repositoryMock{}
	repo.On("GetByID", DefaultCommunicationID).Return(testCommunication, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	telemetry := &telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             c.Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	command := &pushTelemetryCommandMock{
		ID:        string(DefaultCommunicationID),
		Telemetry: telemetry,
	}
	var resCommandIDs []string
	ret := service.pushTelemetryOperation(
		nil,
		pub,
		command,
		func(commandIDs []string) {
			resCommandIDs = commandIDs
		},
	)

	expectEvent := c.TelemetryUpdatedEvent{
		CommunicationID: DefaultCommunicationID,
		Telemetry: c.TelemetrySnapshot{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           5.0,
			Armed:             c.Armed,
			FlightMode:        "NONE",
			X:                 6.0,
			Y:                 7.0,
			Z:                 8.0,
			W:                 9.0,
		},
	}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events, []interface{}{expectEvent})
	a.Equal(resCommandIDs, []string{string(DefaultCommunicationCommandID1), string(DefaultCommunicationCommandID2)})
}
