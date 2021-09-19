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

// TestAdapterQuaternion .
func TestAdapterQuaternion(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient = &telemetryServiceClientQuaternionMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeAttitudeQuaternion", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterQuaternionInternal(ctx, supportMock, telemetryMock)

	var expectReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient = &telemetryServiceClientQuaternionMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterQuaternion .
func TestErrorWhenAdapterQuaternion(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeAttitudeQuaternion", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterQuaternionInternal(ctx, supportMock, telemetryMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"quaternion telemetry error: request error"}, supportMock.messages)
}

// TestAdapterQuaternionSubscriber .
func TestAdapterQuaternionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_telemetry.AttitudeQuaternionResponse{
		AttitudeQuaternion: &mavsdk_rpc_telemetry.Quaternion{
			X: 11.0,
			Y: 21.0,
			Z: 31.0,
			W: 41.0,
		},
	}
	response2 := &mavsdk_rpc_telemetry.AttitudeQuaternionResponse{
		AttitudeQuaternion: &mavsdk_rpc_telemetry.Quaternion{
			X: 12.0,
			Y: 22.0,
			Z: 32.0,
			W: 42.0,
		},
	}

	receiverMock := &telemetryServiceClientQuaternionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterQuaternionSubscriber(receiverMock, supportMock)

	var resQuaternions []*model.Quaternion
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resQuaternions = append(resQuaternions, a)
	}

	expectQuaternions := []*model.Quaternion{
		{
			X: 11.0,
			Y: 21.0,
			Z: 31.0,
			W: 41.0,
		},
		{
			X: 12.0,
			Y: 22.0,
			Z: 32.0,
			W: 42.0,
		},
	}
	a.Equal(expectQuaternions, resQuaternions)
	a.Equal([]string{"quaternion receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterQuaternionSubscriber .
func TestReceiveErrorWhenAdapterQuaternionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientQuaternionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterQuaternionSubscriber(receiverMock, supportMock)

	var resQuaternions []*model.Quaternion
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resQuaternions = append(resQuaternions, a)
	}

	a.Len(resQuaternions, 0)
	a.Equal([]string{"quaternion receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterQuaternionSubscriber .
func TestCloseErrorWhenAdapterQuaternionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientQuaternionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterQuaternionSubscriber(receiverMock, supportMock)

	var resQuaternions []*model.Quaternion
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resQuaternions = append(resQuaternions, a)
	}

	a.Len(resQuaternions, 0)
	a.Equal([]string{"quaternion receive finish", "quaternion telemetry error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterQuaternionSubscriber .
func TestReceiveAndCloseErrorWhenAdapterQuaternionSubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientQuaternionMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterQuaternionSubscriber(receiverMock, supportMock)

	var resQuaternions []*model.Quaternion
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resQuaternions = append(resQuaternions, a)
	}

	a.Len(resQuaternions, 0)
	a.Equal([]string{"quaternion receive error: receive error", "quaternion telemetry error: close error"}, supportMock.messages)
}
