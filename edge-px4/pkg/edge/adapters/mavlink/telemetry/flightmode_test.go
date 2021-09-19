package mavlink

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_telemetry "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/telemetry"
)

// TestAdapterFlightMode .
func TestAdapterFlightMode(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient = &telemetryServiceClientFlightModeMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeFlightMode", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterFlightModeInternal(ctx, supportMock, telemetryMock)

	var expectReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient = &telemetryServiceClientFlightModeMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterFlightMode .
func TestErrorWhenAdapterFlightMode(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeFlightMode", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterFlightModeInternal(ctx, supportMock, telemetryMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"flightMode telemetry error: request error"}, supportMock.messages)
}

// TestAdapterFlightModeSubscriber .
func TestAdapterFlightModeSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_telemetry.FlightModeResponse{
		FlightMode: mavsdk_rpc_telemetry.FlightMode_HOLD,
	}
	response2 := &mavsdk_rpc_telemetry.FlightModeResponse{
		FlightMode: mavsdk_rpc_telemetry.FlightMode_LAND,
	}

	receiverMock := &telemetryServiceClientFlightModeMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterFlightModeSubscriber(receiverMock, supportMock)

	var resFlightModes []*model.FlightMode
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resFlightModes = append(resFlightModes, a)
	}

	expectFlightModes := []*model.FlightMode{
		{
			FlightMode: mavsdk_rpc_telemetry.FlightMode_HOLD.String(),
		},
		{
			FlightMode: mavsdk_rpc_telemetry.FlightMode_LAND.String(),
		},
	}
	a.Equal(expectFlightModes, resFlightModes)
	a.Equal([]string{"flightMode receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterFlightModeSubscriber .
func TestReceiveErrorWhenAdapterFlightModeSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientFlightModeMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterFlightModeSubscriber(receiverMock, supportMock)

	var resFlightModes []*model.FlightMode
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resFlightModes = append(resFlightModes, a)
	}

	a.Len(resFlightModes, 0)
	a.Equal([]string{"flightMode receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterFlightModeSubscriber .
func TestCloseErrorWhenAdapterFlightModeSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientFlightModeMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterFlightModeSubscriber(receiverMock, supportMock)

	var resFlightModes []*model.FlightMode
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resFlightModes = append(resFlightModes, a)
	}

	a.Len(resFlightModes, 0)
	a.Equal([]string{"flightMode receive finish", "flightMode telemetry error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterFlightModeSubscriber .
func TestReceiveAndCloseErrorWhenAdapterFlightModeSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientFlightModeMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterFlightModeSubscriber(receiverMock, supportMock)

	var resFlightModes []*model.FlightMode
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resFlightModes = append(resFlightModes, a)
	}

	a.Len(resFlightModes, 0)
	a.Equal([]string{"flightMode receive error: receive error", "flightMode telemetry error: close error"}, supportMock.messages)
}
