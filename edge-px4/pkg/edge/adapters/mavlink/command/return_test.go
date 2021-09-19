package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_action "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/action"
)

// TestAdapterReturn .
func TestAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.ReturnToLaunchResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_SUCCESS,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterReturnInternal(ctx, actionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterReturn .
func TestRequestErrorWhenAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterReturnInternal(ctx, actionMock)

	a.Equal("rtl command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterReturn .
func TestResponseErrorWhenAdapterReturn(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_action.ReturnToLaunchResponse{
		ActionResult: &mavsdk_rpc_action.ActionResult{
			Result: mavsdk_rpc_action.ActionResult_BUSY,
		},
	}
	actionMock := &actionServiceClientMock{}
	actionMock.On("ReturnToLaunch", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterReturnInternal(ctx, actionMock)

	a.Equal("rtl command error: no rtl command success", ret.Error())
}
