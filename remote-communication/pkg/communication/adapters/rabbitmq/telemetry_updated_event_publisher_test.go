package rabbitmq

import (
	"errors"
	"remote-communication/pkg/communication/domain/communication"
	"remote-communication/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestPublishTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	event := communication.TelemetryUpdatedEvent{
		CommunicationID: DefaultCommunicationID,
		Telemetry: communication.TelemetrySnapshot{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            communication.Armed,
			FlightMode:       "NONE",
			X:                6.0,
			Y:                7.0,
			Z:                8.0,
			W:                9.0,
		},
	}

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishTelemetryUpdatedEvent(chMock, event)

	expectPb := skysign_proto.TelemetryUpdatedEvent{
		CommunicationId: string(DefaultCommunicationID),
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            communication.Armed,
			FlightMode:       "NONE",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}
	expectBin, _ := proto.Marshal(&expectPb)

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.Equal(chMock.message, expectBin)
}

func TestFanoutExchangeDeclareErrorWhenPublishTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	event := communication.TelemetryUpdatedEvent{
		CommunicationID: DefaultCommunicationID,
		Telemetry: communication.TelemetrySnapshot{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            communication.Armed,
			FlightMode:       "NONE",
			X:                6.0,
			Y:                7.0,
			Z:                8.0,
			W:                9.0,
		},
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	ret := PublishTelemetryUpdatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenPublishTelemetryUpdatedEvent(t *testing.T) {
	a := assert.New(t)

	event := communication.TelemetryUpdatedEvent{
		CommunicationID: DefaultCommunicationID,
		Telemetry: communication.TelemetrySnapshot{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            communication.Armed,
			FlightMode:       "NONE",
			X:                6.0,
			Y:                7.0,
			Z:                8.0,
			W:                9.0,
		},
	}

	errPub := errors.New("publish error")

	chMock := &channelMockPublish{}
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	ret := PublishTelemetryUpdatedEvent(chMock, event)

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}
