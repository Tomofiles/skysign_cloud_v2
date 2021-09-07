package mavlink

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_telemetry "edge-px4/pkg/protos/telemetry"
)

// TestAdapterArmed .
func TestAdapterArmed(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient = &telemetryServiceClientArmedMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeArmed", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterArmedInternal(ctx, supportMock, telemetryMock)

	var expectReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient = &telemetryServiceClientArmedMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterArmed .
func TestErrorWhenAdapterArmed(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeArmed", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterArmedInternal(ctx, supportMock, telemetryMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"armed telemetry error: request error"}, supportMock.messages)
}

// TestAdapterArmedSubscriber .
func TestAdapterArmedSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_telemetry.ArmedResponse{
		IsArmed: true,
	}
	response2 := &mavsdk_rpc_telemetry.ArmedResponse{
		IsArmed: false,
	}

	receiverMock := &telemetryServiceClientArmedMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterArmedSubscriber(receiverMock, supportMock)

	var resArmeds []*model.Armed
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resArmeds = append(resArmeds, a)
	}

	expectArmeds := []*model.Armed{
		{
			Armed: true,
		},
		{
			Armed: false,
		},
	}
	a.Equal(expectArmeds, resArmeds)
	a.Equal([]string{"armed receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterArmedSubscriber .
func TestReceiveErrorWhenAdapterArmedSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientArmedMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterArmedSubscriber(receiverMock, supportMock)

	var resArmeds []*model.Armed
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resArmeds = append(resArmeds, a)
	}

	a.Len(resArmeds, 0)
	a.Equal([]string{"armed receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterArmedSubscriber .
func TestCloseErrorWhenAdapterArmedSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientArmedMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterArmedSubscriber(receiverMock, supportMock)

	var resArmeds []*model.Armed
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resArmeds = append(resArmeds, a)
	}

	a.Len(resArmeds, 0)
	a.Equal([]string{"armed receive finish", "armed telemetry error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterArmedSubscriber .
func TestReceiveAndCloseErrorWhenAdapterArmedSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientArmedMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterArmedSubscriber(receiverMock, supportMock)

	var resArmeds []*model.Armed
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resArmeds = append(resArmeds, a)
	}

	a.Len(resArmeds, 0)
	a.Equal([]string{"armed receive error: receive error", "armed telemetry error: close error"}, supportMock.messages)
}
