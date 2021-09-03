package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// TestAdapterTakeoff .
func TestAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.TakeoffResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterTakeoffInternal(ctx, supportMock, actionMock)

	a.Nil(ret)
	a.Empty(supportMock.message)
}

// TestRequestErrorWhenAdapterTakeoff .
func TestRequestErrorWhenAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterTakeoffInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("takeoff command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterTakeoff .
func TestResponseErrorWhenAdapterTakeoff(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.TakeoffResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("Takeoff", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterTakeoffInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrTakeoffCommand)
	a.Equal("takeoff command error: no takeoff command success", supportMock.message)
}
