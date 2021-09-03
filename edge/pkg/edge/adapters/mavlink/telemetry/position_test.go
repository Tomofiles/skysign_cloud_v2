package mavlink

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// TestAdapterPosition .
func TestAdapterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient = &telemetryServiceClientPositionMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribePosition", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterPositionInternal(ctx, supportMock, telemetryMock)

	var expectReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient = &telemetryServiceClientPositionMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterPosition .
func TestErrorWhenAdapterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribePosition", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterPositionInternal(ctx, supportMock, telemetryMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"position telemetry error: request error"}, supportMock.messages)
}

// TestAdapterPositionSubscriber .
func TestAdapterPositionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_telemetry.PositionResponse{
		Position: &mavsdk_rpc_telemetry.Position{
			LatitudeDeg:       11.0,
			LongitudeDeg:      21.0,
			AbsoluteAltitudeM: 31.0,
			RelativeAltitudeM: 41.0,
		},
	}
	response2 := &mavsdk_rpc_telemetry.PositionResponse{
		Position: &mavsdk_rpc_telemetry.Position{
			LatitudeDeg:       12.0,
			LongitudeDeg:      22.0,
			AbsoluteAltitudeM: 32.0,
			RelativeAltitudeM: 42.0,
		},
	}

	receiverMock := &telemetryServiceClientPositionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterPositionSubscriber(receiverMock, supportMock)

	var resPositions []*edge.Position
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resPositions = append(resPositions, a)
	}

	expectPositions := []*edge.Position{
		{
			Latitude:         11.0,
			Longitude:        21.0,
			Altitude:         31.0,
			RelativeAltitude: 41.0,
		},
		{
			Latitude:         12.0,
			Longitude:        22.0,
			Altitude:         32.0,
			RelativeAltitude: 42.0,
		},
	}
	a.Equal(expectPositions, resPositions)
	a.Equal([]string{"position receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterPositionSubscriber .
func TestReceiveErrorWhenAdapterPositionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientPositionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterPositionSubscriber(receiverMock, supportMock)

	var resPositions []*edge.Position
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resPositions = append(resPositions, a)
	}

	a.Len(resPositions, 0)
	a.Equal([]string{"position receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterPositionSubscriber .
func TestCloseErrorWhenAdapterPositionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientPositionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterPositionSubscriber(receiverMock, supportMock)

	var resPositions []*edge.Position
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resPositions = append(resPositions, a)
	}

	a.Len(resPositions, 0)
	a.Equal([]string{"position receive finish", "position telemetry error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterPositionSubscriber .
func TestReceiveAndCloseErrorWhenAdapterPositionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientPositionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterPositionSubscriber(receiverMock, supportMock)

	var resPositions []*edge.Position
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resPositions = append(resPositions, a)
	}

	a.Len(resPositions, 0)
	a.Equal([]string{"position receive error: receive error", "position telemetry error: close error"}, supportMock.messages)
}
