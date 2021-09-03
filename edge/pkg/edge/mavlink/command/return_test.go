package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "edge/pkg/protos/action"
)

// TestAdapterReturn .
func TestAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.ReturnToLaunchResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterReturnInternal(ctx, supportMock, actionMock)

	a.Nil(ret)
	a.Empty(supportMock.message)
}

// TestRequestErrorWhenAdapterReturn .
func TestRequestErrorWhenAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterReturnInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("rtl command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterReturn .
func TestResponseErrorWhenAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_action.ReturnToLaunchResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterReturnInternal(ctx, supportMock, actionMock)

	a.Equal(ret, ErrReturnCommand)
	a.Equal("rtl command error: no rtl command success", supportMock.message)
}
