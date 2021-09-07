package mavlink

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_core "edge-px4/pkg/protos/core"
)

// TestAdapterConnectionState .
func TestAdapterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient = &coreServiceClientConnectionStateMock{}

	coreMock := &coreServiceClientMock{}
	coreMock.On("SubscribeConnectionState", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterConnectionStateInternal(ctx, supportMock, coreMock)

	var expectReceiver mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient = &coreServiceClientConnectionStateMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterConnectionState .
func TestErrorWhenAdapterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	coreMock := &coreServiceClientMock{}
	coreMock.On("SubscribeConnectionState", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterConnectionStateInternal(ctx, supportMock, coreMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"connectionState core error: request error"}, supportMock.messages)
}

// TestAdapterConnectionStateSubscriber .
func TestAdapterConnectionStateSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_core.ConnectionStateResponse{
		ConnectionState: &mavsdk_rpc_core.ConnectionState{
			Uuid: uint64(1234567890),
		},
	}
	response2 := &mavsdk_rpc_core.ConnectionStateResponse{
		ConnectionState: &mavsdk_rpc_core.ConnectionState{
			Uuid: uint64(1234567899),
		},
	}

	receiverMock := &coreServiceClientConnectionStateMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterConnectionStateSubscriber(receiverMock, supportMock)

	var resConnectionStates []*model.ConnectionState
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resConnectionStates = append(resConnectionStates, a)
	}

	expectConnectionStates := []*model.ConnectionState{
		{
			VehicleID: "1234567890",
		},
		{
			VehicleID: "1234567899",
		},
	}
	a.Equal(expectConnectionStates, resConnectionStates)
	a.Equal([]string{"connectionState receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterConnectionStateSubscriber .
func TestReceiveErrorWhenAdapterConnectionStateSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &coreServiceClientConnectionStateMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterConnectionStateSubscriber(receiverMock, supportMock)

	var resConnectionStates []*model.ConnectionState
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resConnectionStates = append(resConnectionStates, a)
	}

	a.Len(resConnectionStates, 0)
	a.Equal([]string{"connectionState receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterConnectionStateSubscriber .
func TestCloseErrorWhenAdapterConnectionStateSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &coreServiceClientConnectionStateMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterConnectionStateSubscriber(receiverMock, supportMock)

	var resConnectionStates []*model.ConnectionState
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resConnectionStates = append(resConnectionStates, a)
	}

	a.Len(resConnectionStates, 0)
	a.Equal([]string{"connectionState receive finish", "connectionState core error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterConnectionStateSubscriber .
func TestReceiveAndCloseErrorWhenAdapterConnectionStateSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &coreServiceClientConnectionStateMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterConnectionStateSubscriber(receiverMock, supportMock)

	var resConnectionStates []*model.ConnectionState
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resConnectionStates = append(resConnectionStates, a)
	}

	a.Len(resConnectionStates, 0)
	a.Equal([]string{"connectionState receive error: receive error", "connectionState core error: close error"}, supportMock.messages)
}
