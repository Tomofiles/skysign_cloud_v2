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

// TestAdapterVelocity .
func TestAdapterVelocity(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	var response mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient = &telemetryServiceClientVelocityMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeGroundSpeedNed", mock.Anything, mock.Anything).Return(response, nil)

	receiver, ret := AdapterVelocityInternal(ctx, supportMock, telemetryMock)

	var expectReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient = &telemetryServiceClientVelocityMock{}

	a.Nil(ret)
	a.Equal(expectReceiver, receiver)
	a.Empty(supportMock.messages)
}

// TestErrorWhenAdapterVelocity .
func TestErrorWhenAdapterVelocity(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	telemetryMock := &telemetryServiceClientMock{}
	telemetryMock.On("SubscribeGroundSpeedNed", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	receiver, ret := AdapterVelocityInternal(ctx, supportMock, telemetryMock)

	a.Nil(receiver)
	a.Equal(ErrRequest, ret)
	a.Equal([]string{"velocity telemetry error: request error"}, supportMock.messages)
}

// TestAdapterVelocitySubscriber .
func TestAdapterVelocitySubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	response1 := &mavsdk_rpc_telemetry.GroundSpeedNedResponse{
		GroundSpeedNed: &mavsdk_rpc_telemetry.SpeedNed{
			VelocityNorthMS: 11.0,
			VelocityEastMS:  21.0,
			VelocityDownMS:  31.0,
		},
	}
	response2 := &mavsdk_rpc_telemetry.GroundSpeedNedResponse{
		GroundSpeedNed: &mavsdk_rpc_telemetry.SpeedNed{
			VelocityNorthMS: 12.0,
			VelocityEastMS:  22.0,
			VelocityDownMS:  32.0,
		},
	}

	receiverMock := &telemetryServiceClientVelocityMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(response1, nil)
	receiverMock.On("Recv", 2).Return(response2, nil)
	receiverMock.On("Recv", 3).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterVelocitySubscriber(receiverMock, supportMock)

	var resVelocitys []*model.Velocity
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resVelocitys = append(resVelocitys, a)
	}

	expectVelocitys := []*model.Velocity{
		{
			NorthMS: 11.0,
			EastMS:  21.0,
			DownMS:  31.0,
		},
		{
			NorthMS: 12.0,
			EastMS:  22.0,
			DownMS:  32.0,
		},
	}
	a.Equal(expectVelocitys, resVelocitys)
	a.Equal([]string{"velocity receive finish"}, supportMock.messages)
}

// TestReceiveErrorWhenAdapterVelocitySubscriber .
func TestReceiveErrorWhenAdapterVelocitySubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientVelocityMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(nil)

	stream := AdapterVelocitySubscriber(receiverMock, supportMock)

	var resVelocitys []*model.Velocity
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resVelocitys = append(resVelocitys, a)
	}

	a.Len(resVelocitys, 0)
	a.Equal([]string{"velocity receive error: receive error"}, supportMock.messages)
}

// TestCloseErrorWhenAdapterVelocitySubscriber .
func TestCloseErrorWhenAdapterVelocitySubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientVelocityMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, io.EOF)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterVelocitySubscriber(receiverMock, supportMock)

	var resVelocitys []*model.Velocity
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resVelocitys = append(resVelocitys, a)
	}

	a.Len(resVelocitys, 0)
	a.Equal([]string{"velocity receive finish", "velocity telemetry error: close error"}, supportMock.messages)
}

// TestReceiveAndCloseErrorWhenAdapterVelocitySubscriber .
func TestReceiveAndCloseErrorWhenAdapterVelocitySubscriber(t *testing.T) {
	a := assert.New(t)

	supportMock := &supportMock{}

	receiverMock := &telemetryServiceClientVelocityMock{
		i: 1,
	}
	receiverMock.On("Recv", 1).Return(nil, ErrReceive)
	receiverMock.On("CloseSend").Return(ErrClose)

	stream := AdapterVelocitySubscriber(receiverMock, supportMock)

	var resVelocitys []*model.Velocity
	for {
		a, ok := <-stream
		if !ok {
			break
		}
		resVelocitys = append(resVelocitys, a)
	}

	a.Len(resVelocitys, 0)
	a.Equal([]string{"velocity receive error: receive error", "velocity telemetry error: close error"}, supportMock.messages)
}
